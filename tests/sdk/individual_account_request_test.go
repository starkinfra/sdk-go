package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualAccountRequest "github.com/starkinfra/sdk-go/starkinfra/individualaccountrequest"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIndividualAccountRequestPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests, err := IndividualAccountRequest.Create(Example.IndividualAccountRequest(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
		assert.NotNil(t, request.Status)
		assert.Equal(t, "individual", request.AccountType)
	}
}

func TestIndividualAccountRequestPostAddressIsObject(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests := Example.IndividualAccountRequest()

	for _, request := range requests {
		assert.NotEmpty(t, request.Address.Street)
		assert.NotEmpty(t, request.Address.Number)
		assert.NotEmpty(t, request.Address.Neighborhood)
		assert.NotEmpty(t, request.Address.City)
		assert.NotEmpty(t, request.Address.State)
		assert.NotEmpty(t, request.Address.ZipCode)
	}

	created, err := IndividualAccountRequest.Create(requests, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range created {
		assert.NotNil(t, request.Id)
		assert.NotEmpty(t, request.Address.Street)
		assert.NotEmpty(t, request.Address.City)
	}
}

func TestIndividualAccountRequestQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, errorChannel := IndividualAccountRequest.Query(params, nil)
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

func TestIndividualAccountRequestPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, cursor, err := IndividualAccountRequest.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIndividualAccountRequestInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var requestList []IndividualAccountRequest.IndividualAccountRequest

	requests, errorChannel := IndividualAccountRequest.Query(paramsQuery, nil)
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
		getRequest, err := IndividualAccountRequest.Get(request.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getRequest.Id)
		assert.Contains(t, []string{"approved", "created", "denied", "processing", "updated"}, getRequest.Status)
		assert.NotNil(t, getRequest.Created)
	}

	assert.Equal(t, limit, len(requestList))
}

func TestIndividualAccountRequestUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var requestList []IndividualAccountRequest.IndividualAccountRequest

	requests, errorChannel := IndividualAccountRequest.Query(paramsQuery, nil)
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

	var patchData = map[string]interface{}{}
	patchData["name"] = "Tony Stark"

	updated, err := IndividualAccountRequest.Update(requestList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, updated.Id)
	assert.Equal(t, "Tony Stark", updated.Name)
}

func TestIndividualAccountRequestUpdateAddressObject(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var requestList []IndividualAccountRequest.IndividualAccountRequest

	requests, errorChannel := IndividualAccountRequest.Query(paramsQuery, nil)
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

	var patchData = map[string]interface{}{}
	patchData["address"] = map[string]interface{}{
		"street":       "Rua do Estilo Barroco",
		"number":       "648",
		"neighborhood": "Santo Amaro",
		"city":         "Sao Paulo",
		"state":        "SP",
		"zipCode":      "05724005",
	}

	updated, err := IndividualAccountRequest.Update(requestList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, updated.Id)
	assert.Equal(t, "Santo Amaro", updated.Address.Neighborhood)
}

func TestIndividualAccountRequestPostInvalidName(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests := Example.IndividualAccountRequest()
	requests[0].Name = ""

	_, err := IndividualAccountRequest.Create(requests, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountRequestPostInvalidTaxId(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests := Example.IndividualAccountRequest()
	requests[0].TaxId = "000.000.000-00"

	_, err := IndividualAccountRequest.Create(requests, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountRequestPostInvalidAddress(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests := Example.IndividualAccountRequest()
	requests[0].Address = IndividualAccountRequest.Address{
		Street: "Rua do Estilo Barroco",
	}

	_, err := IndividualAccountRequest.Create(requests, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountRequestPostInvalidIncome(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests := Example.IndividualAccountRequest()
	requests[0].Income = -1

	_, err := IndividualAccountRequest.Create(requests, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountRequestUpdateInvalidStatus(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var requestList []IndividualAccountRequest.IndividualAccountRequest

	requests, errorChannel := IndividualAccountRequest.Query(paramsQuery, nil)
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

	var patchData = map[string]interface{}{}
	patchData["status"] = "not-a-real-status"

	_, err := IndividualAccountRequest.Update(requestList[0].Id, patchData, nil)
	assert.NotNil(t, err.Errors)
}

func TestIndividualAccountRequestGetNotFound(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	_, err := IndividualAccountRequest.Get("0", nil)
	assert.NotNil(t, err.Errors)
}
