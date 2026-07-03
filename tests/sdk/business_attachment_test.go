package sdk

import (
	"encoding/base64"
	"github.com/starkinfra/sdk-go/starkinfra"
	BusinessAttachment "github.com/starkinfra/sdk-go/starkinfra/businessattachment"
	BusinessIdentity "github.com/starkinfra/sdk-go/starkinfra/businessidentity"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusinessAttachmentPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []BusinessIdentity.BusinessIdentity

	identities, errorChannel := BusinessIdentity.Query(nil, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case identity, ok := <-identities:
			if !ok {
				break loop
			}
			identityList = append(identityList, identity)
		}
	}

	content, _ := base64.StdEncoding.DecodeString("iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAYAAAAfFcSJAAAADUlEQVR42mP8/5+hHgAHggJ/PchI7wAAAABJRU5ErkJggg==")
	attachments, err := BusinessAttachment.Create(Example.BusinessAttachment(identityList[0].Id, "articles-of-incorporation.png", content), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, attachment := range attachments {
		assert.NotNil(t, attachment.Id)
	}
}

func TestBusinessAttachmentQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	attachments, errorChannel := BusinessAttachment.Query(params, nil)
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
			assert.NotNil(t, attachment.Id)
		}
	}
}

func TestBusinessAttachmentPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	attachments, cursor, err := BusinessAttachment.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, attachment := range attachments {
		assert.NotNil(t, attachment.Id)
	}

	assert.NotNil(t, cursor)
}

func TestBusinessAttachmentGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var attachmentList []BusinessAttachment.BusinessAttachment

	attachments, errorChannel := BusinessAttachment.Query(paramsQuery, nil)
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
		var paramsGet = map[string]interface{}{}
		paramsGet["expand"] = []string{"content"}
		getAttachment, err := BusinessAttachment.Get(attachment.Id, paramsGet, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getAttachment.Id)
	}
	assert.Equal(t, limit, len(attachmentList))
}

func TestBusinessAttachmentGetWithExpand(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var attachmentList []BusinessAttachment.BusinessAttachment

	attachments, errorChannel := BusinessAttachment.Query(paramsQuery, nil)
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
		var expand = map[string]interface{}{}
		expand["expand"] = []string{"content"}
		getAttachment, err := BusinessAttachment.Get(attachment.Id, expand, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getAttachment.Id)
		assert.NotEmpty(t, getAttachment.Content)
	}
}

func TestBusinessAttachmentCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "created"

	var attachmentList []BusinessAttachment.BusinessAttachment

	attachments, errorChannel := BusinessAttachment.Query(paramsQuery, nil)
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

	attachment, err := BusinessAttachment.Cancel(attachmentList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, attachment.Id)
	assert.Equal(t, "canceled", attachment.Status)
}
