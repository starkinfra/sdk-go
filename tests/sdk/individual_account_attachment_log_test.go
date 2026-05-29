package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualAccountAttachmentLog "github.com/starkinfra/sdk-go/starkinfra/individualaccountattachment/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndividualAccountAttachmentLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["types"] = []string{"created"}
	params["attachmentIds"] = []string{"5189530608992256"}

	logs, errorChannel := IndividualAccountAttachmentLog.Query(params, nil)
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

func TestIndividualAccountAttachmentLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["attachmentIds"] = []string{"5189530608992256"}

	logs, cursor, err := IndividualAccountAttachmentLog.Page(params, nil)
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

func TestIndividualAccountAttachmentLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var logList []IndividualAccountAttachmentLog.Log

	logs, errorChannel := IndividualAccountAttachmentLog.Query(paramsQuery, nil)
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
		getLog, err := IndividualAccountAttachmentLog.Get(log.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getLog.Id)
		assert.NotNil(t, getLog.Attachment.Id)
	}

	assert.Equal(t, limit, len(logList))
}
