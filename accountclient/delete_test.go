package accountclient_test

import (
	"testing"

	"github.com/Rosalita/form3/accountclient"
	"github.com/stretchr/testify/assert"
)

func TestAccountClientCanDelete(t *testing.T) {

	//invalidAccountID := "1"
	validAccountID := "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc"
	
	var tests = []struct {
		accountID  string
		version int
		response   string
		statusCode int
	}{
		{validAccountID, 0, `{"error_message":"id is not a valid uuid"}`, 201},
	}
	for _, test := range tests {

		accountClient := accountclient.New()
		response, statusCode := accountClient.Delete(test.accountID, test.version)

		assert.Equal(t, test.response, response)
		assert.Equal(t, test.statusCode, statusCode)
	}
}
