package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"sort"
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
	"github.com/rescale-labs/scaleshift/api/src/kubernetes"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/queue"
	"github.com/rescale-labs/scaleshift/api/src/rescale"
	"golang.org/x/sync/errgroup"
	coreV1 "k8s.io/api/core/v1"
)

func jobRoute(api *operations.ScaleShiftAPI) {
	api.JobGetJobsHandler = job.GetJobsHandlerFunc(getJobs)
	api.JobGetJobDetailHandler = job.GetJobDetailHandlerFunc(getJobDetail)
	api.JobGetJobLogsHandler = job.GetJobLogsHandlerFunc(getJobLogs)
	api.JobGetJobFilesHandler = job.GetJobFilesHandlerFunc(getJobFiles)
	api.JobPostNewJobHandler = job.PostNewJobHandlerFunc(postNewJob)
	api.JobModifyJobHandler = job.ModifyJobHandlerFunc(modifyJob)
	api.JobDeleteJobHandler = job.DeleteJobHandlerFunc(deleteJob)
}

func getJobs(params job.GetJobsParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()
	return job.NewGetJobsOK().WithPayload(jobs(ctx, creds, nil))
}

const timeFormat = "2006-01-02 15:04:05.999999999 -0700 MST"

func jobs(ctx context.Context, creds *auth.Credentials, ID *string) []*models.Job {
	result := []*models.Job{}
	if jobs, err := db.GetJobs(); err == nil {
		for _, j := range jobs {
			if ID != nil {
				if !strings.EqualFold(swag.StringValue(ID), j.JobID) {
					continue
				}
			}
			status := swag.String(j.Status)
			var externalLink string
			var ended time.Time

			switch db.JobStatus(j.Status) {
			case db.JobStatusK8sStarted:
				k8sPod, e := kubernetes.PodStatus(creds.Base.K8sConfig, j.JobID, "default")
				if k8sPod == nil || e != nil {
					break
				}
				switch k8sPod.Status.Phase {
				case coreV1.PodPending:
					status = db.JobStatusK8sPending.Value()
				case coreV1.PodRunning:
					status = db.JobStatusK8sRunning.Value()
				case coreV1.PodSucceeded:
					status = db.JobStatusK8sSucceeded.Value()
					if k8sJob, e1 := kubernetes.JobStatus(creds.Base.K8sConfig, j.JobID, "default"); k8sJob != nil && e1 == nil {
						if candidate, e2 := time.Parse(timeFormat, k8sJob.CompletionTime.String()); e2 == nil {
							ended = candidate
						}
					}
				case coreV1.PodFailed:
					status = db.JobStatusK8sFailed.Value()
					if k8sJob, e1 := kubernetes.JobStatus(creds.Base.K8sConfig, j.JobID, "default"); k8sJob != nil && e1 == nil {
						if candidate, e2 := time.Parse(timeFormat, k8sJob.CompletionTime.String()); e2 == nil {
							ended = candidate
						}
					}
				case coreV1.PodUnknown:
					status = db.JobStatusUnknown.Value()
				}
			case db.JobStatusRescaleStarted:
				rStatus, e := rescale.Status(ctx, creds.Base.RescaleKey, j.TargetID)
				if rStatus == nil || e != nil {
					break
				}
				externalLink = fmt.Sprintf(
					"%s/jobs/%s/status/",
					config.Config.RescaleEndpoint,
					j.TargetID)
				switch rStatus.Status {
				case rescale.JobStatusPending, rescale.JobStatusQueued,
					rescale.JobStatusWait4Cls, rescale.JobStatusWaitQueue:
					status = db.JobStatusRescaleStarted.Value()
				case rescale.JobStatusStarted, rescale.JobStatusValidated,
					rescale.JobStatusExecuting:
					status = db.JobStatusRescaleRunning.Value()
				case rescale.JobStatusCompleted:
					status = db.JobStatusRescaleSucceed.Value()
					if rStatus.StatusDate != nil {
						ended = *rStatus.StatusDate
					}
				case rescale.JobStatusStopping, rescale.JobStatusForceStop:
					status = db.JobStatusRescaleFailed.Value()
					if rStatus.StatusDate != nil {
						ended = *rStatus.StatusDate
					}
				}
			}
			tmp := &models.Job{
				ID:           swag.String(j.JobID),
				Platform:     db.PlatformType(j.Platform).String(),
				Status:       swag.StringValue(status),
				Image:        j.DockerImage,
				Mounts:       strings.Split(j.Workspaces, ","),
				Commands:     strings.Split(j.Commands, ","),
				ExternalLink: externalLink,
				Started:      strfmt.DateTime(j.CreatedAt),
				Ended:        strfmt.DateTime(ended),
			}
			log.Debug(fmt.Sprintf("%+v", tmp), nil, nil)
			result = append(result, tmp)
		}
	}
	return result
}

