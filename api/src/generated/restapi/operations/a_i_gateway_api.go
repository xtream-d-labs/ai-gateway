// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/security"
	"github.com/go-openapi/spec"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/xtream-d-labs/ai-gateway/api/src/auth"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/app"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/app_errors"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/image"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/job"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/notebook"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/repository"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/rescale"
	"github.com/xtream-d-labs/ai-gateway/api/src/generated/restapi/operations/workspace"
)

// NewAIGatewayAPI creates a new AIGateway instance
func NewAIGatewayAPI(spec *loads.Document) *AIGatewayAPI {
	return &AIGatewayAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		PreServerShutdown:   func() {},
		ServerShutdown:      func() {},
		spec:                spec,
		useSwaggerUI:        false,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,

		JSONConsumer: runtime.JSONConsumer(),

		JSONProducer: runtime.JSONProducer(),

		ImageDeleteImageHandler: image.DeleteImageHandlerFunc(func(params image.DeleteImageParams) middleware.Responder {
			return middleware.NotImplemented("operation image.DeleteImage has not yet been implemented")
		}),
		JobDeleteJobHandler: job.DeleteJobHandlerFunc(func(params job.DeleteJobParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.DeleteJob has not yet been implemented")
		}),
		NotebookDeleteNotebookHandler: notebook.DeleteNotebookHandlerFunc(func(params notebook.DeleteNotebookParams) middleware.Responder {
			return middleware.NotImplemented("operation notebook.DeleteNotebook has not yet been implemented")
		}),
		WorkspaceDeleteWorkspaceHandler: workspace.DeleteWorkspaceHandlerFunc(func(params workspace.DeleteWorkspaceParams) middleware.Responder {
			return middleware.NotImplemented("operation workspace.DeleteWorkspace has not yet been implemented")
		}),
		AppErrorsGetAppErrorsHandler: app_errors.GetAppErrorsHandlerFunc(func(params app_errors.GetAppErrorsParams) middleware.Responder {
			return middleware.NotImplemented("operation app_errors.GetAppErrors has not yet been implemented")
		}),
		AppGetConfigurationsHandler: app.GetConfigurationsHandlerFunc(func(params app.GetConfigurationsParams) middleware.Responder {
			return middleware.NotImplemented("operation app.GetConfigurations has not yet been implemented")
		}),
		AppGetEndpointsHandler: app.GetEndpointsHandlerFunc(func(params app.GetEndpointsParams) middleware.Responder {
			return middleware.NotImplemented("operation app.GetEndpoints has not yet been implemented")
		}),
		NotebookGetIPythonNotebooksHandler: notebook.GetIPythonNotebooksHandlerFunc(func(params notebook.GetIPythonNotebooksParams) middleware.Responder {
			return middleware.NotImplemented("operation notebook.GetIPythonNotebooks has not yet been implemented")
		}),
		ImageGetImagesHandler: image.GetImagesHandlerFunc(func(params image.GetImagesParams) middleware.Responder {
			return middleware.NotImplemented("operation image.GetImages has not yet been implemented")
		}),
		JobGetJobDetailHandler: job.GetJobDetailHandlerFunc(func(params job.GetJobDetailParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.GetJobDetail has not yet been implemented")
		}),
		JobGetJobFilesHandler: job.GetJobFilesHandlerFunc(func(params job.GetJobFilesParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.GetJobFiles has not yet been implemented")
		}),
		JobGetJobLogsHandler: job.GetJobLogsHandlerFunc(func(params job.GetJobLogsParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.GetJobLogs has not yet been implemented")
		}),
		JobGetJobsHandler: job.GetJobsHandlerFunc(func(params job.GetJobsParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.GetJobs has not yet been implemented")
		}),
		RepositoryGetNgcImagesHandler: repository.GetNgcImagesHandlerFunc(func(params repository.GetNgcImagesParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation repository.GetNgcImages has not yet been implemented")
		}),
		RepositoryGetNgcRepositoriesHandler: repository.GetNgcRepositoriesHandlerFunc(func(params repository.GetNgcRepositoriesParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation repository.GetNgcRepositories has not yet been implemented")
		}),
		NotebookGetNotebookDetailsHandler: notebook.GetNotebookDetailsHandlerFunc(func(params notebook.GetNotebookDetailsParams) middleware.Responder {
			return middleware.NotImplemented("operation notebook.GetNotebookDetails has not yet been implemented")
		}),
		NotebookGetNotebooksHandler: notebook.GetNotebooksHandlerFunc(func(params notebook.GetNotebooksParams) middleware.Responder {
			return middleware.NotImplemented("operation notebook.GetNotebooks has not yet been implemented")
		}),
		RepositoryGetRemoteImagesHandler: repository.GetRemoteImagesHandlerFunc(func(params repository.GetRemoteImagesParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation repository.GetRemoteImages has not yet been implemented")
		}),
		RepositoryGetRemoteRepositoriesHandler: repository.GetRemoteRepositoriesHandlerFunc(func(params repository.GetRemoteRepositoriesParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation repository.GetRemoteRepositories has not yet been implemented")
		}),
		RescaleGetRescaleApplicationHandler: rescale.GetRescaleApplicationHandlerFunc(func(params rescale.GetRescaleApplicationParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation rescale.GetRescaleApplication has not yet been implemented")
		}),
		RescaleGetRescaleApplicationVersionHandler: rescale.GetRescaleApplicationVersionHandlerFunc(func(params rescale.GetRescaleApplicationVersionParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation rescale.GetRescaleApplicationVersion has not yet been implemented")
		}),
		RescaleGetRescaleCoreTypesHandler: rescale.GetRescaleCoreTypesHandlerFunc(func(params rescale.GetRescaleCoreTypesParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation rescale.GetRescaleCoreTypes has not yet been implemented")
		}),
		AppGetVersionsHandler: app.GetVersionsHandlerFunc(func(params app.GetVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation app.GetVersions has not yet been implemented")
		}),
		WorkspaceGetWorkspacesHandler: workspace.GetWorkspacesHandlerFunc(func(params workspace.GetWorkspacesParams) middleware.Responder {
			return middleware.NotImplemented("operation workspace.GetWorkspaces has not yet been implemented")
		}),
		JobModifyJobHandler: job.ModifyJobHandlerFunc(func(params job.ModifyJobParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.ModifyJob has not yet been implemented")
		}),
		NotebookModifyNotebookHandler: notebook.ModifyNotebookHandlerFunc(func(params notebook.ModifyNotebookParams) middleware.Responder {
			return middleware.NotImplemented("operation notebook.ModifyNotebook has not yet been implemented")
		}),
		AppPostConfigurationsHandler: app.PostConfigurationsHandlerFunc(func(params app.PostConfigurationsParams) middleware.Responder {
			return middleware.NotImplemented("operation app.PostConfigurations has not yet been implemented")
		}),
		ImagePostNewImageHandler: image.PostNewImageHandlerFunc(func(params image.PostNewImageParams) middleware.Responder {
			return middleware.NotImplemented("operation image.PostNewImage has not yet been implemented")
		}),
		JobPostNewJobHandler: job.PostNewJobHandlerFunc(func(params job.PostNewJobParams, principal *auth.Principal) middleware.Responder {
			return middleware.NotImplemented("operation job.PostNewJob has not yet been implemented")
		}),
		NotebookPostNewNotebookHandler: notebook.PostNewNotebookHandlerFunc(func(params notebook.PostNewNotebookParams) middleware.Responder {
			return middleware.NotImplemented("operation notebook.PostNewNotebook has not yet been implemented")
		}),
		AppPostNewSessionHandler: app.PostNewSessionHandlerFunc(func(params app.PostNewSessionParams) middleware.Responder {
			return middleware.NotImplemented("operation app.PostNewSession has not yet been implemented")
		}),

		// Applies when the "Authorization" header is set
		APIAuthorizerAuth: func(token string) (*auth.Principal, error) {
			return nil, errors.NotImplemented("api key auth (api-authorizer) Authorization from header param [Authorization] has not yet been implemented")
		},
		// default authorizer is authorized meaning no requests are blocked
		APIAuthorizer: security.Authorized(),
	}
}

