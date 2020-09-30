package controllers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/scaleshift/scaleshift/api/src/auth"
	"github.com/scaleshift/scaleshift/api/src/config"
	"github.com/scaleshift/scaleshift/api/src/generated/models"
	"github.com/scaleshift/scaleshift/api/src/generated/restapi/operations"
	"github.com/scaleshift/scaleshift/api/src/generated/restapi/operations/app"
	"github.com/scaleshift/scaleshift/api/src/log"
	"github.com/scaleshift/scaleshift/api/src/reg/registry"
	"github.com/scaleshift/scaleshift/api/src/reg/repoutils"
)

func sessionRoute(api *operations.ScaleShiftAPI) {
	api.AppPostNewSessionHandler = app.PostNewSessionHandlerFunc(postNewSession)
}

func postNewSession(params app.PostNewSessionParams) middleware.Responder {
	creds := auth.FindCredentials(swag.StringValue(params.Body.DockerUsername))
	creds.Base.MustSignedIn = config.MustSignInToDockerRegistry()
	creds.Base.DockerPassword = params.Body.DockerPassword.String()
	creds.Base.UsePrivateRegistry = isFilled(creds.Base.DockerUsername, creds.Base.DockerPassword)

	// Check its credentials
	config, err := repoutils.GetAuthConfig(
		creds.Base.DockerUsername,
		creds.Base.DockerPassword,
		config.Config.DockerRegistryEndpoint,
	)
	if err != nil {
		code := http.StatusBadRequest
		return app.NewPostNewSessionDefault(code).WithPayload(newerror(code))
	}
	if _, err = registry.NewInsecure(config, true); err != nil {
		code := http.StatusBadRequest
		log.Debug("NewInsecure@postNewSession", err, nil)
		return app.NewPostNewSessionDefault(code).WithPayload(newerror(code))
	}

	// creds to JWT
	jwt, err := creds.ToSession().ToJWT()
	if err != nil {
		code := http.StatusBadRequest
		log.Debug("NewInsecure@postNewSession", err, nil)
		return app.NewPostNewSessionDefault(code).WithPayload(newerror(code))
	}
	// Store the result
	if err = creds.Save(); err != nil {
		log.Warn("creds.Save@postConfigurations", err, nil)
	}
	return app.NewPostNewSessionCreated().WithPayload(&models.Session{
		Token: swag.String(jwt),
	})
}
