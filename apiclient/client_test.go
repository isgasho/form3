package apiclient_test

import (
	"testing"

	"github.com/Rosalita/form3/apiclient"
	"github.com/stretchr/testify/assert"
)

func TestCanInstantiateAnAPIClient(t *testing.T) {

	client := apiclient.New()
	assert.IsType(t, &apiclient.APIClient{}, client)
}
