package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixChargeback "github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixChargebackPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	chargebacks, err := PixChargeback.Create(Example.PixChargeback("E35547753202201201450oo8sDGca066"), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, chargeback := range chargebacks {
		assert.NotNil(t, chargeback.Id)
		fmt.Println(chargeback.Id)
	}
}

func TestPixChargebackQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	chargebacks := PixChargeback.Query(params, nil)
	for chargeback := range chargebacks {
		assert.NotNil(t, chargeback.Id)
		fmt.Println(chargeback.Id)
	}
}

func TestPixChargebackPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	chargebacks, cursor, err := PixChargeback.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, chargeback := range chargebacks {
		assert.NotNil(t, chargeback.Id)
		fmt.Println(chargeback.Id)
	}
	fmt.Println(cursor)
}

func TestPixChargebackGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var chargebackList []PixChargeback.PixChargeback
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	chargebacks := PixChargeback.Query(paramsQuery, nil)
	for chargeback := range chargebacks {
		chargebackList = append(chargebackList, chargeback)
	}

	chargeback, err := PixChargeback.Get(chargebackList[rand.Intn(len(chargebackList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, chargeback.Id)
	fmt.Println(chargeback.Id)
}

func TestPixChargebackInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var chargebackList []PixChargeback.PixChargeback
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	chargebacks := PixChargeback.Query(paramsQuery, nil)
	for chargeback := range chargebacks {
		chargebackList = append(chargebackList, chargeback)
	}

	var patchData = map[string]interface{}{}
	patchData["result"] = "rejected"
	patchData["rejectionReason"] = "noBalance"

	chargeback, err := PixChargeback.Update(chargebackList[rand.Intn(len(chargebackList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, chargeback.Id)
	fmt.Println(chargeback.Id)
}

func TestPixChargebackDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var chargebackList []PixChargeback.PixChargeback
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	chargebacks := PixChargeback.Query(paramsQuery, nil)
	for chargeback := range chargebacks {
		chargebackList = append(chargebackList, chargeback)
	}

	chargeback, err := PixChargeback.Cancel(chargebackList[rand.Intn(len(chargebackList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, chargeback.Id)
	fmt.Println(chargeback.Id)
}
