package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	Event "github.com/starkinfra/sdk-go/starkinfra/event"
	Attempt "github.com/starkinfra/sdk-go/starkinfra/event/attempt"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestEventGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var eventList []Event.Event
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	events := Event.Query(paramsQuery, nil)
	for event := range events {
		eventList = append(eventList, event)
	}

	event, err := Event.Get(eventList[rand.Intn(len(eventList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, event.Id)
}

func TestEventQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	events := Event.Query(params, nil)
	for event := range events {
		assert.NotNil(t, event.Id)
	}
}

func TestEventPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 2
	params["isDelivered"] = false

	events, cursor, err := Event.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, event := range events {
		assert.NotNil(t, event.Id)
		fmt.Println(event.Id)
	}

	fmt.Println(cursor)
}

func TestEventDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var eventList []Event.Event
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	events := Event.Query(paramsQuery, nil)
	for event := range events {
		eventList = append(eventList, event)
	}

	event, err := Event.Delete(eventList[rand.Intn(len(eventList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, event.Id)
	fmt.Println(event.IsDelivered)
}

func TestEventUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["isDelivered"] = false

	var eventList []Event.Event
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	events := Event.Query(paramsQuery, nil)
	for event := range events {
		eventList = append(eventList, event)
	}

	event, err := Event.Update(eventList[rand.Intn(len(eventList))].Id, true, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, event.Id)
	fmt.Println(event.Id)
}

func TestEventAttemptQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 100

	attempts := Attempt.Query(params, nil)
	for attempt := range attempts {
		assert.NotNil(t, attempt.Id)
		fmt.Println(attempt.Id)
	}
}

func TestEventAttemptGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var attemptList []Attempt.Attempt
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	attempts := Attempt.Query(paramsQuery, nil)
	for attempt := range attempts {
		attemptList = append(attemptList, attempt)
	}

	attempt, err := Attempt.Get(attemptList[rand.Intn(len(attemptList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, attempt.Id)
	fmt.Println(attempt.Id)
}

func TestEventAttemptPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 2

	attempts, cursor, err := Attempt.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, attempt := range attempts {
		assert.NotNil(t, attempt.Id)
		fmt.Println(attempt.Id)
	}

	fmt.Println(cursor)
}

func TestEventParse(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"event\": {\"created\": \"2022-02-15T20:45:09.852878+00:00\", \"id\": \"5015597159022592\", \"log\": {\"created\": \"2022-02-15T20:45:09.436621+00:00\", \"errors\": [{\"code\": \"insufficientFunds\", \"message\": \"Amount of funds available is not sufficient to cover the specified transfer\"}], \"id\": \"5288053467774976\", \"request\": {\"amount\": 1000, \"bankCode\": \"34052649\", \"cashAmount\": 0, \"cashierBankCode\": \"\", \"cashierType\": \"\", \"created\": \"2022-02-15T20:45:08.210009+00:00\", \"description\": \"For saving my life\", \"endToEndId\": \"E34052649202201272111u34srod1a91\", \"externalId\": \"141322efdgber1ecd1s342341321\", \"fee\": 0, \"flow\": \"out\", \"id\": \"5137269514043392\", \"initiatorTaxId\": \"\", \"method\": \"manual\", \"receiverAccountNumber\": \"000001\", \"receiverAccountType\": \"checking\", \"receiverBankCode\": \"00000001\", \"receiverBranchCode\": \"0001\", \"receiverKeyId\": \"\", \"receiverName\": \"Jamie Lennister\", \"receiverTaxId\": \"45.987.245/0001-92\", \"reconciliationId\": \"\", \"senderAccountNumber\": \"000000\", \"senderAccountType\": \"checking\", \"senderBankCode\": \"34052649\", \"senderBranchCode\": \"0000\", \"senderName\": \"tyrion Lennister\", \"senderTaxId\": \"012.345.678-90\", \"status\": \"failed\", \"tags\": [], \"updated\": \"2022-02-15T20:45:09.436661+00:00\"}, \"type\": \"failed\"}, \"subscription\": \"pix-request.out\", \"workspaceId\": \"5692908409716736\"}}"
	signature := "\"MEYCIQD0oFxFQX0fI6B7oqjwLhkRhkDjrOiD86wguEKWdzkJbgIhAPNGUUdlNpYBe+npOaHa9WJopzy3WJYl8XJG6f4ek2R/\""
	parsed := Event.Parse(content, signature, nil)
	fmt.Println(parsed)
}
