package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualDocument "github.com/starkinfra/sdk-go/starkinfra/individualdocument"
	IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"math/rand"
	"testing"
)

func TestIndividualDocumentPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []IndividualIdentity.IndividualIdentity
	identities := IndividualIdentity.Query(nil, nil)
	for identity := range identities {
		assert.NotNil(t, identity.Id)
		identityList = append(identityList, identity)
	}

	bytesFront, _ := ioutil.ReadFile("../utils/identity/identity-front-face.png")
	frontDocuments, errFront := IndividualDocument.Create(Example.IndividualDocument(identityList[0].Id, "identity-front", bytesFront), nil)
	if errFront.Errors != nil {
		for _, e := range errFront.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, frontDocument := range frontDocuments {
		assert.NotNil(t, frontDocument.Id)
		fmt.Println(frontDocument.Id, frontDocument.Status)
	}

	bytesBack, _ := ioutil.ReadFile("../utils/identity/identity-back-face.png")
	backDocuments, errBack := IndividualDocument.Create(Example.IndividualDocument(identityList[0].Id, "identity-back", bytesBack), nil)
	if errBack.Errors != nil {
		for _, e := range errBack.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, backDocument := range backDocuments {
		assert.NotNil(t, backDocument.Id)
		fmt.Println(backDocument.Id, backDocument.Status)
	}

	bytesSelfie, _ := ioutil.ReadFile("../utils/identity/walter-white.png")
	selfieDocuments, errSelfie := IndividualDocument.Create(Example.IndividualDocument(identityList[0].Id, "selfie", bytesSelfie), nil)
	if errSelfie.Errors != nil {
		for _, e := range errSelfie.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, selfieDocument := range selfieDocuments {
		assert.NotNil(t, selfieDocument.Id)
		fmt.Println(selfieDocument.Id, selfieDocument.Status)
	}

	identity, err := IndividualIdentity.Update(identityList[0].Id, "processing", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.Equal(t, "processing", identity.Status)
}

func TestIndividualDocumentGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var documentList []IndividualDocument.IndividualDocument
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	documents := IndividualDocument.Query(paramsQuery, nil)
	for document := range documents {
		documentList = append(documentList, document)
	}

	document, err := IndividualDocument.Get(documentList[rand.Intn(len(documentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, document.Id)
	fmt.Println(document.Id)
}

func TestIndividualDocumentQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "created"

	documents := IndividualDocument.Query(params, nil)
	for document := range documents {
		assert.NotNil(t, document.Id)
	}
}

func TestIndividualDocumentPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	documents, cursor, err := IndividualDocument.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, document := range documents {
		assert.NotNil(t, document.Id)
		fmt.Println(document.Id)
	}

	fmt.Println(cursor)
}
