package apiclient

import (
	"net/url"
	"net/http"
	"time"
)

var (
	defaultBaseURL, _  = url.Parse("http://localhost:8080/v1/organisation")
)

type httpClientIface interface {
	Do(req *http.Request) (*http.Response, error)
}

//APIClient is the type used to interface with the Accounts API
type APIClient struct {
	BaseURL   *url.URL
	HTTPClient httpClientIface
}

// New creates a new instance of an APIclient
func New() *APIClient {
	return &APIClient{
		BaseURL: defaultBaseURL, 
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}
