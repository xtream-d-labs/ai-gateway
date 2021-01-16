package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/docker/docker/api/types"
	docker "github.com/docker/docker/client"
	"github.com/go-openapi/swag"
	"github.com/xtream-d-labs/ai-gateway/api/src/auth"
	"github.com/xtream-d-labs/ai-gateway/api/src/config"
	"github.com/xtream-d-labs/ai-gateway/api/src/db"
	"github.com/xtream-d-labs/ai-gateway/api/src/kubernetes"
	"github.com/xtream-d-labs/ai-gateway/api/src/lib"
	"github.com/xtream-d-labs/ai-gateway/api/src/log"
	"github.com/xtream-d-labs/ai-gateway/api/src/rescale"
)

func init() {
	db.SetupQueue(worker, fallback)
}

// SubmitPullImageJob puts a new job
func SubmitPullImageJob(imageName, authConfig, principal string) error {
	bytes, err := json.Marshal(job{
		Action: actionPullImage,
		Arg1:   imageName,
		Arg2:   authConfig,
		Arg3:   principalName(principal),
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
		Arg5:   principalName(builderName),
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
		Arg3:   principalName(builderName),
		Arg4:   k8sConfig,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// PushJobDockerImage pushes a docker image
func PushJobDockerImage(jobID, imageName, dockerCredential, k8sConfig, principal string) error {
	bytes, err := json.Marshal(job{
		Action: actionPushJobImg,
		Arg1:   jobID,
		Arg2:   imageName,
		Arg3:   dockerCredential,
		Arg4:   k8sConfig,
		Arg5:   principalName(principal),
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
		Arg4:   principalName(builderName),
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// SubmitToRescaleJob puts a new job
func SubmitToRescaleJob(id, rescaleConfig, simg, principal string) error {
	bytes, err := json.Marshal(job{
		Action: actionSendRescale,
		Arg1:   id,
		Arg2:   rescaleConfig,
		Arg3:   simg,
		Arg4:   principalName(principal),
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

func principalName(candidate string) string {
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
		// Arg3: Principal
		if err = pullImage(ctx, j.Arg1, j.Arg2); err != nil {
			db.ImageError(
				j.Arg3, "Could not pull the specified image",
				db.ImageActionPulling, j.Arg1, err,
			)
			return err
		}
	case actionWrapJupyter:
		// Arg1: DockerImageName
		// Arg2: DockerImageID
		// Arg3: WorkspaceHostDirectory
		// Arg4: WrappedDockerImageID
		// Arg5: Principal
		wrappedImage := swag.String(j.Arg4)
		if swag.IsZero(j.Arg4) {
			wrappedImage, err = lib.WrapWithJupyterNotebook(ctx, j.Arg2, j.Arg1, j.Arg5)
			if err != nil {
				db.OnlineError(
					j.Arg5, "Failed to wrap the image with jupyter notebook",
					db.OnlineActionWrapJupyter, err,
				)
				return err
			}
		}
		workdir := lib.DetectImageWorkDir(ctx, swag.StringValue(wrappedImage))
		ID, e1 := lib.RunJupyterNotebook(ctx, j.Arg1, swag.StringValue(wrappedImage), j.Arg3, workdir)
		if e1 != nil {
			db.OnlineError(
				j.Arg5, "Failed to start a jupyter notebook",
				db.OnlineActionRunJupyter, e1,
			)
			return e1
		}
		if err = db.RemoveBuildingImage(j.Arg1); err != nil {
			log.Warn("Failed to remove a building image", err, nil)
		}
		log.Debug("Run JupyterNotebook!", nil, &log.Map{
			"id": swag.StringValue(ID),
		})
	case actionBuildJobImg:
		// Arg1: ID
		// Arg2: DockerCredential
		// Arg3: Principal
		// Arg4: KubernetesConfig
		name, e2 := lib.BuildJobImage(ctx, j.Arg1, j.Arg3, true)
		if e2 != nil {
			db.JobError(
				j.Arg3, "Failed to build an image for the task",
				db.JobActionBuilding, j.Arg1, e2,
			)
			return e2
		}
		if err = db.UpdateJob(
			j.Arg1,
			db.JobActionBuilding,
			db.JobActionPushing,
			db.JobStatusImagePushing,
			nil,
		); err != nil {
			db.OnlineError(
				j.Arg3, "Failed to update the task status",
				db.OnlineActionUpdateStatus, err,
			)
			return err
		}
		if err = PushJobDockerImage(j.Arg1, swag.StringValue(name), j.Arg2, j.Arg4, j.Arg3); err != nil {
			db.OnlineError(
				j.Arg3, "Failed to update the task status",
				db.OnlineActionUpdateStatus, err,
			)
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
		// Arg5: Principal
		job, e3 := db.GetJob(j.Arg1)
		if e3 != nil {
			db.OnlineError(
				j.Arg5, "Failed to retrieve the task configuration",
				db.OnlineActionSelectData, e3,
			)
			return e3
		}
		if err = lib.PushJobImage(ctx, j.Arg2, j.Arg3); err != nil {
			db.JobError(
				j.Arg5, "Failed to push the image to cloud",
				db.JobActionPushing, j.Arg1, err,
			)
			return err
		}
		log.Debug("Pushed the Job Image!", nil, &log.Map{
			"id":   job.JobID,
			"name": j.Arg2,
		})
		if err = lib.DeleteImage(ctx, j.Arg2); err != nil {
			log.Warn("Failed to remove the task image", err, nil)
		}
		if err = kubernetes.CreateJob(j.Arg4, j.Arg1, "default", j.Arg2, job.CPU, job.GPU); err != nil {
			db.JobError(
				j.Arg5, "Failed to create a kubernetes job",
				db.JobActionKubernetes, j.Arg1, err,
			)
			return err
		}
		// Change this job's status
		if err = db.UpdateJob(
			job.JobID,
			db.JobActionPushing,
			db.JobActionKubernetes,
			db.JobStatusK8sStarted,
			nil,
		); err != nil {
			db.OnlineError(
				j.Arg5, "Failed to update the task status",
				db.OnlineActionUpdateStatus, err,
			)
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
		// Arg4: Principal
		name, e4 := lib.BuildJobImage(ctx, j.Arg1, j.Arg4, false)
		if e4 != nil {
			db.ImageError(
				j.Arg4, "Failed to build a docker image for the task",
				db.ImageActionBuilding, swag.StringValue(name), e4,
			)
			return e4
		}
		simg, e5 := lib.ConvertToSingularityImage(j.Arg1, swag.StringValue(name))
		if e5 != nil {
			db.OnlineError(
				j.Arg4, "Failed to convert the docker image to singularity",
				db.OnlineActionConvertToSig, e5,
			)
			return e5
		}
		if err = lib.DeleteImage(ctx, swag.StringValue(name)); err != nil {
			log.Warn("Failed to remove a temporary docker image", err, nil)
		}
		if err = db.UpdateJob(
			j.Arg1,
			db.JobActionBuilding,
			db.JobActionRescale,
			db.JobStatusRescaleSend,
			nil,
		); err != nil {
			db.OnlineError(
				j.Arg4, "Failed to update the task status",
				db.OnlineActionUpdateStatus, err,
			)
			return err
		}
		if err = SubmitToRescaleJob(j.Arg1, j.Arg3, swag.StringValue(simg), j.Arg4); err != nil {
			db.OnlineError(
				j.Arg4, "Failed to update the task status",
				db.OnlineActionUpdateStatus, err,
			)
			return err
		}
		log.Debug("Built Singularity Image!", nil, nil)

	case actionSendRescale:
		// Arg1: ID
		// Arg2: RescaleConfig
		// Arg3: SingularityImage
		// Arg4: Principal
		// Upload the singularity image
		job, e3 := db.GetJob(j.Arg1)
		if e3 != nil {
			db.OnlineError(
				j.Arg4, "Failed to retrieve the task configuration",
				db.OnlineActionSelectData, e3,
			)
			return e3
		}
		meta, err := rescale.Upload(ctx, j.Arg2, j.Arg3)
		if err != nil {
			db.JobError(
				j.Arg4, "Failed to upload files to Rescale",
				db.JobActionRescale, j.Arg1, err,
			)
			return err
		}
		// Create a new job
		input := rescale.JobInput{
			Name: strings.TrimRight(filepath.Base(j.Arg3), ".simg"),
			JobAnalyses: []rescale.JobInputAnalyse{{
				Command: fmt.Sprintf("singularity run %s\nrm -rf %s", meta.Name, meta.Name),
				InputFiles: []rescale.JobInputFile{{
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
			db.JobError(
				j.Arg4, "Failed to create a Rescale job",
				db.JobActionRescale, j.Arg1, err,
			)
			return err
		}
		// Submit the job
		if err = rescale.Submit(ctx, j.Arg2, swag.StringValue(jobID)); err != nil {
			db.JobError(
				j.Arg4, "Failed to start the Rescale job",
				db.JobActionRescale, j.Arg1, err,
			)
			return err
		}
		// Change this job's status
		if err = db.UpdateJob(
			j.Arg1,
			db.JobActionRescale,
			db.JobActionRescale,
			db.JobStatusRescaleStarted,
			jobID,
		); err != nil {
			db.OnlineError(
				j.Arg5, "Failed to update the task status",
				db.OnlineActionUpdateStatus, err,
			)
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
	cli, err := docker.NewClientWithOpts(docker.FromEnv)
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
