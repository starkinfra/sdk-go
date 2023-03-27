package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingHolderPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	holders, err := IssuingHolder.Create(Example.IssuingHolder(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, holder := range holders {
		assert.NotNil(t, holder.Id)
		fmt.Printf("%+v", holder)
	}
}

func TestIssuingHolderQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var holderList []IssuingHolder.IssuingHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := IssuingHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	for holder := range holders {
		assert.NotNil(t, holder.Id)
		fmt.Printf("%+v", holder)
	}
}

func TestIssuingHolderPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10
	//params["expand"] = "rules"

	holders, cursor, err := IssuingHolder.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, holder := range holders {
		assert.NotNil(t, holder.Id)
		fmt.Printf("%+v\n", holder)
	}

	fmt.Println(cursor)
}

func TestIssuingHolderGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var holderList []IssuingHolder.IssuingHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := IssuingHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	holder, err := IssuingHolder.Get(holderList[rand.Intn(len(holderList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, holder.Id)
	fmt.Println(holder.Id)
}

func TestIssuingHolderUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var holderList []IssuingHolder.IssuingHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := IssuingHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	var patchData = map[string]interface{}{}
	patchData["name"] = "Tony Stark"

	holder, err := IssuingHolder.Update(holderList[rand.Intn(len(holderList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, holder.Id)
	fmt.Println(holder.Id)
}

func TestIssuingHolderCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var holderList []IssuingHolder.IssuingHolder
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	holders := IssuingHolder.Query(paramsQuery, nil)
	for holder := range holders {
		holderList = append(holderList, holder)
	}

	holder, err := IssuingHolder.Cancel(holderList[rand.Intn(len(holderList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, holder.Id)
	fmt.Println(holder.Id)
}