/*AIGatewayAPI A platform for machine learning & high performance computing
 */
type AIGatewayAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler
	useSwaggerUI    bool

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for the following mime types:
	//   - application/json
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for the following mime types:
	//   - application/json
	JSONProducer runtime.Producer

	// APIAuthorizerAuth registers a function that takes a token and returns a principal
	// it performs authentication based on an api key Authorization provided in the header
	APIAuthorizerAuth func(string) (*auth.Principal, error)

	// APIAuthorizer provides access control (ACL/RBAC/ABAC) by providing access to the request and authenticated principal
	APIAuthorizer runtime.Authorizer

	// ImageDeleteImageHandler sets the operation handler for the delete image operation
	ImageDeleteImageHandler image.DeleteImageHandler
	// JobDeleteJobHandler sets the operation handler for the delete job operation
	JobDeleteJobHandler job.DeleteJobHandler
	// NotebookDeleteNotebookHandler sets the operation handler for the delete notebook operation
	NotebookDeleteNotebookHandler notebook.DeleteNotebookHandler
	// WorkspaceDeleteWorkspaceHandler sets the operation handler for the delete workspace operation
	WorkspaceDeleteWorkspaceHandler workspace.DeleteWorkspaceHandler
	// AppErrorsGetAppErrorsHandler sets the operation handler for the get app errors operation
	AppErrorsGetAppErrorsHandler app_errors.GetAppErrorsHandler
	// AppGetConfigurationsHandler sets the operation handler for the get configurations operation
	AppGetConfigurationsHandler app.GetConfigurationsHandler
	// AppGetEndpointsHandler sets the operation handler for the get endpoints operation
	AppGetEndpointsHandler app.GetEndpointsHandler
	// NotebookGetIPythonNotebooksHandler sets the operation handler for the get i python notebooks operation
	NotebookGetIPythonNotebooksHandler notebook.GetIPythonNotebooksHandler
	// ImageGetImagesHandler sets the operation handler for the get images operation
	ImageGetImagesHandler image.GetImagesHandler
	// JobGetJobDetailHandler sets the operation handler for the get job detail operation
	JobGetJobDetailHandler job.GetJobDetailHandler
	// JobGetJobFilesHandler sets the operation handler for the get job files operation
	JobGetJobFilesHandler job.GetJobFilesHandler
	// JobGetJobLogsHandler sets the operation handler for the get job logs operation
	JobGetJobLogsHandler job.GetJobLogsHandler
	// JobGetJobsHandler sets the operation handler for the get jobs operation
	JobGetJobsHandler job.GetJobsHandler
	// RepositoryGetNgcImagesHandler sets the operation handler for the get ngc images operation
	RepositoryGetNgcImagesHandler repository.GetNgcImagesHandler
	// RepositoryGetNgcRepositoriesHandler sets the operation handler for the get ngc repositories operation
	RepositoryGetNgcRepositoriesHandler repository.GetNgcRepositoriesHandler
	// NotebookGetNotebookDetailsHandler sets the operation handler for the get notebook details operation
	NotebookGetNotebookDetailsHandler notebook.GetNotebookDetailsHandler
	// NotebookGetNotebooksHandler sets the operation handler for the get notebooks operation
	NotebookGetNotebooksHandler notebook.GetNotebooksHandler
	// RepositoryGetRemoteImagesHandler sets the operation handler for the get remote images operation
	RepositoryGetRemoteImagesHandler repository.GetRemoteImagesHandler
	// RepositoryGetRemoteRepositoriesHandler sets the operation handler for the get remote repositories operation
	RepositoryGetRemoteRepositoriesHandler repository.GetRemoteRepositoriesHandler
	// RescaleGetRescaleApplicationHandler sets the operation handler for the get rescale application operation
	RescaleGetRescaleApplicationHandler rescale.GetRescaleApplicationHandler
	// RescaleGetRescaleApplicationVersionHandler sets the operation handler for the get rescale application version operation
	RescaleGetRescaleApplicationVersionHandler rescale.GetRescaleApplicationVersionHandler
	// RescaleGetRescaleCoreTypesHandler sets the operation handler for the get rescale core types operation
	RescaleGetRescaleCoreTypesHandler rescale.GetRescaleCoreTypesHandler
	// AppGetVersionsHandler sets the operation handler for the get versions operation
	AppGetVersionsHandler app.GetVersionsHandler
	// WorkspaceGetWorkspacesHandler sets the operation handler for the get workspaces operation
	WorkspaceGetWorkspacesHandler workspace.GetWorkspacesHandler
	// JobModifyJobHandler sets the operation handler for the modify job operation
	JobModifyJobHandler job.ModifyJobHandler
	// NotebookModifyNotebookHandler sets the operation handler for the modify notebook operation
	NotebookModifyNotebookHandler notebook.ModifyNotebookHandler
	// AppPostConfigurationsHandler sets the operation handler for the post configurations operation
	AppPostConfigurationsHandler app.PostConfigurationsHandler
	// ImagePostNewImageHandler sets the operation handler for the post new image operation
	ImagePostNewImageHandler image.PostNewImageHandler
	// JobPostNewJobHandler sets the operation handler for the post new job operation
	JobPostNewJobHandler job.PostNewJobHandler
	// NotebookPostNewNotebookHandler sets the operation handler for the post new notebook operation
	NotebookPostNewNotebookHandler notebook.PostNewNotebookHandler
	// AppPostNewSessionHandler sets the operation handler for the post new session operation
	AppPostNewSessionHandler app.PostNewSessionHandler
	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// PreServerShutdown is called before the HTTP(S) server is shutdown
	// This allows for custom functions to get executed before the HTTP(S) server stops accepting traffic
	PreServerShutdown func()

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// UseRedoc for documentation at /docs
func (o *AIGatewayAPI) UseRedoc() {
	o.useSwaggerUI = false
}

