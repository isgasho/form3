package apiclient_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rosalita/form3/apiclient"
	"github.com/stretchr/testify/assert"
)

var ( // Global variables referenced by tests requiring pointers to ints
	zero = 0
	one  = 1
	two  = 2
)

func TestAPIClientCanList(t *testing.T) {

	//1. Given valid parameters for a List request
	params := apiclient.ListParams{PageNum: &zero, PageSize: &one}

	//2. Given a test server has been configured to respond "ok" to a Fetch() request for that accountID
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//4. Then the testserver asserts that the correct request has been received
		assert.Equal(t, "/account?page%5Bnumber%5D=%00&page%5Bsize%5D=%01", req.URL.String())
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, http.NoBody, req.Body)

		rw.Write([]byte("ok"))
	}))

	testServerURL, _ := url.Parse(testServer.URL)

	apiClient := apiclient.New()
	apiClient.BaseURL = testServerURL

	//3. When a the APIClient makes a List request
	response, err := apiClient.List(params)

	//5. And the APIClient receives a response from the API
	assert.Equal(t, "ok", response)
	assert.Equal(t, nil, err)
}

func TestListCanHandleHTTPErrors(t *testing.T) {

	//1. Given valid parameters for a List request
	params := apiclient.ListParams{PageNum: &zero, PageSize: &one}

	//2. When the HTTPClient is mocked to return an error
	apiClient := apiclient.New()
	apiClient.HTTPClient = MockHttpClient{}

	//3. and APIClient makes a List request 
	response, err := apiClient.List(params)

	//4. Then the response is empty
	assert.Equal(t, "", response)

	//5. And the err is an error containing the mocked error message
	assert.NotEqual(t, nil, err)
	assert.Equal(t, "This is a mocked error", err.Error())
}


func TestEncodeOptionalQueryParametersString(t *testing.T) {

	//1. Given some parameters
	emptyParams := apiclient.ListParams{}
	pageNumZero := apiclient.ListParams{PageNum: &zero}
	pageNumOne := apiclient.ListParams{PageNum: &one}
	pageNumTwo := apiclient.ListParams{PageNum: &two}
	pageSizeZero := apiclient.ListParams{PageSize: &zero}
	pageSizeOne := apiclient.ListParams{PageSize: &one}
	pageSizeTwo := apiclient.ListParams{PageSize: &two}
	pageNumAndSize := apiclient.ListParams{PageNum: &zero, PageSize: &one}

	var tests = []struct {
		params      apiclient.ListParams
		queryString string
	}{
		{emptyParams, ""},
		{pageNumZero, "page%5Bnumber%5D=%00"},
		{pageNumOne, "page%5Bnumber%5D=%01"},
		{pageNumTwo, "page%5Bnumber%5D=%02"},
		{pageSizeZero, "page%5Bsize%5D=%00"},
		{pageSizeOne, "page%5Bsize%5D=%01"},
		{pageSizeTwo, "page%5Bsize%5D=%02"},
		{pageNumAndSize, "page%5Bnumber%5D=%00&page%5Bsize%5D=%01"},
	}
	for _, test := range tests {

		//2. When the parameters are encoded into a query string
		queryString := apiclient.EncodeOptionalQueryParameters(test.params)

		//3. Then the correct query string is returned
		assert.Equal(t, test.queryString, queryString)
	}
}
