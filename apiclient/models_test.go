package apiclient_test

import (
	"encoding/json"
	"testing"

	"github.com/Rosalita/form3/apiclient"
	"github.com/stretchr/testify/assert"
)

func TestAccountDataUnmarshalJSON(t *testing.T) {
	rawJSON := `{` +
		`"data":{` +
		`"type":"A",` +
		`"id":"B",` +
		`"organisation_id":"C",` +
		`"Attributes":{"country":"D",` +
		`"base_currency":"E",` +
		`"account_number":"F",` +
		`"bank_id":"G",` +
		`"bank_id_code":"H",` +
		`"bic":"I",` +
		`"iban":"J",` +
		`"title":"K",` +
		`"first_name":"L",` +
		`"bank_account_name":"M",` +
		`"alternative_bank_account_names":["N","O"],` +
		`"account_classification":"P",` +
		`"joint_account":true,` +
		`"account_matching_opt_out":true,` +
		`"secondary_identification":"Q"` +
		`}}}`

	expectedAccountData := apiclient.AccountData{
		Data: apiclient.Account{
			AccountType:    "A",
			ID:             "B",
			OrganisationID: "C",
			Attributes: apiclient.AccountAttributes{
				Country:                     "D",
				BaseCurrency:                "E",
				AccountNumber:               "F",
				BankID:                      "G",
				BankIDCode:                  "H",
				Bic:                         "I",
				Iban:                        "J",
				Title:                       "K",
				FirstName:                   "L",
				BankAccountName:             "M",
				AlternativeBankAccountNames: []string{"N", "O"},
				AccountClassification:       "P",
				JointAccount:                true,
				AccountMatchingOptOut:       true,
				SecondaryIdentification:     "Q",
			},
		},
	}

	var accountData apiclient.AccountData
	err := json.Unmarshal([]byte(rawJSON), &accountData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAccountData, accountData)
}

func TestAccountDataMarshalJson(t *testing.T) {
	AccountData := apiclient.AccountData{
		Data: apiclient.Account{
			AccountType:    "A",
			ID:             "B",
			OrganisationID: "C",
			Attributes: apiclient.AccountAttributes{
				Country:                     "D",
				BaseCurrency:                "E",
				AccountNumber:               "F",
				BankID:                      "G",
				BankIDCode:                  "H",
				Bic:                         "I",
				Iban:                        "J",
				Title:                       "K",
				FirstName:                   "L",
				BankAccountName:             "M",
				AlternativeBankAccountNames: []string{"N", "O"},
				AccountClassification:       "P",
				JointAccount:                true,
				AccountMatchingOptOut:       true,
				SecondaryIdentification:     "Q",
			},
		},
	}

	expectedJSON := `{` +
		`"data":{` +
		`"type":"A",` +
		`"id":"B",` +
		`"organisation_id":"C",` +
		`"Attributes":{"country":"D",` +
		`"base_currency":"E",` +
		`"account_number":"F",` +
		`"bank_id":"G",` +
		`"bank_id_code":"H",` +
		`"bic":"I",` +
		`"iban":"J",` +
		`"title":"K",` +
		`"first_name":"L",` +
		`"bank_account_name":"M",` +
		`"alternative_bank_account_names":["N","O"],` +
		`"account_classification":"P",` +
		`"joint_account":true,` +
		`"account_matching_opt_out":true,` +
		`"secondary_identification":"Q"` +
		`}}}`

	json, err := json.Marshal(&AccountData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJSON, string(json))

}

func TestAccountListDataUnmarshalJSON(t *testing.T) {
	rawJSON := `{"data":[{"attributes":{"account_classification":"Personal",` +
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

	expectedAccountListData := apiclient.AccountListData(
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

	var accountListData apiclient.AccountListData
	err := json.Unmarshal([]byte(rawJSON), &accountListData)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedAccountListData, accountListData)
}
