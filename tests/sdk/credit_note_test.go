package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCreditNotePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes, err := CreditNote.Create(Example.CreditNote(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditNoteGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var noteList []CreditNote.CreditNote
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	notes := CreditNote.Query(paramsQuery, nil)
	for note := range notes {
		noteList = append(noteList, note)
	}

	note, err := CreditNote.Get(noteList[rand.Intn(len(noteList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, note.Id)
}

func TestCreditNoteQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "created"

	notes := CreditNote.Query(params, nil)
	for note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditNotePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	notes, cursor, err := CreditNote.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
	fmt.Println(cursor)
}

func TestCreditNoteCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var noteList []CreditNote.CreditNote
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	notes := CreditNote.Query(paramsQuery, nil)
	for note := range notes {
		noteList = append(noteList, note)
	}

	note, err := CreditNote.Cancel(noteList[rand.Intn(len(noteList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, note.Id)
}
