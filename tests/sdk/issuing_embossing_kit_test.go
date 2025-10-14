package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingEmbossingKit "github.com/starkinfra/sdk-go/starkinfra/issuingembossingkit"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingEmbossingKitQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	kits, errorChannel := IssuingEmbossingKit.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case kit, ok := <-kits:
			if !ok {
				break loop
			}
			assert.NotNil(t, kit.Id)
		}
	}
}

func TestIssuingEmbossingKitPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	kits, cursor, err := IssuingEmbossingKit.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, kit := range kits {
		assert.NotNil(t, kit.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingEmbossingKitGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var kitList []IssuingEmbossingKit.IssuingEmbossingKit

	kits, errorChannel := IssuingEmbossingKit.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case kit, ok := <-kits:
			if !ok {
				break loop
			}
			kitList = append(kitList, kit)
		}
	}

	for _, kit := range kitList {
		getKit, err := IssuingEmbossingKit.Get(kit.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getKit.Id)
	}
}
