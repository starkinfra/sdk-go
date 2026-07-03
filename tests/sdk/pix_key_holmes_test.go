package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixKeyHolmes "github.com/starkinfra/sdk-go/starkinfra/pixkeyholmes"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixKeyHolmesPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	holmes, err := PixKeyHolmes.Create(Example.PixKeyHolmes(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, sherlock := range holmes {
		assert.NotNil(t, sherlock.Id)
		assert.NotEmpty(t, sherlock.Status)
		assert.NotNil(t, sherlock.Created)
	}
}

func TestPixKeyHolmesQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = []string{"solved", "solving"}
	params["tags"] = []string{"war", "stark"}

	holmes, errorChannel := PixKeyHolmes.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case sherlock, ok := <-holmes:
			if !ok {
				break loop
			}
			assert.NotNil(t, sherlock.Id)
		}
	}
}

func TestPixKeyHolmesPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	holmes, cursor, err := PixKeyHolmes.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, sherlock := range holmes {
		assert.NotNil(t, sherlock.Id)
	}

	assert.NotNil(t, cursor)
}
