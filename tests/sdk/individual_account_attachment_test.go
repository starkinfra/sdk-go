package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualAccountAttachment "github.com/starkinfra/sdk-go/starkinfra/individualaccountattachment"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestIndividualAccountAttachmentPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	attachments, err := IndividualAccountAttachment.Create(Example.IndividualAccountAttachment(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, attachment := range attachments {
		assert.NotNil(t, attachment.Id)
		assert.NotNil(t, attachment.Status)
	}
}

func TestIndividualAccountAttachmentPostEncodesDataUrl(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	attachments := Example.IndividualAccountAttachment()

	assert.NotEmpty(t, attachments[0].Content)
	assert.NotEmpty(t, attachments[0].ContentType)

	_, err := IndividualAccountAttachment.Create(attachments, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.True(t, strings.HasPrefix(attachments[0].Content, "data:image/png;base64,"))
	assert.Equal(t, "", attachments[0].ContentType)
}

func TestIndividualAccountAttachmentPostTypeEnum(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	validTypes := []string{"drivers-license-front", "drivers-license-back", "identity-front", "identity-back"}

	attachments := Example.IndividualAccountAttachment()
	for _, attachment := range attachments {
		assert.Contains(t, validTypes, attachment.Type)
	}
}

func TestIndividualAccountAttachmentQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	attachments, errorChannel := IndividualAccountAttachment.Query(params, nil)
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

func TestIndividualAccountAttachmentPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	attachments, cursor, err := IndividualAccountAttachment.Page(params, nil)
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

func TestIndividualAccountAttachmentInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var attachmentList []IndividualAccountAttachment.IndividualAccountAttachment

	attachments, errorChannel := IndividualAccountAttachment.Query(paramsQuery, nil)
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
		getAttachment, err := IndividualAccountAttachment.Get(attachment.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getAttachment.Id)
		assert.NotNil(t, getAttachment.Created)
	}

	assert.Equal(t, limit, len(attachmentList))
}

func TestIndividualAccountAttachmentCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	created, errCreate := IndividualAccountAttachment.Create(Example.IndividualAccountAttachment(), nil)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	canceled, err := IndividualAccountAttachment.Cancel(created[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, canceled.Id)
	assert.Equal(t, "deleted", canceled.Status)
}

func TestIndividualAccountAttachmentCancelIdempotent(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	created, errCreate := IndividualAccountAttachment.Create(Example.IndividualAccountAttachment(), nil)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	_, errFirst := IndividualAccountAttachment.Cancel(created[0].Id, nil)
	if errFirst.Errors != nil {
		for _, e := range errFirst.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	canceledAgain, err := IndividualAccountAttachment.Cancel(created[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("expected idempotent cancel to succeed, got code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, canceledAgain.Id)
	assert.Equal(t, "deleted", canceledAgain.Status)
}

func TestIndividualAccountAttachmentPostInvalidType(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	attachments := Example.IndividualAccountAttachment()
	attachments[0].Type = "not-a-real-type"

	_, err := IndividualAccountAttachment.Create(attachments, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountAttachmentPostInvalidContent(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	attachments := Example.IndividualAccountAttachment()
	attachments[0].Content = ""

	_, err := IndividualAccountAttachment.Create(attachments, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountAttachmentPostInvalidContentType(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	attachments := Example.IndividualAccountAttachment()
	attachments[0].ContentType = ""

	_, err := IndividualAccountAttachment.Create(attachments, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountAttachmentPostNotFoundParent(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	attachments := Example.IndividualAccountAttachment()
	attachments[0].AccountRequestId = "0"

	_, err := IndividualAccountAttachment.Create(attachments, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountAttachmentGetNotFound(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	_, err := IndividualAccountAttachment.Get("0", nil)
	assert.NotNil(t, err.Errors)
}
