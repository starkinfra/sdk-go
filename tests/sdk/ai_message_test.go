package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiAgent "github.com/starkinfra/sdk-go/starkinfra/aiagent"
	AiChat "github.com/starkinfra/sdk-go/starkinfra/aichat"
	AiMessage "github.com/starkinfra/sdk-go/starkinfra/aimessage"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

const systemMessageAnswer = `{"id":"5079418695319552","chatId":"5632499082330112","sender":"system","text":"Hello!","speech":"Hello!","metadata":{"order_id":"123"},"model":"bender-1.0","created":"2026-10-01T14:28:02.653375+00:00"}`

func collectMessages(t *testing.T, params map[string]interface{}) []AiMessage.AiMessage {
	found := []AiMessage.AiMessage{}
	messages, errorChannel := AiMessage.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
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
	starkinfra.User = utils.ExampleProject
	agent, errs := AiAgent.Create(Example.ExampleAiAgent(nil), nil)
	if errs.Errors != nil {
		t.Fatalf("code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
	}
	defer AiAgent.Delete([]string{agent.Id}, nil)
	chat, errs := AiChat.Create(Example.ExampleAiChat(agent.Id), nil)
	if errs.Errors != nil {
		t.Fatalf("code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
	}
	defer AiChat.Delete([]string{chat.Id}, nil)
	posted, errs := AiMessage.Create(Example.ExampleAiMessage(chat.Id), []string{"chatName"}, nil)
	if errs.Errors != nil {
		t.Fatalf("code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
	}

	t.Run("test_go_create_returns_the_user_message_and_the_answer", func(t *testing.T) {
		messages := posted
		if !assert.Len(t, messages, 2) {
			return
		}
		assert.Equal(t, "user", messages[0].Sender)
		assert.Equal(t, "system", messages[1].Sender)
		for _, message := range messages {
			assert.Equal(t, chat.Id, message.ChatId)
			assert.NotNil(t, message.Created)
			assert.NotEmpty(t, message.ChatName)
		}
	})

	t.Run("test_go_query_returns_the_whole_history", func(t *testing.T) {
		found := collectMessages(t, map[string]interface{}{"chatId": chat.Id})
		assert.ElementsMatch(t, messageIds(posted), messageIds(found))
	})

	t.Run("test_go_query_with_limit_stops_at_the_limit", func(t *testing.T) {
		found := collectMessages(t, map[string]interface{}{"chatId": chat.Id, "limit": 1})
		assert.Len(t, found, 1)
	})

	t.Run("test_go_page_returns_a_cursor_that_leads_to_the_next_page", func(t *testing.T) {
		chatId := chat.Id
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
		_, err := AiMessage.Create(AiMessage.AiMessage{ChatId: "0000000000000000", Text: "hi"}, nil, nil)
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})

	t.Run("test_go_query_and_page_without_chat_id_read_the_workspace_history", func(t *testing.T) {
		found := collectMessages(t, map[string]interface{}{"limit": 3})
		assert.NotEmpty(t, found)
		page, _, err := AiMessage.Page(map[string]interface{}{"limit": 3}, nil)
		assert.Nil(t, err.Errors)
		assert.NotEmpty(t, page)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiMessage.Page(map[string]interface{}{"chatId": chat.Id, "limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})
}
