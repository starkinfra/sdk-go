package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixFraud "github.com/starkinfra/sdk-go/starkinfra/pixfraud"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixFraudPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	frauds, err := PixFraud.Create(Example.PixFraud(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, fraud := range frauds {
		assert.NotNil(t, fraud.Id)
	}
}

func TestPixFraudQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	frauds, errorChannel := PixFraud.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case fraud, ok := <-frauds:
			if !ok {
				break loop
			}
			assert.NotNil(t, fraud.Id)
		}
	}
}

func TestPixFraudPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	frauds, cursor, err := PixFraud.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, fraud := range frauds {
		assert.NotNil(t, fraud.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixFraudInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var fraudList []PixFraud.PixFraud

	frauds, errorChannel := PixFraud.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case fraud, ok := <-frauds:
			if !ok {
				break loop
			}
			fraudList = append(fraudList, fraud)
		}
	}

	for _, fraud := range fraudList {
		getFraud, err := PixFraud.Get(fraud.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getFraud.Id)
	}

	assert.Equal(t, limit, len(fraudList))
}
