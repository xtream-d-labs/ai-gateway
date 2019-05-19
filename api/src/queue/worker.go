package queue

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/rescale"
	batchV1 "k8s.io/api/batch/v1"
	coreV1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func init() {
	db.SetupQueue(worker, fallback)
}

// SubmitPullImageJob puts a new job
func SubmitPullImageJob(imageName, authConfig string) error {
	bytes, err := json.Marshal(job{
		Action: actionPullImage,
		Arg1:   imageName,
		Arg2:   authConfig,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// SubmitBuildImageJob puts a new job
func SubmitBuildImageJob(imageName, imageID, workspace, wrappedImageID, builderName string) error {
	bytes, err := json.Marshal(job{
		Action: actionWrapJupyter,
		Arg1:   imageName,
		Arg2:   imageID,
		Arg3:   workspace,
		Arg4:   wrappedImageID,
		Arg5:   builderName,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// BuildJobDockerImage puts a new job
func BuildJobDockerImage(ID, dockerCredential, builderName, k8sConfig string) error {
	bytes, err := json.Marshal(job{
		Action: actionBuildJobImg,
		Arg1:   ID,
		Arg2:   dockerCredential,
		Arg3:   builderName,
		Arg4:   k8sConfig,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// PushJobDockerImage pushes a docker image
func PushJobDockerImage(jobID, imageName, dockerCredential, k8sConfig string) error {
	bytes, err := json.Marshal(job{
		Action: actionPushJobImg,
		Arg1:   jobID,
		Arg2:   imageName,
		Arg3:   dockerCredential,
		Arg4:   k8sConfig,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// ApplyKubernetesJob puts a new job
func ApplyKubernetesJob(jobID, imageName, k8sConfig string) error {
	bytes, err := json.Marshal(job{
		Action: actionApplyK8s,
		Arg1:   jobID,
		Arg2:   imageName,
		Arg3:   k8sConfig,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// BuildSingularityImageJob puts a new job
func BuildSingularityImageJob(ID, dockerCredential, rescaleConfig, builderName string) error {
	bytes, err := json.Marshal(job{
		Action: actionSingularity,
		Arg1:   ID,
		Arg2:   dockerCredential,
		Arg3:   rescaleConfig,
		Arg4:   builderName,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// SubmitToRescaleJob puts a new job
func SubmitToRescaleJob(ID, rescaleConfig, simg string) error {
	bytes, err := json.Marshal(job{
		Action: actionSendRescale,
		Arg1:   ID,
		Arg2:   rescaleConfig,
		Arg3:   simg,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

type actionType int

const (
	actionPullImage = iota
	actionWrapJupyter
	actionBuildJobImg
	actionPushJobImg
	actionApplyK8s
	actionSingularity
	actionSendRescale
)

type job struct {
	Action actionType
	Arg1   string
	Arg2   string
	Arg3   string
	Arg4   string
	Arg5   string
}

func worker(name string) (err error) {
	var j job
	if err = json.Unmarshal([]byte(name), &j); err != nil {
		log.Error("Worker 'Unmarshal arguments' failed", err, nil)
		return err
	}
	ctx := context.Background()

	switch j.Action {
	case actionPullImage:
		// Arg1: DockerImageName
		// Arg2: AuthConfig
		if err = pullImage(ctx, j.Arg1, j.Arg2); err != nil {
			log.Error("Worker failed at PullImage", err, nil)
			return err
		}
	case actionWrapJupyter:
		// Arg1: DockerImageName
		// Arg2: DockerImageID
		// Arg3: WorkspaceHostDirectory
		// Arg4: WrappedDockerImageID
		// Arg5: BuildersName
		wrappedImage := swag.String(j.Arg4)
		if swag.IsZero(j.Arg4) {
			wrappedImage, err = lib.WrapWithJupyterNotebook(ctx, j.Arg2, j.Arg1, j.Arg5)
			if err != nil {
				log.Error("Worker failed at WrapImageWithJupyterNotebook", err, nil)
				return err
			}
		}
		workdir := lib.DetectImageWorkDir(ctx, swag.StringValue(wrappedImage))
		ID, e1 := lib.RunJupyterNotebook(ctx, j.Arg1, swag.StringValue(wrappedImage), j.Arg3, workdir)
		if e1 != nil {
			log.Error("Worker failed at RunJupyterNotebook", err, nil)
			return e1
		}
		if err = db.RemoveBuildingImage(j.Arg1); err != nil {
			return err
		}
		log.Debug("Run JupyterNotebook!", nil, &log.Map{
			"id": swag.StringValue(ID),
		})
	case actionBuildJobImg:
		// Arg1: ID
		// Arg2: DockerCredential
		// Arg3: BuildersName
		// Arg4: KubernetesConfig
		name, e2 := lib.BuildJobImage(ctx, j.Arg1, j.Arg3)
		if e2 != nil {
			log.Error("Worker failed at BuildJobImage", err, nil)
			return e2
		}
		err = db.UpdateJob(j.Arg1, db.BuildingJobStoreKey, db.PushingJobStoreKey, db.PushingJob)
		if err != nil {
			log.Error("Worker failed at UpdateJob@actionBuildJobImg", err, nil)
			return err
		}
		err = PushJobDockerImage(j.Arg1, swag.StringValue(name), j.Arg2, j.Arg4)
		if err != nil {
			log.Error("Worker failed at PushJobDockerImage", err, nil)
			return err
		}
		log.Debug("Built Job Image!", nil, &log.Map{
			"id":   j.Arg1,
			"name": swag.StringValue(name),
		})
	case actionPushJobImg:
		// Arg1: JobID
		// Arg2: ImageName
		// Arg3: DockerCredential
		// Arg4: KubernetesConfig
		err = lib.PushJobImage(ctx, j.Arg2, j.Arg3)
		if err != nil {
			log.Error("Worker failed at PushJobImage", err, nil)
			return err
		}
		err = lib.DeleteImage(ctx, j.Arg2)
		if err != nil {
			log.Error("Worker failed at DeleteImage", err, nil)
			return err
		}
		err = db.UpdateJob(j.Arg1, db.PushingJobStoreKey, db.KubernetesJobStoreKey, db.KubernetesJob)
		if err != nil {
			log.Error("Worker failed at UpdateJob@actionPushJobImg", err, nil)
			return err
		}
		err = ApplyKubernetesJob(j.Arg1, j.Arg2, j.Arg4)
		if err != nil {
			log.Error("Worker failed at ApplyKubernetesJob", err, nil)
			return err
		}
		log.Debug("Pushed the Job Image!", nil, &log.Map{
			"id":   j.Arg1,
			"name": j.Arg2,
		})
	case actionApplyK8s:
		// Arg1: JobID
		// Arg2: ImageName
		// Arg3: KubernetesConfig
		job, e3 := db.GetJob(j.Arg1)
		if e3 != nil {
			log.Error("Worker failed at GetJob@actionApplyK8s", err, nil)
			return e3
		}
		// Apply a k8s job
		config, e4 := clientcmd.RESTConfigFromKubeConfig([]byte(j.Arg3))
		if e4 != nil {
			return e4
		}
		clientset, e5 := kubernetes.NewForConfig(config)
		if e5 != nil {
			return e5
		}
		pod := coreV1.PodTemplateSpec{
			Spec: coreV1.PodSpec{
				Containers: []coreV1.Container{
					coreV1.Container{
						Name:  "main",
						Image: j.Arg2,
						Resources: coreV1.ResourceRequirements{
							Requests: coreV1.ResourceList{
								"cpu": resource.MustParse(fmt.Sprintf("%d", job.CPU)),
							},
						},
					},
				},
				RestartPolicy: coreV1.RestartPolicyNever,
			},
		}
		if job.GPU > 0 {
			// https://kubernetes.io/docs/tasks/manage-gpus/scheduling-gpus/
			pod.Spec.Containers[0].Resources.Limits = coreV1.ResourceList{
				"nvidia.com/gpu": resource.MustParse(fmt.Sprintf("%d", job.GPU)),
			}
		}
		_, err = clientset.BatchV1().Jobs("default").Create(&batchV1.Job{
			ObjectMeta: metav1.ObjectMeta{
				Name: fmt.Sprintf("job-%s", time.Now().Format("060102150405")),
			},
			Spec: batchV1.JobSpec{
				Completions:  swag.Int32(1),
				Parallelism:  swag.Int32(1),
				BackoffLimit: swag.Int32(1),
				Template:     pod,
			},
		})
		if err != nil {
			return err
		}
		// Change this job's status
		err = db.UpdateJob(j.Arg1, db.KubernetesJobStoreKey, db.KubernetesJobStoreKey, db.K8sJobStart)
		if err != nil {
			log.Error("Worker failed at UpdateJob@actionApplyK8s", err, nil)
			return err
		}
		log.Debug("Job submitted!", nil, &log.Map{
			"id":   j.Arg1,
			"name": j.Arg2,
			"cmds": job.Commands,
		})
	case actionSingularity:
		// Arg1: ID
		// Arg2: DockerCredential
		// Arg3: RescaleConfig
		// Arg4: BuildersName
		simg, err := lib.BuildSingularityImage(j.Arg1, j.Arg2, j.Arg4)
		if err != nil {
			log.Error("Worker failed at BuildSingularityImage", err, nil)
			return err
		}
		err = db.UpdateJob(j.Arg1, db.BuildingJobStoreKey, db.RescaleJobStoreKey, db.RescaleSend)
		if err != nil {
			log.Error("Worker failed at UpdateJob", err, nil)
			return err
		}
		err = SubmitToRescaleJob(j.Arg1, j.Arg3, swag.StringValue(simg))
		if err != nil {
			log.Error("Worker failed at SubmitToRescaleJob", err, nil)
			return err
		}
		log.Debug("Built Singularity Image!", nil, &log.Map{
			"simg": swag.StringValue(simg),
		})
	case actionSendRescale:
		// Arg1: ID
		// Arg2: RescaleConfig
		// Arg3: SingularityImage
		job, err := db.GetJob(j.Arg1)
		if err != nil {
			log.Error("Worker failed at GetJob@SendRescaleJob", err, nil)
			return err
		}
		// Upload the singularity image
		body, contentType, err := loadSingularityFile(j.Arg3)
		if err != nil {
			log.Error("Worker failed at loadSingularityFile@SendRescaleJob", err, nil)
			return err
		}
		meta, err := rescale.Upload(j.Arg2, body, contentType)
		if err != nil {
			log.Error("Worker failed at Upload@SendRescaleJob", err, nil)
			return err
		}
		// Create a new job
		input := rescale.JobInput{
			Name: strings.TrimRight(filepath.Base(j.Arg3), ".simg"),
			JobAnalyses: []rescale.JobInputAnalyse{rescale.JobInputAnalyse{
				Command: fmt.Sprintf("singularity run %s\nrm -rf %s", meta.Name, meta.Name),
				InputFiles: []rescale.JobInputFile{rescale.JobInputFile{
					ID:         meta.ID,
					Decompress: true,
				}},
				Analysis: rescale.JobAnalyse{
					Code:    rescale.ApplicationSingularity,
					Version: config.Config.RescaleSingularityVer,
				},
				Hardware: rescale.JobHardware{
					Type: "compute",
					CoreType: rescale.JobCoreType{
						Code: job.CoreType,
					},
					Slots:        1,
					CoresPerSlot: int(job.Cores),
					WallTime:     config.Config.RescaleJobWallTime,
				},
			}},
			JobVariables:     []string{},
			IsLowPriority:    false,
			IsTemplateDryRun: false,
		}
		jobID, err := rescale.CreateJob(j.Arg2, input)
		if err != nil {
			log.Error("Worker failed at CreateJob@SendRescaleJob", err, nil)
			return err
		}
		// Submit the job
		_, err = rescale.Submit(j.Arg2, swag.StringValue(jobID))
		if err != nil {
			log.Error("Worker failed at Submit@SendRescaleJob", err, nil)
			return err
		}
		// Change this job's status
		err = db.UpdateJob(j.Arg1, db.RescaleJobStoreKey, db.RescaleJobStoreKey, db.RescaleStart)
		if err != nil {
			log.Error("Worker failed at UpdateJob@SendRescaleJob", err, nil)
			return err
		}
		log.Debug("Job submitted!", nil, &log.Map{
			"file": meta.ID,
			"cmds": job.Commands,
			"core": job.CoreType,
		})
	}
	return nil
}

func fallback(name string) error {
	var j job
	err := json.Unmarshal([]byte(name), &j)
	if err != nil {
		log.Error("Worker 'Unmarshal arguments' failed", err, nil)
		return err
	}
	switch j.Action {
	case actionPullImage:
		return db.RemovePullingImage(j.Arg1)
	case actionWrapJupyter:
		return db.RemoveBuildingImage(j.Arg1)
	case actionSingularity:
		return db.RemoveBuildingJobImagesJobs(j.Arg1)
	case actionSendRescale:
		return db.RemoveRescaleJob(j.Arg1)
	}
	return nil
}

// Pull a new Docker image
func pullImage(ctx context.Context, imageName, authConfig string) error {
	log.Info("Worker triggers 'pullImage'", nil, &log.Map{
		"image": imageName,
	})
	cli, err := docker.NewEnvClient()
	if err != nil {
		return err
	}
	defer cli.Close()

	options := types.ImagePullOptions{}
	if strings.HasPrefix(imageName, config.Config.DockerRegistryHostName) {
		options.RegistryAuth = authConfig
	}
	if strings.HasPrefix(imageName, config.Config.NgcRegistryHostName) {
		options.RegistryAuth = authConfig
	}
	reader, err := cli.ImagePull(ctx, imageName, options)
	if err != nil {
		return err
	}
	defer reader.Close()
	_, err = io.Copy(ioutil.Discard, reader) // wait for its done
	if err != nil {
		return err
	}
	return db.RemovePullingImage(imageName)
}

func loadSingularityFile(path string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, "", err
	}
	file, err := os.Open(path)
	if err != nil {
		return nil, "", err
	}
	defer file.Close()

	if _, err = io.Copy(part, file); err != nil {
		return nil, "", err
	}
	if err = writer.Close(); err != nil {
		return nil, "", err
	}
	return body, writer.FormDataContentType(), nil
}
