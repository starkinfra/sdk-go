package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	IssuingEmbossingRequest "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
	IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingEmbossingRequestPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var holderList []IssuingHolder.IssuingHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := IssuingHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	cards, errCreate := IssuingCard.Create(Example.IssuingCardEmbossing(holderList[0]), nil, nil)
	if errCreate.Errors != nil {
		for _, erro := range errCreate.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", erro.Code, erro.Message))
		}
	}

	requests, err := IssuingEmbossingRequest.Create(Example.IssuingEmbossingRequest(cards[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, request := range requests {
		fmt.Printf("%+v\n", request)
		assert.NotNil(t, request.Id)
		fmt.Println(request.Id)
	}
}

func TestIssuingEmbossingRequestQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	requests := IssuingEmbossingRequest.Query(params, nil)
	for request := range requests {
		assert.NotNil(t, request.Id)
		fmt.Println(request.Id)
	}
}

func TestIssuingEmbossingRequestPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests, cursor, err := IssuingEmbossingRequest.Page(nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
		fmt.Println(request.Id)
	}
	fmt.Println(cursor)
}

func TestIssuingEmbossingRequestGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var requestList []IssuingEmbossingRequest.IssuingEmbossingRequest
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	requests := IssuingEmbossingRequest.Query(paramsQuery, nil)
	for request := range requests {
		requestList = append(requestList, request)
	}

	request, err := IssuingEmbossingRequest.Get(requestList[rand.Intn(len(requestList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, request.Id)
	fmt.Println(request.Id)
}
