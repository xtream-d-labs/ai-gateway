package lib

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/docker/docker/api/types"
	util "github.com/xtream-d-labs/ai-gateway/api/src/http"
	"github.com/xtream-d-labs/ai-gateway/api/src/log"
)

/**
 * VMWare/Harbor API
 * @see https://raw.githubusercontent.com/vmware/harbor/master/docs/swagger.yaml
 */

// HarborProject defines the repo information
type HarborProject struct {
	ID   int64  `json:"project_id"`
	Name string `json:"name"`
}

// HarborRepository defines the repo information
type HarborRepository struct {
	ProjectID   int64  `json:"project_id"`
	ProjectName string `json:"project_name"`
	Public      bool   `json:"project_public"`
	Name        string `json:"repository_name"`
	Tags        int64  `json:"tags_count"`
}

// HarborSerchResult defines the response of /api/search
type HarborSerchResult struct {
	Projects     []*HarborProject    `json:"project"`
	Repositories []*HarborRepository `json:"repository"`
}

// HarborRepositories returns harbor repos
func HarborRepositories(ctx context.Context, auth types.AuthConfig) ([]*HarborRepository, error) {
	headers := http.Header{}
	headers.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(
		auth.Username+":"+auth.Password,
	)))
	resp, err := util.HTTPSend(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/api/search", auth.ServerAddress),
		nil, nil, headers, 0)
	if err != nil {
		log.Error("util.HTTPSend", err, nil)
		return nil, err
	}
	obj := HarborSerchResult{}
	if err := json.Unmarshal(resp, &obj); err != nil {
		return nil, err
	}
	return obj.Repositories, nil
}
