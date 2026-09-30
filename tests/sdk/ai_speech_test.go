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

const speechAnswer = `{"id":"5646488461901824","voiceId":"5632499082330112","text":"Short test.","status":"success","audio":"SUQzBAAAAAAA","errors":[],"created":"2026-10-01T14:28:06.942491+00:00","updated":"2026-10-01T14:28:07.605185+00:00"}`

func collectSpeeches(t *testing.T, params map[string]interface{}) []AiSpeech.AiSpeech {
	found := []AiSpeech.AiSpeech{}
	speeches, errorChannel := AiSpeech.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
		case speech, ok := <-speeches:
			if !ok {
				return found
			}
			found = append(found, speech)
		}
	}
}

func firstFinishedSpeech(t *testing.T) AiSpeech.AiSpeech {
	starkinfra.User = utils.ExampleProject
	for _, speech := range collectSpeeches(t, nil) {
		if speech.Status == "success" {
			return speech
		}
	}
	t.Skip("the workspace has no finished speech to read")
	return AiSpeech.AiSpeech{}
}

func firstReadyVoice(t *testing.T) AiVoice.AiVoice {
	starkinfra.User = utils.ExampleProject
	for _, voice := range collectVoices(t, nil) {
		if voice.Status == "success" {
			return voice
		}
	}
	t.Skip("the workspace has no voice ready to speak")
	return AiVoice.AiVoice{}
}

func TestAiSpeechLive(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_query_leaves_the_audio_out", func(t *testing.T) {
		for _, speech := range collectSpeeches(t, map[string]interface{}{"limit": 5}) {
			assert.NotEmpty(t, speech.Id)
			assert.Empty(t, speech.Audio)
			assert.NotNil(t, speech.Created)
		}
	})

	t.Run("test_go_query_with_limit", func(t *testing.T) {
		firstFinishedSpeech(t)
		assert.Len(t, collectSpeeches(t, map[string]interface{}{"limit": 1}), 1)
	})

	t.Run("test_go_page_returns_the_entities_and_walks_to_the_end", func(t *testing.T) {
		seen := 0
		cursor := ""
		for {
			params := map[string]interface{}{"limit": 5}
			if cursor != "" {
				params["cursor"] = cursor
			}
			speeches, next, err := AiSpeech.Page(params, nil)
			assert.Nil(t, err.Errors)
			assert.LessOrEqual(t, len(speeches), 5)
			seen += len(speeches)
			if next == "" {
				break
			}
			cursor = next
		}
		assert.GreaterOrEqual(t, seen, 0)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiSpeech.Page(map[string]interface{}{"limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})

	t.Run("test_go_get_returns_the_audio_and_expands_the_voice_name", func(t *testing.T) {
		speech := firstFinishedSpeech(t)
		fetched, err := AiSpeech.Get(speech.Id, nil, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, speech.Id, fetched.Id)
		assert.NotEmpty(t, fetched.Audio)
		assert.Empty(t, fetched.VoiceName)
		expanded, err := AiSpeech.Get(speech.Id, map[string]interface{}{"expand": []string{"voiceName"}}, nil)
		assert.Nil(t, err.Errors)
		assert.NotEmpty(t, expanded.VoiceName)
	})

	t.Run("test_go_create_synthesizes_the_text_with_a_ready_voice", func(t *testing.T) {
		voice := firstReadyVoice(t)
		speech, err := AiSpeech.Create(AiSpeech.AiSpeech{VoiceId: voice.Id, Text: "Short test."}, nil)
		assert.Nil(t, err.Errors)
		assert.NotEmpty(t, speech.Id)
		assert.Equal(t, voice.Id, speech.VoiceId)
		assert.Equal(t, "Short test.", speech.Text)
		assert.NotNil(t, speech.Created)
	})

	t.Run("test_go_get_unknown_id_returns_input_errors", func(t *testing.T) {
		_, err := AiSpeech.Get("0000000000000000", nil, nil)
		assertInputError(t, err)
	})
}

func TestAiSpeechAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_sends_only_the_creatable_fields", func(t *testing.T) {
		recorded := answerInOrder(t, `{"speech":`+speechAnswer+`}`)
		created, err := AiSpeech.Create(AiSpeech.AiSpeech{VoiceId: "5632499082330112", Text: "Short test."}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "POST", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-speech"), request.Url)
		assert.Equal(t, []string{"text", "voiceId"}, bodyKeys(t, request.Body))
		assert.Equal(t, "5632499082330112", created.VoiceId)
		assert.Equal(t, "SUQzBAAAAAAA", created.Audio)
		assert.Equal(t, "success", created.Status)
	})

	t.Run("test_go_get_reads_the_speech_key_and_sends_expand", func(t *testing.T) {
		recorded := answerInOrder(t, `{"speech":`+speechAnswer+`}`)
		speech, err := AiSpeech.Get("5646488461901824", map[string]interface{}{"expand": []string{"voiceName"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "5646488461901824", speech.Id)
		assert.Equal(t, "SUQzBAAAAAAA", speech.Audio)
		assert.True(t, strings.HasSuffix((*recorded)[0].Url, "/v2/ai-speech/5646488461901824?expand=voiceName"), (*recorded)[0].Url)
	})

	t.Run("test_go_page_reads_the_speeches_key_and_returns_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("speeches", "next-page", []string{speechAnswer}), pageAnswer("speeches", "", []string{speechAnswer}))
		first, cursor, err := AiSpeech.Page(map[string]interface{}{"limit": 1, "expand": []string{"voiceName"}}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		assert.Equal(t, "5646488461901824", first[0].Id)
		assert.Equal(t, "next-page", cursor)
		assert.Contains(t, (*recorded)[0].Url, "/v2/ai-speech?")
		assert.Contains(t, (*recorded)[0].Url, "limit=1")
		assert.Contains(t, (*recorded)[0].Url, "expand=voiceName")
		_, cursor, err = AiSpeech.Page(map[string]interface{}{"cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_follows_the_cursor_through_empty_pages", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			pageAnswer("speeches", "first", []string{speechAnswer}),
			pageAnswer("speeches", "second", nil),
			pageAnswer("speeches", "", []string{speechAnswer}),
		)
		found := collectSpeeches(t, nil)
		assert.Len(t, found, 2)
		if assert.Len(t, *recorded, 3) {
			assert.NotContains(t, (*recorded)[0].Url, "cursor")
			assert.NotContains(t, (*recorded)[0].Url, "limit")
			assert.Contains(t, (*recorded)[1].Url, "cursor=first")
			assert.Contains(t, (*recorded)[2].Url, "cursor=second")
		}
	})

	t.Run("test_go_query_with_limit_150_asks_100_then_50_and_stops", func(t *testing.T) {
		recorded := answerInOrder(
			t,
			pageAnswer("speeches", "first", repeated(speechAnswer, 100)),
			pageAnswer("speeches", "second", repeated(speechAnswer, 50)),
		)
		found := collectSpeeches(t, map[string]interface{}{"limit": 150, "expand": []string{"voiceName"}})
		assert.Len(t, found, 150)
		if assert.Len(t, *recorded, 2) {
			assert.Contains(t, (*recorded)[0].Url, "limit=100")
			assert.Contains(t, (*recorded)[1].Url, "limit=50")
			assert.Contains(t, (*recorded)[1].Url, "expand=voiceName")
		}
	})
}
