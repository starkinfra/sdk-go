package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualIdentityLog "github.com/starkinfra/sdk-go/starkinfra/individualidentity/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIndividualIdentityLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	logs := IndividualIdentityLog.Query(paramsQuery, nil)
	for log := range logs {
		fmt.Println(log)
	}
}

func TestIndividualIdentityLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	logs, cursor, err := IndividualIdentityLog.Page(params, nil)
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

func TestIndividualIdentityLogInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var logsList []IndividualIdentityLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	logs := IndividualIdentityLog.Query(paramsQuery, nil)
	for log := range logs {
		logsList = append(logsList, log)
	}

	log, err := IndividualIdentityLog.Get(logsList[rand.Intn(len(logsList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, log.Id)
	fmt.Println(log.Id)
}
