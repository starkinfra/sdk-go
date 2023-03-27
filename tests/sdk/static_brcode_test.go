package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	StaticBrcode "github.com/starkinfra/sdk-go/starkinfra/staticbrcode"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestStaticBrcodePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	brcodes, err := StaticBrcode.Create(Example.StaticBrcode(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
		fmt.Println(brcode.Id)
	}
}

func TestStaticBrcodeQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	brcodes := StaticBrcode.Query(params, nil)
	for brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
	}
}

func TestStaticBrcodePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	brcodes, cursor, err := StaticBrcode.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
		fmt.Println(brcode.Uuid)
		fmt.Println(brcode.Id)
	}
	fmt.Println(cursor)
}

func TestStaticBrcodeInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var brcodeList []StaticBrcode.StaticBrcode
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	brcodes := StaticBrcode.Query(paramsQuery, nil)
	for brcode := range brcodes {
		brcodeList = append(brcodeList, brcode)
	}

	brcode, err := StaticBrcode.Get(brcodeList[rand.Intn(len(brcodeList))].Uuid, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, brcode.Id)
	fmt.Println(brcode.Id)
}
