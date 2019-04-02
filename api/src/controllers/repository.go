package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/pottava/ngc-registry-api/app/ngc/registry"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/repository"
	"github.com/rescale-labs/scaleshift/api/src/log"
)

func repositoryRoute(api *operations.ScaleShiftAPI) {
	api.RepositoryGetNgcRepositoriesHandler = repository.GetNgcRepositoriesHandlerFunc(getNgcRepositories)
	api.RepositoryGetNgcImagesHandler = repository.GetNgcImagesHandlerFunc(getNgcImages)
}

const repositriesCacheKey = "cached-repositries"

func getNgcRepositories(params repository.GetNgcRepositoriesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusForbidden
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	result := []*models.Repository{}
	if bytes, e := db.GetValueSimple(repositriesCacheKey); e == nil {
		json.Unmarshal(bytes, &result)
		if len(result) > 0 {
			return repository.NewGetRepositoriesOK().WithPayload(result)
		}
	}
	my, err := registry.Me(creds.NgcSession)
	if err != nil {
		log.Error("Me@getRepositories", err, nil)
		code := http.StatusBadRequest
		return repository.NewGetRepositoriesDefault(code).WithPayload(newerror(code))
	}
	repositries, err := registry.Repositries(creds.NgcSession, my.PriorNamespace(), true, true, true)
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
		return txn.SetWithTTL([]byte(repositriesCacheKey), bytes, 1*time.Hour)
	})
	return repository.NewGetRepositoriesOK().WithPayload(result)
}

func getNgcImages(params repository.GetNgcImagesParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.NgcSession) {
		code := http.StatusForbidden
		return repository.NewGetNgcImagesDefault(code).WithPayload(newerror(code))
	}
	images, err := registry.Images(creds.NgcSession, params.Namespace, params.ID)
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
