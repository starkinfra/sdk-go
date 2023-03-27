package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixKeyPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	key, err := PixKey.Create(Example.PixKey(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	fmt.Println(key)
	assert.NotNil(t, key.Id)
	fmt.Println(key.Id)
}

func TestPixKeyQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 2

	keys := PixKey.Query(params, nil)
	for key := range keys {
		assert.NotNil(t, key.Id)
		fmt.Println(key.Id)
	}
}

func TestPixKeyPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 100

	keys, cursor, err := PixKey.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, key := range keys {
		assert.NotNil(t, key.Id)
		fmt.Println(key.Id)
	}
	fmt.Println(cursor)
}

func TestPixKeyInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var keyList []PixKey.PixKey
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	keys := PixKey.Query(paramsQuery, nil)
	for key := range keys {
		keyList = append(keyList, key)
	}

	key, err := PixKey.Get(keyList[rand.Intn(len(keyList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, key.Id)
	fmt.Println(key.Id)
}

func TestPixKeyInfoDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var keyList []PixKey.PixKey
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	keys := PixKey.Query(paramsQuery, nil)
	for key := range keys {
		keyList = append(keyList, key)
	}

	key, err := PixKey.Cancel(keyList[rand.Intn(len(keyList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, key.Id)
	fmt.Println(key.Id)
}

func TestPixKeyInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var keyList []PixKey.PixKey
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	keys := PixKey.Query(paramsQuery, nil)
	for key := range keys {
		keyList = append(keyList, key)
	}

	var patchData = map[string]interface{}{}
	patchData["reason"] = "branchTransfer"
	patchData["accountType"] = "savings"

	key, err := PixKey.Update(keyList[rand.Intn(len(keyList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, key.Id)
	fmt.Println(key.Id)
}
