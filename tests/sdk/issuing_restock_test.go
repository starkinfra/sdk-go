package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingRestock "github.com/starkinfra/sdk-go/starkinfra/issuingrestock"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingRestockPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	restocks, err := IssuingRestock.Create(Example.IssuingRestock(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, restock := range restocks {
		assert.NotNil(t, restock.Id)
	}
}

func TestIssuingRestockQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	restocks, errorChannel := IssuingRestock.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case restock, ok := <-restocks:
			if !ok {
				break loop
			}
			assert.NotNil(t, restock.Id)
		}
	}
}

func TestIssuingRestockPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	restocks, cursor, err := IssuingRestock.Page(nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, restock := range restocks {
		assert.NotNil(t, restock.Id)
	}
	assert.NotNil(t, cursor)
}

func TestIssuingRestockGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var restockList []IssuingRestock.IssuingRestock

	restocks, errorChannel := IssuingRestock.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case restock, ok := <-restocks:
			if !ok {
				break loop
			}
			restockList = append(restockList, restock)
		}
	}

	for _, restock := range restockList {
		getRestock, err := IssuingRestock.Get(restock.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getRestock.Id)
	}

	assert.Equal(t, limit, len(restockList))
}
