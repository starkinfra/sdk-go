package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingStock "github.com/starkinfra/sdk-go/starkinfra/issuingstock"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingStockQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	stocks := IssuingStock.Query(nil, nil)
	for stock := range stocks {
		assert.NotNil(t, stock.Id)
		fmt.Println(stock.Id)
	}
}

func TestIssuingStockPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	stocks, cursor, err := IssuingStock.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, stock := range stocks {
		assert.NotNil(t, stock.Id)
		fmt.Println(stock.Id)
	}

	fmt.Println(cursor)
}

func TestIssuingStockInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var stockList []IssuingStock.IssuingStock
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	stocks := IssuingStock.Query(paramsQuery, nil)
	for stock := range stocks {
		stockList = append(stockList, stock)
	}

	stock, err := IssuingStock.Get(stockList[rand.Intn(len(stockList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, stock.Id)
	fmt.Println(stock.Id)
}
