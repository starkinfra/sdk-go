package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingPurchaseLog "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingPurchaseLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs := IssuingPurchaseLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestIssuingPurchaseLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := IssuingPurchaseLog.Page(params, nil)
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

func TestIssuingPurchaseLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []IssuingPurchaseLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := IssuingPurchaseLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := IssuingPurchaseLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
	fmt.Println(log.Id)
}
