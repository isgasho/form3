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
	params := apiclient.ListParams{PageNum: &zero, PageSize: &two}

	expectedAccountList := apiclient.AccountListData(
		apiclient.AccountListData{
			Data: []apiclient.Account{
				apiclient.Account{
					AccountType:    "accounts",
					ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
					OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Attributes: apiclient.AccountAttributes{Country: "GB",
						BaseCurrency:                "GBP",
						AccountNumber:               "41426819",
						BankID:                      "400300",
						BankIDCode:                  "GBDSC",
						Bic:                         "NWBKGB22",
						Iban:                        "GB11NWBK40030041426819",
						Title:                       "Ms",
						FirstName:                   "Samantha",
						BankAccountName:             "Samantha Holder",
						AlternativeBankAccountNames: []string{"Sam Holder"},
						AccountClassification:       "Personal",
						JointAccount:                false,
						AccountMatchingOptOut:       false,
						SecondaryIdentification:     "A1B2C3D4"}},
				apiclient.Account{AccountType: "accounts",
					ID:             "cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
					OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
					Attributes: apiclient.AccountAttributes{Country: "GB",
						BaseCurrency:                "GBP",
						AccountNumber:               "41426819",
						BankID:                      "400300",
						BankIDCode:                  "GBDSC",
						Bic:                         "NWBKGB22",
						Iban:                        "GB11NWBK40030041426819",
						Title:                       "Ms",
						FirstName:                   "Samantha",
						BankAccountName:             "Samantha Holder",
						AlternativeBankAccountNames: []string{"Sam Holder"},
						AccountClassification:       "Personal",
						JointAccount:                false,
						AccountMatchingOptOut:       false,
						SecondaryIdentification:     "A1B2C3D4"}}},
			Links: apiclient.PageLinks{
				First: "/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",
				Last:  "/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",
				Self:  "/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02",
			}})

	//2. Given a test server has been configured to respond "ok" to a Fetch() request for that accountID
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//4. Then the testserver asserts that the correct request has been received
		assert.Equal(t, "/v1/organisation/accounts?page%5Bnumber%5D=%00&page%5Bsize%5D=%02", req.URL.String())
		assert.Equal(t, "GET", req.Method)
		assert.Equal(t, http.NoBody, req.Body)

		responseJSON := `{"data":[{"attributes":{"account_classification":"Personal",` +
			`"account_matching_opt_out":false,` +
			`"account_number":"41426819",` +
			`"alternative_bank_account_names":["Sam Holder"],` +
			`"bank_account_name":"Samantha Holder",` +
			`"bank_id":"400300",` +
			`"bank_id_code":"GBDSC",` +
			`"base_currency":"GBP",` +
			`"bic":"NWBKGB22",` +
			`"country":"GB",` +
			`"first_name":"Samantha",` +
			`"iban":"GB11NWBK40030041426819",` +
			`"joint_account":false,` +
			`"secondary_identification":"A1B2C3D4",` +
			`"title":"Ms"},` +
			`"created_on":"2020-01-15T21:41:09.508Z",` +
			`"id":"bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
			`"modified_on":"2020-01-15T21:41:09.508Z",` +
			`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
			`"type":"accounts",` +
			`"version":0},` +
			`{"attributes":{"account_classification":"Personal",` +
			`"account_matching_opt_out":false,` +
			`"account_number":"41426819",` +
			`"alternative_bank_account_names":["Sam Holder"],` +
			`"bank_account_name":"Samantha Holder",` +
			`"bank_id":"400300",` +
			`"bank_id_code":"GBDSC",` +
			`"base_currency":"GBP",` +
			`"bic":"NWBKGB22",` +
			`"country":"GB",` +
			`"first_name":"Samantha",` +
			`"iban":"GB11NWBK40030041426819",` +
			`"joint_account":false,` +
			`"secondary_identification":"A1B2C3D4",` +
			`"title":"Ms"},` +
			`"created_on":"2020-01-16T20:01:25.633Z",` +
			`"id":"cd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
			`"modified_on":"2020-01-16T20:01:25.633Z",` +
			`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
			`"type":"accounts",` +
			`"version":0}],` +
			`"links":{"first":"/v1/organisation/accounts?page%5Bnumber%5D=first\u0026page%5Bsize%5D=%02",` +
			`"last":"/v1/organisation/accounts?page%5Bnumber%5D=last\u0026page%5Bsize%5D=%02",` +
			`"self":"/v1/organisation/accounts?page%5Bnumber%5D=%00\u0026page%5Bsize%5D=%02"}}`

		rw.Write([]byte(responseJSON))
	}))

	testServerURL, _ := url.Parse(testServer.URL)

	apiClient := apiclient.New()
	apiClient.BaseURL = testServerURL

	//3. When a the APIClient makes a List request
	response, err := apiClient.List(params)

	//5. And the APIClient receives a response from the API
	assert.Equal(t, expectedAccountList, response)
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
	assert.Equal(t, apiclient.AccountListData{}, response)

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
