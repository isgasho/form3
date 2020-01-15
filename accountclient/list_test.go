package accountclient_test

import (
	"testing"

	"github.com/Rosalita/form3/accountclient"
	"github.com/stretchr/testify/assert"
)

func TestAccountClientCanList(t *testing.T) {


	var tests = []struct {
		response   string
		statusCode int
		params accountclient.ListParams
	}{
		{`{"error_message":"id is not a valid uuid"}`, 200, accountclient.ListParams{}}, 
		{`{"error_message":"id is not a valid uuid"}`, 200, accountclient.ListParams{PageNum: "0"}}, 
		{`{"error_message":"id is not a valid uuid"}`, 200, accountclient.ListParams{PageSize: "1"}}, 
		{`{"error_message":"id is not a valid uuid"}`, 200, accountclient.ListParams{FilterAttr: "country", FilterValue: "GB"}},

	}
	for _, test := range tests {

		accountClient := accountclient.New()
		response, statusCode := accountClient.List(test.params)

		assert.Equal(t, test.response, response)
		assert.Equal(t, test.statusCode, statusCode)
	}
}
