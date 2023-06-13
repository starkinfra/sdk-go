package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingEmbossingKit "github.com/starkinfra/sdk-go/starkinfra/issuingembossingkit"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingEmbossingKitQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	kits := IssuingEmbossingKit.Query(params, nil)
	for kit := range kits {
		assert.NotNil(t, kit.Id)
		fmt.Printf("%+v\n", kit)
	}
}

func TestIssuingEmbossingKitPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	kits, cursor, err := IssuingEmbossingKit.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, kit := range kits {
		assert.NotNil(t, kit.Id)
		fmt.Println(kit.Id)
	}

	fmt.Println(cursor)
}

func TestIssuingEmbossingKitGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var kitList []IssuingEmbossingKit.IssuingEmbossingKit
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	kits := IssuingEmbossingKit.Query(paramsQuery, nil)
	for kit := range kits {
		kitList = append(kitList, kit)
	}

	kit, err := IssuingEmbossingKit.Get(kitList[rand.Intn(len(kitList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, kit.Id)
	fmt.Println(kit.Id)
}
