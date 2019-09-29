package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"docker.io/go-docker/api/types"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/notebook"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/queue"
)

func notebookRoute(api *operations.ScaleShiftAPI) {
	api.NotebookGetNotebooksHandler = notebook.GetNotebooksHandlerFunc(getNotebooks)
	api.NotebookPostNewNotebookHandler = notebook.PostNewNotebookHandlerFunc(postNewNotebook)
	api.NotebookGetNotebookDetailsHandler = notebook.GetNotebookDetailsHandlerFunc(getNotebookDetails)
	api.NotebookModifyNotebookHandler = notebook.ModifyNotebookHandlerFunc(modifyNewNotebook)
	api.NotebookDeleteNotebookHandler = notebook.DeleteNotebookHandlerFunc(deleteNewNotebook)
	api.NotebookGetIpythonNotebooksHandler = notebook.GetIpythonNotebooksHandlerFunc(getIpythonNotebooks)
}

func getNotebooks(params notebook.GetNotebooksParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return notebook.NewGetNotebooksDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{
		Quiet: true,
		All:   true,
	})
	if err != nil {
		log.Error("ContainerList@getNotebooks", err, nil)
		code := http.StatusBadRequest
		return notebook.NewGetNotebooksDefault(code).WithPayload(newerror(code))
	}
	result := []*models.Notebook{}
	for _, container := range containers {
		name := ""
		if len(container.Names) > 0 {
			name = container.Names[0]
		}
		isJupyter := false
		if as, ok1 := container.Labels["com.rescale.scaleshift.image.built-as"]; ok1 {
			if _, ok2 := container.Labels["com.rescale.scaleshift.container.publish"]; ok2 {
				if strings.EqualFold(as, "jupyter-notebook") {
					isJupyter = true
				}
			}
		}
		if !isJupyter {
			continue
		}
		image, port, started := lib.ContainerAttrs(container.Labels)
		result = append(result, &models.Notebook{
			ID:      swag.String(container.ID),
			Name:    swag.String(name),
			Image:   swag.String(image),
			State:   container.State,
			Port:    port,
			Started: strfmt.DateTime(started),
		})
	}
	// building Images
	if images, err := db.FindImages(db.ImageActionBuilding); err == nil {
		for idx, image := range images {
			result = append(result, &models.Notebook{
				ID:      swag.String(fmt.Sprintf("building-image-%d", idx)),
				Name:    swag.String("-"),
				Image:   swag.String(image.Tag),
				State:   "creating",
				Port:    0,
				Started: strfmt.DateTime(image.CreatedAt),
			})
		}
	}
	return notebook.NewGetNotebooksOK().WithPayload(result)
}

func postNewNotebook(params notebook.PostNewNotebookParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return notebook.NewPostNewNotebookDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	ctx := context.Background()
	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		log.Error("ImageList@postNewNotebook", err, nil)
		code := http.StatusBadRequest
		return notebook.NewPostNewNotebookDefault(code).WithPayload(newerror(code))
	}
	name := swag.StringValue(params.Body.Image)
	workspace := params.Body.Workspace
	id := ""
	for _, image := range images {
		if id != "" {
			break
		}
		for _, tag := range image.RepoTags {
			if strings.EqualFold(name, tag) {
				id = image.ID
				break
			}
		}
	}
	wrappedImageID := ""
	for _, image := range images {
		fromID := image.Labels["com.rescale.scaleshift.image.original"]
		fromNm := image.Labels["com.rescale.scaleshift.image.built-on"]
		if strings.EqualFold(name, fromID) || strings.EqualFold(name, fromNm) {
			wrappedImageID = image.ID
			break
		}
	}
	username := auth.Anonymous
	if sess, err := auth.RetrieveSession(params.HTTPRequest); err == nil && sess != nil {
		username = sess.DockerUsername
	}
	image := &db.Image{
		Tag:    name,
		Action: string(db.ImageActionBuilding),
		Owner:  username,
	}
	if err := image.Create(); err != nil {
		log.Error("image.Create@postNewNotebook", err, nil)
		code := http.StatusInternalServerError
		return notebook.NewPostNewNotebookDefault(code).WithPayload(newerror(code))
	}
	builder := "unknown"
	if sess, err := auth.RetrieveSession(params.HTTPRequest); err == nil && sess != nil {
		builder = sess.DockerUsername
	}
	if err := queue.SubmitBuildImageJob(name, id, workspace, wrappedImageID, builder); err != nil {
		log.Error("SubmitBuildImageJob@postNewNotebook", err, nil)
		code := http.StatusInternalServerError
		return notebook.NewPostNewNotebookDefault(code).WithPayload(newerror(code))
	}
	return notebook.NewPostNewNotebookCreated()
}

var rToken *regexp.Regexp

func init() {
	rToken = regexp.MustCompile(`\?token=([0-9a-zA-Z]+)`)
}

