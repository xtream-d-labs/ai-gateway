package controllers

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"docker.io/go-docker/api/types"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/xtream-d-labs/ai-gateway/api/src/config"
	"github.com/xtream-d-labs/ai-gateway/api/src/db"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/models"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/workspace"
	"github.com/xtream-d-labs/ai-gateway/api/src/log"
)

func workspaceRoute(api *operations.AIGatewayAPI) {
	api.WorkspaceGetWorkspacesHandler = workspace.GetWorkspacesHandlerFunc(getWorkspaces)
	api.WorkspaceDeleteWorkspaceHandler = workspace.DeleteWorkspaceHandlerFunc(deleteWorkspace)
}

func getWorkspaces(params workspace.GetWorkspacesParams) middleware.Responder {
	files, err := ioutil.ReadDir(config.Config.WorkspaceContainerDir)
	if err != nil {
		log.Error("ReadDir@getWorkspaces", err, nil)
		code := http.StatusInternalServerError
		return workspace.NewGetWorkspacesDefault(code).WithPayload(newerror(code))
	}
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return workspace.NewGetWorkspacesDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Quiet: true,
		All:   true,
	})
	if err != nil {
		log.Error("ContainerList@getWorkspaces", err, nil)
		code := http.StatusBadRequest
		return workspace.NewGetWorkspacesDefault(code).WithPayload(newerror(code))
	}
	jobs, err := db.GetJobs()
	if err != nil {
		log.Error("GetJobs@getWorkspaces", err, nil)
		code := http.StatusBadRequest
		return workspace.NewGetWorkspacesDefault(code).WithPayload(newerror(code))
	}
	payload := []*models.Workspace{}
	for _, f := range files {
		if f.IsDir() {
			payload = append(payload, &models.Workspace{
				Notebooks:    findNotebooks(f.Name(), containers),
				Jobs:         findJobs(f.Name(), jobs),
				Path:         swag.String(f.Name()),
				AbsolutePath: filepath.Join(config.Config.WorkspaceHostDir, f.Name()),
			})
		}
	}
	return workspace.NewGetWorkspacesOK().WithPayload(payload)
}

func findNotebooks(path string, containers []types.Container) []string {
	notebooks := []string{}
	for _, container := range containers {
		isJupyter := false
		if as, ok1 := container.Labels["com.aigateway.image.built-as"]; ok1 {
			if _, ok2 := container.Labels["com.aigateway.container.publish"]; ok2 {
				if strings.EqualFold(as, "jupyter-notebook") {
					isJupyter = true
				}
			}
		}
		if !isJupyter {
			continue
		}
		for _, mount := range container.Mounts {
			src := strings.TrimPrefix(strings.Replace(mount.Source, config.Config.WorkspaceHostDir, "", -1), "/")
			if path == src {
				name := ""
				if len(container.Names) > 0 {
					name = strings.TrimPrefix(container.Names[0], "/")
				}
				notebooks = append(notebooks, name)
				break
			}
		}
	}
	return notebooks
}

func findJobs(path string, jobs []*db.Job) []string {
	result := []string{}
	for _, job := range jobs {
		for _, workspace := range strings.Split(job.Workspaces, ",") {
			if strings.EqualFold(path, workspace) {
				result = append(result, job.JobID)
				break
			}
		}
	}
	return result
}

func deleteWorkspace(params workspace.DeleteWorkspaceParams) middleware.Responder {
	path := filepath.Join(config.Config.WorkspaceContainerDir, swag.StringValue(params.Body.Path))
	err := os.RemoveAll(path)
	if err != nil {
		log.Error("RemoveAll@deleteWorkspace", err, nil)
		code := http.StatusInternalServerError
		return workspace.NewDeleteWorkspaceDefault(code).WithPayload(newerror(code))
	}
	return workspace.NewDeleteWorkspaceNoContent()
}
