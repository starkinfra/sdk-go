package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingRestock "github.com/starkinfra/sdk-go/starkinfra/issuingrestock"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingRestockPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	restocks, err := IssuingRestock.Create(Example.IssuingRestock(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, restock := range restocks {
		assert.NotNil(t, restock.Id)
		fmt.Println(restock.Id)
	}
}

func TestIssuingRestockQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	restocks := IssuingRestock.Query(params, nil)
	for restock := range restocks {
		assert.NotNil(t, restock.Id)
		fmt.Println(restock.Id)
	}
}

func TestIssuingRestockPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	restocks, cursor, err := IssuingRestock.Page(nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, restock := range restocks {
		assert.NotNil(t, restock.Id)
		fmt.Println(restock.Id)
	}
	fmt.Println(cursor)
}

func TestIssuingRestockGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var restockList []IssuingRestock.IssuingRestock
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	restocks := IssuingRestock.Query(paramsQuery, nil)
	for restock := range restocks {
		restockList = append(restockList, restock)
	}

	restock, err := IssuingRestock.Get(restockList[rand.Intn(len(restockList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, restock.Id)
	fmt.Println(restock.Id)
}
