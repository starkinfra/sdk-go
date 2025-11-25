package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingBalance "github.com/starkinfra/sdk-go/starkinfra/issuingbalance"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingBalanceQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	balance, err := IssuingBalance.Get(nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, balance.Id)
}
