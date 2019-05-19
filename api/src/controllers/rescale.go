package controllers

import (
	"net/http"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/rescale"
	"github.com/rescale-labs/scaleshift/api/src/log"
	api "github.com/rescale-labs/scaleshift/api/src/rescale"
)

func rescaleRoute(api *operations.ScaleShiftAPI) {
	api.RescaleGetRescaleCoreTypesHandler = rescale.GetRescaleCoreTypesHandlerFunc(getCoreTypes)
	api.RescaleGetRescaleApplicationHandler = rescale.GetRescaleApplicationHandlerFunc(getApplication)
	api.RescaleGetRescaleApplicationVersionHandler = rescale.GetRescaleApplicationVersionHandlerFunc(getApplicationVersion)
}

// curl -s -H "Authorization: Bearer xxx" "http://localhost:8080/api/v1/rescale/coretypes/"
// curl -s -H "Authorization: Bearer xxx" "http://localhost:8080/api/v1/rescale/coretypes/?app_ver=cpu:cheap"
// curl -s -H "Authorization: Bearer xxx" "http://localhost:8080/api/v1/rescale/coretypes/?app_ver=user_included_singularity_container:3.2.0"
func getCoreTypes(params rescale.GetRescaleCoreTypesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.Base.RescaleKey) {
		code := http.StatusForbidden
		return rescale.NewGetRescaleCoreTypesDefault(code).WithPayload(newerror(code))
	}
	// Restrict core types by some application
	var limitedCoreTypes []string
	if params.AppVer != nil {
		if strings.EqualFold(swag.StringValue(params.AppVer), "gpu:volta") { // nolint:gocritic
			limitedCoreTypes = []string{"dolomite"}

		} else if strings.EqualFold(swag.StringValue(params.AppVer), "cpu:cheap") {
			limitedCoreTypes = []string{"emerald"}

		} else {
			appver := strings.Split(swag.StringValue(params.AppVer), ":")
			if len(appver) != 2 {
				code := http.StatusBadRequest
				return rescale.NewGetRescaleCoreTypesDefault(code).WithPayload(newerror(code))
			}
			app, err := api.Analyses(creds.Base.RescaleKey, appver[0])
			if err != nil {
				log.Error("analyses@getCoreTypes", err, nil)
				code := http.StatusBadRequest
				return rescale.NewGetRescaleCoreTypesDefault(code).WithPayload(newerror(code))
			}
			for _, ver := range app.Versions {
				if ver.MustBeRequested || !strings.EqualFold(ver.VersionCode, appver[1]) {
					continue
				}
				limitedCoreTypes = ver.CoreTypes
			}
		}
	}
	// Retrieve Rescale core types
	cores, err := api.CoreTypes(creds.Base.RescaleKey, nil, nil)
	if err != nil {
		log.Error("CoreTypes@getCoreTypes", err, nil)
		code := http.StatusBadRequest
		return rescale.NewGetRescaleCoreTypesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.RescaleCoreType{}
	for _, core := range cores {
		if core.MustBeRequested {
			continue
		}
		if limitedCoreTypes != nil && !contains(limitedCoreTypes, core.Code) {
			continue
		}
		resources := []*models.RescaleCoreTypeResources{}
		for idx, cpunum := range core.Cores {
			var gpunum int64
			if len(core.GPUCounts) > idx {
				gpunum = int64(core.GPUCounts[idx])
			}
			// If MinGPUs is specified, filter resources
			if params.MinGpus != nil {
				if gpunum < swag.Int64Value(params.MinGpus) {
					continue
				}
			}
			resources = append(resources, &models.RescaleCoreTypeResources{
				Cores: swag.Int64(int64(cpunum)),
				Gpus:  swag.Int64(gpunum),
			})
		}
		if len(resources) == 0 {
			continue
		}
		result = append(result, &models.RescaleCoreType{
			Code:         swag.String(core.Code),
			Name:         swag.String(core.Name),
			Processor:    core.ProcessorInfo,
			BaseClock:    swag.StringValue(core.BaseClockSpeed),
			Interconnect: core.IO,
			Resources:    resources,
		})
	}
	return rescale.NewGetRescaleCoreTypesOK().WithPayload(result)
}

func contains(coretypes []string, coretype string) bool {
	for _, ct := range coretypes {
		if strings.EqualFold(ct, coretype) {
			return true
		}
	}
	return false
}

// curl -s -H "Authorization: Bearer xxx" "http://localhost:8080/api/v1/rescale/applications/singularity/"
func getApplication(params rescale.GetRescaleApplicationParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.Base.RescaleKey) {
		code := http.StatusForbidden
		return rescale.NewGetRescaleApplicationDefault(code).WithPayload(newerror(code))
	}
	app, err := api.Analyses(creds.Base.RescaleKey, api.ApplicationSingularity) // params.Code
	if err != nil {
		log.Error("analyses@getApplication", err, nil)
		code := http.StatusBadRequest
		return rescale.NewGetRescaleApplicationDefault(code).WithPayload(newerror(code))
	}
	versions := []*models.RescaleApplicationVersion{}
	for _, ver := range app.Versions {
		if ver.MustBeRequested {
			continue
		}
		versions = append(versions, &models.RescaleApplicationVersion{
			ID:        swag.String(ver.ID),
			Code:      swag.String(ver.VersionCode),
			Version:   swag.String(ver.Version),
			Coretypes: ver.CoreTypes,
		})
	}
	payload := &models.RescaleApplication{
		Code:     swag.String(app.Code),
		Versions: versions,
	}
	return rescale.NewGetRescaleApplicationOK().WithPayload(payload)
}

// curl -s -H "Authorization: Bearer xxx" "http://localhost:8080/api/v1/rescale/applications/singularity/2.5.1/"
func getApplicationVersion(params rescale.GetRescaleApplicationVersionParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.Base.RescaleKey) {
		code := http.StatusForbidden
		return rescale.NewGetRescaleApplicationVersionDefault(code).WithPayload(newerror(code))
	}
	app, err := api.Analyses(creds.Base.RescaleKey, api.ApplicationSingularity) // params.Code
	if err != nil {
		log.Error("analyses@getApplicationVersion", err, nil)
		code := http.StatusBadRequest
		return rescale.NewGetRescaleApplicationVersionDefault(code).WithPayload(newerror(code))
	}
	for _, ver := range app.Versions {
		if ver.MustBeRequested || !strings.EqualFold(ver.VersionCode, params.Version) {
			continue
		}
		return rescale.NewGetRescaleApplicationVersionOK().WithPayload(&models.RescaleApplicationVersion{
			ID:        swag.String(ver.ID),
			Code:      swag.String(ver.VersionCode),
			Version:   swag.String(ver.Version),
			Coretypes: ver.CoreTypes,
		})
	}
	code := http.StatusNotFound
	return rescale.NewGetRescaleApplicationVersionDefault(code).WithPayload(newerror(code))
}
