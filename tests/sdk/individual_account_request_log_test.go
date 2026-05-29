package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualAccountRequestLog "github.com/starkinfra/sdk-go/starkinfra/individualaccountrequest/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndividualAccountRequestLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["types"] = []string{"created"}
	params["accountRequestIds"] = []string{"5189530608992256"}

	logs, errorChannel := IndividualAccountRequestLog.Query(params, nil)
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

func TestIndividualAccountRequestLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, cursor, err := IndividualAccountRequestLog.Page(params, nil)
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

func TestIndividualAccountRequestLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var logList []IndividualAccountRequestLog.Log

	logs, errorChannel := IndividualAccountRequestLog.Query(paramsQuery, nil)
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
		getLog, err := IndividualAccountRequestLog.Get(log.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getLog.Id)
		assert.NotNil(t, getLog.Request.Id)
	}

	assert.Equal(t, limit, len(logList))
}
