package sdk

import (
	"strings"
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiKnowledgeBase "github.com/starkinfra/sdk-go/starkinfra/aiknowledgebase"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
)

const knowledgeBaseAnswer = `{"id":"6767676767676767","name":"Public Documentation","rootUrl":"https://docs.starkinfra.com","isRecursive":true,"status":"success","tags":["support"],"created":"2022-01-01T00:00:00.000000+00:00","updated":"2022-01-02T00:00:00.000000+00:00"}`

func collectKnowledgeBases(t *testing.T, params map[string]interface{}) []AiKnowledgeBase.AiKnowledgeBase {
	found := []AiKnowledgeBase.AiKnowledgeBase{}
	knowledgeBases, errorChannel := AiKnowledgeBase.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
		case knowledgeBase, ok := <-knowledgeBases:
			if !ok {
				return found
			}
			found = append(found, knowledgeBase)
		}
	}
}

func TestAiKnowledgeBaseLive(t *testing.T) {
	fixtures.load(t)
	knowledgeBase := fixtures.knowledgeBase

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

func TestAiKnowledgeBaseAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_hosts_groups_pages_by_host", func(t *testing.T) {
		recorded := answerInOrder(t, `{"hosts": {"docs.starkinfra.com": [{"originalUrl": "https://docs.starkinfra.com/get-started", "status": "success", "storageUrl": "https://storage.googleapis.com/ai-knowledge/6767676767676767/get-started.md"}]}}`)
		hosts, err := AiKnowledgeBase.Hosts("6767676767676767", nil)
		assert.Nil(t, err.Errors)
		assert.True(t, strings.HasSuffix((*recorded)[0].Url, "/v2/ai-knowledge-base/6767676767676767/hosts"), (*recorded)[0].Url)
		assert.Equal(t, "GET", (*recorded)[0].Method)
		assert.Equal(t, map[string][]AiKnowledgeBase.HostPage{
			"docs.starkinfra.com": {{
				OriginalUrl: "https://docs.starkinfra.com/get-started",
				Status:      "success",
				StorageUrl:  "https://storage.googleapis.com/ai-knowledge/6767676767676767/get-started.md",
			}},
		}, hosts)
	})

	t.Run("test_go_delete_sends_ids_in_the_query_string_and_returns_the_deleted_objects", func(t *testing.T) {
		recorded := answerInOrder(t, `{"knowledgeBases": [`+knowledgeBaseAnswer+`]}`)
		deleted, err := AiKnowledgeBase.Delete([]string{"6767676767676767", "6767676767676768"}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-knowledge-base?ids=6767676767676767%2C6767676767676768"), request.Url)
		assert.Equal(t, "DELETE", request.Method)
		assert.Equal(t, "", request.Body)
		assert.Len(t, deleted, 1)
		assert.Equal(t, "6767676767676767", deleted[0].Id)
		assert.Equal(t, "Public Documentation", deleted[0].Name)
	})

	t.Run("test_go_create_sends_only_the_creatable_fields", func(t *testing.T) {
		recorded := answerInOrder(t, `{"knowledgeBase": `+knowledgeBaseAnswer+`}`)
		notRecursive := false
		returned := AiKnowledgeBase.AiKnowledgeBase{
			Name:        "Public Documentation",
			RootUrl:     "https://docs.starkinfra.com",
			IsRecursive: &notRecursive,
			Tags:        []string{"support"},
		}
		created, err := AiKnowledgeBase.Create(returned, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "POST", (*recorded)[0].Method)
		assert.True(t, strings.HasSuffix((*recorded)[0].Url, "/v2/ai-knowledge-base"), (*recorded)[0].Url)
		assert.Equal(t, []string{"isRecursive", "name", "rootUrl", "tags"}, bodyKeys(t, (*recorded)[0].Body))
		assert.Equal(t, false, bodyOf(t, (*recorded)[0].Body)["isRecursive"])
		assert.Equal(t, "6767676767676767", created.Id)
		assert.NotNil(t, created.Created)
	})

	t.Run("test_go_create_sends_an_empty_tags_list", func(t *testing.T) {
		recorded := answerInOrder(t, `{"knowledgeBase": `+knowledgeBaseAnswer+`}`)
		_, err := AiKnowledgeBase.Create(AiKnowledgeBase.AiKnowledgeBase{Name: "Docs", RootUrl: "https://docs.starkinfra.com", Tags: []string{}}, nil)
		assert.Nil(t, err.Errors)
		sent := bodyOf(t, (*recorded)[0].Body)
		assert.Equal(t, []interface{}{}, sent["tags"])
		assert.NotContains(t, sent, "isRecursive")
	})

	t.Run("test_go_get_reads_the_knowledge_base_key", func(t *testing.T) {
		recorded := answerInOrder(t, `{"knowledgeBase": `+knowledgeBaseAnswer+`}`)
		knowledgeBase, err := AiKnowledgeBase.Get("6767676767676767", nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "6767676767676767", knowledgeBase.Id)
		assert.True(t, strings.HasSuffix((*recorded)[0].Url, "/v2/ai-knowledge-base/6767676767676767"), (*recorded)[0].Url)
	})

	t.Run("test_go_update_sends_only_the_given_fields", func(t *testing.T) {
		recorded := answerInOrder(t, `{"knowledgeBase": `+knowledgeBaseAnswer+`}`)
		_, err := AiKnowledgeBase.Update("6767676767676767", map[string]interface{}{"name": "Renamed", "tags": []string{}}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "PATCH", (*recorded)[0].Method)
		assert.Equal(t, map[string]interface{}{"name": "Renamed", "tags": []interface{}{}}, bodyOf(t, (*recorded)[0].Body))
	})

	t.Run("test_go_page_sends_ids_and_filters_and_returns_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("knowledgeBases", "next-page", []string{knowledgeBaseAnswer}), pageAnswer("knowledgeBases", "", []string{knowledgeBaseAnswer}))
		params := map[string]interface{}{"ids": []string{"1", "2"}, "name": "docs", "status": "success", "limit": 1}
		first, cursor, err := AiKnowledgeBase.Page(params, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		assert.Equal(t, "next-page", cursor)
		url := (*recorded)[0].Url
		assert.Contains(t, url, "ids=1%2C2")
		assert.Contains(t, url, "name=docs")
		assert.Contains(t, url, "status=success")
		assert.Contains(t, url, "limit=1")
		_, cursor, err = AiKnowledgeBase.Page(map[string]interface{}{"cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_follows_the_cursor_through_empty_pages", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			pageAnswer("knowledgeBases", "first", []string{knowledgeBaseAnswer}),
			pageAnswer("knowledgeBases", "second", nil),
			pageAnswer("knowledgeBases", "", []string{knowledgeBaseAnswer}),
		)
		found := collectKnowledgeBases(t, map[string]interface{}{"ids": []string{"1", "2"}})
		assert.Len(t, found, 2)
		if assert.Len(t, *recorded, 3) {
			assert.NotContains(t, (*recorded)[0].Url, "cursor")
			assert.Contains(t, (*recorded)[0].Url, "ids=1%2C2")
			assert.Contains(t, (*recorded)[1].Url, "cursor=first")
			assert.Contains(t, (*recorded)[2].Url, "cursor=second")
		}
	})

	t.Run("test_go_query_with_limit_150_asks_100_then_50_and_stops", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			pageAnswer("knowledgeBases", "first", repeated(knowledgeBaseAnswer, 100)),
			pageAnswer("knowledgeBases", "second", repeated(knowledgeBaseAnswer, 50)),
		)
		found := collectKnowledgeBases(t, map[string]interface{}{"limit": 150})
		assert.Len(t, found, 150)
		if assert.Len(t, *recorded, 2) {
			assert.Contains(t, (*recorded)[0].Url, "limit=100")
			assert.Contains(t, (*recorded)[1].Url, "limit=50")
			assert.Contains(t, (*recorded)[1].Url, "cursor=first")
		}
	})

	t.Run("test_go_query_does_not_change_the_callers_params", func(t *testing.T) {
		answerInOrder(t, pageAnswer("knowledgeBases", "first", []string{knowledgeBaseAnswer}), pageAnswer("knowledgeBases", "", nil))
		params := map[string]interface{}{"limit": 5}
		collectKnowledgeBases(t, params)
		assert.Equal(t, map[string]interface{}{"limit": 5}, params)
	})
}
