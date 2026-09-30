package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiAgent "github.com/starkinfra/sdk-go/starkinfra/aiagent"
	AiKnowledgeBase "github.com/starkinfra/sdk-go/starkinfra/aiknowledgebase"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

func collectAgents(t *testing.T, params map[string]interface{}) []AiAgent.AiAgent {
	found := []AiAgent.AiAgent{}
	agents, errorChannel := AiAgent.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		case agent, ok := <-agents:
			if !ok {
				return found
			}
			found = append(found, agent)
		}
	}
}

func TestAiAgentLive(t *testing.T) {
	starkinfra.User = utils.ExampleProject
	knowledgeBase, errs := AiKnowledgeBase.Create(Example.ExampleAiKnowledgeBase(), nil)
	if errs.Errors != nil {
		t.Fatalf("code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
	}
	defer AiKnowledgeBase.Delete([]string{knowledgeBase.Id}, nil)
	agent, errs := AiAgent.Create(Example.ExampleAiAgent([]string{knowledgeBase.Id}), nil)
	if errs.Errors != nil {
		t.Fatalf("code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
	}
	defer AiAgent.Delete([]string{agent.Id}, nil)
	knowledgeBaseId := knowledgeBase.Id

	t.Run("test_go_create_returns_the_agent_with_the_schema_keys_in_camel_case", func(t *testing.T) {
		assert.NotEmpty(t, agent.Id)
		assert.Equal(t, "bender-1.0", agent.Model)
		assert.Equal(t, []string{knowledgeBaseId}, agent.KnowledgeBaseIds)
		assert.Contains(t, agent.MetadataSchema, "orderId")
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
		assert.Contains(t, renamed.MetadataSchema, "orderId")
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
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})

	t.Run("test_go_get_unknown_id_returns_input_errors", func(t *testing.T) {
		_, err := AiAgent.Get("0000000000000000", nil, nil)
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})
}
