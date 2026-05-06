package sdk

import (
	"testing"
	"time"
	"github.com/starkinfra/sdk-go/starkinfra"
	"github.com/starkinfra/sdk-go/starkinfra/event"
	PixPullRequest "github.com/starkinfra/sdk-go/starkinfra/pixpullrequest"
	PixPullSubscription "github.com/starkinfra/sdk-go/starkinfra/pixpullsubscription"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

func TestPixPullRequestPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	subscription, err := PixPullSubscription.Create(Example.PixPullSubscription(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	time.Sleep(10 * time.Second)

	requests, err := PixPullRequest.Create(Example.PixPullRequest(subscription[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
	}
}

func TestPixPullRequestQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, errorChannel := PixPullRequest.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case request, ok := <-requests:
			if !ok {
				break loop
			}
			assert.NotNil(t, request.Id)
		}
	}
}

func TestPixPullRequestPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	requests, cursor, err := PixPullRequest.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
	}

	assert.NotNil(t, cursor)
}

func TestPixPullRequestInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var requestList []PixPullRequest.PixPullRequest

	requests, errorChannel := PixPullRequest.Query(paramsQuery, nil)
	
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case request, ok := <-requests:
			if !ok {
				break loop
			}
			requestList = append(requestList, request)
		}
	}

	for _, request := range requestList {
		getRequest, err := PixPullRequest.Get(request.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getRequest.Id)
	}

	assert.Equal(t, limit, len(requestList))
}

func TestPixPullRequestPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	subscription, err := PixPullSubscription.Create(Example.PixPullSubscription(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	time.Sleep(10 * time.Second)

	requests, err := PixPullRequest.Create(Example.PixPullRequest(subscription[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	patchData := map[string]interface{}{
		"status": "scheduled",
	}

	updatedRequest, err := PixPullRequest.Update(requests[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, updatedRequest.Id)
}

func TestPixPullRequestCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	subscription, err := PixPullSubscription.Create(Example.PixPullSubscription(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	time.Sleep(10 * time.Second)

	requests, err := PixPullRequest.Create(Example.PixPullRequest(subscription[0].Id), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	canceledRequest, err := PixPullRequest.Cancel(requests[0].Id, "accountClosed", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, canceledRequest.Id)
}

func TestPixPullRequestParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"event\": {\"created\": \"2026-03-17T22:17:48.687366+00:00\", \"id\": \"5980132964564992\", \"log\": {\"created\": \"2026-03-17T22:17:44.741312+00:00\", \"description\": \"The Pix Pull Request was created in Stark Infra.\", \"errors\": [], \"id\": \"4777799707525120\", \"reason\": \"\", \"request\": {\"amount\": 79562, \"attemptType\": \"default\", \"created\": \"2026-03-17T22:17:44.727124+00:00\", \"description\": \"Monthly fare\", \"due\": \"2026-03-18T19:17:44.382949+00:00\", \"endToEndId\": \"E32160637202617031917FXbuEOeqxTE\", \"flow\": \"out\", \"id\": \"5859939668983808\", \"receiverAccountNumber\": \"00000000\", \"receiverAccountType\": \"payment\", \"receiverBankCode\": \"32160637\", \"receiverBranchCode\": \"\", \"receiverName\": \"Stark Bank\", \"receiverTaxId\": \"39.908.427/0001-28\", \"reconciliationId\": \"20260317191744.382994-03001917VKqeyyGMWvK\", \"senderBankCode\": null, \"senderFinalName\": \"STARK SCD S.A.\", \"senderFinalTaxId\": \"39.908.427/0001-28\", \"senderTaxId\": \"99.999.919/9999-79\", \"status\": \"created\", \"subscriptionBacenId\": \"RR321606372026170319175775651\", \"subscriptionId\": \"6366699370577920\", \"tags\": [], \"updated\": \"2026-03-17T22:17:45.560279+00:00\"}, \"type\": \"created\"}, \"subscription\": \"pix-pull-request\", \"workspaceId\": \"4828094443552768\"}}"
	validSignature := "MEUCIQDPci6mVcRQUqQazbol04cYvz8Ffuhh0birk4+8jSUH4AIgKlLhIH5zKzu+4jQlyabvSJin+8+5kJKiJpoqSQPCITg="

	parsed, err := event.Parse(content, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, parsed.Id)
}

func TestPixPullRequestParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"event\": {\"created\": \"2026-03-17T22:17:48.687366+00:00\", \"id\": \"5980132964564992\", \"log\": {\"created\": \"2026-03-17T22:17:44.741312+00:00\", \"description\": \"The Pix Pull Request was created in Stark Infra.\", \"errors\": [], \"id\": \"4777799707525120\", \"reason\": \"\", \"request\": {\"amount\": 79562, \"attemptType\": \"default\", \"created\": \"2026-03-17T22:17:44.727124+00:00\", \"description\": \"Monthly fare\", \"due\": \"2026-03-18T19:17:44.382949+00:00\", \"endToEndId\": \"E32160637202617031917FXbuEOeqxTE\", \"flow\": \"out\", \"id\": \"5859939668983808\", \"receiverAccountNumber\": \"00000000\", \"receiverAccountType\": \"payment\", \"receiverBankCode\": \"32160637\", \"receiverBranchCode\": \"\", \"receiverName\": \"Stark Bank\", \"receiverTaxId\": \"39.908.427/0001-28\", \"reconciliationId\": \"20260317191744.382994-03001917VKqeyyGMWvK\", \"senderBankCode\": null, \"senderFinalName\": \"STARK SCD S.A.\", \"senderFinalTaxId\": \"39.908.427/0001-28\", \"senderTaxId\": \"99.999.919/9999-79\", \"status\": \"created\", \"subscriptionBacenId\": \"RR321606372026170319175775651\", \"subscriptionId\": \"6366699370577920\", \"tags\": [], \"updated\": \"2026-03-17T22:17:45.560279+00:00\"}, \"type\": \"created\"}, \"subscription\": \"pix-pull-request\", \"workspaceId\": \"4828094443552768\"}}"
	invalidSignature := "MEUCIQDPci6mVcRQUqQazbol04cYvz8Ffuhh0bIrk4+8jSUH4AIgKlLhIH5zKzu+4jQlyabvSJin+8+5kJKiJpoqSQPCITg="

	_, err := event.Parse(content, invalidSignature, nil)
	
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}
