package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditHolmesLog "github.com/starkinfra/sdk-go/starkinfra/creditholmes/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditHolmesLogLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 20

	logs := CreditHolmesLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
	}
}

func TestCreditHolmesLogLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := CreditHolmesLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}
	fmt.Println(cursor)
}

func TestCreditHolmesLogLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []CreditHolmesLog.Log
	var paramsQuery = map[string]interface{}{}

	logs := CreditHolmesLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := CreditHolmesLog.Get(logList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
}
