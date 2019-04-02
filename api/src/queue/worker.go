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

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
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
		Arg5:   builderName,
	})
	if err != nil {
		return err
	}
	return db.Enqueue(string(bytes))
}

// BuildSingularityImageJob puts a new job
func BuildSingularityImageJob(ID, ngcConfig, rescaleConfig, builderName string) error {
	bytes, err := json.Marshal(job{
		Action: actionSingularity,
		Arg1:   ID,
		Arg2:   ngcConfig,
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
		ID, err := lib.RunJupyterNotebook(ctx, j.Arg1, swag.StringValue(wrappedImage), j.Arg3, workdir)
		if err != nil {
			log.Error("Worker failed at RunJupyterNotebook", err, nil)
			return err
		}
		if err = db.RemoveBuildingImage(j.Arg1); err != nil {
			return err
		}
		log.Debug("Run JupyterNotebook!", nil, &log.Map{
			"id": swag.StringValue(ID),
		})
	case actionSingularity:
		// Arg1: ID
		// Arg2: NgcConfig
		// Arg3: RescaleConfig
		// Arg4: BuildersName
		simg, err := lib.BuildSingularityImage(j.Arg1, j.Arg2, j.Arg4)
		if err != nil {
			log.Error("Worker failed at BuildSingularityImage", err, nil)
			return err
		}
		err = db.UpdateJob(j.Arg1, db.SingularityStoreKey, db.RescaleJobStoreKey, db.RescaleSend)
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
		application := rescale.ApplicationSingularity
		if strings.EqualFold(job.CoreType, "dolomite") {
			application = rescale.ApplicationSingularityMPI
		}
		input := rescale.JobInput{
			Name: strings.TrimRight(filepath.Base(j.Arg3), ".simg"),
			JobAnalyses: []rescale.JobInputAnalyse{rescale.JobInputAnalyse{
				Command: fmt.Sprintf("singularity run %s", meta.Name),
				InputFiles: []rescale.JobInputFile{rescale.JobInputFile{
					ID:         meta.ID,
					Decompress: true,
				}},
				Analysis: rescale.JobAnalyse{
					Code:    application,
					Version: "2.5.1", // FIXME
				},
				Hardware: rescale.JobHardware{
					Type: "compute",
					CoreType: rescale.JobCoreType{
						Code: job.CoreType,
					},
					Slots:        1,   // FIXME
					CoresPerSlot: 1,   // FIXME CoresPerSlot: fmt.Sprintf("%d", job.Cores),
					WallTime:     750, // TODO
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
		return db.RemoveSingularityJobs(j.Arg1)
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
