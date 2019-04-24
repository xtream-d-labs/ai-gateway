package controllers

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	docker "docker.io/go-docker"
	"docker.io/go-docker/api/types"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	ngc "github.com/pottava/ngc-registry-api/app/ngc/registry"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/app"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/reg/registry"
	"github.com/rescale-labs/scaleshift/api/src/reg/repoutils"
	"github.com/rescale-labs/scaleshift/api/src/rescale"
	"golang.org/x/sync/errgroup"
	"golang.org/x/xerrors"
)

func configRoute(api *operations.ScaleShiftAPI) {
	api.AppGetConfigurationsHandler = app.GetConfigurationsHandlerFunc(getConfigurations)
	api.AppPostConfigurationsHandler = app.PostConfigurationsHandlerFunc(postConfigurations)
}

func getConfigurations(params app.GetConfigurationsParams) middleware.Responder {
	if sess, err := auth.RetriveSession(params.HTTPRequest); err == nil && sess != nil {
		creds := auth.FindCredentials(sess.DockerUsername)
		return app.NewGetConfigurationsOK().WithPayload(&models.Configuration{
			MustSignedIn:       config.MustSignInToDockerRegistry(),
			UsePrivateRegistry: creds.Base.UsePrivateRegistry,
			DockerRegistry:     config.Config.DockerRegistryEndpoint,
			DockerHostname:     config.Config.DockerRegistryHostName,
			DockerUsername:     creds.Base.DockerUsername,
			DockerPassword:     hideChars(creds.Base.DockerPassword, 3),
			UseNgc:             creds.Base.UseNgc,
			NgcEmail:           creds.Base.NgcEmail,
			NgcPassword:        hideChars(creds.Base.NgcPassword, 3),
			NgcApikey:          hideChars(creds.Base.NgcApikey, 5),
			UseRescale:         creds.Base.UseRescale,
			RescalePlatform:    config.Config.RescaleEndpoint,
			RescaleKey:         hideChars(config.Config.RescaleAPIToken, 5),
			UseK8s:             creds.Base.UseK8s,
			K8sConfig:          creds.Base.K8sConfig,
		})
	}
	registry := ""
	if config.Config.DockerRegistryEndpoint != repoutils.DefaultDockerRegistry {
		registry = config.Config.DockerRegistryEndpoint
	}
	hostname := ""
	if config.Config.DockerRegistryHostName != "docker.io" {
		hostname = config.Config.DockerRegistryHostName
	}
	return app.NewGetConfigurationsOK().WithPayload(&models.Configuration{
		MustSignedIn:       config.MustSignInToDockerRegistry(),
		DockerRegistry:     registry,
		DockerHostname:     hostname,
		UsePrivateRegistry: "no",
		UseNgc:             "no",
		UseRescale:         "no",
		RescalePlatform:    config.Config.RescaleEndpoint,
		UseK8s:             "no",
	})
}

func hideChars(value string, num int) string {
	if len(value) > num {
		return value[0:num] + strings.Repeat("*", len(value)-num)
	}
	return value
}

