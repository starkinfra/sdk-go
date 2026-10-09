package sdk

import (
	"strings"
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiChat "github.com/starkinfra/sdk-go/starkinfra/aichat"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

const chatAnswer = `{"id":"5632499082330112","agentId":"5740688905863168","title":"Greeting","tags":["whatsapp"],"context":{"name":"Ana"},"updated":"2026-10-01T14:28:02.652375+00:00"}`

func collectChats(t *testing.T, params map[string]interface{}) []AiChat.AiChat {
	found := []AiChat.AiChat{}
	chats, errorChannel := AiChat.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
		case chat, ok := <-chats:
			if !ok {
				return found
			}
			found = append(found, chat)
		}
	}
}

func TestAiChatLive(t *testing.T) {
	fixtures.load(t)
	chat := fixtures.chat
	agent := fixtures.agent

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
		assertInputError(t, err)
	})

	t.Run("test_go_get_unknown_id_returns_input_errors", func(t *testing.T) {
		_, err := AiChat.Get("0000000000000000", nil, nil)
		assertInputError(t, err)
	})
}

func TestAiChatAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_sends_only_the_creatable_fields_with_tags_and_context_as_written", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chat":`+chatAnswer+`}`)
		returned := AiChat.AiChat{
			AgentId:   "5740688905863168",
			Title:     "Greeting",
			Tags:      []string{"whatsapp"},
			Context:   map[string]interface{}{"customer_name": "Ana & Bia <vip>", "city": "São Paulo", "nested": map[string]interface{}{"is_vip": true}},
			Id:        "5632499082330112",
			AgentName: "Support assistant",
		}
		created, err := AiChat.Create(returned, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "POST", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-chat"), request.Url)
		assert.Equal(t, []string{"agentId", "context", "tags", "title"}, bodyKeys(t, request.Body))
		sent := bodyOf(t, request.Body)
		assert.Equal(t, []interface{}{"whatsapp"}, sent["tags"])
		assert.Equal(t, map[string]interface{}{
			"customer_name": "Ana & Bia <vip>",
			"city":          "São Paulo",
			"nested":        map[string]interface{}{"is_vip": true},
		}, sent["context"])
		assert.Equal(t, "5632499082330112", created.Id)
		assert.Equal(t, []string{"whatsapp"}, created.Tags)
		assert.Equal(t, map[string]interface{}{"name": "Ana"}, created.Context)
		assert.NotNil(t, created.Updated)
	})

	t.Run("test_go_create_without_optionals_sends_only_the_agent", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chat":`+chatAnswer+`}`)
		_, err := AiChat.Create(AiChat.AiChat{AgentId: "5740688905863168"}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, []string{"agentId"}, bodyKeys(t, (*recorded)[0].Body))
	})

	t.Run("test_go_update_names_all_four_keys_and_sends_null_for_the_absent", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chat":`+chatAnswer+`}`)
		_, err := AiChat.Update("5632499082330112", map[string]interface{}{"title": "Greeting", "id": "ignored"}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "PATCH", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-chat/5632499082330112"), request.Url)
		assert.Equal(t, map[string]interface{}{"title": "Greeting", "agentId": nil, "tags": nil, "context": nil}, bodyOf(t, request.Body))
	})

	t.Run("test_go_update_without_fields_sends_four_nulls", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chat":`+chatAnswer+`}`)
		_, err := AiChat.Update("5632499082330112", nil, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, map[string]interface{}{"title": nil, "agentId": nil, "tags": nil, "context": nil}, bodyOf(t, (*recorded)[0].Body))
	})

	t.Run("test_go_update_sends_empty_values_and_context_keys_as_written", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chat":`+chatAnswer+`}`, `{"chat":`+chatAnswer+`}`)
		_, err := AiChat.Update("5632499082330112", map[string]interface{}{"title": "", "tags": []string{}, "context": map[string]interface{}{}}, nil)
		assert.Nil(t, err.Errors)
		sent := bodyOf(t, (*recorded)[0].Body)
		assert.Equal(t, "", sent["title"])
		assert.Equal(t, []interface{}{}, sent["tags"])
		assert.Equal(t, map[string]interface{}{}, sent["context"])
		_, err = AiChat.Update("5632499082330112", map[string]interface{}{"context": map[string]interface{}{"order_id": "1", "isVip": true}}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, map[string]interface{}{"order_id": "1", "isVip": true}, bodyOf(t, (*recorded)[1].Body)["context"])
	})

	t.Run("test_go_get_sends_expand_in_the_query_string", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chat":`+chatAnswer+`}`)
		chat, err := AiChat.Get("5632499082330112", map[string]interface{}{"expand": []string{"agentName"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "5740688905863168", chat.AgentId)
		assert.True(t, strings.HasSuffix((*recorded)[0].Url, "/v2/ai-chat/5632499082330112?expand=agentName"), (*recorded)[0].Url)
	})

	t.Run("test_go_page_sends_tags_comma_separated_and_returns_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("chats", "next-page", []string{chatAnswer}), pageAnswer("chats", "", []string{chatAnswer}))
		first, cursor, err := AiChat.Page(map[string]interface{}{"tags": []string{"a", "b"}, "limit": 1, "expand": []string{"agentName"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		assert.Equal(t, "next-page", cursor)
		assert.Equal(t, []string{"whatsapp"}, first[0].Tags)
		url := (*recorded)[0].Url
		assert.Contains(t, url, "tags=a%2Cb")
		assert.Contains(t, url, "limit=1")
		assert.Contains(t, url, "expand=agentName")
		_, cursor, err = AiChat.Page(map[string]interface{}{"cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_follows_the_cursor_and_sends_tags", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("chats", "next-page", []string{chatAnswer}), pageAnswer("chats", "", []string{chatAnswer}))
		found := collectChats(t, map[string]interface{}{"tags": []string{"a", "b"}})
		assert.Len(t, found, 2)
		assert.Contains(t, (*recorded)[0].Url, "tags=a%2Cb")
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_delete_sends_ids_in_the_query_string_and_returns_the_deleted_objects", func(t *testing.T) {
		recorded := answerInOrder(t, `{"chats":[`+chatAnswer+`]}`)
		deleted, err := AiChat.Delete([]string{"5632499082330112"}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "DELETE", request.Method)
		assert.Equal(t, "", request.Body)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-chat?ids=5632499082330112"), request.Url)
		assert.Equal(t, "5632499082330112", deleted[0].Id)
	})
}
