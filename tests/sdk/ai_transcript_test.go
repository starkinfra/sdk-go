package sdk

import (
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiSpeech "github.com/starkinfra/sdk-go/starkinfra/aispeech"
	AiTranscript "github.com/starkinfra/sdk-go/starkinfra/aitranscript"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

func collectTranscripts(t *testing.T, params map[string]interface{}) []AiTranscript.AiTranscript {
	found := []AiTranscript.AiTranscript{}
	transcripts, errorChannel := AiTranscript.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		case transcript, ok := <-transcripts:
			if !ok {
				return found
			}
			found = append(found, transcript)
		}
	}
}

func TestAiTranscriptLive(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_query_and_page", func(t *testing.T) {
		speech := firstFinishedSpeech(t)
		audio, err := AiSpeech.Get(speech.Id, nil, nil)
		assert.Nil(t, err.Errors)
		transcript, err := AiTranscript.Create(AiTranscript.AiTranscript{Audio: audio.Audio}, nil)
		assert.Nil(t, err.Errors)
		assert.NotEmpty(t, transcript.Id)
		assert.NotNil(t, transcript.Created)
		page, _, err := AiTranscript.Page(map[string]interface{}{"limit": 1}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, page, 1)
		assert.Len(t, collectTranscripts(t, map[string]interface{}{"limit": 1}), 1)
	})

	t.Run("test_go_page_with_an_invalid_limit_returns_the_api_error", func(t *testing.T) {
		for _, limit := range []int{0, -1, 101} {
			_, _, err := AiTranscript.Page(map[string]interface{}{"limit": limit}, nil)
			if assert.NotEmpty(t, err.Errors, "limit %d", limit) {
				assert.Equal(t, "invalidLimit", err.Errors[0].Code)
			}
		}
	})

	t.Run("test_go_create_without_audio_returns_input_errors", func(t *testing.T) {
		_, err := AiTranscript.Create(AiTranscript.AiTranscript{}, nil)
		if assert.NotEmpty(t, err.Errors) {
			assert.NotEqual(t, "internalServerError", err.Errors[0].Code)
		}
	})
}
