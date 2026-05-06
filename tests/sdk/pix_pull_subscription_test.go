package sdk

import (
	"testing"
	"github.com/starkinfra/sdk-go/starkinfra"
	"github.com/starkinfra/sdk-go/starkinfra/event"
	PixPullSubscription "github.com/starkinfra/sdk-go/starkinfra/pixpullsubscription"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

func TestPixPullSubscriptionPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	subscriptions, err := PixPullSubscription.Create(Example.PixPullSubscription(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, subscription := range subscriptions {
		assert.NotNil(t, subscription.Id)
	}
}

func TestPixPullSubscriptionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	subscriptions, errorChannel := PixPullSubscription.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case subscription, ok := <-subscriptions:
			if !ok {
				break loop
			}
			assert.NotNil(t, subscription.Id)
		}
	}
}

func TestPixPullSubscriptionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	subscriptions, cursor, err := PixPullSubscription.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, subscription := range subscriptions {
		assert.NotNil(t, subscription.Id)
	}

	assert.NotNil(t, cursor)
}

func TestPixPullSubscriptionInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var subscriptionList []PixPullSubscription.PixPullSubscription

	subscriptions, errorChannel := PixPullSubscription.Query(paramsQuery, nil)
	
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case subscription, ok := <-subscriptions:
			if !ok {
				break loop
			}
			subscriptionList = append(subscriptionList, subscription)
		}
	}

	for _, subscription := range subscriptionList {
		getSubscription, err := PixPullSubscription.Get(subscription.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getSubscription.Id)
	}

	assert.Equal(t, limit, len(subscriptionList))
}

func TestPixPullSubscriptionPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	subscriptions, err := PixPullSubscription.Create(Example.PixPullSubscription(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	patchData := map[string]interface{}{
		"status": "active",
	}

	updatedSubscription, err := PixPullSubscription.Update(subscriptions[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, updatedSubscription.Id)
}

func TestPixPullSubscriptionCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	subscriptions, err := PixPullSubscription.Create(Example.PixPullSubscription(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	canceledSubscription, err := PixPullSubscription.Cancel(subscriptions[0].Id, "accountClosed", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, canceledSubscription.Id)
}

func TestPixPullSubscriptionParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"event\": {\"created\": \"2026-03-17T20:24:02.006080+00:00\", \"id\": \"5739991880695808\", \"log\": {\"created\": \"2026-03-17T20:23:58.050406+00:00\", \"errors\": [], \"id\": \"5340798381981696\", \"reason\": \"\", \"subscription\": {\"amount\": 52064, \"amountMinLimit\": 0, \"bacenId\": \"RR321606372026170317231564231\", \"created\": \"2026-03-17T20:23:57.255567+00:00\", \"description\": \"A Lannister always pays his debts\", \"due\": \"2026-04-17T02:59:59.999000+00:00\", \"externalId\": \"606512134\", \"flow\": \"out\", \"id\": \"5656970050666496\", \"installmentEnd\": \"\", \"installmentStart\": \"2026-03-18T02:59:59.999999+00:00\", \"interval\": \"month\", \"pullRetryLimit\": 3, \"receiverBankCode\": \"32160637\", \"receiverName\": \"Stark Bank\", \"receiverTaxId\": \"39.908.427/0001-28\", \"referenceCode\": \"36135971\", \"senderAccountNumber\": \"55213\", \"senderBankCode\": null, \"senderBranchCode\": \"356\", \"senderCityCode\": \"\", \"senderFinalName\": \"STARK SCD S.A.\", \"senderFinalTaxId\": \"39.908.427/0001-28\", \"senderTaxId\": \"99.999.919/9999-79\", \"status\": \"created\", \"tags\": [], \"type\": \"push\", \"updated\": \"2026-03-17T20:23:58.050421+00:00\"}, \"type\": \"delivering\"}, \"subscription\": \"pix-pull-subscription\", \"workspaceId\": \"4828094443552768\"}}"
	validSignature := "MEUCIQCCZWR4+JYoDNENLnRbSCGGZf+atOaG4q8jWB3ADgc+DQIgIZ1LuXLZ06pke2qzaMNTlDLwcriuH+S3ve1aTQeqNK0="

	parsed, err := event.Parse(content, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, parsed.Id)
}

func TestPixPullSubscriptionParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"event\":{\"created\":\"2026-03-17T20:24:02.006080+00:00\",\"id\":\"5739991880695808\",\"log\":{\"created\":\"2026-03-17T20:23:58.050406+00:00\",\"errors\":[],\"id\":\"5340798381981696\",\"reason\":\"\",\"subscription\":{\"amount\":52064,\"amountMinLimit\":0,\"bacenId\":\"RR321606372026170317231564231\",\"created\":\"2026-03-17T20:23:57.255567+00:00\",\"description\":\"A Lannister always pays his debts\",\"due\":\"2026-04-17T02:59:59.999000+00:00\",\"externalId\":\"606512134\",\"flow\":\"out\",\"id\":\"5656970050666496\",\"installmentEnd\":\"\",\"installmentStart\":\"2026-03-18T02:59:59.999999+00:00\",\"interval\":\"month\",\"pullRetryLimit\":3,\"receiverBankCode\":\"32160637\",\"receiverName\":\"Stark Bank\",\"receiverTaxId\":\"39.908.427/0001-28\",\"referenceCode\":\"36135971\",\"senderAccountNumber\":\"55213\",\"senderBankCode\":null,\"senderBranchCode\":\"356\",\"senderCityCode\":\"\",\"senderFinalName\":\"STARK SCD S.A.\",\"senderFinalTaxId\":\"39.908.427/0001-28\",\"senderTaxId\":\"99.999.919/9999-79\",\"status\":\"created\",\"tags\":[],\"type\":\"push\",\"updated\":\"2026-03-17T20:23:58.050421+00:00\"},\"type\":\"delivering\"},\"subscription\":\"pix-pull-subscription\",\"workspaceId\":\"4828094443552768\"}}"
	invalidSignature := "MEUCIQCCZWR4+JYoDNENLnRbSCGGZf+atOaG4q8jWB3ADgc+DQIgIZ1LuXLZ06pke2qzaMNTlDLwcriuH+S3ve1aTQEqNK0="

	_, err := event.Parse(content, invalidSignature, nil)
	
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}
