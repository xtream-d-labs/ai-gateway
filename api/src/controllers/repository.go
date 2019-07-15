package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	ngc "github.com/pottava/ngc-registry-api/app/ngc/registry"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/repository"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/reg/registry"
	"github.com/rescale-labs/scaleshift/api/src/reg/repoutils"
)

func repositoryRoute(api *operations.ScaleShiftAPI) {
	api.RepositoryGetRemoteRepositoriesHandler = repository.GetRemoteRepositoriesHandlerFunc(getRepositories)
	api.RepositoryGetRemoteImagesHandler = repository.GetRemoteImagesHandlerFunc(getRemoteImages)
	api.RepositoryGetNgcRepositoriesHandler = repository.GetNgcRepositoriesHandlerFunc(getNgcRepositories)
	api.RepositoryGetNgcImagesHandler = repository.GetNgcImagesHandlerFunc(getNgcImages)
}

const repositriesNgcCacheKey = "cached-ngc-repositries"

func getRepositories(params repository.GetRemoteRepositoriesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	payload := []*models.Repository{}
	ctx := params.HTTPRequest.Context()

	// @see https://stackoverflow.com/questions/37082826/insufficient-scope-when-attempting-to-get-docker-hub-catalog#answer-37649824
	if config.Config.DockerRegistryEndpoint == repoutils.DefaultDockerRegistry {
		if cfg, e1 := repoutils.GetAuthConfig(
			creds.Base.DockerUsername,
			creds.Base.DockerPassword,
			"https://index.docker.io",
		); e1 == nil {
			if reg, e2 := registry.New(cfg, false); e2 == nil {
				if repos, e3 := reg.Search(fmt.Sprintf("/v1/search?q=%s&n=100", creds.Base.DockerUsername)); e3 == nil {
					for _, repo := range repos {
						payload = append(payload, &models.Repository{
							Namespace:   swag.String(""),
							Name:        swag.String(repo.Name),
							Description: repo.Description,
						})
					}
					return repository.NewGetRemoteRepositoriesOK().WithPayload(payload)
				}
			}
		}
	}
	if cfg, e1 := repoutils.GetAuthConfig(
		creds.Base.DockerUsername,
		creds.Base.DockerPassword,
		config.Config.DockerRegistryEndpoint,
	); e1 == nil {
		if reg, e2 := registry.NewInsecure(cfg, false); e2 == nil {
			if catalogs, e3 := reg.Catalog("/v2/_catalog"); e3 == nil {
				for _, catalog := range catalogs {
					payload = append(payload, &models.Repository{
						Namespace:   swag.String(config.Config.DockerRegistryHostName),
						Name:        swag.String(catalog),
						Description: "",
					})
				}
				return repository.NewGetRemoteRepositoriesOK().WithPayload(payload)
			}
		}
		// Trying VMWare/Harbor API
		if repositories, e2 := lib.HarborRepositories(ctx, cfg); e2 == nil {
			for _, repo := range repositories {
				payload = append(payload, &models.Repository{
					Namespace:   swag.String(config.Config.DockerRegistryHostName),
					Name:        swag.String(repo.Name),
					Description: "",
				})
			}
			return repository.NewGetRemoteRepositoriesOK().WithPayload(payload)
		}
	}
	code := http.StatusNotFound
	return repository.NewGetRemoteRepositoriesDefault(code).WithPayload(newerror(code))
}

func getRemoteImages(params repository.GetRemoteImagesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	registryEndpoint := config.Config.DockerRegistryEndpoint
	if !swag.IsZero(creds.Base.DockerRegistry) {
		registryEndpoint = creds.Base.DockerRegistry
	}
	payload := []*models.Image{}

	if config, e1 := repoutils.GetAuthConfig(
		creds.Base.DockerUsername,
		creds.Base.DockerPassword,
		registryEndpoint,
	); e1 == nil {
		if reg, e2 := registry.NewInsecure(config, false); e2 == nil {
			if tags, e3 := reg.Tags(params.ID); e3 == nil {
				for _, tag := range tags {
					payload = append(payload, &models.Image{
						ID:       swag.String(params.ID),
						RepoTags: []string{tag},
					})
				}
			}
		}
	}
	return repository.NewGetRemoteImagesOK().WithPayload(payload)
}

func getNgcRepositories(params repository.GetNgcRepositoriesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusUnauthorized
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.Repository{}
	if bytes, e := db.GetValueSimple(repositriesNgcCacheKey); e == nil {
		json.Unmarshal(bytes, &result)
		if len(result) > 0 {
			return repository.NewGetRepositoriesOK().WithPayload(result)
		}
	}
	my, err := ngc.Me(creds.NgcSession)
	if err != nil {
		log.Error("Me@getRepositories", err, nil)
		code := http.StatusBadRequest
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	repositries, err := ngc.Repositries(creds.NgcSession, my.PriorNamespace(), true, true, true)
	if err != nil {
		log.Error("Repositries@getRepositories", err, nil)
		code := http.StatusBadRequest
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	for _, repositry := range repositries {
		result = append(result, &models.Repository{
			Namespace:   swag.String(repositry.Namespace),
			Name:        swag.String(repositry.Name),
			Description: repositry.Description,
		})
	}
	db.SetValue(func(txn *badger.Txn) error {
		bytes, err := json.Marshal(result)
		if err != nil {
			return err
		}
		return txn.SetWithTTL([]byte(repositriesNgcCacheKey), bytes, 1*time.Hour)
	})
	return repository.NewGetRepositoriesOK().WithPayload(result)
}

func getNgcImages(params repository.GetNgcImagesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusForbidden
		return repository.NewGetNgcImagesDefault(code).WithPayload(newerror(code))
	}
	images, err := ngc.Images(creds.NgcSession, params.Namespace, params.ID)
	if err != nil {
		log.Error("Images@getNgcImages", err, nil)
		code := http.StatusBadRequest
		return repository.NewGetNgcImagesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.NgcImage{}
	for _, image := range images {
		var updated strfmt.DateTime
		if candidate, err := time.Parse(time.RFC3339Nano, image.UpdatedDate); err == nil {
			updated = strfmt.DateTime(candidate)
		}
		result = append(result, &models.NgcImage{
			Tag:     swag.String(image.Tag),
			Size:    swag.Int64(image.Size),
			Updated: &updated,
		})
	}
	return repository.NewGetNgcImagesOK().WithPayload(result)
}
