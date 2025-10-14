package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingWithdrawal "github.com/starkinfra/sdk-go/starkinfra/issuingwithdrawal"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingWithdrawalPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	withdrawal, err := IssuingWithdrawal.Create(Example.IssuingWithdrawal(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, withdrawal.Id)
}

func TestIssuingWithdrawalQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	withdrawals, errorChannel := IssuingWithdrawal.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case withdrawal, ok := <-withdrawals:
			if !ok {
				break loop
			}
			assert.NotNil(t, withdrawal.Id)
		}
	}
}

func TestIssuingWithdrawalPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	withdrawals, cursor, err := IssuingWithdrawal.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, withdrawal := range withdrawals {
		assert.NotNil(t, withdrawal.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingWithdrawalGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var withdrawalList []IssuingWithdrawal.IssuingWithdrawal

	withdrawals, errorChannel := IssuingWithdrawal.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case withdrawal, ok := <-withdrawals:
			if !ok {
				break loop
			}
			withdrawalList = append(withdrawalList, withdrawal)
		}
	}

	for _, withdrawal := range withdrawalList {
		getWithdrawal, err := IssuingWithdrawal.Get(withdrawal.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getWithdrawal.Id)
	}

	assert.Equal(t, limit, len(withdrawalList))
}
