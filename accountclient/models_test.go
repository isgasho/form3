package accountclient_test

import (
    "encoding/json"
    "testing"

    "github.com/Rosalita/form3/accountclient"
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

    expectedAccountData := accountclient.AccountData{
        Data: accountclient.Account{
            AccountType:    "A",
            ID:             "B",
            OrganisationID: "C",
            Attributes: accountclient.AccountAttributes{
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

    var accountData accountclient.AccountData
    err := json.Unmarshal([]byte(rawJSON), &accountData)
    assert.Equal(t, nil, err)
    assert.Equal(t, expectedAccountData, accountData)
}

func TestAccountDataMarshalJson(t *testing.T) {
    AccountData := accountclient.AccountData{
        Data: accountclient.Account{
            AccountType:    "A",
            ID:             "B",
            OrganisationID: "C",
            Attributes: accountclient.AccountAttributes{
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

    expectedJson := `{` +
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
    assert.Equal(t, expectedJson, string(json))

}
