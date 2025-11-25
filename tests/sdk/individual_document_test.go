package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualDocument "github.com/starkinfra/sdk-go/starkinfra/individualdocument"
	IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIndividualDocumentPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []IndividualIdentity.IndividualIdentity

	identities, errorChannel := IndividualIdentity.Query(nil, nil)
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

	bytesFront, _ := os.ReadFile("../utils/identity/identity-front-face.png")
	frontDocuments, errFront := IndividualDocument.Create(Example.IndividualDocument(identityList[0].Id, "identity-front", bytesFront), nil)
	if errFront.Errors != nil {
		for _, e := range errFront.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, frontDocument := range frontDocuments {
		assert.NotNil(t, frontDocument.Id)
	}

	bytesBack, _ := os.ReadFile("../utils/identity/identity-back-face.png")
	backDocuments, errBack := IndividualDocument.Create(Example.IndividualDocument(identityList[0].Id, "identity-back", bytesBack), nil)
	if errBack.Errors != nil {
		for _, e := range errBack.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, backDocument := range backDocuments {
		assert.NotNil(t, backDocument.Id)
	}

	bytesSelfie, _ := os.ReadFile("../utils/identity/walter-white.png")
	selfieDocuments, errSelfie := IndividualDocument.Create(Example.IndividualDocument(identityList[0].Id, "selfie", bytesSelfie), nil)
	if errSelfie.Errors != nil {
		for _, e := range errSelfie.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, selfieDocument := range selfieDocuments {
		assert.NotNil(t, selfieDocument.Id)
	}

	identity, err := IndividualIdentity.Update(identityList[0].Id, "processing", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.Equal(t, "processing", identity.Status)
}

func TestIndividualDocumentQueryAndGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var documentList []IndividualDocument.IndividualDocument

	documents, errorChannel := IndividualDocument.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case document, ok := <-documents:
			if !ok {
				break loop
			}
			documentList = append(documentList, document)
		}
	}

	for _, document := range documentList {
		getDocument, err := IndividualDocument.Get(document.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}

		assert.NotNil(t, getDocument.Id)
	}
	assert.Equal(t, limit, len(documentList))
}

func TestIndividualDocumentQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "created"

	documents, errorChannel := IndividualDocument.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case document, ok := <-documents:
			if !ok {
				break loop
			}
			assert.NotNil(t, document.Id)
		}
	}
}

func TestIndividualDocumentPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	documents, cursor, err := IndividualDocument.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, document := range documents {
		assert.NotNil(t, document.Id)
	}

	assert.NotNil(t, cursor)
}
