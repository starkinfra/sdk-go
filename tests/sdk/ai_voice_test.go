package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiSpeech "github.com/starkinfra/sdk-go/starkinfra/aispeech"
	AiVoice "github.com/starkinfra/sdk-go/starkinfra/aivoice"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func collectVoices(t *testing.T, params map[string]interface{}) []AiVoice.AiVoice {
	found := []AiVoice.AiVoice{}
	voices, errorChannel := AiVoice.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
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
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})
}
