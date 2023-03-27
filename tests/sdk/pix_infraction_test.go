package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
	PixInfractionLog "github.com/starkinfra/sdk-go/starkinfra/pixinfraction/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixInfractionPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	infractions, err := PixInfraction.Create(Example.PixInfraction(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, infraction := range infractions {
		assert.NotNil(t, infraction.Id)
		fmt.Println(infraction.Id)
	}
}

func TestPixInfractionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	infractions := PixInfraction.Query(params, nil)
	for infraction := range infractions {
		assert.NotNil(t, infraction.Id)
		fmt.Println(infraction.Id)
	}
}

func TestPixInfractionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	infractions, cursor, err := PixInfraction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, infraction := range infractions {
		assert.NotNil(t, infraction.Id)
		fmt.Println(infraction.Id)
	}
	fmt.Println(cursor)
}

func TestPixInfractionInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var infractionList []PixInfraction.PixInfraction
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	infractions := PixInfraction.Query(paramsQuery, nil)
	for infraction := range infractions {
		infractionList = append(infractionList, infraction)
	}

	infraction, err := PixInfraction.Get(infractionList[rand.Intn(len(infractionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, infraction.Id)
	fmt.Println(infraction.Id)
}

func TestPixInfractionInfoDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var infractionList []PixInfractionLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	infractions := PixInfractionLog.Query(paramsQuery, nil)
	for infraction := range infractions {
		infractionList = append(infractionList, infraction)
	}

	infraction, err := PixInfraction.Cancel(infractionList[rand.Intn(len(infractionList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, infraction.Id)
	fmt.Println(infraction.Id)
}

func TestPixInfractionInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var infractionList []PixInfractionLog.Log
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	infractions := PixInfractionLog.Query(paramsQuery, nil)
	for infraction := range infractions {
		infractionList = append(infractionList, infraction)
	}

	var patchData = map[string]interface{}{}
	patchData["result"] = "agreed"
	patchData["analysis"] = "Upon investigation fraud was confirmed."

	infraction, err := PixInfraction.Update(infractionList[rand.Intn(len(infractionList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, infraction.Id)
	fmt.Println(infraction.Id)
}
