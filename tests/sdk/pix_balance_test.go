package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixBalance "github.com/starkinfra/sdk-go/starkinfra/pixbalance"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixBalanceGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	balance := PixBalance.Get(nil)
	assert.NotNil(t, balance.Id)
	fmt.Println(balance.Id)
}
