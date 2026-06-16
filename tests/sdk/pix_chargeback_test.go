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

	referenceId, ok := deriveSentReferenceId(t)
	if !ok {
		t.Skip("no successful PixRequest in sandbox to derive a reference_id; skipping create happy-path")
	}

	chargebacks, err := PixChargeback.Create(Example.PixChargeback(referenceId), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			if e.Code == "invalidDispute" || e.Code == "invalidReferenceId" || e.Code == "repeatedReferenceId" {
				t.Skipf("no fresh eligible transaction in sandbox (%s); skipping create happy-path", e.Code)
			}
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
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

	assert.LessOrEqual(t, len(chargebackList), limit)
}

func TestPixChargebackReturnOnlyFields(t *testing.T) {

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
		var isMonitoringRequired bool = chargeback.IsMonitoringRequired
		_ = isMonitoringRequired
		var disputeId string = chargeback.DisputeId
		var reversalAccountNumber string = chargeback.ReversalAccountNumber
		var reversalAccountType string = chargeback.ReversalAccountType
		var reversalBankCode string = chargeback.ReversalBankCode
		var reversalBranchCode string = chargeback.ReversalBranchCode
		var reversalTaxId string = chargeback.ReversalTaxId
		_ = disputeId
		_ = reversalAccountNumber
		_ = reversalAccountType
		_ = reversalBankCode
		_ = reversalBranchCode
		_ = reversalTaxId
	}
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

	if len(chargebackList) == 0 {
		t.Skip("no PixChargeback in sandbox; skipping patch happy-path")
	}

	var patchData = map[string]interface{}{}
	patchData["result"] = "rejected"
	patchData["rejectionReason"] = "noBalance"

	chargeback, err := PixChargeback.Update(chargebackList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			if e.Code == "invalidChargeback" {
				t.Skip("no patchable PixChargeback in sandbox (entity cannot be closed); skipping patch happy-path")
			}
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	assert.NotNil(t, chargeback.Id)
	assert.Equal(t, "rejected", chargeback.Result)
}

func TestPixChargebackCancel(t *testing.T) {

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

	if len(chargebackList) == 0 {
		t.Skip("no PixChargeback in sandbox; skipping cancel happy-path")
	}

	chargeback, err := PixChargeback.Cancel(chargebackList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			if e.Code == "invalidCancellationStatus" {
				t.Skip("no cancelable PixChargeback in sandbox (entity in non-cancelable status); skipping cancel happy-path")
			}
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	assert.NotNil(t, chargeback.Id)
	assert.Equal(t, "canceled", chargeback.Status)
}
