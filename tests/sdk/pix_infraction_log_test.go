package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixInfractionLog "github.com/starkinfra/sdk-go/starkinfra/pixinfraction/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixInfractionLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs := PixInfractionLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		fmt.Println(log.Id)
	}
}

func TestPixInfractionLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := PixInfractionLog.Page(params, nil)
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

func TestPixInfractionLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []PixInfractionLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := PixInfractionLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	log, err := PixInfractionLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
	fmt.Println(log.Id)
}
