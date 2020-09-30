package repoutils

import (
	"strings"

	"docker.io/go-docker/api/types"
	"github.com/docker/distribution/reference"
	"github.com/scaleshift/scaleshift/api/src/log"
)

const (
	// DefaultDockerRegistry is the default docker registry address.
	DefaultDockerRegistry = "https://registry-1.docker.io"

	latestTagSuffix = ":latest"
)

// GetAuthConfig returns the docker registry AuthConfig.
// Optionally takes in the authentication values, otherwise pulls them from the
// docker config file.
func GetAuthConfig(username, password, registry string) (types.AuthConfig, error) {
	if username != "" && password != "" && registry != "" {
		return types.AuthConfig{
			Username:      username,
			Password:      password,
			ServerAddress: registry,
		}, nil
	}

	// Don't use any authentication.
	// We should never get here.
	log.Info("Not using any authentication", nil, nil)
	return types.AuthConfig{}, nil
}

// GetRepoAndRef parses the repo name and reference.
func GetRepoAndRef(image string) (repo, ref string, err error) {
	if image == "" {
		return "", "", reference.ErrNameEmpty
	}

	image = addLatestTagSuffix(image)

	var parts []string
	if strings.Contains(image, "@") {
		parts = strings.Split(image, "@")
	} else if strings.Contains(image, ":") {
		parts = strings.Split(image, ":")
	}

	repo = parts[0]
	if len(parts) > 1 {
		ref = parts[1]
	}

	return repo, ref, nil
}

// addLatestTagSuffix adds :latest to the image if it does not have a tag
func addLatestTagSuffix(image string) string {
	if !strings.Contains(image, ":") {
		return image + latestTagSuffix
	}
	return image
}
