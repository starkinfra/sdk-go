package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixChargebackPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	chargebacks, err := PixChargeback.Create(Example.PixChargeback("E35547753202201201450oo8sDGca066"), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, chargeback := range chargebacks {
		assert.NotNil(t, chargeback.Id)
	}
}

func TestPixChargebackQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	chargebacks, errorChannel := PixChargeback.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case chargeback, ok := <-chargebacks:
			if !ok {
				break loop
			}
			assert.NotNil(t, chargeback.Id)
		}
	}
}

func TestPixChargebackPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 3
	var params = map[string]interface{}{}
	params["limit"] = limit

	chargebacks, cursor, err := PixChargeback.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, chargeback := range chargebacks {
		assert.NotNil(t, chargeback.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixChargebackGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var chargebackList []PixChargeback.PixChargeback

	chargebacks, errorChannel := PixChargeback.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case chargeback, ok := <-chargebacks:
			if !ok {
				break loop
			}
			chargebackList = append(chargebackList, chargeback)
		}
	}

	for _, chargeback := range chargebackList {
		getChargeback, err := PixChargeback.Get(chargeback.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getChargeback.Id)
	}

	assert.Equal(t, limit, len(chargebackList))
}

func TestPixChargebackInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var chargebackList []PixChargeback.PixChargeback

	chargebacks, errorChannel := PixChargeback.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case chargeback, ok := <-chargebacks:
			if !ok {
				break loop
			}
			chargebackList = append(chargebackList, chargeback)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["result"] = "rejected"
	patchData["rejectionReason"] = "noBalance"

	chargeback, err := PixChargeback.Update(chargebackList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, chargeback.Id)
	assert.Equal(t, chargeback.Result, "rejected")
}

func TestPixChargebackDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var chargebackList []PixChargeback.PixChargeback

	chargebacks, errorChannel := PixChargeback.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case chargeback, ok := <-chargebacks:
			if !ok {
				break loop
			}
			chargebackList = append(chargebackList, chargeback)
		}
	}

	chargeback, err := PixChargeback.Cancel(chargebackList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, chargeback.Id)
	assert.Equal(t, chargeback.Status, "canceled")
}
