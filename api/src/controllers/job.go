package controllers

import (
	"context"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"docker.io/go-docker/api/types/mount"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/job"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/queue"
)

func jobRoute(api *operations.ScaleShiftAPI) {
	api.JobGetJobsHandler = job.GetJobsHandlerFunc(getJobs)
	api.JobPostNewJobHandler = job.PostNewJobHandlerFunc(postNewJob)
	api.JobModifyJobHandler = job.ModifyJobHandlerFunc(modifyJob)
	api.JobDeleteJobHandler = job.DeleteJobHandlerFunc(deleteJob)
}

func getJobs(params job.GetJobsParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	payload := []*models.Job{}
	if jobs, err := db.GetJobs(); err == nil {
		for _, job := range jobs {
			payload = append(payload, &models.Job{
				ID:       swag.String(job.ID),
				Status:   job.Status,
				Image:    job.DockerImage,
				Mounts:   job.Workspaces,
				Commands: job.Commands,
				Started:  strfmt.DateTime(job.Started),
			})
		}
	}
	return job.NewGetJobsOK().WithPayload(payload)
}

func postNewJob(params job.PostNewJobParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.Base.K8sConfig) && swag.IsZero(creds.Base.RescaleKey) {
		code := http.StatusForbidden
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	// FIXME allow only emerald yet, dolomite will come soon
	if params.Body.Coretype != "emerald" {
		code := http.StatusNotAcceptable
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}

	ctx := context.Background()
	container, err := cli.ContainerInspect(ctx, params.Body.NotebookID)
	if err != nil {
		log.Error("ContainerInspect@postNewJob", err, nil)
		code := http.StatusBadRequest
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	mounts := []mount.Mount{}
	workspaces := []string{}
	src := ""
	for _, mnt := range container.Mounts {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: mnt.Source,
			Target: mnt.Destination,
		})
		workspaces = append(workspaces, strings.TrimLeft(strings.Replace(
			mnt.Source, config.Config.WorkspaceHostDir, "", -1), "/"))
		src = strings.Replace(mnt.Source, config.Config.WorkspaceHostDir,
			config.Config.WorkspaceContainerDir, 1)
	}
	ipynb := params.Body.EntrypointFile
	if ipynb != "none" {
		nb := lib.ParseIPython(filepath.Join(src, ipynb))
		cmd := "python"
		if strings.EqualFold(nb.Meta.KernelSpec.Lang, "bash") {
			cmd = "script"
		}
		workdir := lib.DetectImageWorkDir(ctx, container.Image)
		if err := lib.ConvertNotebook(ctx, container.Image, ipynb, cmd, workdir, mounts); err != nil {
			log.Error("ConvertNotebook@postNewJob", err, nil)
			code := http.StatusBadRequest
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
		ipynb = strings.Replace(ipynb, ".ipynb", nb.Meta.Lang.FileExt, -1)
	}
	commands := []string{}
	for _, cmd := range params.Body.Commands {
		switch {
		case strings.HasSuffix(ipynb, ".py"):
			cmd = strings.Replace(cmd, "<converted-notebook.py>",
				filepath.Join("/workspace", ipynb), -1)
		case strings.HasSuffix(ipynb, ".sh"):
			if strings.EqualFold(cmd, "python") {
				cmd = "bash"
			}
			cmd = strings.Replace(cmd, "<converted-notebook.py>",
				filepath.Join("/workspace", ipynb), -1)
		}
		if command := strings.TrimSpace(cmd); command != "" {
			commands = append(commands, command)
		}
	}
	if len(commands) < 1 {
		code := http.StatusBadRequest
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	platform := db.PlatformKubernetes
	if params.Body.PlatformID == models.PostNewJobParamsBodyPlatformIDRescale {
		platform = db.PlatformRescale
	}
	image, _, _ := lib.ContainerAttrs(container.Config.Labels)
	newjob := &db.Job{
		Platform:    platform,
		ID:          uuid.New().String(),
		Status:      db.BuildingJob,
		DockerImage: image,
		PythonFile:  ipynb,
		Workspaces:  workspaces,
		Commands:    commands,
		CPU:         params.Body.CPU,
		Memory:      params.Body.Mem,
		GPU:         params.Body.Gpu,
		CoreType:    params.Body.Coretype,
		Cores:       params.Body.Cores,
		Started:     time.Now(),
	}
	if err := db.SetJobMeta(newjob); err != nil {
		log.Error("SetJobMeta@postNewJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	credential := ""
	if strings.HasPrefix(image, config.Config.DockerRegistryHostName) {
		credential = creds.Base.DockerPassword
	}
	if strings.HasPrefix(image, config.Config.NgcRegistryHostName) {
		credential = creds.Base.NgcApikey
	}
	switch params.Body.PlatformID {
	case models.PostNewJobParamsBodyPlatformIDKubernetes:
		config.Config.DockerRegistryUserName = creds.Base.DockerUsername
		if err := queue.BuildJobDockerImage(
			newjob.ID,
			creds.Base.DockerPassword,
			principal.Username,
			creds.Base.K8sConfig,
		); err != nil {
			log.Error("BuildDockerJobImage@postNewJob", err, nil)
			code := http.StatusInternalServerError
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
	case models.PostNewJobParamsBodyPlatformIDRescale:
		if err := queue.BuildSingularityImageJob(
			newjob.ID,
			credential,
			creds.Base.RescaleKey,
			principal.Username,
		); err != nil {
			log.Error("BuildSingularityImageJob@postNewJob", err, nil)
			code := http.StatusInternalServerError
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
	}
	return job.NewPostNewJobCreated().WithPayload(&models.PostNewJobCreatedBody{
		ID: newjob.ID,
	})
}

func modifyJob(params job.ModifyJobParams, principal *auth.Principal) middleware.Responder {
	switch params.Body.Status { // nolint:gocritic
	case models.ModifyJobParamsBodyStatusStopped:
		// TODO implement
	}
	return job.NewModifyJobOK()
}

func deleteJob(params job.DeleteJobParams, principal *auth.Principal) middleware.Responder {
	j, err := db.GetJob(params.ID)
	if err != nil {
		log.Error("GetJob@deleteJob", err, nil)
		code := http.StatusBadRequest
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	err = db.RemoveRescaleJob(j.ID)
	if err != nil {
		log.Error("RemoveRescaleJob@deleteJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	err = os.RemoveAll(filepath.Join(config.Config.SingImgContainerDir, j.ID))
	if err != nil {
		log.Error("RemoveAll@deleteJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	return job.NewDeleteJobNoContent()
}
