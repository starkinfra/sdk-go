package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditHolmesLog "github.com/starkinfra/sdk-go/starkinfra/creditholmes/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditHolmesLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 20
	var params = map[string]interface{}{}
	params["limit"] = limit

	var logsList []CreditHolmesLog.Log

	logs, errorChannel := CreditHolmesLog.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			assert.NotNil(t, log.Id)
			logsList = append(logsList, log)
		}
	}
	assert.Equal(t, limit, len(logsList))
}

func TestCreditHolmesLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, _, err := CreditHolmesLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestCreditHolmesLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var logList []CreditHolmesLog.Log

	logs, errorChannel := CreditHolmesLog.Query(paramsQuery, nil)

	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			logList = append(logList, log)
		}
	}
	for _, log := range logList {
		assert.NotNil(t, log.Id)
	}

	assert.Equal(t, limit, len(logList))
}
