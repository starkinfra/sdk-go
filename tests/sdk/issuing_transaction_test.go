package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingTransaction "github.com/starkinfra/sdk-go/starkinfra/issuingtransaction"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingTransactionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	transactions, errorChannel := IssuingTransaction.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			assert.NotNil(t, transaction.Id)
		}
	}
}

func TestIssuingTransactionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	transactions, cursor, err := IssuingTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
	assert.NotNil(t, cursor)
}

func TestIssuingTransactionGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var transactionList []IssuingTransaction.IssuingTransaction

	transactions, errorChannel := IssuingTransaction.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			transactionList = append(transactionList, transaction)
		}
	}

	for _, transaction := range transactionList {
		getTransaction, err := IssuingTransaction.Get(transaction.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getTransaction.Id)
	}

	assert.Equal(t, limit, len(transactionList))
}
