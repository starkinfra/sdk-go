package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingPurchaseLog "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingPurchaseLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, errorChannel := IssuingPurchaseLog.Query(params, nil)
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
		}
	}
}

func TestIssuingPurchaseLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, cursor, err := IssuingPurchaseLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}
	assert.NotNil(t, cursor)
}

func TestIssuingPurchaseLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var logList []IssuingPurchaseLog.Log

	logs, errorChannel := IssuingPurchaseLog.Query(paramsQuery, nil)
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
		getLog, err := IssuingPurchaseLog.Get(log.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getLog.Id)
	}

	assert.Equal(t, limit, len(logList))
}
