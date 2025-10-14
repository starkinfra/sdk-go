package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditHolmes "github.com/starkinfra/sdk-go/starkinfra/creditholmes"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditHolmesPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes, err := CreditHolmes.Create(Example.CreditHolmes(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditHolmesGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 5
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var sherlockList []CreditHolmes.CreditHolmes

	sherlocks, errorChannel := CreditHolmes.Query(paramsQuery, nil)

	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case sherlock, ok := <-sherlocks:
			if !ok {
				break loop
			}
			sherlockList = append(sherlockList, sherlock)
		}
	}

	for _, sherlock := range sherlockList {
		assert.NotNil(t, sherlock.Id)
	}

	assert.Equal(t, limit, len(sherlockList))
}

func TestCreditHolmesQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 5
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	notes, errorChannel := CreditHolmes.Query(paramsQuery, nil)

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

func TestCreditHolmesPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	notes, _, err := CreditHolmes.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
}
