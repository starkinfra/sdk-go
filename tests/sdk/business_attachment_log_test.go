package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	BusinessAttachmentLog "github.com/starkinfra/sdk-go/starkinfra/businessattachment/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusinessAttachmentLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	logs, errorChannel := BusinessAttachmentLog.Query(paramsQuery, nil)
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

func TestBusinessAttachmentLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := BusinessAttachmentLog.Page(params, nil)
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

func TestBusinessAttachmentLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var attachmentList []BusinessAttachmentLog.Log

	attachments, errorChannel := BusinessAttachmentLog.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case attachment, ok := <-attachments:
			if !ok {
				break loop
			}
			attachmentList = append(attachmentList, attachment)
		}
	}

	for _, attachment := range attachmentList {
		getAttachment, err := BusinessAttachmentLog.Get(attachment.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getAttachment.Id)
	}
	assert.Equal(t, limit, len(attachmentList))
}
