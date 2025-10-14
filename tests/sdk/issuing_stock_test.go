package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingStock "github.com/starkinfra/sdk-go/starkinfra/issuingstock"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingStockQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	stocks, errorChannel := IssuingStock.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case stock, ok := <-stocks:
			if !ok {
				break loop
			}
			assert.NotNil(t, stock.Id)
		}
	}
}

func TestIssuingStockPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 3
	var params = map[string]interface{}{}
	params["limit"] = limit

	stocks, cursor, err := IssuingStock.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, stock := range stocks {
		assert.NotNil(t, stock.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingStockInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 2
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var stockList []IssuingStock.IssuingStock

	stocks, errorChannel := IssuingStock.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case stock, ok := <-stocks:
			if !ok {
				break loop
			}
			stockList = append(stockList, stock)
		}
	}

	var expand = map[string]interface{}{}
	expand["expand"] = "balance"

	for _, stock := range stockList {
		getStock, err := IssuingStock.Get(stock.Id, expand, nil)
		if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
			assert.NotNil(t, getStock.Id)
	}

	assert.Equal(t, limit, len(stockList))
}
