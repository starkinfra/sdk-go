package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	Webhook "github.com/starkinfra/sdk-go/starkinfra/webhook"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestWebhookPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	webhook, err := Webhook.Create(Example.Webhook(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, webhook.Id)
	fmt.Println(webhook.Id)
}

func TestWebhookGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var webhookList []Webhook.Webhook
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	webhooks := Webhook.Query(paramsQuery, nil)
	for webhook := range webhooks {
		webhookList = append(webhookList, webhook)
	}

	webhook, err := Webhook.Get(webhookList[rand.Intn(len(webhookList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, webhook.Id)
	fmt.Println(webhook.Id)
}

func TestWebhookQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	webhooks := Webhook.Query(nil, nil)
	for webhook := range webhooks {
		assert.NotNil(t, webhook.Id)
		fmt.Println(webhook.Id)
	}
}

func TestWebhookPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	webhooks, cursor, err := Webhook.Page(nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, webhook := range webhooks {
		assert.NotNil(t, webhook.Id)
		fmt.Println(webhook.Id)
	}
	fmt.Println(cursor)
}

func TestWebhookDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var webhookList []Webhook.Webhook
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	webhooks := Webhook.Query(paramsQuery, nil)
	for webhook := range webhooks {
		webhookList = append(webhookList, webhook)
	}

	webhook, err := Webhook.Delete(webhookList[rand.Intn(len(webhookList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, webhook.Id)
	fmt.Println(webhook.Id)
}
