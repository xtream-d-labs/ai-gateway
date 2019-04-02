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

// curl -s -H "Authorization: Token xxx" http://localhost:9000/api/v1/coretypes/
/**
 * CPU: http://localhost:9000/api/v1/coretypes/?app_ver=singularity%3A2.5.1
 *  - Emerald (emerald):      AWS c5 instances (Intel Xeon Platinum P-8124, Skylake)
 *  - Onyx (hpc-3):           AWS c4 instances (Intel Xeon E5-2666 v3, Haswell)
 *  - Nickel (hpc-plus):      AWS c3 instances (Intel Xeon E5-2680 v2, Ivy Bridge)
 *  - Melanite (melanite):    AWS m5 instances (Intel Xeon Platinum 8175, Skylake)
 *  - Titanium (titanium):    AWS m4 instances (Intel Xeon E5-2676 v3, Haswell)
 *  - Marble (standard-plus): AWS m3 instances (Intel Xeon E5-2670 v2, Ivy Bridge)
 *  - Topaz (topaz):          AWS x1 instances (Intel Xeon E7-8880 v3, Haswell)
 *  - Zinc (zinc):            AWS r4 instances (Intel Xeon E5-2686 v4, Broadwell)
 *  - Gold (hi-mem-hpc):      AWS r3 instances (Intel Xeon E5-2670 v2, Ivy Bridge)
 *  - Graphite (graphite):    AWS i3 instances (Intel Xeon E5-2686 v4, Broadwell)
 *  - Quartz (hi-io-plus):    AWS i2 instances (Intel Xeon E5-2670 v2, Ivy Bridge)
 *
 * GPU: http://localhost:9000/api/v1/coretypes/?app_ver=singularity%3Auser-included-cuda-8.0&min_gpus=1
 *  - Dolomite (dolomite):    AWS p3 instances (NVIDIA V100 w/ NVLink, Intel Xeon E5-2686 v4, Broadwell)
 *  - Obsidian (obsidian):    AWS p2 instances (NVIDIA K80, Intel Xeon E5-2676 v3, Haswell)
 */
func getCoreTypes(params rescale.GetRescaleCoreTypesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusForbidden
		return rescale.NewGetRescaleCoreTypesDefault(code).WithPayload(newerror(code))
	}
	// Restrict some application
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
			app, err := analyses(creds.Base.RescaleKey, appver[0])
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

// curl -s -H "Authorization: Token xxx" http://localhost:9000/api/v1/applications/singularity/
func getApplication(params rescale.GetRescaleApplicationParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusForbidden
		return rescale.NewGetRescaleApplicationDefault(code).WithPayload(newerror(code))
	}
	app, err := analyses(creds.Base.RescaleKey, params.Code)
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

// curl -s -H "Authorization: Token xxx" http://localhost:9000/api/v1/applications/singularity/2.5.1/
func getApplicationVersion(params rescale.GetRescaleApplicationVersionParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusForbidden
		return rescale.NewGetRescaleApplicationVersionDefault(code).WithPayload(newerror(code))
	}
	app, err := analyses(creds.Base.RescaleKey, params.Code)
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

func analyses(token, param string) (*api.Application, error) {
	code := api.ApplicationSingularity
	if strings.EqualFold(param, "singularity_mpi") {
		code = api.ApplicationSingularityMPI
	}
	return api.Analyses(token, code)
}
