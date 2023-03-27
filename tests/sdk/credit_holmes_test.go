package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditHolmes "github.com/starkinfra/sdk-go/starkinfra/creditholmes"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestCreditHolmesPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes, err := CreditHolmes.Create(Example.CreditHolmes(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditHolmesGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var sherlockList []CreditHolmes.CreditHolmes
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	sherlocks := CreditHolmes.Query(paramsQuery, nil)
	for sherlock := range sherlocks {
		sherlockList = append(sherlockList, sherlock)
	}

	note, err := CreditHolmes.Get(sherlockList[rand.Intn(len(sherlockList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, note.Id)
}

func TestCreditHolmesQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes := CreditHolmes.Query(nil, nil)
	for note := range notes {
		assert.NotNil(t, note.Id)
	}
}

func TestCreditHolmesPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	notes, cursor, err := CreditHolmes.Page(params, nil)
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
