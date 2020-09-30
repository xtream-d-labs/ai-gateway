// Package controllers defines application's routes.
package controllers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/go-openapi/swag"
	"github.com/scaleshift/scaleshift/api/src/generated/models"
	"github.com/scaleshift/scaleshift/api/src/generated/restapi/operations"
	"github.com/scaleshift/scaleshift/api/src/log"
)

// Routes set API handlers
func Routes(api *operations.ScaleShiftAPI) {
	appRoute(api)
	configRoute(api)
	sessionRoute(api)
	repositoryRoute(api)
	imageRoute(api)
	notebookRoute(api)
	workspaceRoute(api)
	jobRoute(api)
	rescaleRoute(api)
	errorRoute(api)
}

func newerror(code int) *models.Error {
	return &models.Error{
		Code:    swag.String(fmt.Sprintf("%d", code)),
		Message: swag.String(http.StatusText(code)),
	}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func dockerClient(cfg *types.AuthConfig) (*docker.Client, *string, int) {
	cli, err := docker.NewEnvClient()
	if err != nil {
		log.Error("NewEnvClient@dockerClient", err, nil)
		return nil, nil, http.StatusInternalServerError
	}
	if cfg != nil {
		converted, err := json.Marshal(cfg)
		if err != nil {
			log.Error("Marshal@dockerClient", err, nil)
			return nil, nil, http.StatusInternalServerError
		}
		return cli, swag.String(base64.URLEncoding.EncodeToString(converted)), 0
	}
	return cli, nil, 0
}
