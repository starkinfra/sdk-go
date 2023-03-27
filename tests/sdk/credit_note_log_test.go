package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditNoteLog "github.com/starkinfra/sdk-go/starkinfra/creditnote/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCreditNoteLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs := CreditNoteLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestCreditNoteLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := CreditNoteLog.Page(params, nil)
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

func TestCreditNoteLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var noteList []CreditNoteLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	notes := CreditNoteLog.Query(paramsQuery, nil)
	for note := range notes {
		noteList = append(noteList, note)
	}

	log, err := CreditNoteLog.Get(noteList[rand.Intn(len(noteList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
}
