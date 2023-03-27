package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixChargebackLog "github.com/starkinfra/sdk-go/starkinfra/pixchargeback/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixChargebackLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs := PixChargebackLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		fmt.Println(log.Id)
	}
}

func TestPixChargebackLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := PixChargebackLog.Page(params, nil)
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

func TestPixChargebackLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []PixChargebackLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := PixChargebackLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := PixChargebackLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	fmt.Println(log.Id)
}
