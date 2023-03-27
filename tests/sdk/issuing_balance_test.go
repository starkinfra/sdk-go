package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingBalance "github.com/starkinfra/sdk-go/starkinfra/issuingbalance"
	"github.com/starkinfra/sdk-go/tests/utils"
	"testing"
)

func TestIssuingBalanceQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	balance := IssuingBalance.Get(nil)
	fmt.Println(balance)
}