func postConfigurations(params app.PostConfigurationsParams) middleware.Responder {
	creds := auth.FindCredentials(params.Body.DockerUsername)
	if sess, err := auth.RetriveSession(params.HTTPRequest); err == nil && sess != nil {
		creds = auth.FindCredentials(sess.DockerUsername)
	}
	creds.Base.MustSignedIn = config.MustSignInToDockerRegistry()

	// Docker registry
	if params.Body.DockerRegistry != "" {
		creds.Base.DockerRegistry = params.Body.DockerRegistry
	} else {
		creds.Base.DockerRegistry = repoutils.DefaultDockerRegistry
	}
	if params.Body.DockerHostname != "" {
		creds.Base.DockerHostname = params.Body.DockerHostname
	} else {
		creds.Base.DockerHostname = "docker.io"
	}
	creds.Base.DockerUsername = params.Body.DockerUsername
	password := params.Body.DockerPassword.String()
	if password != hideChars(creds.Base.DockerPassword, 3) {
		creds.Base.DockerPassword = password
	}
	creds.Base.UsePrivateRegistry = isFilled(creds.Base.DockerUsername, password)

	// NGC
	creds.Base.NgcEmail = params.Body.NgcEmail
	password = params.Body.NgcPassword.String()
	if password != hideChars(creds.Base.NgcPassword, 3) {
		creds.Base.NgcPassword = password
	}
	if params.Body.NgcApikey != hideChars(creds.Base.NgcApikey, 5) {
		creds.Base.NgcApikey = params.Body.NgcApikey
	}
	creds.Base.UseNgc = isFilled(params.Body.NgcEmail.String(), password, params.Body.NgcApikey)

	// Kubernetes
	if params.Body.K8sConfig != "" {
		creds.Base.K8sConfig = params.Body.K8sConfig
		creds.Base.UseK8s = isFilled(params.Body.K8sConfig)
	}

	// Rescale
	creds.Base.RescalePlatform = params.Body.RescalePlatform
	creds.Base.RescaleKey = ""
	if params.Body.RescaleKey != hideChars(creds.Base.RescaleKey, 5) {
		creds.Base.RescaleKey = params.Body.RescaleKey
	}
	creds.Base.UseRescale = isFilled(creds.Base.RescaleKey)

	eg := errgroup.Group{}

	// Check if you can access to docker registry
	eg.Go(func() error {
		if swag.IsZero(creds.Base.DockerUsername) && swag.IsZero(creds.Base.DockerPassword) {
			return nil
		}
		config, e1 := repoutils.GetAuthConfig(
			creds.Base.DockerUsername,
			creds.Base.DockerPassword,
			creds.Base.DockerRegistry,
		)
		if e1 != nil {
			return xerrors.Errorf("[Docker registry] %s", e1.Error())
		}
		if _, err := registry.NewInsecure(config, false); err != nil {
			return xerrors.Errorf("[Docker registry] %s", err.Error())
		}
		return nil
	})
	// Check if it is able to login to NGC web console
	eg.Go(func() error {
		email := creds.Base.NgcEmail.String()
		if swag.IsZero(email) || swag.IsZero(creds.Base.NgcPassword) {
			return nil
		}
		token, e2 := ngc.Login(email, creds.Base.NgcPassword)
		if e2 != nil {
			return xerrors.Errorf("[NGC Email & password] %s", e2.Error())
		}
		// _, claims, err := ngc.ParseJWT(swag.StringValue(token))
		// if err != nil {
		// 	return err
		// }
		// if err = claims.Valid(); err != nil {
		// 	return err
		// }
		creds.NgcSession = base64.URLEncoding.EncodeToString([]byte(swag.StringValue(token)))
		return nil
	})
	// Check if docker client can login to NGC registry
	eg.Go(func() error {
		if swag.IsZero(creds.Base.NgcApikey) {
			return nil
		}
		cli, e3 := docker.NewEnvClient()
		if e3 != nil {
			return xerrors.Errorf("[NGC API Key] %s", e3.Error())
		}
		defer cli.Close()

		res, e4 := cli.RegistryLogin(context.Background(), types.AuthConfig{
			ServerAddress: config.Config.NgcRegistryHostName,
			Username:      config.Config.NgcRegistryUserName,
			Password:      creds.Base.NgcApikey,
		})
		if e4 != nil {
			return xerrors.Errorf("[NGC API Key] %s", e4.Error())
		}
		if !strings.EqualFold(res.Status, "login succeeded") {
			return xerrors.Errorf("[NGC API Key] Failed to sign in to NGC repos with API key")
		}
		return nil
	})
	// // Check if Kubernetes API returns 200
	eg.Go(func() error {
		return nil
	})
	// Check if Rescale API returns 200
	eg.Go(func() error {
		if swag.IsZero(creds.Base.RescaleKey) {
			return nil
		}
		rescale.SetEndpoint(creds.Base.RescalePlatform)
		rescale.EnableCache(false)
		coretypes, err := rescale.CoreTypes(creds.Base.RescaleKey, nil, nil)
		if err != nil {
			return xerrors.Errorf("[Rescale API Key] %s", err.Error())
		}
		if len(coretypes) == 0 {
			return xerrors.Errorf("[Rescale API Key] Invalid token.")
		}
		rescale.EnableCache(true)
		return nil
	})

	if err := eg.Wait(); err != nil {
		code := http.StatusUnauthorized
		log.Error("Wait@postConfigurations", err, nil)
		return app.NewPostConfigurationsDefault(code).WithPayload(&models.Error{
			Code:    swag.String(fmt.Sprintf("%d", code)),
			Message: swag.String(err.Error()),
		})
	}

	// Store the result
	if err := creds.Save(); err != nil {
		code := http.StatusUnauthorized
		log.Error("SetValue@postConfigurations", err, nil)
		return app.NewPostConfigurationsDefault(code).WithPayload(newerror(code))
	}
	jwt, err := creds.ToSession().ToJWT()
	if err != nil {
		code := http.StatusUnauthorized
		log.Error("ToSession.ToJWT@postConfigurations", err, nil)
		return app.NewPostConfigurationsDefault(code).WithPayload(newerror(code))
	}
	return app.NewPostConfigurationsCreated().WithPayload(&models.Session{
		Token: swag.String(jwt),
	})
}

func isFilled(values ...string) string {
	result := "yes"
	for _, value := range values {
		if swag.IsZero(value) {
			result = "no"
		}
	}
	return result
}
