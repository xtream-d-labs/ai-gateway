package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/kubernetes"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/rescale"
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
		Arg5:   buildersName(builderName),
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// BuildJobDockerImage puts a new job
func BuildJobDockerImage(id, dockerCredential, builderName, k8sConfig string) error {
	bytes, err := json.Marshal(job{
		Action: actionBuildJobImg,
		Arg1:   id,
		Arg2:   dockerCredential,
		Arg3:   buildersName(builderName),
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

// BuildSingularityImageJob puts a new job
func BuildSingularityImageJob(id, dockerCredential, rescaleConfig, builderName string) error {
	bytes, err := json.Marshal(job{
		Action: actionSingularity,
		Arg1:   id,
		Arg2:   dockerCredential,
		Arg3:   rescaleConfig,
		Arg4:   buildersName(builderName),
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// SubmitToRescaleJob puts a new job
func SubmitToRescaleJob(id, rescaleConfig, simg string) error {
	bytes, err := json.Marshal(job{
		Action: actionSendRescale,
		Arg1:   id,
		Arg2:   rescaleConfig,
		Arg3:   simg,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

func buildersName(candidate string) string {
	if candidate == "" {
		return auth.Anonymous
	}
	return candidate
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
			log.Error("Worker failed at RunJupyterNotebook", e1, nil)
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
		name, e2 := lib.BuildJobImage(ctx, j.Arg1, j.Arg3, true)
		if e2 != nil {
			log.Error("Worker failed at BuildJobImage", e2, nil)
			return e2
		}
		err = db.UpdateJob(j.Arg1, db.JobActionBuilding, db.JobActionPushing, db.JobStatusImagePushing, nil)
		if err != nil {
			log.Error("Worker failed at UpdateJob@actionBuildJobImg", err, nil)
			return err
		}
		if err = PushJobDockerImage(j.Arg1, swag.StringValue(name), j.Arg2, j.Arg4); err != nil {
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
		job, e3 := db.GetJob(j.Arg1)
		if e3 != nil {
			log.Error("Worker failed at GetJob@actionPushJobImg", e3, nil)
			return e3
		}
		if err = lib.PushJobImage(ctx, j.Arg2, j.Arg3); err != nil {
			log.Error("Worker failed at PushJobImage", err, nil)
			return err
		}
		log.Debug("Pushed the Job Image!", nil, &log.Map{
			"id":   job.JobID,
			"name": j.Arg2,
		})
		if err = lib.DeleteImage(ctx, j.Arg2); err != nil {
			log.Error("Worker failed at DeleteImage", err, nil)
			return err
		}
		err = kubernetes.CreateJob(j.Arg4, j.Arg1, "default", j.Arg2, job.CPU, job.GPU)
		if err != nil {
			return err
		}
		// Change this job's status
		err = db.UpdateJob(job.JobID, db.JobActionPushing, db.JobActionKubernetes, db.JobStatusK8sStarted, nil)
		if err != nil {
			log.Error("Worker failed at UpdateJob@actionApplyK8s", err, nil)
			return err
		}
		log.Debug("Job submitted!", nil, &log.Map{
			"id":   job.JobID,
			"name": j.Arg2,
			"cmds": job.Commands,
		})
	case actionSingularity:
		// Arg1: ID
		// Arg2: DockerCredential
		// Arg3: RescaleConfig
		// Arg4: BuildersName
		name, e4 := lib.BuildJobImage(ctx, j.Arg1, j.Arg4, false)
		if e4 != nil {
			log.Error("Worker failed at BuildJobImage", e4, nil)
			return e4
		}
		simg, e5 := lib.ConvertToSingularityImage(j.Arg1, swag.StringValue(name))
		if e5 != nil {
			log.Error("Worker failed at BuildSingularityImage", e5, nil)
			return e5
		}
		if err = lib.DeleteImage(ctx, swag.StringValue(name)); err != nil {
			log.Error("Worker failed at DeleteImage", err, nil)
			return err
		}
		err = db.UpdateJob(j.Arg1, db.JobActionBuilding, db.JobActionRescale, db.JobStatusRescaleSend, nil)
		if err != nil {
			log.Error("Worker failed at UpdateJob", err, nil)
			return err
		}
		if err = SubmitToRescaleJob(j.Arg1, j.Arg3, swag.StringValue(simg)); err != nil {
			log.Error("Worker failed at SubmitToRescaleJob", err, nil)
			return err
		}
		log.Debug("Built Singularity Image!", nil, nil)

	case actionSendRescale:
		// Arg1: ID
		// Arg2: RescaleConfig
		// Arg3: SingularityImage
		// Upload the singularity image
		job, e3 := db.GetJob(j.Arg1)
		if e3 != nil {
			log.Error("Worker failed at GetJob@actionSendRescale", e3, nil)
			return e3
		}
		meta, err := rescale.Upload(ctx, j.Arg2, j.Arg3)
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
		jobID, err := rescale.CreateJob(ctx, j.Arg2, input)
		if err != nil {
			log.Error("Worker failed at CreateJob@SendRescaleJob", err, nil)
			return err
		}
		// Submit the job
		if err = rescale.Submit(ctx, j.Arg2, swag.StringValue(jobID)); err != nil {
			log.Error("Worker failed at Submit@SendRescaleJob", err, nil)
			return err
		}
		// Change this job's status
		err = db.UpdateJob(j.Arg1, db.JobActionRescale, db.JobActionRescale, db.JobStatusRescaleStarted, jobID)
		if err != nil {
			log.Error("Worker failed at UpdateJob@SendRescaleJob", err, nil)
			return err
		}
		log.Debug("Job submitted!", nil, &log.Map{
			"ID":   jobID,
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
	case actionBuildJobImg:
		return db.RemoveBuildingJobImagesJobs(j.Arg1)
	case actionPushJobImg:
		return db.RemovePushingJobImageJobs(j.Arg1)
	case actionApplyK8s:
		return db.RemoveJob(j.Arg1)
	case actionSingularity:
		return db.RemoveBuildingJobImagesJobs(j.Arg1)
	case actionSendRescale:
		return db.RemoveJob(j.Arg1)
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
