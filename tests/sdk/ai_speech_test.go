package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiSpeech "github.com/starkinfra/sdk-go/starkinfra/aispeech"
	AiVoice "github.com/starkinfra/sdk-go/starkinfra/aivoice"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func collectSpeeches(t *testing.T, params map[string]interface{}) []AiSpeech.AiSpeech {
	found := []AiSpeech.AiSpeech{}
	speeches, errorChannel := AiSpeech.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
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
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})
}
