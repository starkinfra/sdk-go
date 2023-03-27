package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	Log "github.com/starkinfra/sdk-go/starkinfra/issuingrestock/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingRestockLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	restocks := Log.Query(params, nil)
	for restock := range restocks {
		assert.NotNil(t, restock.Id)
		fmt.Println(restock.Id)
	}
}

func TestIssuingRestockLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	restocks, cursor, err := Log.Page(nil, nil)
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

func TestIssuingRestockLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []Log.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := Log.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	restock, err := Log.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, restock.Id)
	fmt.Println(restock.Id)
}
