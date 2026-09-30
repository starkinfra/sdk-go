package sdk

import (
	"strings"
	"sync"
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiMessage "github.com/starkinfra/sdk-go/starkinfra/aimessage"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

const userMessageAnswer = `{"id":"5642368648740864","chatId":"5632499082330112","sender":"user","text":"Say hello.","speech":"Say hello.","metadata":{},"model":"bender-1.0","created":"2026-10-01T14:28:02.652375+00:00"}`
const systemMessageAnswer = `{"id":"5079418695319552","chatId":"5632499082330112","sender":"system","text":"Hello!","speech":"Hello!","metadata":{"order_id":"123"},"model":"bender-1.0","created":"2026-10-01T14:28:02.653375+00:00"}`

var postedOnce sync.Once
var posted []AiMessage.AiMessage
var postedFailure string

func postedMessages(t *testing.T) []AiMessage.AiMessage {
	fixtures.load(t)
	postedOnce.Do(func() {
		messages, err := AiMessage.Create(Example.ExampleAiMessage(fixtures.chat.Id), []string{"chatName"}, nil)
		if err.Errors != nil {
			postedFailure = err.Errors[0].Code + ": " + err.Errors[0].Message
			return
		}
		posted = messages
	})
	if postedFailure != "" {
		t.Fatalf("could not post the shared message: %s", postedFailure)
	}
	return posted
}

func collectMessages(t *testing.T, params map[string]interface{}) []AiMessage.AiMessage {
	found := []AiMessage.AiMessage{}
	messages, errorChannel := AiMessage.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
		case message, ok := <-messages:
			if !ok {
				return found
			}
			found = append(found, message)
		}
	}
}

func messageIds(messages []AiMessage.AiMessage) []string {
	ids := []string{}
	for _, message := range messages {
		ids = append(ids, message.Id)
	}
	return ids
}

