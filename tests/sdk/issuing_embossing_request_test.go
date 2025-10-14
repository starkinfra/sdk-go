package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	IssuingEmbossingRequest "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
	IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingEmbossingRequestPost(t *testing.T) {

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

	cards, errCreate := IssuingCard.Create(Example.IssuingCardEmbossing(holderList[0]), nil, nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			t.Errorf("code: %s, message: %s", erro.Code, erro.Message)
		}
	}

	requests, err := IssuingEmbossingRequest.Create(Example.IssuingEmbossingRequest(cards[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
	}
}

func TestIssuingEmbossingRequestQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, errorChannel := IssuingEmbossingRequest.Query(params, nil)
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

func TestIssuingEmbossingRequestPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests, _, err := IssuingEmbossingRequest.Page(nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
	}
}

func TestIssuingEmbossingRequestGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var requestList []IssuingEmbossingRequest.IssuingEmbossingRequest

	requests, errorChannel := IssuingEmbossingRequest.Query(paramsQuery, nil)
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
		getRequest, err := IssuingEmbossingRequest.Get(request.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getRequest.Id)
	}

	assert.Equal(t, limit, len(requestList))
}
