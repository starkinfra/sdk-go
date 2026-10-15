package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiKnowledgeBase "github.com/starkinfra/sdk-go/starkinfra/aiknowledgebase"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

func collectKnowledgeBases(t *testing.T, params map[string]interface{}) []AiKnowledgeBase.AiKnowledgeBase {
	found := []AiKnowledgeBase.AiKnowledgeBase{}
	knowledgeBases, errorChannel := AiKnowledgeBase.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		case knowledgeBase, ok := <-knowledgeBases:
			if !ok {
				return found
			}
			found = append(found, knowledgeBase)
		}
	}
}

func TestAiKnowledgeBaseLive(t *testing.T) {
	starkinfra.User = utils.ExampleProject
	knowledgeBase, errs := AiKnowledgeBase.Create(Example.ExampleAiKnowledgeBase(), nil)
	if errs.Errors != nil {
		t.Fatalf("code: %s, message: %s", errs.Errors[0].Code, errs.Errors[0].Message)
	}
	defer AiKnowledgeBase.Delete([]string{knowledgeBase.Id}, nil)

	t.Run("test_go_create_returns_processing_knowledge_base", func(t *testing.T) {
		assert.NotEmpty(t, knowledgeBase.Id)
		assert.Equal(t, "processing", knowledgeBase.Status)
		assert.Equal(t, "https://docs.starkinfra.com", knowledgeBase.RootUrl)
		assert.NotNil(t, knowledgeBase.IsRecursive)
		assert.False(t, *knowledgeBase.IsRecursive)
		assert.Equal(t, []string{"sdk-go", "test"}, knowledgeBase.Tags)
		assert.NotNil(t, knowledgeBase.Created)
	})

	t.Run("test_go_get", func(t *testing.T) {
		fetched, err := AiKnowledgeBase.Get(knowledgeBase.Id, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, knowledgeBase.Id, fetched.Id)
		assert.Equal(t, knowledgeBase.Name, fetched.Name)
	})

	t.Run("test_go_query_filters_by_ids", func(t *testing.T) {
		found := collectKnowledgeBases(t, map[string]interface{}{"ids": []string{knowledgeBase.Id}})
		assert.Len(t, found, 1)
		assert.Equal(t, knowledgeBase.Id, found[0].Id)
	})

	t.Run("test_go_query_filters_by_name_and_status", func(t *testing.T) {
		current, err := AiKnowledgeBase.Get(knowledgeBase.Id, nil)
		assert.Nil(t, err.Errors)
		found := collectKnowledgeBases(t, map[string]interface{}{"name": current.Name, "status": current.Status})
		ids := []string{}
		for _, entity := range found {
			ids = append(ids, entity.Id)
		}
		assert.Contains(t, ids, knowledgeBase.Id)
	})

	t.Run("test_go_query_without_match_is_empty", func(t *testing.T) {
		found := collectKnowledgeBases(t, map[string]interface{}{"name": "no-knowledge-base-has-this-name"})
		assert.Empty(t, found)
	})

	t.Run("test_go_query_with_limit", func(t *testing.T) {
		assert.Len(t, collectKnowledgeBases(t, map[string]interface{}{"limit": 1}), 1)
	})

	t.Run("test_go_page_returns_the_entities_and_walks_to_the_end", func(t *testing.T) {
		seen := 0
		cursor := ""
		for {
			params := map[string]interface{}{"limit": 2}
			if cursor != "" {
				params["cursor"] = cursor
			}
			knowledgeBases, next, err := AiKnowledgeBase.Page(params, nil)
			assert.Nil(t, err.Errors)
			assert.LessOrEqual(t, len(knowledgeBases), 2)
			seen += len(knowledgeBases)
			if next == "" {
				break
			}
			cursor = next
		}
		assert.GreaterOrEqual(t, seen, 1)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiKnowledgeBase.Page(map[string]interface{}{"limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})

	t.Run("test_go_update_changes_name_and_tags_only", func(t *testing.T) {
		own, err := AiKnowledgeBase.Create(Example.ExampleAiKnowledgeBase(), nil)
		assert.Nil(t, err.Errors)
		defer func() {
			_, err := AiKnowledgeBase.Delete([]string{own.Id}, nil)
			assert.Nil(t, err.Errors)
		}()
		updated, err := AiKnowledgeBase.Update(own.Id, map[string]interface{}{"name": "renamed-by-sdk", "tags": []string{"renamed"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "renamed-by-sdk", updated.Name)
		assert.Equal(t, []string{"renamed"}, updated.Tags)
		assert.Equal(t, own.RootUrl, updated.RootUrl)
		cleared, err := AiKnowledgeBase.Update(own.Id, map[string]interface{}{"tags": []string{}}, nil)
		assert.Nil(t, err.Errors)
		assert.Empty(t, cleared.Tags)
		assert.Equal(t, "renamed-by-sdk", cleared.Name)
	})

	t.Run("test_go_delete_returns_the_deleted_knowledge_bases", func(t *testing.T) {
		own, err := AiKnowledgeBase.Create(Example.ExampleAiKnowledgeBase(), nil)
		assert.Nil(t, err.Errors)
		deleted, err := AiKnowledgeBase.Delete([]string{own.Id}, nil)
		assert.Nil(t, err.Errors)
		if assert.Len(t, deleted, 1) {
			assert.Equal(t, own.Id, deleted[0].Id)
		}
	})

	t.Run("test_go_create_with_invalid_root_url_returns_input_errors", func(t *testing.T) {
		_, err := AiKnowledgeBase.Create(AiKnowledgeBase.AiKnowledgeBase{Name: "invalid", RootUrl: "not-a-url"}, nil)
		if assert.Len(t, err.Errors, 1) {
			assert.Equal(t, "invalidRootUrl", err.Errors[0].Code)
		}
	})

	t.Run("test_go_get_unknown_id_returns_input_errors", func(t *testing.T) {
		_, err := AiKnowledgeBase.Get("0000000000000000", nil)
		if assert.Len(t, err.Errors, 1) {
			assert.Equal(t, "invalidKnowledgeBaseId", err.Errors[0].Code)
		}
	})
}
