package apiclient_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Rosalita/form3/apiclient"
	"github.com/stretchr/testify/assert"
)

func TestAPIClientCanCreate(t *testing.T) {

	//1. Given a valid account data
	validAccountData := apiclient.AccountData{
		Data: apiclient.Account{
			AccountType:    "accounts",
			ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
			OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
			Attributes: apiclient.AccountAttributes{
				Country:                     "GB",
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
				SecondaryIdentification:     "A1B2C3D4",
			},
		},
	}

	//2. Given a test server has been configured to respond "ok" to a Create() request with valid data
	testServer := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//4. Then the testserver asserts that the correct request has been received
		assert.Equal(t, "/account", req.URL.String())
		assert.Equal(t, "POST", req.Method)

		expectedRequestBody := `{` +
			`"data":{` +
			`"type":"accounts",` +
			`"id":"bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",` +
			`"organisation_id":"eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",` +
			`"Attributes":{"country":"GB",` +
			`"base_currency":"GBP",` +
			`"account_number":"41426819",` +
			`"bank_id":"400300",` +
			`"bank_id_code":"GBDSC",` +
			`"bic":"NWBKGB22",` +
			`"iban":"GB11NWBK40030041426819",` +
			`"title":"Ms",` +
			`"first_name":"Samantha",` +
			`"bank_account_name":"Samantha Holder",` +
			`"alternative_bank_account_names":["Sam Holder"],` +
			`"account_classification":"Personal",` +
			`"joint_account":false,` +
			`"account_matching_opt_out":false,` +
			`"secondary_identification":"A1B2C3D4"` +
			`}}}`

		body, _ := ioutil.ReadAll(req.Body)
		assert.Equal(t, expectedRequestBody, string(body))

		rw.Write([]byte("ok"))
	}))

	testServerURL, _ := url.Parse(testServer.URL)

	apiClient := apiclient.New()
	apiClient.BaseURL = testServerURL

	//3. When a the APIClient makes a Create request with validAccountData
	response := apiClient.Create(validAccountData)

	//5. And the APIClient receives a response from the API
	assert.Equal(t, "ok", response)
}
