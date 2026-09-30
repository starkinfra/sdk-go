package sdk

import (
	"strings"
	"testing"
	"time"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiAgent "github.com/starkinfra/sdk-go/starkinfra/aiagent"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

const agentAnswer = `{"id":"5740688905863168","name":"Support assistant","model":"bender-1.0","systemPrompt":"Answer in one short sentence.","voiceId":"","knowledgeBaseIds":["5083538508480512"],"metadataSchema":{"order_id":{"type":"string"}},"created":"2026-09-30T15:42:56.879325+00:00","updated":"2026-09-30T15:42:56.879334+00:00"}`

func collectAgents(t *testing.T, params map[string]interface{}) []AiAgent.AiAgent {
	found := []AiAgent.AiAgent{}
	agents, errorChannel := AiAgent.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
		case agent, ok := <-agents:
			if !ok {
				return found
			}
			found = append(found, agent)
		}
	}
}

func TestAiAgentLive(t *testing.T) {
	fixtures.load(t)
	agent := fixtures.agent
	knowledgeBaseId := fixtures.knowledgeBase.Id

	t.Run("test_go_create_returns_the_agent_with_the_schema_keys_as_written", func(t *testing.T) {
		assert.NotEmpty(t, agent.Id)
		assert.Equal(t, "bender-1.0", agent.Model)
		assert.Equal(t, []string{knowledgeBaseId}, agent.KnowledgeBaseIds)
		assert.Contains(t, agent.MetadataSchema, "order_id")
		assert.Len(t, agent.MetadataSchema, 1)
		assert.NotNil(t, agent.Created)
	})

	t.Run("test_go_get_and_expand_knowledge_bases", func(t *testing.T) {
		plain, err := AiAgent.Get(agent.Id, nil, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, agent.Id, plain.Id)
		assert.Nil(t, plain.KnowledgeBases)
		expanded, err := AiAgent.Get(agent.Id, map[string]interface{}{"expand": []string{"knowledgeBases"}}, nil)
		assert.Nil(t, err.Errors)
		if assert.Len(t, expanded.KnowledgeBases, 1) {
			assert.Equal(t, knowledgeBaseId, expanded.KnowledgeBases[0].Id)
		}
	})

	t.Run("test_go_query_and_expand", func(t *testing.T) {
		found := map[string]AiAgent.AiAgent{}
		for _, entity := range collectAgents(t, map[string]interface{}{"expand": []string{"knowledgeBases"}}) {
			found[entity.Id] = entity
		}
		assert.Contains(t, found, agent.Id)
		assert.Len(t, found[agent.Id].KnowledgeBases, 1)
	})

	t.Run("test_go_query_with_limit", func(t *testing.T) {
		assert.Len(t, collectAgents(t, map[string]interface{}{"limit": 1}), 1)
	})

	t.Run("test_go_page_returns_the_entities_and_walks_to_the_end", func(t *testing.T) {
		seen := 0
		cursor := ""
		for {
			params := map[string]interface{}{"limit": 2}
			if cursor != "" {
				params["cursor"] = cursor
			}
			agents, next, err := AiAgent.Page(params, nil)
			assert.Nil(t, err.Errors)
			assert.LessOrEqual(t, len(agents), 2)
			seen += len(agents)
			if next == "" {
				break
			}
			cursor = next
		}
		assert.GreaterOrEqual(t, seen, 1)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiAgent.Page(map[string]interface{}{"limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})

	t.Run("test_go_update_keeps_what_it_was_not_given_and_clears_with_empty_values", func(t *testing.T) {
		own, err := AiAgent.Create(Example.ExampleAiAgent([]string{knowledgeBaseId}), nil)
		assert.Nil(t, err.Errors)
		defer func() {
			_, err := AiAgent.Delete([]string{own.Id}, nil)
			assert.Nil(t, err.Errors)
		}()
		renamed, err := AiAgent.Update(own.Id, map[string]interface{}{"name": "renamed-by-sdk"}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "renamed-by-sdk", renamed.Name)
		assert.Equal(t, []string{knowledgeBaseId}, renamed.KnowledgeBaseIds)
		assert.Equal(t, own.SystemPrompt, renamed.SystemPrompt)
		assert.Contains(t, renamed.MetadataSchema, "order_id")
		cleared, err := AiAgent.Update(
			own.Id,
			map[string]interface{}{
				"systemPrompt":     "",
				"knowledgeBaseIds": []string{},
				"metadataSchema":   map[string]interface{}{},
			},
			nil,
		)
		assert.Nil(t, err.Errors)
		assert.Empty(t, cleared.SystemPrompt)
		assert.Empty(t, cleared.KnowledgeBaseIds)
		assert.Empty(t, cleared.MetadataSchema)
		assert.Equal(t, "renamed-by-sdk", cleared.Name)
	})

	t.Run("test_go_delete_returns_the_deleted_agents", func(t *testing.T) {
		own, err := AiAgent.Create(Example.ExampleAiAgent(nil), nil)
		assert.Nil(t, err.Errors)
		deleted, err := AiAgent.Delete([]string{own.Id}, nil)
		assert.Nil(t, err.Errors)
		if assert.Len(t, deleted, 1) {
			assert.Equal(t, own.Id, deleted[0].Id)
		}
	})

	t.Run("test_go_create_with_invalid_model_returns_input_errors", func(t *testing.T) {
		_, err := AiAgent.Create(AiAgent.AiAgent{Name: "invalid", Model: "gpt"}, nil)
		assertInputError(t, err)
	})

	t.Run("test_go_get_unknown_id_returns_input_errors", func(t *testing.T) {
		_, err := AiAgent.Get("0000000000000000", nil, nil)
		assertInputError(t, err)
	})
}

func TestAiAgentAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_sends_only_the_creatable_fields_and_does_not_touch_the_schema_keys", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		moment := time.Date(2026, 9, 30, 0, 0, 0, 0, time.UTC)
		returned := AiAgent.AiAgent{
			Name:             "Support assistant",
			Model:            "bender-1.0",
			SystemPrompt:     "Be brief.",
			VoiceId:          "5632499082330112",
			KnowledgeBaseIds: []string{"5083538508480512"},
			MetadataSchema: map[string]interface{}{
				"order_id": map[string]interface{}{"type": "string"},
				"isUrgent": map[string]interface{}{"type": "boolean"},
			},
			Id:      "5740688905863168",
			Created: &moment,
			Updated: &moment,
		}
		created, err := AiAgent.Create(returned, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "POST", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-agent"), request.Url)
		assert.Equal(t, []string{"knowledgeBaseIds", "metadataSchema", "model", "name", "systemPrompt", "voiceId"}, bodyKeys(t, request.Body))
		sent := bodyOf(t, request.Body)
		assert.Equal(t, map[string]interface{}{
			"order_id": map[string]interface{}{"type": "string"},
			"isUrgent": map[string]interface{}{"type": "boolean"},
		}, sent["metadataSchema"])
		assert.Equal(t, "5740688905863168", created.Id)
		assert.Equal(t, []string{"5083538508480512"}, created.KnowledgeBaseIds)
		assert.Contains(t, created.MetadataSchema, "order_id")
	})

	t.Run("test_go_create_keeps_an_empty_list_and_drops_empty_strings", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		_, err := AiAgent.Create(AiAgent.AiAgent{Name: "Support assistant", Model: "bender-1.0", KnowledgeBaseIds: []string{}}, nil)
		assert.Nil(t, err.Errors)
		sent := bodyOf(t, (*recorded)[0].Body)
		assert.Equal(t, []interface{}{}, sent["knowledgeBaseIds"])
		assert.Equal(t, []string{"knowledgeBaseIds", "model", "name"}, bodyKeys(t, (*recorded)[0].Body))
	})

	t.Run("test_go_update_names_all_six_keys_sends_null_for_the_absent_and_does_not_read_the_agent", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		_, err := AiAgent.Update("5740688905863168", map[string]interface{}{"name": "Renamed"}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, *recorded, 1)
		request := (*recorded)[0]
		assert.Equal(t, "PATCH", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-agent/5740688905863168"), request.Url)
		assert.Equal(t, map[string]interface{}{
			"name":             "Renamed",
			"model":            nil,
			"systemPrompt":     nil,
			"voiceId":          nil,
			"knowledgeBaseIds": nil,
			"metadataSchema":   nil,
		}, bodyOf(t, request.Body))
	})

	t.Run("test_go_update_sends_empty_values_so_they_clear_the_field", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		patchData := map[string]interface{}{
			"systemPrompt":     "",
			"voiceId":          "",
			"knowledgeBaseIds": []string{},
			"metadataSchema":   map[string]interface{}{},
		}
		_, err := AiAgent.Update("5740688905863168", patchData, nil)
		assert.Nil(t, err.Errors)
		sent := bodyOf(t, (*recorded)[0].Body)
		assert.Equal(t, "", sent["systemPrompt"])
		assert.Equal(t, "", sent["voiceId"])
		assert.Equal(t, []interface{}{}, sent["knowledgeBaseIds"])
		assert.Equal(t, map[string]interface{}{}, sent["metadataSchema"])
		assert.Nil(t, sent["name"])
		assert.Contains(t, sent, "name")
	})

	t.Run("test_go_update_treats_a_nil_list_and_a_nil_value_as_absent", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		var none []string
		_, err := AiAgent.Update("5740688905863168", map[string]interface{}{"knowledgeBaseIds": none, "voiceId": nil}, nil)
		assert.Nil(t, err.Errors)
		sent := bodyOf(t, (*recorded)[0].Body)
		assert.Nil(t, sent["knowledgeBaseIds"])
		assert.Nil(t, sent["voiceId"])
		assert.Len(t, sent, 6)
	})

	t.Run("test_go_update_sends_the_schema_keys_as_written", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		schema := map[string]interface{}{"order_id": map[string]interface{}{"type": "string"}, "isUrgent": map[string]interface{}{"type": "boolean"}}
		_, err := AiAgent.Update("5740688905863168", map[string]interface{}{"metadataSchema": schema}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, schema, bodyOf(t, (*recorded)[0].Body)["metadataSchema"])
	})

	t.Run("test_go_get_sends_expand_in_the_query_string", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agent":`+agentAnswer+`}`)
		_, err := AiAgent.Get("5740688905863168", map[string]interface{}{"expand": []string{"knowledgeBases"}}, nil)
		assert.Nil(t, err.Errors)
		assert.True(t, strings.HasSuffix((*recorded)[0].Url, "/v2/ai-agent/5740688905863168?expand=knowledgeBases"), (*recorded)[0].Url)
	})

	t.Run("test_go_page_returns_the_entities_the_cursor_and_sends_the_params", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("agents", "next-page", []string{agentAnswer}), pageAnswer("agents", "", []string{agentAnswer}))
		first, cursor, err := AiAgent.Page(map[string]interface{}{"limit": 1, "expand": []string{"knowledgeBases"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		assert.Equal(t, "5740688905863168", first[0].Id)
		assert.Equal(t, "next-page", cursor)
		assert.Contains(t, (*recorded)[0].Url, "limit=1")
		assert.Contains(t, (*recorded)[0].Url, "expand=knowledgeBases")
		_, cursor, err = AiAgent.Page(map[string]interface{}{"cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_follows_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("agents", "next-page", []string{agentAnswer}), pageAnswer("agents", "", []string{agentAnswer}))
		found := collectAgents(t, nil)
		assert.Len(t, found, 2)
		assert.Len(t, *recorded, 2)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_delete_sends_ids_in_the_query_string_and_returns_the_deleted_objects", func(t *testing.T) {
		recorded := answerInOrder(t, `{"agents":[`+agentAnswer+`]}`)
		deleted, err := AiAgent.Delete([]string{"5740688905863168", "5740688905863169"}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "DELETE", request.Method)
		assert.Equal(t, "", request.Body)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-agent?ids=5740688905863168%2C5740688905863169"), request.Url)
		assert.Len(t, deleted, 1)
		assert.Equal(t, "5740688905863168", deleted[0].Id)
	})
}
