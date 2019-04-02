package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/app"
)

func appRoute(api *operations.ScaleShiftAPI) {
	api.AppGetVersionsHandler = app.GetVersionsHandlerFunc(getVersions)
	api.AppGetEndpointsHandler = app.GetEndpointsHandlerFunc(getEndpoints)
}

// AppVersion defines application version
type AppVersion struct {
	Version string `json:"version"`
	Commit  string `json:"commit"`
	Date    string `json:"date"`
}

func getVersions(params app.GetVersionsParams) middleware.Responder {
	version := swag.String(config.BuildVersion())
	date := config.BuildDate()
	if resp, err := http.Get("https://s3-ap-northeast-1.amazonaws.com/scaleshift/latest-version.json"); err == nil {
		defer resp.Body.Close()
		if bytes, err := ioutil.ReadAll(resp.Body); err == nil {
			latest := &AppVersion{}
			if err = json.Unmarshal(bytes, latest); err == nil {
				version = swag.String(fmt.Sprintf("%s-%s", latest.Version, latest.Commit))
				date = latest.Date
			}
		}
	}
	return app.NewGetVersionsOK().WithPayload(&models.Versions{
		Current: &models.Version{
			Version:   swag.String(config.BuildVersion()),
			BuildDate: config.BuildDate(),
		},
		Latest: &models.Version{
			Version:   version,
			BuildDate: date,
		},
	})
}

func getEndpoints(params app.GetEndpointsParams) middleware.Responder {
	return app.NewGetEndpointsOK().WithPayload(&models.Endpoints{
		DockerRegistry: config.Config.DockerRegistryEndpoint,
		NgcRegistry:    config.Config.NgcRegistryEndpoint,
		KubernetesAPI:  config.Config.KubernetesAPIEndpoint,
		RescaleAPI:     config.Config.RescaleEndpoint,
	})
}
