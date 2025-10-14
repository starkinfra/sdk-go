package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	Webhook "github.com/starkinfra/sdk-go/starkinfra/webhook"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWebhookPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	webhook, err := Webhook.Create(Example.Webhook(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, webhook.Id)
}

func TestWebhookGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var webhookList []Webhook.Webhook

	webhooks, errorChannel := Webhook.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case webhook, ok := <-webhooks:
			if !ok {
				break loop
			}
			webhookList = append(webhookList, webhook)
		}
	}

	for _, webhook := range webhookList {
		getWebhook, err := Webhook.Get(webhook.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getWebhook.Id)
	}

	assert.Equal(t, limit, len(webhookList))
}

func TestWebhookQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	webhooks, errorChannel := Webhook.Query(nil, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case webhook, ok := <-webhooks:
			if !ok {
				break loop
			}
			assert.NotNil(t, webhook.Id)
		}
	}
}

func TestWebhookPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 3
	var params = map[string]interface{}{}
	params["limit"] = limit

	webhooks, cursor, err := Webhook.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, webhook := range webhooks {
		assert.NotNil(t, webhook.Id)
	}
	assert.NotNil(t, cursor)
}

func TestWebhookDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var webhookList []Webhook.Webhook

	webhooks, errorChannel := Webhook.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case webhook, ok := <-webhooks:
			if !ok {
				break loop
			}
			webhookList = append(webhookList, webhook)
		}
	}

	webhook, err := Webhook.Delete(webhookList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, webhook.Id)
}
