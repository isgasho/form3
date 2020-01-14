package accountclient_test

import (
	"testing"

	"github.com/Rosalita/form3/accountclient"
	"github.com/stretchr/testify/assert"
)

func TestAccountClientCanFetch(t *testing.T) {

	invalidAccountID := "1"

	var tests = []struct {
		accountID  string
		response   string
		statusCode int
	}{
		{invalidAccountID, `{"error_message":"id is not a valid uuid"}`, 400},
	}
	for _, test := range tests {

		accountClient := accountclient.New()
		response, statusCode := accountClient.Fetch(test.accountID)

		assert.Equal(t, test.response, response)
		assert.Equal(t, test.statusCode, statusCode)
	}
}
