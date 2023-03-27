package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingWithdrawal "github.com/starkinfra/sdk-go/starkinfra/issuingwithdrawal"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingWithdrawalPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	withdrawal, err := IssuingWithdrawal.Create(Example.IssuingWithdrawal(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, withdrawal.Id)
	fmt.Println(withdrawal.Id)

}

func TestIssuingWithdrawalQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	withdrawals := IssuingWithdrawal.Query(params, nil)
	for withdrawal := range withdrawals {
		assert.NotNil(t, withdrawal.Id)
		fmt.Println(withdrawal.Id)
	}
}

func TestIssuingWithdrawalPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	withdrawals, cursor, err := IssuingWithdrawal.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, withdrawal := range withdrawals {
		assert.NotNil(t, withdrawal.Id)
		fmt.Println(withdrawal.Id)
	}

	fmt.Println(cursor)
}

func TestIssuingWithdrawalGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var withdrawalList []IssuingWithdrawal.IssuingWithdrawal
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	withdrawals := IssuingWithdrawal.Query(paramsQuery, nil)
	for withdrawal := range withdrawals {
		withdrawalList = append(withdrawalList, withdrawal)
	}

	withdrawal, err := IssuingWithdrawal.Get(withdrawalList[rand.Intn(len(withdrawalList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, withdrawal.Id)
	fmt.Println(withdrawal.Id)
}
