package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualDocumentLog "github.com/starkinfra/sdk-go/starkinfra/individualdocument/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIndividualDocumentLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	logs := IndividualDocumentLog.Query(nil, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		fmt.Println(log.Id)
	}
}

func TestIndividualDocumentLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := IndividualDocumentLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}

	fmt.Println(cursor)
}

func TestIndividualDocumentLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var documentList []IndividualDocumentLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	documents := IndividualDocumentLog.Query(paramsQuery, nil)
	for document := range documents {
		documentList = append(documentList, document)
	}

	log, err := IndividualDocumentLog.Get(documentList[rand.Intn(len(documentList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
	fmt.Println(log.Id)
}
