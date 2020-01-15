package accountclient_test

import (
	"testing"

	"github.com/Rosalita/form3/accountclient"
	"github.com/stretchr/testify/assert"
)

func TestAccountClientCanCreate(t *testing.T) {

	validAccount := accountclient.AccountData{
		Data: accountclient.Account{
            AccountType:    "accounts",
            ID:             "bd27e265-9605-4b4b-a0e5-3003ea9cc4dc",
            OrganisationID: "eb0bd6f5-c3f5-44b2-b677-acd23cdde73c",
            Attributes: accountclient.AccountAttributes{
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

	var tests = []struct {
		account  accountclient.AccountData
		response   string
		statusCode int
	}{
		{validAccount, `{"error_message":"id is not a valid uuid"}`, 201},
	}
	for _, test := range tests {

		accountClient := accountclient.New()
		response, statusCode := accountClient.Create(test.account)

		assert.Equal(t, test.response, response)
		assert.Equal(t, test.statusCode, statusCode)
	}
}
