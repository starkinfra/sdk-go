package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingHolderPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	holders, err := IssuingHolder.Create(Example.IssuingHolder(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, holder := range holders {
		assert.NotNil(t, holder.Id)
	}
}

func TestIssuingHolderQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	holders, errorChannel := IssuingHolder.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case holder, ok := <-holders:
			if !ok {
				break loop
			}
			assert.NotNil(t, holder.Id)
		}
	}
}

func TestIssuingHolderPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	//params["expand"] = "rules"

	holders, cursor, err := IssuingHolder.Page(paramsQuery, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, holder := range holders {
		assert.NotNil(t, holder.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingHolderGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var holderList []IssuingHolder.IssuingHolder

	holders, errorChannel := IssuingHolder.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case holder, ok := <-holders:
			if !ok {
				break loop
			}
			assert.NotNil(t, holder.Id)
			holderList = append(holderList, holder)
		}
	}

	for _, holder := range holderList {
		getHolder, err := IssuingHolder.Get(holder.Id, nil, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getHolder.Id)
	}

	assert.Equal(t, limit, len(holderList))
}

func TestIssuingHolderUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var holderList []IssuingHolder.IssuingHolder

	holders, errorChannel := IssuingHolder.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case holder, ok := <-holders:
			if !ok {
				break loop
			}
			holderList = append(holderList, holder)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["name"] = "Tony Stark"

	holder, err := IssuingHolder.Update(holderList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, holder.Id)
	assert.Equal(t, "Tony Stark", holder.Name)
}

func TestIssuingHolderCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var holderList []IssuingHolder.IssuingHolder

	holders, errorChannel := IssuingHolder.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case holder, ok := <-holders:
			if !ok {
				break loop
			}
			holderList = append(holderList, holder)
		}
	}

	holder, err := IssuingHolder.Cancel(holderList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, holder.Id)
	assert.Equal(t, "canceled", holder.Status)
}
