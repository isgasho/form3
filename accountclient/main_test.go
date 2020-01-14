package accountclient_test

import (
	"testing"

	"github.com/Rosalita/form3/accountclient"
	"github.com/stretchr/testify/assert"
)

func TestCanInstantiateAnAccountClient(t *testing.T) {

	accountClient := accountclient.New()
	assert.IsType(t, &accountclient.Client{}, accountClient)
}
