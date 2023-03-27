package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	Log "github.com/starkinfra/sdk-go/starkinfra/issuingstock/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingStockLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	logs := Log.Query(nil, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		fmt.Println(log.Id)
	}
}

func TestIssuingStockLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := Log.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
		fmt.Println(log.Id)
	}

	fmt.Println(cursor)
}

func TestIssuingStockLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []Log.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := Log.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := Log.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
	fmt.Println(log.Id)
}
