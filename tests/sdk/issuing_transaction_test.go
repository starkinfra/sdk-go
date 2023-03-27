package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingTransaction "github.com/starkinfra/sdk-go/starkinfra/issuingtransaction"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingTransactionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	transactions := IssuingTransaction.Query(params, nil)
	for transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
}

func TestIssuingTransactionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	transactions, cursor, err := IssuingTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
		fmt.Println(transaction.Id)
	}
	fmt.Println(cursor)
}

func TestIssuingTransactionGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var transactionList []IssuingTransaction.IssuingTransaction
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	transactions := IssuingTransaction.Query(paramsQuery, nil)
	for transaction := range transactions {
		transactionList = append(transactionList, transaction)
	}

	transaction, err := IssuingTransaction.Get(transactionList[rand.Intn(len(transactionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, transaction.Id)
	fmt.Println(transaction.Id)
}
