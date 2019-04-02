package config

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/kelseyhightower/envconfig"
)

// for compile flags
var (
	version = "dev"
	commit  string
	date    = "---"
)

// API props
const (
	ProjectName = "ScaleShift"
	ProjectPath = "github.com/rescale-labs/scaleshift/api"
)

// Config can be set via environment variables
type config struct { // nolint:maligned
	APIVersion             string   `envconfig:"API_VERSION" default:"dev"`
	APIEndpoint            string   `envconfig:"API_ENDPOINT" default:"http://localhost:9000"`
	MustSignIn             bool     `envconfig:"MUST_SIGN_IN" default:"false"`
	JwtIssuer              string   `envconfig:"JWT_ISSUER" default:"owner"`
	JwtAudience            string   `envconfig:"JWT_AUDIENCE" default:"scaleshift.io"`
	JwtExpiration          int      `envconfig:"JWT_EXPIRATION" default:"86400"`
	JwsPrivateKey          string   `envconfig:"JWS_PRIVATE_KEY" default:"/certs/private.pem"`
	JwsPublicKey           string   `envconfig:"JWS_PUBLIC_KEY" default:"/certs/public.pem"`
	DockerRegistryEndpoint string   `envconfig:"DOCKER_REGISTRY_ENDPOINT" default:"https://registry-1.docker.io"`
	DockerRegistryHostName string   `envconfig:"DOCKER_REGISTRY_HOST_NAME" default:"docker.io"`
	DockerRegistryUserName string   `envconfig:"DOCKER_REGISTRY_USER_NAME"`
	NgcRegistryEndpoint    string   `envconfig:"NGC_REGISTRY_ENDPOINT" default:"https://registry.nvidia.com"`
	NgcRegistryHostName    string   `envconfig:"NGC_REGISTRY_HOST_NAME" default:"nvcr.io"`
	NgcRegistryUserName    string   `envconfig:"NGC_REGISTRY_USER_NAME" default:"$oauthtoken"`
	JupyterImageNamespace  string   `envconfig:"JUPYTER_IMAGE_NAMESPACE" default:"ss-jupyter"`
	JupyterMinimumPort     uint16   `envconfig:"JUPYTER_MINIMUM_PORT" default:"30000"`
	KubernetesAPIEndpoint  string   `envconfig:"KUBERNETES_API_ENDPOINT"`
	KubernetesConfig       string   `envconfig:"KUBERNETES_CONFIG"`
	RescaleEndpoint        string   `envconfig:"RESCALE_ENDPOINT" default:"https://platform.rescale.com"`
	RescaleAPIToken        string   `envconfig:"RESCALE_API_TOKEN"`
	AccessLog              bool     `envconfig:"ACCESS_LOG" default:"true"`
	LogLevel               string   `envconfig:"LOG_LEVEL" default:"warn"`
	LogFormat              string   `envconfig:"LOG_FORMAT" default:"default"`
	AllowCORS              bool     `envconfig:"ALLOW_CORS" default:"true"`
	SecuredTransport       bool     `envconfig:"SECURED_TRANSPORT" default:"false"`
	ContentEncoding        bool     `envconfig:"CONTENT_ENCODING" default:"false"`
	ImagesToBeIgnored      []string `envconfig:"IMAGES_TOBE_IGNORED" default:"scaleshift/,ss-jupyter/,amazon-ecs-,<none>"`
	DatabaseDir            string   `envconfig:"DATABASE_CNTR_DIR" default:"/tmp/badger"`
	WorkspaceHostDir       string   `envconfig:"WORKSPACE_HOST_DIR"`
	WorkspaceContainerDir  string   `envconfig:"WORKSPACE_CNTR_DIR" default:"/tmp/work"`
	SingImg                string   `envconfig:"SINGULARITY_IMAGE" default:"scaleshift/singularity:2.6-d2s"`
	SingImgHostPath        string   `envconfig:"SINGULARITY_HOST_DIR"`
	SingImgContainerDir    string   `envconfig:"SINGULARITY_CNTR_DIR" default:"/tmp/simg"`
}

// Config represents its configurations
var Config *config

func init() {
	Config = &config{}
	envconfig.MustProcess("ss", Config)
	Config.APIVersion = version
	if len(commit) > 0 && len(date) > 0 {
		Config.APIVersion = fmt.Sprintf("%s-%s (built at %s)", version, commit, date)
	}
}

// BuildDate returns the version of this app
func BuildVersion() string {
	if len(commit) > 0 {
		return fmt.Sprintf("%s-%s", version, commit)
	}
	return version
}

// BuildDate returns the date this app was built
func BuildDate() string {
	return date
}

// MustSignInToDockerRegistry returns string bool
func MustSignInToDockerRegistry() *string {
	if Config.MustSignIn {
		return swag.String("yes")
	}
	return swag.String("no")
}