func TestAiMessageLive(t *testing.T) {
	chat := func(t *testing.T) string {
		postedMessages(t)
		return fixtures.chat.Id
	}

	t.Run("test_go_create_returns_the_user_message_and_the_answer", func(t *testing.T) {
		messages := postedMessages(t)
		if !assert.Len(t, messages, 2) {
			return
		}
		assert.Equal(t, "user", messages[0].Sender)
		assert.Equal(t, "system", messages[1].Sender)
		for _, message := range messages {
			assert.Equal(t, fixtures.chat.Id, message.ChatId)
			assert.NotNil(t, message.Created)
			assert.NotEmpty(t, message.ChatName)
		}
	})

	t.Run("test_go_query_returns_the_whole_history", func(t *testing.T) {
		found := collectMessages(t, map[string]interface{}{"chatId": chat(t)})
		assert.ElementsMatch(t, messageIds(posted), messageIds(found))
	})

	t.Run("test_go_query_with_limit_stops_at_the_limit", func(t *testing.T) {
		found := collectMessages(t, map[string]interface{}{"chatId": chat(t), "limit": 1})
		assert.Len(t, found, 1)
	})

	t.Run("test_go_page_returns_a_cursor_that_leads_to_the_next_page", func(t *testing.T) {
		chatId := chat(t)
		first, cursor, err := AiMessage.Page(map[string]interface{}{"chatId": chatId, "limit": 1}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		if !assert.NotEmpty(t, cursor) {
			return
		}
		second, _, err := AiMessage.Page(map[string]interface{}{"chatId": chatId, "limit": 1, "cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, second, 1)
		assert.NotEqual(t, first[0].Id, second[0].Id)
	})

	t.Run("test_go_create_in_an_unknown_chat_returns_input_errors", func(t *testing.T) {
		starkinfra.User = utils.ExampleProject
		_, err := AiMessage.Create(AiMessage.AiMessage{ChatId: "0000000000000000", Text: "hi"}, nil, nil)
		assertInputError(t, err)
	})

	t.Run("test_go_query_and_page_without_chat_id_read_the_workspace_history", func(t *testing.T) {
		chat(t)
		starkinfra.User = utils.ExampleProject
		found := collectMessages(t, map[string]interface{}{"limit": 3})
		assert.NotEmpty(t, found)
		page, _, err := AiMessage.Page(map[string]interface{}{"limit": 3}, nil)
		assert.Nil(t, err.Errors)
		assert.NotEmpty(t, page)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		starkinfra.User = utils.ExampleProject
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiMessage.Page(map[string]interface{}{"chatId": fixtures.chat.Id, "limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})
}

func TestAiMessageAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_sends_expand_in_the_query_string_and_not_in_the_body", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chatName":"Greeting","messages":[`+userMessageAnswer+`,`+systemMessageAnswer+`]}`)
		messages, err := AiMessage.Create(
			AiMessage.AiMessage{ChatId: "5632499082330112", Text: "Say hello.", Model: "prime-1.0", Id: "ignored", Sender: "user", ChatName: "ignored"},
			[]string{"chatName"},
			nil,
		)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "POST", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-message?expand=chatName"), request.Url)
		assert.Equal(t, []string{"chatId", "model", "text"}, bodyKeys(t, request.Body))
		if assert.Len(t, messages, 2) {
			assert.Equal(t, "user", messages[0].Sender)
			assert.Equal(t, "system", messages[1].Sender)
			assert.Equal(t, "Greeting", messages[0].ChatName)
			assert.Equal(t, "Greeting", messages[1].ChatName)
			assert.Equal(t, map[string]interface{}{"order_id": "123"}, messages[1].Metadata)
		}
	})

	t.Run("test_go_create_without_expand_has_no_query_string_and_no_chat_name", func(t *testing.T) {
		recorded := answerInOrder(t, `{"messages":[`+userMessageAnswer+`,`+systemMessageAnswer+`]}`)
		messages, err := AiMessage.Create(AiMessage.AiMessage{ChatId: "5632499082330112", Text: "Say hello."}, nil, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-message"), request.Url)
		assert.Equal(t, []string{"chatId", "text"}, bodyKeys(t, request.Body))
		assert.Empty(t, messages[0].ChatName)
	})

	t.Run("test_go_query_follows_the_cursor_until_it_runs_out", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			`{"cursor":"next-page","messages":[`+userMessageAnswer+`]}`,
			`{"cursor":null,"messages":[`+systemMessageAnswer+`]}`,
		)
		found := collectMessages(t, map[string]interface{}{"chatId": "5632499082330112"})
		assert.Equal(t, []string{"5642368648740864", "5079418695319552"}, messageIds(found))
		if assert.Len(t, *recorded, 2) {
			assert.Contains(t, (*recorded)[0].Url, "/v2/ai-message?")
			assert.Contains(t, (*recorded)[0].Url, "chatId=5632499082330112")
			assert.NotContains(t, (*recorded)[0].Url, "cursor")
			assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
		}
	})

	t.Run("test_go_query_stops_at_the_limit_without_asking_for_another_page", func(t *testing.T) {
		recorded := answerInOrder(t, `{"cursor":"next-page","messages":[`+userMessageAnswer+`]}`)
		found := collectMessages(t, map[string]interface{}{"chatId": "5632499082330112", "limit": 1})
		assert.Len(t, found, 1)
		assert.Len(t, *recorded, 1)
		assert.Contains(t, (*recorded)[0].Url, "limit=1")
	})

	t.Run("test_go_page_returns_the_cursor_and_an_empty_one_on_the_last_page", func(t *testing.T) {
		answerInOrder(
			t,
			`{"cursor":"next-page","messages":[`+userMessageAnswer+`]}`,
			`{"cursor":null,"messages":[`+systemMessageAnswer+`]}`,
		)
		first, cursor, err := AiMessage.Page(map[string]interface{}{"chatId": "5632499082330112", "limit": 1}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "next-page", cursor)
		assert.Len(t, first, 1)
		_, cursor, err = AiMessage.Page(map[string]interface{}{"chatId": "5632499082330112", "cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
	})

	t.Run("test_go_page_without_chat_id_sends_no_chat_id", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("messages", "", []string{userMessageAnswer}))
		messages, cursor, err := AiMessage.Page(map[string]interface{}{"limit": 5}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, messages, 1)
		assert.Equal(t, "", cursor)
		assert.Len(t, *recorded, 1)
		assert.NotContains(t, (*recorded)[0].Url, "chatId")
		assert.Contains(t, (*recorded)[0].Url, "limit=5")
	})

	t.Run("test_go_query_without_params_sends_no_chat_id", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("messages", "", []string{userMessageAnswer}))
		found := collectMessages(t, nil)
		assert.Len(t, found, 1)
		assert.NotContains(t, (*recorded)[0].Url, "chatId")
	})

	t.Run("test_go_query_with_limit_150_asks_100_then_50", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			pageAnswer("messages", "next-page", repeated(userMessageAnswer, 100)),
			pageAnswer("messages", "another-page", repeated(userMessageAnswer, 50)),
		)
		found := collectMessages(t, map[string]interface{}{"limit": 150})
		assert.Len(t, found, 150)
		if assert.Len(t, *recorded, 2) {
			assert.Contains(t, (*recorded)[0].Url, "limit=100")
			assert.Contains(t, (*recorded)[1].Url, "limit=50")
		}
	})
}
