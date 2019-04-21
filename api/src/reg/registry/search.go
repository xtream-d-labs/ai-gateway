package registry

import (
	"fmt"
	"strings"
)

type SearchResult struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	StarCount   int    `json:"star_count"`
	IsTrusted   bool   `json:"is_trusted"`
	IsAutomated bool   `json:"is_automated"`
	IsOfficial  bool   `json:"is_official"`
}

type searchResponse struct {
	NumResults int            `json:"num_results"`
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	NumPages   int            `json:"num_pages"`
	Results    []SearchResult `json:"results"`
}

// Search returns the repositories in a registry.
func (r *Registry) Search(user string) ([]SearchResult, error) {
	if user == "" {
		user = "/v1/search"
	}
	if !strings.Contains(user, "n=") {
		user += fmt.Sprintf("&n=%d", 250)
	}
	uri := r.url(user)
	r.Logf("registry.search url=%s", uri)

	var response searchResponse
	if _, err := r.getJSON(uri, &response, false); err != nil {
		return nil, err
	}
	return response.Results, nil
}
