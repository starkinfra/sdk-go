package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditNotePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes, err := CreditNote.Create(Example.CreditNote(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditNoteQueryAndGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var noteList []CreditNote.CreditNote

	notes, errorChannel := CreditNote.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case note, ok := <-notes:
			if !ok {
				break loop
			}
			noteList = append(noteList, note)
		}
	}

	for _, note := range noteList {
		getNote, err := CreditNote.Get(note.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getNote.Id)
	}

	assert.Equal(t, limit, len(noteList))
}

func TestCreditNoteQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "created"

	notes, errorChannel := CreditNote.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case note, ok := <-notes:
			if !ok {
				break loop
			}
			assert.NotNil(t, note.Id)
		}
	}
}

func TestCreditNotePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	notes, _, err := CreditNote.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditNoteCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "created"
	
	var noteList []CreditNote.CreditNote

	notes, errorChannel := CreditNote.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case note, ok := <-notes:
			if !ok {
				break loop
			}
			noteList = append(noteList, note)
		}
	}

	note, err := CreditNote.Cancel(noteList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, note.Id)
}
