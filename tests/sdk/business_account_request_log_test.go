package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	BusinessAccountRequestLog "github.com/starkinfra/sdk-go/starkinfra/businessaccountrequest/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusinessAccountRequestLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var logsList []BusinessAccountRequestLog.Log

	logs, errorChannel := BusinessAccountRequestLog.Query(paramsQuery, nil)
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

func TestBusinessAccountRequestLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := BusinessAccountRequestLog.Page(params, nil)
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

func TestBusinessAccountRequestLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var logsList []BusinessAccountRequestLog.Log

	logs, errorChannel := BusinessAccountRequestLog.Query(paramsQuery, nil)
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
			logsList = append(logsList, log)
		}
	}

	for _, log := range logsList {
		getLog, err := BusinessAccountRequestLog.Get(log.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getLog.Id)
	}

	assert.Equal(t, limit, len(logsList))
}
