package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixClaimLog "github.com/starkinfra/sdk-go/starkinfra/pixclaim/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixClaimLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs := PixClaimLog.Query(params, nil)
	for log := range logs {
		assert.NotNil(t, log.Id)
		fmt.Println(log.Id)
	}
}

func TestPixClaimLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, cursor, err := PixClaimLog.Page(params, nil)
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

func TestPixClaimLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logList []PixClaimLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := PixClaimLog.Query(paramsQuery, nil)
	for log := range logs {
		logList = append(logList, log)
	}

	claim, err := PixClaimLog.Get(logList[rand.Intn(len(logList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, claim.Id)
	fmt.Println(claim.Id)
}