// UseSwaggerUI for documentation at /docs
func (o *AIGatewayAPI) UseSwaggerUI() {
	o.useSwaggerUI = true
}

// SetDefaultProduces sets the default produces media type
func (o *AIGatewayAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AIGatewayAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AIGatewayAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AIGatewayAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AIGatewayAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AIGatewayAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AIGatewayAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AIGatewayAPI
func (o *AIGatewayAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.APIAuthorizerAuth == nil {
		unregistered = append(unregistered, "AuthorizationAuth")
	}

	if o.ImageDeleteImageHandler == nil {
		unregistered = append(unregistered, "image.DeleteImageHandler")
	}
	if o.JobDeleteJobHandler == nil {
		unregistered = append(unregistered, "job.DeleteJobHandler")
	}
	if o.NotebookDeleteNotebookHandler == nil {
		unregistered = append(unregistered, "notebook.DeleteNotebookHandler")
	}
	if o.WorkspaceDeleteWorkspaceHandler == nil {
		unregistered = append(unregistered, "workspace.DeleteWorkspaceHandler")
	}
	if o.AppErrorsGetAppErrorsHandler == nil {
		unregistered = append(unregistered, "app_errors.GetAppErrorsHandler")
	}
	if o.AppGetConfigurationsHandler == nil {
		unregistered = append(unregistered, "app.GetConfigurationsHandler")
	}
	if o.AppGetEndpointsHandler == nil {
		unregistered = append(unregistered, "app.GetEndpointsHandler")
	}
	if o.NotebookGetIPythonNotebooksHandler == nil {
		unregistered = append(unregistered, "notebook.GetIPythonNotebooksHandler")
	}
	if o.ImageGetImagesHandler == nil {
		unregistered = append(unregistered, "image.GetImagesHandler")
	}
	if o.JobGetJobDetailHandler == nil {
		unregistered = append(unregistered, "job.GetJobDetailHandler")
	}
	if o.JobGetJobFilesHandler == nil {
		unregistered = append(unregistered, "job.GetJobFilesHandler")
	}
	if o.JobGetJobLogsHandler == nil {
		unregistered = append(unregistered, "job.GetJobLogsHandler")
	}
	if o.JobGetJobsHandler == nil {
		unregistered = append(unregistered, "job.GetJobsHandler")
	}
	if o.RepositoryGetNgcImagesHandler == nil {
		unregistered = append(unregistered, "repository.GetNgcImagesHandler")
	}
	if o.RepositoryGetNgcRepositoriesHandler == nil {
		unregistered = append(unregistered, "repository.GetNgcRepositoriesHandler")
	}
	if o.NotebookGetNotebookDetailsHandler == nil {
		unregistered = append(unregistered, "notebook.GetNotebookDetailsHandler")
	}
	if o.NotebookGetNotebooksHandler == nil {
		unregistered = append(unregistered, "notebook.GetNotebooksHandler")
	}
	if o.RepositoryGetRemoteImagesHandler == nil {
		unregistered = append(unregistered, "repository.GetRemoteImagesHandler")
	}
	if o.RepositoryGetRemoteRepositoriesHandler == nil {
		unregistered = append(unregistered, "repository.GetRemoteRepositoriesHandler")
	}
	if o.RescaleGetRescaleApplicationHandler == nil {
		unregistered = append(unregistered, "rescale.GetRescaleApplicationHandler")
	}
	if o.RescaleGetRescaleApplicationVersionHandler == nil {
		unregistered = append(unregistered, "rescale.GetRescaleApplicationVersionHandler")
	}
	if o.RescaleGetRescaleCoreTypesHandler == nil {
		unregistered = append(unregistered, "rescale.GetRescaleCoreTypesHandler")
	}
	if o.AppGetVersionsHandler == nil {
		unregistered = append(unregistered, "app.GetVersionsHandler")
	}
	if o.WorkspaceGetWorkspacesHandler == nil {
		unregistered = append(unregistered, "workspace.GetWorkspacesHandler")
	}
	if o.JobModifyJobHandler == nil {
		unregistered = append(unregistered, "job.ModifyJobHandler")
	}
	if o.NotebookModifyNotebookHandler == nil {
		unregistered = append(unregistered, "notebook.ModifyNotebookHandler")
	}
	if o.AppPostConfigurationsHandler == nil {
		unregistered = append(unregistered, "app.PostConfigurationsHandler")
	}
	if o.ImagePostNewImageHandler == nil {
		unregistered = append(unregistered, "image.PostNewImageHandler")
	}
	if o.JobPostNewJobHandler == nil {
		unregistered = append(unregistered, "job.PostNewJobHandler")
	}
	if o.NotebookPostNewNotebookHandler == nil {
		unregistered = append(unregistered, "notebook.PostNewNotebookHandler")
	}
	if o.AppPostNewSessionHandler == nil {
		unregistered = append(unregistered, "app.PostNewSessionHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AIGatewayAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AIGatewayAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {
	result := make(map[string]runtime.Authenticator)
	for name := range schemes {
		switch name {
		case "api-authorizer":
			scheme := schemes[name]
			result[name] = o.APIKeyAuthenticator(scheme.Name, scheme.In, func(token string) (interface{}, error) {
				return o.APIAuthorizerAuth(token)
			})

		}
	}
	return result
}

// Authorizer returns the registered authorizer
func (o *AIGatewayAPI) Authorizer() runtime.Authorizer {
	return o.APIAuthorizer
}

// ConsumersFor gets the consumers for the specified media types.
// MIME type parameters are ignored here.
func (o *AIGatewayAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {
	result := make(map[string]runtime.Consumer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONConsumer
		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result
}

// ProducersFor gets the producers for the specified media types.
// MIME type parameters are ignored here.
func (o *AIGatewayAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {
	result := make(map[string]runtime.Producer, len(mediaTypes))
	for _, mt := range mediaTypes {
		switch mt {
		case "application/json":
			result["application/json"] = o.JSONProducer
		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result
}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AIGatewayAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the a i gateway API
func (o *AIGatewayAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AIGatewayAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened
	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/images"] = image.NewDeleteImage(o.context, o.ImageDeleteImageHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/jobs/{id}"] = job.NewDeleteJob(o.context, o.JobDeleteJobHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/notebooks/{id}"] = notebook.NewDeleteNotebook(o.context, o.NotebookDeleteNotebookHandler)
	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/workspaces"] = workspace.NewDeleteWorkspace(o.context, o.WorkspaceDeleteWorkspaceHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/errors"] = app_errors.NewGetAppErrors(o.context, o.AppErrorsGetAppErrorsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/configurations"] = app.NewGetConfigurations(o.context, o.AppGetConfigurationsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/endpoints"] = app.NewGetEndpoints(o.context, o.AppGetEndpointsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notebooks/{id}/ipynbs"] = notebook.NewGetIPythonNotebooks(o.context, o.NotebookGetIPythonNotebooksHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/images"] = image.NewGetImages(o.context, o.ImageGetImagesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/jobs/{id}"] = job.NewGetJobDetail(o.context, o.JobGetJobDetailHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/jobs/{id}/files"] = job.NewGetJobFiles(o.context, o.JobGetJobFilesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/jobs/{id}/logs"] = job.NewGetJobLogs(o.context, o.JobGetJobLogsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/jobs"] = job.NewGetJobs(o.context, o.JobGetJobsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/nvidia/repositories/{namespace}/images/{id}"] = repository.NewGetNgcImages(o.context, o.RepositoryGetNgcImagesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/nvidia/repositories"] = repository.NewGetNgcRepositories(o.context, o.RepositoryGetNgcRepositoriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notebooks/{id}"] = notebook.NewGetNotebookDetails(o.context, o.NotebookGetNotebookDetailsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/notebooks"] = notebook.NewGetNotebooks(o.context, o.NotebookGetNotebooksHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/remote-images/{id}"] = repository.NewGetRemoteImages(o.context, o.RepositoryGetRemoteImagesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/repositories"] = repository.NewGetRemoteRepositories(o.context, o.RepositoryGetRemoteRepositoriesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rescale/applications/{code}"] = rescale.NewGetRescaleApplication(o.context, o.RescaleGetRescaleApplicationHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rescale/applications/{code}/{version}"] = rescale.NewGetRescaleApplicationVersion(o.context, o.RescaleGetRescaleApplicationVersionHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rescale/coretypes"] = rescale.NewGetRescaleCoreTypes(o.context, o.RescaleGetRescaleCoreTypesHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/versions"] = app.NewGetVersions(o.context, o.AppGetVersionsHandler)
	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/workspaces"] = workspace.NewGetWorkspaces(o.context, o.WorkspaceGetWorkspacesHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/jobs/{id}"] = job.NewModifyJob(o.context, o.JobModifyJobHandler)
	if o.handlers["PATCH"] == nil {
		o.handlers["PATCH"] = make(map[string]http.Handler)
	}
	o.handlers["PATCH"]["/notebooks/{id}"] = notebook.NewModifyNotebook(o.context, o.NotebookModifyNotebookHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/configurations"] = app.NewPostConfigurations(o.context, o.AppPostConfigurationsHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/images"] = image.NewPostNewImage(o.context, o.ImagePostNewImageHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/jobs"] = job.NewPostNewJob(o.context, o.JobPostNewJobHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/notebooks"] = notebook.NewPostNewNotebook(o.context, o.NotebookPostNewNotebookHandler)
	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/sessions"] = app.NewPostNewSession(o.context, o.AppPostNewSessionHandler)
}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AIGatewayAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	if o.useSwaggerUI {
		return o.context.APIHandlerSwaggerUI(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AIGatewayAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AIGatewayAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AIGatewayAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}

// AddMiddlewareFor adds a http middleware to existing handler
func (o *AIGatewayAPI) AddMiddlewareFor(method, path string, builder middleware.Builder) {
	um := strings.ToUpper(method)
	if path == "/" {
		path = ""
	}
	o.Init()
	if h, ok := o.handlers[um][path]; ok {
		o.handlers[method][path] = builder(h)
	}
}
