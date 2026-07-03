package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	Ledger "github.com/starkinfra/sdk-go/starkinfra/ledger"
	LedgerTransaction "github.com/starkinfra/sdk-go/starkinfra/ledgertransaction"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func queryOneLedgerId(t *testing.T) string {

	var params = map[string]interface{}{}
	params["limit"] = 1

	ledgers, errorChannel := Ledger.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case ledger, ok := <-ledgers:
			if !ok {
				break loop
			}
			return ledger.Id
		}
	}
	return ""
}

func TestLedgerTransactionPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	ledgerId := queryOneLedgerId(t)

	transactions, err := LedgerTransaction.Create(Example.LedgerTransaction(ledgerId), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}
}

func TestLedgerTransactionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	ledgerId := queryOneLedgerId(t)

	var params = map[string]interface{}{}
	params["ledgerId"] = ledgerId
	params["limit"] = 5

	transactions, errorChannel := LedgerTransaction.Query(params, nil)
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

func TestLedgerTransactionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	ledgerId := queryOneLedgerId(t)

	var params = map[string]interface{}{}
	params["ledgerId"] = ledgerId
	params["limit"] = 1

	transactions, cursor, err := LedgerTransaction.Page(params, nil)
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

func TestLedgerTransactionGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	ledgerId := queryOneLedgerId(t)

	var params = map[string]interface{}{}
	params["ledgerId"] = ledgerId
	params["limit"] = 10

	var transactionList []LedgerTransaction.LedgerTransaction

	transactions, errorChannel := LedgerTransaction.Query(params, nil)
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
		getTransaction, err := LedgerTransaction.Get(transaction.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getTransaction.Id)
	}
}
