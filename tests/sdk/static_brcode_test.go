package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	StaticBrcode "github.com/starkinfra/sdk-go/starkinfra/staticbrcode"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStaticBrcodePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	brcodes, err := StaticBrcode.Create(Example.StaticBrcode(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
	}
}

func TestStaticBrcodeQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	brcodes, errorChannel := StaticBrcode.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			assert.NotNil(t, brcode.Id)
		}
	}
}

func TestStaticBrcodePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 3
	var params = map[string]interface{}{}
	params["limit"] = limit

	brcodes, cursor, err := StaticBrcode.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
	}
	assert.NotNil(t, cursor)
}

func TestStaticBrcodeInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var brcodeList []StaticBrcode.StaticBrcode

	brcodes, errorChannel := StaticBrcode.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			brcodeList = append(brcodeList, brcode)
		}
	}

	for _, brcode := range brcodeList {
		getBrcode, err := StaticBrcode.Get(brcode.Uuid, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getBrcode.Id)
	}

	assert.Equal(t, limit, len(brcodeList))
}
