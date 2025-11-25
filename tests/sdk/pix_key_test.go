package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixKey "github.com/starkinfra/sdk-go/starkinfra/pixkey"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixKeyPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	key, err := PixKey.Create(Example.PixKey(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, key.Id)
}

func TestPixKeyQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	keys, errorChannel := PixKey.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case key, ok := <-keys:
			if !ok {
				break loop
			}
			assert.NotNil(t, key.Id)
		}
	}
}

func TestPixKeyPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	keys, cursor, err := PixKey.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, key := range keys {
		assert.NotNil(t, key.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixKeyInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var keyList []PixKey.PixKey

	keys, errorChannel := PixKey.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case key, ok := <-keys:
			if !ok {
				break loop
			}
			keyList = append(keyList, key)
		}
	}

	paramsGet := map[string]interface{}{}
	paramsGet["payerId"] = "422.791.690-96"

	for _, key := range keyList {
		getKey, err := PixKey.Get(key.Id, paramsGet, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getKey.Id)
	}

	assert.Equal(t, limit, len(keyList))
}

func TestPixKeyInfoDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var keyList []PixKey.PixKey

	keys, errorChannel := PixKey.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case key, ok := <-keys:
			if !ok {
				break loop
			}
			keyList = append(keyList, key)
		}
	}

	key, err := PixKey.Cancel(keyList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, key.Id)
}

func TestPixKeyInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var keyList []PixKey.PixKey

	keys, errorChannel := PixKey.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case key, ok := <-keys:
			if !ok {
				break loop
			}
			keyList = append(keyList, key)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["reason"] = "branchTransfer"
	patchData["accountType"] = "savings"

	key, err := PixKey.Update(keyList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, key.Id)
}
