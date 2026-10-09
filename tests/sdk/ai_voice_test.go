package sdk

import (
	"strings"
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiSpeech "github.com/starkinfra/sdk-go/starkinfra/aispeech"
	AiVoice "github.com/starkinfra/sdk-go/starkinfra/aivoice"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

const voiceAnswer = `{"id":"5631671361601536","name":"Helena","description":"Calm voice","language":"portuguese","gender":"female","status":"processing","errors":[],"created":"2026-10-01T14:28:24.566332+00:00","updated":"2026-10-01T14:28:24.566342+00:00"}`

func collectVoices(t *testing.T, params map[string]interface{}) []AiVoice.AiVoice {
	found := []AiVoice.AiVoice{}
	voices, errorChannel := AiVoice.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
		case voice, ok := <-voices:
			if !ok {
				return found
			}
			found = append(found, voice)
		}
	}
}

func TestAiVoiceLive(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_query_page_and_delete", func(t *testing.T) {
		speech := firstFinishedSpeech(t)
		audio, err := AiSpeech.Get(speech.Id, nil, nil)
		assert.Nil(t, err.Errors)
		voice, err := AiVoice.Create(AiVoice.AiVoice{Audio: audio.Audio, Name: "sdk-go-voice", Language: "portuguese", Gender: "female"}, nil)
		assert.Nil(t, err.Errors)
		assert.NotEmpty(t, voice.Id)
		assert.Equal(t, "sdk-go-voice", voice.Name)
		assert.NotNil(t, voice.Created)
		defer func() {
			deleted, err := AiVoice.Delete([]string{voice.Id}, nil)
			assert.Nil(t, err.Errors)
			if assert.Len(t, deleted, 1) {
				assert.Equal(t, voice.Id, deleted[0].Id)
			}
		}()
		ids := []string{}
		for _, entity := range collectVoices(t, nil) {
			ids = append(ids, entity.Id)
		}
		assert.Contains(t, ids, voice.Id)
		page, _, err := AiVoice.Page(map[string]interface{}{"limit": 1}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, page, 1)
	})

	t.Run("test_go_query_with_limit", func(t *testing.T) {
		firstReadyVoice(t)
		assert.Len(t, collectVoices(t, map[string]interface{}{"limit": 1}), 1)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiVoice.Page(map[string]interface{}{"limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})

	t.Run("test_go_create_without_audio_returns_input_errors", func(t *testing.T) {
		_, err := AiVoice.Create(AiVoice.AiVoice{Name: "no-audio"}, nil)
		assertInputError(t, err)
	})
}

func TestAiVoiceAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_sends_the_given_fields", func(t *testing.T) {
		recorded := answerInOrder(t, `{"voice":`+voiceAnswer+`}`)
		created, err := AiVoice.Create(
			AiVoice.AiVoice{Audio: "SUQzBAAAAAAA", Name: "Helena", Description: "Calm voice", Language: "portuguese", Gender: "female"},
			nil,
		)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "POST", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-voice"), request.Url)
		assert.Equal(t, []string{"audio", "description", "gender", "language", "name"}, bodyKeys(t, request.Body))
		assert.Equal(t, "5631671361601536", created.Id)
		assert.Equal(t, "processing", created.Status)
		assert.Equal(t, []string{}, created.Errors)
		assert.NotNil(t, created.Created)
	})

	t.Run("test_go_create_drops_the_optional_fields_left_empty", func(t *testing.T) {
		recorded := answerInOrder(t, `{"voice":`+voiceAnswer+`}`)
		_, err := AiVoice.Create(AiVoice.AiVoice{Audio: "SUQzBAAAAAAA"}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, []string{"audio"}, bodyKeys(t, (*recorded)[0].Body))
	})

	t.Run("test_go_page_returns_the_entities_and_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("voices", "next-page", []string{voiceAnswer}), pageAnswer("voices", "", []string{voiceAnswer}))
		first, cursor, err := AiVoice.Page(map[string]interface{}{"limit": 1}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		assert.Equal(t, "5631671361601536", first[0].Id)
		assert.Equal(t, "next-page", cursor)
		assert.Contains(t, (*recorded)[0].Url, "limit=1")
		_, cursor, err = AiVoice.Page(map[string]interface{}{"cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_follows_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("voices", "next-page", []string{voiceAnswer}), pageAnswer("voices", "", []string{voiceAnswer}))
		found := collectVoices(t, nil)
		assert.Len(t, found, 2)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_with_limit_150_asks_100_then_50", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			pageAnswer("voices", "first", repeated(voiceAnswer, 100)),
			pageAnswer("voices", "second", repeated(voiceAnswer, 50)),
		)
		found := collectVoices(t, map[string]interface{}{"limit": 150})
		assert.Len(t, found, 150)
		if assert.Len(t, *recorded, 2) {
			assert.Contains(t, (*recorded)[0].Url, "limit=100")
			assert.Contains(t, (*recorded)[1].Url, "limit=50")
		}
	})

	t.Run("test_go_delete_sends_ids_in_the_query_string_and_returns_the_deleted_objects", func(t *testing.T) {
		recorded := answerInOrder(t, `{"voices":[`+voiceAnswer+`]}`)
		deleted, err := AiVoice.Delete([]string{"5631671361601536", "5631671361601537"}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "DELETE", request.Method)
		assert.Equal(t, "", request.Body)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-voice?ids=5631671361601536%2C5631671361601537"), request.Url)
		assert.Len(t, deleted, 1)
		assert.Equal(t, "5631671361601536", deleted[0].Id)
	})
}