func getNotebookDetails(params notebook.GetNotebookDetailsParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return notebook.NewGetNotebookDetailsDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	ctx := context.Background()
	container, err := cli.ContainerInspect(ctx, params.ID)
	if err != nil {
		log.Error("ContainerInspect@getNotebookDetails", err, nil)
		code := http.StatusBadRequest
		return notebook.NewGetNotebookDetailsDefault(code).WithPayload(newerror(code))
	}
	_, port, started := lib.ContainerAttrs(container.Config.Labels)
	mounts := []string{}
	for _, mount := range container.Mounts {
		mounts = append(mounts, fmt.Sprintf("%s:%s",
			strings.TrimPrefix(strings.Replace(mount.Source, config.Config.WorkspaceHostDir, "", -1), "/"),
			mount.Destination,
		))
	}
	out, err := lib.LogsOfDockerContainer(ctx, cli, params.ID)
	if err != nil {
		log.Error("LogsOfDockerContainer@getNotebookDetails", err, nil)
		code := http.StatusInternalServerError
		return notebook.NewGetNotebookDetailsDefault(code).WithPayload(newerror(code))
	}
	var token string
	if candidates := rToken.FindStringSubmatch(swag.StringValue(out)); len(candidates) > 1 {
		token = candidates[1]
	}
	var ended time.Time
	if candidate, err := time.Parse(time.RFC3339, container.State.FinishedAt); err == nil {
		ended = candidate
	}
	result := &models.NotebookDetail{
		ID:      swag.String(container.ID),
		Name:    container.Name,
		State:   container.State.Status,
		Port:    port,
		Started: strfmt.DateTime(started),
		Ended:   strfmt.DateTime(ended),
		Token:   swag.String(token),
		Mounts:  mounts,
	}
	return notebook.NewGetNotebookDetailsOK().WithPayload(result)
}

func modifyNewNotebook(params notebook.ModifyNotebookParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return notebook.NewModifyNotebookDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	ctx := context.Background()
	switch params.Body.Status {
	case models.ModifyNotebookParamsBodyStatusStarted:
		// TODO implement client side
		if err := cli.ContainerStart(ctx, params.ID, types.ContainerStartOptions{}); err != nil {
			log.Error("ContainerStart@modifyNewNotebook", err, nil)
			code := http.StatusBadRequest
			return notebook.NewModifyNotebookDefault(code).WithPayload(newerror(code))
		}
	case models.ModifyNotebookParamsBodyStatusStopped:
		timeout := 15 * time.Second
		if err := cli.ContainerStop(ctx, params.ID, &timeout); err != nil {
			log.Error("ContainerStop@modifyNewNotebook", err, nil)
			code := http.StatusBadRequest
			return notebook.NewModifyNotebookDefault(code).WithPayload(newerror(code))
		}
	}
	return notebook.NewModifyNotebookOK()
}

func deleteNewNotebook(params notebook.DeleteNotebookParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return notebook.NewDeleteNotebookDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	ctx := context.Background()
	if err := cli.ContainerRemove(ctx, params.ID, types.ContainerRemoveOptions{
		RemoveVolumes: false,
		Force:         true,
	}); err != nil {
		log.Error("ContainerRemove@deleteNewNotebook", err, nil)
		code := http.StatusBadRequest
		return notebook.NewDeleteNotebookDefault(code).WithPayload(newerror(code))
	}
	return notebook.NewDeleteNotebookNoContent()
}

var (
	nonascii = regexp.MustCompile(`[^a-zA-Z0-9._\-/]`)
	ipynbs   = regexp.MustCompile(`.*\.ipynb`)
)

func getIpythonNotebooks(params notebook.GetIpythonNotebooksParams) middleware.Responder {
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return notebook.NewGetIpythonNotebooksDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	ctx := context.Background()
	container, err := cli.ContainerInspect(ctx, params.ID)
	if err != nil {
		log.Error("ContainerInspect@getIpythonNotebooks", err, nil)
		code := http.StatusBadRequest
		return notebook.NewGetIpythonNotebooksDefault(code).WithPayload(newerror(code))
	}
	src := ""
	for _, mnt := range container.Mounts {
		src = strings.Replace(mnt.Source, config.Config.WorkspaceHostDir,
			config.Config.WorkspaceContainerDir, -1)
	}
	result := []*models.IpythonNotebook{}
	for _, file := range fildFiles(src) {
		name := nonascii.ReplaceAllString(strings.Replace(file, src, "", -1), "")
		result = append(result, &models.IpythonNotebook{
			Name: swag.String(strings.TrimLeft(name, "/")),
		})
	}
	return notebook.NewGetIpythonNotebooksOK().WithPayload(result)
}

func fildFiles(dir string) []string {
	files := []string{}
	if err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() && strings.Contains(path, ".ipynb_checkpoints") {
			return filepath.SkipDir
		}
		if info.IsDir() && strings.EqualFold(info.Name(), "bin") {
			return filepath.SkipDir
		}
		if info.IsDir() && strings.EqualFold(info.Name(), "lib") {
			return filepath.SkipDir
		}
		if info.IsDir() {
			return nil
		}
		if ipynbs.MatchString(info.Name()) {
			files = append(files, path)
		}
		return nil
	}); err != nil {
		return []string{}
	}
	return files
}
