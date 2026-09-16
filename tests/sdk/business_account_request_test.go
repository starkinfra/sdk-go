package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	BusinessAccountRequest "github.com/starkinfra/sdk-go/starkinfra/businessaccountrequest"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusinessAccountRequestPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests, err := BusinessAccountRequest.Create(Example.BusinessAccountRequest(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
		assert.Equal(t, "business", request.AccountType)
		for _, owner := range request.Owners {
			assert.NotNil(t, owner.TaxId)
		}
	}
}

func TestBusinessAccountRequestQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, errorChannel := BusinessAccountRequest.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case request, ok := <-requests:
			if !ok {
				break loop
			}
			assert.NotNil(t, request.Id)
		}
	}
}

func TestBusinessAccountRequestPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var ids = map[string]bool{}
	var cursor string

	for i := 0; i < 2; i++ {
		var params = map[string]interface{}{}
		params["limit"] = 2
		params["cursor"] = cursor

		requests, nextCursor, err := BusinessAccountRequest.Page(params, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}

		for _, request := range requests {
			assert.NotNil(t, request.Id)
			ids[request.Id] = true
		}

		cursor = nextCursor
		if cursor == "" {
			break
		}
	}

	assert.Equal(t, 4, len(ids))
}

func TestBusinessAccountRequestGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var requestList []BusinessAccountRequest.BusinessAccountRequest

	requests, errorChannel := BusinessAccountRequest.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case request, ok := <-requests:
			if !ok {
				break loop
			}
			requestList = append(requestList, request)
		}
	}

	for _, request := range requestList {
		getRequest, err := BusinessAccountRequest.Get(request.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getRequest.Id)
		for _, owner := range getRequest.Owners {
			assert.NotNil(t, owner.Status)
		}
	}
	assert.Equal(t, limit, len(requestList))
}
