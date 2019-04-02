package rescale

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sort"
	"time"

	"github.com/dgraph-io/badger"
	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/log"
)

var (
	v3       string
	useCache bool
)

func init() {
	v3 = fmt.Sprintf("%s/api/v3", config.Config.RescaleEndpoint)
	useCache = true
}

// SetEndpoint sets Rescale API endpoint
func SetEndpoint(endpoint string) {
	v3 = fmt.Sprintf("%s/api/v3", endpoint)
}

// EnableCache sets cache availability
func EnableCache(enabled bool) {
	useCache = enabled
}

// CoreTypes returns supported core types
func CoreTypes(token string, page, pageSize *int64) ([]*CoreType, error) {
	if useCache {
		if bytes, e := db.GetValueSimple(coretypesCacheKey); e == nil {
			result := []*CoreType{}
			json.Unmarshal(bytes, &result)
			if len(result) > 0 {
				return result, nil
			}
		}
	}
	ctx := context.Background()

	// send http request
	query := url.Values{}
	if page == nil {
		page = swag.Int64(1)
	}
	if pageSize == nil {
		pageSize = swag.Int64(999)
	}
	query.Set("page", fmt.Sprintf("%d", swag.Int64Value(page)))
	query.Set("page_size", fmt.Sprintf("%d", swag.Int64Value(pageSize)))

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))

	resp, err := send(ctx, "GET", fmt.Sprintf("%s/coretypes/", v3), query, nil, headers)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := struct{ Results []*CoreType }{}
	if err := json.Unmarshal(resp, &obj); err != nil {
		return nil, err
	}
	sort.Slice(obj.Results, func(i, j int) bool {
		return obj.Results[i].DisplayOrder < obj.Results[j].DisplayOrder
	})
	db.SetValue(func(txn *badger.Txn) error {
		bytes, err := json.Marshal(obj.Results)
		if err != nil {
			return err
		}
		return txn.SetWithTTL([]byte(coretypesCacheKey), bytes, 1*time.Hour)
	})
	return obj.Results, nil
}

// Application codes
const (
	ApplicationSingularity    = "user_included_singularity_container"
	ApplicationSingularityMPI = "user_included_singularity_container_mpi"

	coretypesCacheKey = "cached-coretypes"
)

// nolint:misspell
// Analyses returns supported applications
func Analyses(token, code string) (*Application, error) {
	ctx := context.Background()

	// send http request
	query := url.Values{}
	query.Set("code", code)

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := send(ctx, "GET", fmt.Sprintf("%s/analyses/", v3), query, nil, headers)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := struct{ Results []*Application }{}
	if err := json.Unmarshal(resp, &obj); err != nil {
		return nil, err
	}
	if len(obj.Results) == 0 {
		return nil, nil
	}
	return obj.Results[0], nil
}

// Upload uploads specified files
func Upload(token string, body io.Reader, contentType string) (*UploadedFile, error) {
	ctx := context.Background()

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	headers.Add("Content-Type", contentType)
	resp, err := send(ctx, "POST", fmt.Sprintf("%s/files/contents/", v3), nil, body, headers)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := &UploadedFile{}
	if err := json.Unmarshal(resp, obj); err != nil {
		return nil, err
	}
	return obj, nil
}

// CreateJob creates a new job
func CreateJob(token string, input JobInput) (*string, error) {
	ctx := context.Background()

	raw, _ := json.Marshal(input)
	log.Debug("Rescale submit a job", nil, &log.Map{
		"body": string(raw),
	})
	body := bytes.NewBuffer(raw)

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	headers.Add("Content-Type", "application/json")
	resp, err := send(ctx, "POST", fmt.Sprintf("%s/jobs/", v3), nil, body, headers)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := &JobStatus{}
	if err := json.Unmarshal(resp, obj); err != nil {
		return nil, err
	}
	return swag.String(obj.ID), nil
}

// Submit submits a job
func Submit(token, ID string) (*string, error) {
	ctx := context.Background()

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := send(ctx, "POST", fmt.Sprintf("%s/jobs/%s/submit/", v3, ID), nil, nil, headers)
	if err != nil {
		return nil, err
	}
	log.Info(string(resp), nil, nil)
	return nil, nil
}
