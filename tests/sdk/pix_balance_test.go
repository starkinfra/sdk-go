package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixBalance "github.com/starkinfra/sdk-go/starkinfra/pixbalance"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixBalanceGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	balance, err := PixBalance.Get(nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, balance.Id)
}
