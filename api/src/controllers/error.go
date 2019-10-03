package controllers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/app_errors"
)

func errorRoute(api *operations.ScaleShiftAPI) {
	api.AppErrorsGetAppErrorsHandler = app_errors.GetAppErrorsHandlerFunc(getAppErrors)
}

func getAppErrors(params app_errors.GetAppErrorsParams) middleware.Responder {
	owner := auth.Anonymous
	if sess, err := auth.RetrieveSession(params.HTTPRequest); err == nil && sess != nil {
		owner = sess.DockerUsername
	}
	payload := []*models.AppError{}
	if errors, err := db.FindErrors(owner); err == nil {
		for _, e := range errors {
			condition := ""
			if !swag.IsZero(e.ImageAction) {
				condition = swag.StringValue(e.ImageTag)
			}
			if !swag.IsZero(e.JobAction) {
				condition = swag.StringValue(e.JobID)
			}
			payload = append(payload, &models.AppError{
				Caption:   swag.String(e.Caption),
				Condition: condition,
				Detail:    swag.StringValue(e.ErrorMessage),
				Owner:     e.Owner,
				OccursAt:  strfmt.DateTime(e.CreatedAt),
			})
		}
	}
	return app_errors.NewGetAppErrorsOK().WithPayload(payload)
}
