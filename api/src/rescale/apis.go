package rescale

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/go-openapi/swag"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	util "github.com/rescale-labs/scaleshift/api/src/http"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"golang.org/x/net/context/ctxhttp"
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
func CoreTypes(ctx context.Context, token string, page, pageSize *int64) ([]*CoreType, error) {
	if useCache {
		if bytes, e := db.GetCache(coretypesCacheKey); e == nil {
			result := []*CoreType{}
			json.Unmarshal(bytes, &result)
			if len(result) > 0 {
				return result, nil
			}
		}
	}
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

	resp, err := util.HTTPSend(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/coretypes/", v3),
		query, nil, headers, 0)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := struct{ Results []*CoreType }{}
	if err = json.Unmarshal(resp, &obj); err != nil {
		return nil, err
	}
	sort.Slice(obj.Results, func(i, j int) bool {
		return obj.Results[i].DisplayOrder < obj.Results[j].DisplayOrder
	})
	bytes, err := json.Marshal(obj.Results)
	if err != nil {
		return nil, err
	}
	duration := 1 * time.Hour
	if err = db.SetCache(coretypesCacheKey, bytes, &duration); err != nil {
		return nil, err
	}
	return obj.Results, nil
}

// Application codes
const (
	ApplicationSingularity = "user_included_singularity_container"
	coretypesCacheKey      = "cached-coretypes"
)

// nolint:misspell
// Analyses returns supported applications
func Analyses(ctx context.Context, token, code string) (*Application, error) {
	query := url.Values{}
	query.Set("code", code)

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := util.HTTPSend(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/analyses/", v3),
		query, nil, headers, 0)
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
func Upload(ctx context.Context, token, path string) (*UploadedFile, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}
	out, in := io.Pipe()
	writer := multipart.NewWriter(in)

	done := make(chan error)
	var resp *http.Response
	go func() {
		req, e := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/files/contents/", v3), out)
		if e != nil {
			done <- e
			return
		}
		req.ContentLength = contentLength(filepath.Base(path)) + stat.Size()
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", token))
		req.Header.Set("Content-Type", writer.FormDataContentType())
		resp, err = ctxhttp.Do(ctx, http.DefaultClient, req)
		done <- err
	}()

	part, err := writer.CreateFormFile("file", filepath.Base(path))
	if err != nil {
		return nil, err
	}
	if _, err = io.Copy(part, f); err != nil {
		return nil, err
	}
	// writer & in should be closed in order to notify http client to close the connection
	if err = writer.Close(); err != nil {
		log.Debug("at writer.Close", err, nil)
	}
	if err = in.Close(); err != nil {
		log.Debug("at pIn.Close", err, nil)
	}
	// Wait for the HTTP request done
	if err = <-done; err != nil {
		return nil, err
	}
	// Format the result
	obj := &UploadedFile{}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		if err = json.Unmarshal(body, obj); err != nil {
			return nil, err
		}
	}
	return obj, nil
}

func contentLength(path string) int64 {
	var body bytes.Buffer
	writer := multipart.NewWriter(&body)
	defer writer.Close()
	if _, err := writer.CreateFormFile("file", path); err != nil {
		return 0
	}
	boundary := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", writer.Boundary()))
	return int64(body.Len()) + int64(boundary.Len())
}

// CreateJob creates a new job
func CreateJob(ctx context.Context, token string, input JobInput) (*string, error) {
	raw, _ := json.Marshal(input)
	body := bytes.NewBuffer(raw)

	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	headers.Add("Content-Type", "application/json")
	resp, err := util.HTTPSend(ctx, http.MethodPost, fmt.Sprintf("%s/jobs/", v3), nil, body, headers, 0)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := &Job{}
	if err := json.Unmarshal(resp, obj); err != nil {
		return nil, err
	}
	return swag.String(obj.ID), nil
}

// Submit submits a job
func Submit(ctx context.Context, token, jobID string) error {
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	_, err := util.HTTPSend(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/jobs/%s/submit/", v3, jobID),
		nil, nil, headers, 0)
	if err != nil {
		return err
	}
	return nil
}

// Status retrieve the job status
func Status(ctx context.Context, token, jobID string) (*JobStatus, error) {
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := util.HTTPSend(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/jobs/%s/statuses/", v3, jobID),
		nil, nil, headers, 0)
	if err != nil {
		return nil, err
	}
	// parse http response body
	obj := &JobStatuses{}
	if err := json.Unmarshal(resp, obj); err != nil {
		return nil, err
	}
	if len(obj.Results) == 0 {
		return nil, fmt.Errorf("Cannot find the specified job")
	}
	obj.Sort()
	return obj.Results[0], nil
}

// Log the log structure
type Log struct {
	Time time.Time
	Log  string
}

var (
	reNewline = regexp.MustCompile(`\r\n|\r|\n`)
	reDateLog = regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}Z)\]: (.*)`)
	reAnsiSeq = regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)
)

// Logs returns output of a job
func Logs(ctx context.Context, token, jobID string) ([]*Log, error) {
	files, err := OutputFiles(ctx, token, jobID)
	if err != nil {
		return nil, err
	}
	result := []*Log{}
	for _, file := range files.Results {
		if !strings.EqualFold(file.Name, "process_output.log") {
			continue
		}
		headers := http.Header{}
		headers.Add("Authorization", fmt.Sprintf("Token %s", token))
		resp, err := util.HTTPSend(
			ctx,
			http.MethodGet,
			file.DownloadURL,
			nil, nil, headers, 0)
		if err != nil {
			return nil, err
		}
		for _, line := range reNewline.Split(string(resp), -1) {
			values := reDateLog.FindStringSubmatch(line)
			if len(values) < 3 {
				continue
			}
			value, e := time.Parse(time.RFC3339, values[1])
			if e != nil {
				continue
			}
			if values[2] == "" {
				continue
			}
			result = append(result, &Log{
				Time: value,
				Log:  reAnsiSeq.ReplaceAllString(values[2], ""),
			})
		}
		sort.Slice(result, func(i, j int) bool {
			return result[i].Time.Before(result[j].Time)
		})
		return result, nil
	}
	return nil, fmt.Errorf("Cannot find logs of the specified job")
}

// outputFiles returns output files information
func OutputFiles(ctx context.Context, token, jobID string) (*Files, error) {
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	resp, err := util.HTTPSend(
		ctx,
		http.MethodGet,
		fmt.Sprintf("%s/jobs/%s/files/", v3, jobID),
		nil, nil, headers, 0)
	if err != nil {
		return nil, err
	}
	// parse http response body
	files := &Files{}
	if err := json.Unmarshal(resp, files); err != nil {
		return nil, err
	}
	return files, nil
}

// Stop the specified job
func Stop(ctx context.Context, token, jobID string) error {
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	_, err := util.HTTPSend(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s/jobs/%s/stop/", v3, jobID),
		nil, nil, headers, 0)
	if err != nil {
		return err
	}
	return nil
}

// Delete the specified job
func Delete(ctx context.Context, token, jobID string) error {
	headers := http.Header{}
	headers.Add("Authorization", fmt.Sprintf("Token %s", token))
	_, err := util.HTTPSend(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("%s/jobs/%s/", v3, jobID),
		nil, nil, headers, 0)
	if err != nil {
		return err
	}
	return nil
}
