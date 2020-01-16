package apiclient

import (
	"net/url"
	"net/http"
	"time"
)

var (
	defaultBaseURL, _  = url.Parse("http://localhost:8080/v1/organisation")
)

//APIClient is the type used to interface with the Accounts API
type APIClient struct {
	BaseURL   *url.URL
	httpClient *http.Client
}

// New creates a new instance of an APIclient
func New() *APIClient {
	return &APIClient{
		BaseURL: defaultBaseURL, 
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
