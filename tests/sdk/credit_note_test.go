package sdk

import (
	"encoding/json"
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

func TestCreditNoteRuleConstruct(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	rule := CreditNote.Rule{
		Key:   "invoiceCreationMode",
		Value: "scheduled",
	}

	assert.Equal(t, "invoiceCreationMode", rule.Key)
	assert.Equal(t, "scheduled", rule.Value)
}

func TestCreditNoteRulesConstruct(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	empty := CreditNote.CreditNote{}
	assert.Len(t, empty.Rules, 0)

	note := CreditNote.CreditNote{
		Rules: []CreditNote.Rule{
			{
				Key:   "invoiceCreationMode",
				Value: "scheduled",
			},
		},
	}
	assert.IsType(t, []CreditNote.Rule{}, note.Rules)
	assert.Len(t, note.Rules, 1)
	assert.Equal(t, "invoiceCreationMode", note.Rules[0].Key)
	assert.Equal(t, "scheduled", note.Rules[0].Value)

	marshaled, errMarshal := json.Marshal(note)
	assert.Nil(t, errMarshal)

	var roundTripped CreditNote.CreditNote
	errUnmarshal := json.Unmarshal(marshaled, &roundTripped)
	assert.Nil(t, errUnmarshal)

	assert.IsType(t, []CreditNote.Rule{}, roundTripped.Rules)
	assert.Len(t, roundTripped.Rules, 1)
	assert.Equal(t, "invoiceCreationMode", roundTripped.Rules[0].Key)
	assert.Equal(t, "scheduled", roundTripped.Rules[0].Value)
}

func TestCreditNotePostWithRules(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes, err := CreditNote.Create(Example.CreditNote(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
		assert.Len(t, note.Rules, 1)
		assert.Equal(t, "invoiceCreationMode", note.Rules[0].Key)
		assert.Equal(t, "scheduled", note.Rules[0].Value)
	}
}

func TestCreditNotePostWithInlineRule(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes := Example.CreditNote()
	notes[0].Rules = []CreditNote.Rule{
		{
			Key:   "invoiceCreationMode",
			Value: "scheduled",
		},
	}

	created, err := CreditNote.Create(notes, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, note := range created {
		assert.NotNil(t, note.Id)
		assert.Len(t, note.Rules, 1)
		assert.Equal(t, "invoiceCreationMode", note.Rules[0].Key)
		assert.Equal(t, "scheduled", note.Rules[0].Value)
	}
}

func TestCreditNoteDebtorWorkspaceId(t *testing.T) {

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
		assert.IsType(t, "", getNote.DebtorWorkspaceId)
	}
}
