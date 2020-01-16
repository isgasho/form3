package apiclient_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rosalita/form3/apiclient"
	"github.com/stretchr/testify/assert"
)

func TestAPIClientCanFetch(t *testing.T) {

	//1. Given a valid accountID
	accountID := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"

	//2. Given a test server has been configured to respond "ok" to a Fetch() request for that accountID
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//4. Then the testserver asserts that the correct request has been received
		assert.Equal(t, "/account/ad27e265-9605-4b4b-a0e5-3003ea9cc4dc", req.URL.String())
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, http.NoBody, req.Body)

		rw.Write([]byte("ok"))
	}))

	testServerURL, _ := url.Parse(testServer.URL)

	apiClient := apiclient.New()
	apiClient.BaseURL = testServerURL

	//3. When a the APIClient makes a Fetch request for the accountID
	response := apiClient.Fetch(accountID)

	//5. And the APIClient receives a response from the API
	assert.Equal(t, "ok", response)
}