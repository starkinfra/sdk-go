package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	Ledger "github.com/starkinfra/sdk-go/starkinfra/ledger"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLedgerPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	ledgers, err := Ledger.Create(Example.Ledger(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, ledger := range ledgers {
		assert.NotNil(t, ledger.Id)
	}
}

func TestLedgerQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

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
			assert.NotNil(t, ledger.Id)
		}
	}
}

func TestLedgerPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	ledgers, cursor, err := Ledger.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, ledger := range ledgers {
		assert.NotNil(t, ledger.Id)
	}

	assert.NotNil(t, cursor)
}

func TestLedgerGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var ledgerList []Ledger.Ledger

	ledgers, errorChannel := Ledger.Query(paramsQuery, nil)
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
			ledgerList = append(ledgerList, ledger)
		}
	}

	for _, ledger := range ledgerList {
		getLedger, err := Ledger.Get(ledger.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getLedger.Id)
	}
	assert.Equal(t, limit, len(ledgerList))
}

func TestLedgerUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var ledgerList []Ledger.Ledger

	ledgers, errorChannel := Ledger.Query(paramsQuery, nil)
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
			ledgerList = append(ledgerList, ledger)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["tags"] = []string{"account/123", "updated"}

	ledger, err := Ledger.Update(ledgerList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, ledger.Id)
}
