package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiAgent "github.com/starkinfra/sdk-go/starkinfra/aiagent"
	AiChat "github.com/starkinfra/sdk-go/starkinfra/aichat"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

func collectChats(t *testing.T, params map[string]interface{}) []AiChat.AiChat {
	found := []AiChat.AiChat{}
	chats, errorChannel := AiChat.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		case chat, ok := <-chats:
			if !ok {
				return found
			}
			found = append(found, chat)
		}
	}
}

func TestAiChatLive(t *testing.T) {
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

	t.Run("test_go_create_returns_the_chat_with_tags_and_context", func(t *testing.T) {
		assert.NotEmpty(t, chat.Id)
		assert.Equal(t, agent.Id, chat.AgentId)
		assert.NotEmpty(t, chat.Title)
		assert.Equal(t, []string{"sdk-go", "test"}, chat.Tags)
		assert.Equal(t, "Ana", chat.Context["name"])
		assert.NotNil(t, chat.Updated)
	})

	t.Run("test_go_get_and_expand_agent_name", func(t *testing.T) {
		plain, err := AiChat.Get(chat.Id, nil, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, chat.Id, plain.Id)
		assert.Empty(t, plain.AgentName)
		expanded, err := AiChat.Get(chat.Id, map[string]interface{}{"expand": []string{"agentName"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, agent.Name, expanded.AgentName)
	})

	t.Run("test_go_query_expand_and_tags", func(t *testing.T) {
		found := map[string]AiChat.AiChat{}
		params := map[string]interface{}{"expand": []string{"agentName"}, "tags": []string{"sdk-go", "other"}}
		for _, entity := range collectChats(t, params) {
			found[entity.Id] = entity
		}
		assert.Equal(t, agent.Name, found[chat.Id].AgentName)
		none := collectChats(t, map[string]interface{}{"tags": []string{"no-chat-has-this-tag"}})
		assert.Empty(t, none)
	})

	t.Run("test_go_page_returns_the_entities_and_walks_to_the_end", func(t *testing.T) {
		seen := 0
		cursor := ""
		for {
			params := map[string]interface{}{"limit": 2}
			if cursor != "" {
				params["cursor"] = cursor
			}
			chats, next, err := AiChat.Page(params, nil)
			assert.Nil(t, err.Errors)
			assert.LessOrEqual(t, len(chats), 2)
			seen += len(chats)
			if next == "" {
				break
			}
			cursor = next
		}
		assert.GreaterOrEqual(t, seen, 1)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiChat.Page(map[string]interface{}{"limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})

	t.Run("test_go_update_changes_what_is_given_and_clears_with_empty_values", func(t *testing.T) {
		own, err := AiChat.Create(Example.ExampleAiChat(agent.Id), nil)
		assert.Nil(t, err.Errors)
		defer func() {
			_, err := AiChat.Delete([]string{own.Id}, nil)
			assert.Nil(t, err.Errors)
		}()
		updated, err := AiChat.Update(own.Id, map[string]interface{}{"title": "renamed-by-sdk"}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "renamed-by-sdk", updated.Title)
		assert.Equal(t, agent.Id, updated.AgentId)
		assert.Equal(t, own.Tags, updated.Tags)
		assert.Equal(t, own.Context, updated.Context)
		changed, err := AiChat.Update(
			own.Id,
			map[string]interface{}{"tags": []string{"vip"}, "context": map[string]interface{}{"order_id": "123"}},
			nil,
		)
		assert.Nil(t, err.Errors)
		assert.Equal(t, []string{"vip"}, changed.Tags)
		assert.Equal(t, map[string]interface{}{"order_id": "123"}, changed.Context)
		cleared, err := AiChat.Update(own.Id, map[string]interface{}{"tags": []string{}, "context": map[string]interface{}{}}, nil)
		assert.Nil(t, err.Errors)
		assert.Empty(t, cleared.Tags)
		assert.Empty(t, cleared.Context)
		assert.Equal(t, "renamed-by-sdk", cleared.Title)
	})

	t.Run("test_go_delete_returns_the_deleted_chats", func(t *testing.T) {
		own, err := AiChat.Create(Example.ExampleAiChat(agent.Id), nil)
		assert.Nil(t, err.Errors)
		deleted, err := AiChat.Delete([]string{own.Id}, nil)
		assert.Nil(t, err.Errors)
		if assert.Len(t, deleted, 1) {
			assert.Equal(t, own.Id, deleted[0].Id)
		}
	})

	t.Run("test_go_create_with_unknown_agent_returns_input_errors", func(t *testing.T) {
		_, err := AiChat.Create(AiChat.AiChat{AgentId: "0000000000000000"}, nil)
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})

	t.Run("test_go_get_unknown_id_returns_input_errors", func(t *testing.T) {
		_, err := AiChat.Get("0000000000000000", nil, nil)
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})
}
