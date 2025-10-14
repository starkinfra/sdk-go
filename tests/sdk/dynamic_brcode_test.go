package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	DynamicBrcode "github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDynamicBrcodePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	brcodes, err := DynamicBrcode.Create(Example.DynamicBrcode(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
	}
}

func TestDynamicBrcodeQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 20
	var params = map[string]interface{}{}
	params["limit"] = limit

	var brcodeList []DynamicBrcode.DynamicBrcode

	brcodes, errorChannel := DynamicBrcode.Query(params, nil)
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
	assert.Equal(t, limit, len(brcodeList))
}

func TestDynamicBrcodePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	brcodes, _, err := DynamicBrcode.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, brcode := range brcodes {
		assert.NotNil(t, brcode.Id)
	}
}

func TestDynamicBrcodeQueryAndGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var brcodeList []DynamicBrcode.DynamicBrcode

	brcodes, errorChannel := DynamicBrcode.Query(paramsQuery, nil)
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
		getBrcode, err := DynamicBrcode.Get(brcode.Uuid, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getBrcode.Id)
	}
	assert.Equal(t, limit, len(brcodeList))
}

func TestDynamicBrcodeParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	uuid := "21f174ab942843eb90837a5c3135dfd6"
	validSignature := "MEYCIQC+Ks0M54DPLEbHIi0JrMiWbBFMRETe/U2vy3gTiid3rAIhANMmOaxT03nx2bsdo+vg6EMhWGzdphh90uBH9PY2gJdd"

	_, err := DynamicBrcode.Verify(uuid, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
}

func TestDynamicBrcodeParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	uuid := "21f174ab942843eb90837a5c3135dfd6"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	_, err := DynamicBrcode.Verify(uuid, invalidSignature, nil)
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestDynamicBrcodeResponseDue(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var due = map[string]interface{}{}
	due["version"] = 1
	due["created"] = "2022-03-10T10:30:00+00:00"
	due["due"] = "2022-07-15"
	due["expiration"] = 1000000
	due["keyId"] = "+5511989898989"
	due["status"] = "paid"
	due["reconciliationId"] = "b77f5236-7ab9-4487-9f95-66ee6eaf1781"
	due["nominalAmount"] = 100
	due["senderName"] = "Anthony Edward Stark"
	due["senderTaxId"] = "012.345.678-90"
	due["receiverName"] = "Jamie Lannister"
	due["receiverTaxId"] = "20.018.183/0001-8"
	due["receiverStreetLine"] = "Av. Paulista, 200"
	due["receiverCity"] = "Sao Paulo"
	due["receiverStateCode"] = "SP"
	due["receiverZipCode"] = "01234-567"
	due["fine"] = 64
	due["interest"] = 0.8
	due["discounts"] = []map[string]interface{}{
		{
			"percentage": 5,
			"due":        "2022-03-10T10:30:00+00:00",
		},
		{
			"percentage": 1,
			"due":        "2022-03-10T10:30:00+00:00",
		},
	}
	due["description"] = "Response Due Golang Test"

	response := DynamicBrcode.ResponseDue(due)
	assert.NotNil(t, response)
}

func TestDynamicBrcodeResponseInstant(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var instant = map[string]interface{}{}
	instant["version"] = 1
	instant["created"] = "2022-07-15"
	instant["keyId"] = "+5511989898989"
	instant["status"] = "paid"
	instant["reconciliationId"] = "b77f5236-7ab9-4487-9f95-66ee6eaf1781"
	instant["amount"] = 100

	response := DynamicBrcode.ResponseInstant(instant)
	assert.NotNil(t, response)
}
