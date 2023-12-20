package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixFraud "github.com/starkinfra/sdk-go/starkinfra/pixfraud"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixFraudPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	frauds, err := PixFraud.Create(Example.PixFraud(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, fraud := range frauds {
		assert.NotNil(t, fraud.Id)
		fmt.Println(fraud.Id)
	}
}

func TestPixFraudQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	frauds := PixFraud.Query(params, nil)
	for fraud := range frauds {
		assert.NotNil(t, fraud.Id)
		fmt.Println(fraud.Id)
	}
}

func TestPixFraudPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	frauds, cursor, err := PixFraud.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, fraud := range frauds {
		assert.NotNil(t, fraud.Id)
		fmt.Println(fraud.Id)
	}
	fmt.Println(cursor)
}

func TestPixFraudInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var fraudList []PixFraud.PixFraud
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	frauds := PixFraud.Query(paramsQuery, nil)
	for fraud := range frauds {
		fraudList = append(fraudList, fraud)
	}

	fraud, err := PixFraud.Get(fraudList[rand.Intn(len(fraudList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, fraud.Id)
	fmt.Println(fraud.Id)
}
