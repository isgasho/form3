package apiclient_test

import (
	"net/http"
	"errors"
)

// MockHttpClient is used for mocking error responses from the http package
type MockHttpClient struct {
}

//Do is mocked to always return an error
func (m MockHttpClient) Do(req *http.Request) (*http.Response, error) {
	err := errors.New("This is a mocked error")
	return &http.Response{}, err
}