func getJobDetail(params job.GetJobDetailParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	payload := &models.JobDetail{}
	eg := errgroup.Group{}
	eg.Go(func() error {
		if jobs := jobs(ctx, creds, swag.String(params.ID)); len(jobs) > 0 {
			payload.Job = *jobs[0]
		}
		return nil
	})
	eg.Go(func() error {
		payload.JobLogs = models.JobLogs{
			Logs: jobLogs(ctx, creds, params.ID),
		}
		return nil
	})
	eg.Go(func() error {
		payload.JobFiles = models.JobFiles{
			APIToken: creds.Base.RescaleKey,
			Files:    jobFiles(ctx, creds, params.ID),
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		code := http.StatusUnauthorized
		log.Error("Wait@getJobDetail", err, nil)
		return job.NewGetJobDetailDefault(code).WithPayload(&models.Error{
			Code:    swag.String(fmt.Sprintf("%d", code)),
			Message: swag.String(err.Error()),
		})
	}
	return job.NewGetJobDetailOK().WithPayload(payload)
}

func getJobLogs(params job.GetJobLogsParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	return job.NewGetJobLogsOK().WithPayload(&models.JobLogs{
		Logs: jobLogs(ctx, creds, params.ID),
	})
}

func jobLogs(ctx context.Context, creds *auth.Credentials, ID string) []*models.JobLog {
	result := []*models.JobLog{}
	jobs, err := db.GetJobs()
	if err != nil {
		return result
	}
	for _, j := range jobs {
		switch db.JobStatus(j.Status) {
		case db.JobStatusK8sStarted:
			k8sLog, e := kubernetes.Logs(creds.Base.K8sConfig, j.JobID, "default")
			if k8sLog == nil || e != nil {
				break
			}
			for _, logs := range k8sLog {
				result = append(result, &models.JobLog{
					Time: strfmt.DateTime(logs.Time),
					Log:  swag.String(logs.Log),
				})
			}
		case db.JobStatusRescaleStarted:
			rescaleLog, e := rescale.Logs(ctx, creds.Base.RescaleKey, j.TargetID)
			if rescaleLog == nil || e != nil {
				break
			}
			for _, logs := range rescaleLog {
				result = append(result, &models.JobLog{
					Time: strfmt.DateTime(logs.Time),
					Log:  swag.String(logs.Log),
				})
			}
		}
	}
	sort.Slice(result, func(i, j int) bool {
		return time.Time(result[i].Time).Before(time.Time(result[j].Time))
	})
	return result
}

func getJobFiles(params job.GetJobFilesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	return job.NewGetJobFilesOK().WithPayload(&models.JobFiles{
		APIToken: creds.Base.RescaleKey,
		Files:    jobFiles(ctx, creds, params.ID),
	})
}

func jobFiles(ctx context.Context, creds *auth.Credentials, ID string) []*models.JobFile {
	result := []*models.JobFile{}
	jobs, err := db.GetJobs()
	if err != nil {
		return result
	}
	for _, j := range jobs {
		switch db.JobStatus(j.Status) {
		case db.JobStatusK8sStarted:
			// There is no such things

		case db.JobStatusRescaleStarted:
			// Uncommnet if the Rescale API allows CORS
			// rFiles, e := rescale.OutputFiles(ctx, creds.Base.RescaleKey, j.TargetID)
			// if rFiles == nil || e != nil {
			// 	break
			// }
			// for _, file := range rFiles.Results {
			// 	result = append(result, &models.JobFile{
			// 		Name:        swag.String(file.Name),
			// 		Size:        swag.Int64(file.Size),
			// 		DownloadURL: file.DownloadURL,
			// 	})
			// }
		}
	}
	return result
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
		Platform:    int(platform),
		JobID:       uuid.New().String(),
		Action:      string(db.JobActionBuilding),
		Status:      string(db.JobStatusImageBuilding),
		DockerImage: image,
		PythonFile:  ipynb,
		Workspaces:  strings.Join(workspaces, ","),
		Commands:    strings.Join(commands, ","),
		CPU:         params.Body.CPU,
		Memory:      params.Body.Mem,
		GPU:         params.Body.Gpu,
		CoreType:    params.Body.Coretype,
		Cores:       params.Body.Cores,
	}
	if err := newjob.Create(); err != nil {
		log.Error("job.Create@postNewJob", err, nil)
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
			newjob.JobID,
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
			newjob.JobID,
			credential,
			creds.Base.RescaleKey,
			principal.Username,
		); err != nil {
			log.Error("BuildSingularityImageJob@postNewJob", err, nil)
			code := http.StatusInternalServerError
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
	}
	return job.NewPostNewJobCreated().WithPayload(&job.PostNewJobCreatedBody{
		ID: newjob.JobID,
	})
}

func modifyJob(params job.ModifyJobParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	j, err := db.GetJob(params.ID)
	if err != nil {
		log.Error("GetJob@modifyJob", err, nil)
		code := http.StatusBadRequest
		return job.NewModifyJobDefault(code).WithPayload(newerror(code))
	}
	switch params.Body.Status { // nolint:gocritic
	case models.ModifyJobParamsBodyStatusStopped:
		switch db.JobStatus(j.Status) {
		case db.JobStatusK8sStarted:
			// There is no proper method
		case db.JobStatusRescaleStarted:
			if swag.IsZero(creds.Base.RescaleKey) {
				code := http.StatusForbidden
				return job.NewModifyJobDefault(code).WithPayload(newerror(code))
			}
			if e := rescale.Stop(ctx, creds.Base.RescaleKey, j.TargetID); e != nil {
				log.Error("StopRescaleJob@modifyJob", e, nil)
			}
			if e := db.UpdateJob(j.JobID, db.JobActionRescale, db.JobActionRescale, db.JobStatusRescaleFailed, nil); e != nil {
				log.Error("UpdateJob@modifyJob", e, nil)
			}
		}
	}
	return job.NewModifyJobOK()
}

func deleteJob(params job.DeleteJobParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	j, err := db.GetJob(params.ID)
	if err != nil {
		log.Error("GetJob@deleteJob", err, nil)
		code := http.StatusBadRequest
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	switch db.JobStatus(j.Status) {
	case db.JobStatusK8sStarted, db.JobStatusK8sPending,
		db.JobStatusK8sRunning, db.JobStatusK8sSucceeded,
		db.JobStatusK8sFailed:
		if swag.IsZero(creds.Base.K8sConfig) {
			code := http.StatusForbidden
			return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
		}
		if e := kubernetes.DeleteJob(creds.Base.K8sConfig, j.JobID, "default"); e != nil {
			log.Error("DeleteKubernetesJob@deleteJob", e, nil)
		}
	case db.JobStatusRescaleStarted, db.JobStatusRescaleRunning,
		db.JobStatusRescaleSucceed, db.JobStatusRescaleFailed:
		if swag.IsZero(creds.Base.RescaleKey) {
			code := http.StatusForbidden
			return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
		}
		if e := rescale.Delete(ctx, creds.Base.RescaleKey, j.TargetID); e != nil {
			log.Error("DeleteRescaleJob@deleteJob", e, nil)
		}
	}
	err = db.RemoveJob(j.JobID)
	if err != nil {
		log.Error("RemoveJob@deleteJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	err = os.RemoveAll(filepath.Join(config.Config.SingImgContainerDir, j.JobID))
	if err != nil {
		log.Error("RemoveAll@deleteJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	return job.NewDeleteJobNoContent()
}
