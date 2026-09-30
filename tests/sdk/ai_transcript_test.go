package sdk

import (
	"strings"
	"testing"

	"github.com/starkinfra/sdk-go/starkinfra"
	AiSpeech "github.com/starkinfra/sdk-go/starkinfra/aispeech"
	AiTranscript "github.com/starkinfra/sdk-go/starkinfra/aitranscript"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

const transcriptAnswer = `{"id":"5147403464212480","text":"This is a short recording used to test the transcription service.","status":"success","errors":[],"created":"2026-10-01T14:28:04.482326+00:00","updated":"2026-10-01T14:28:05.752389+00:00"}`

func collectTranscripts(t *testing.T, params map[string]interface{}) []AiTranscript.AiTranscript {
	found := []AiTranscript.AiTranscript{}
	transcripts, errorChannel := AiTranscript.Query(params, nil)
	for {
		select {
		case err := <-errorChannel:
			failOnErrors(t, err)
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
		assertInputError(t, err)
	})
}

func TestAiTranscriptAtTheHttpBoundary(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	t.Run("test_go_create_sends_the_audio", func(t *testing.T) {
		recorded := answerInOrder(t, `{"transcript":`+transcriptAnswer+`}`)
		created, err := AiTranscript.Create(AiTranscript.AiTranscript{Audio: "SUQzBAAAAAAA"}, nil)
		assert.Nil(t, err.Errors)
		request := (*recorded)[0]
		assert.Equal(t, "POST", request.Method)
		assert.True(t, strings.HasSuffix(request.Url, "/v2/ai-transcript"), request.Url)
		assert.Equal(t, []string{"audio"}, bodyKeys(t, request.Body))
		assert.Equal(t, "This is a short recording used to test the transcription service.", created.Text)
		assert.Equal(t, "success", created.Status)
		assert.NotNil(t, created.Updated)
	})

	t.Run("test_go_page_returns_the_entities_and_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("transcripts", "next-page", []string{transcriptAnswer}), pageAnswer("transcripts", "", []string{transcriptAnswer}))
		first, cursor, err := AiTranscript.Page(map[string]interface{}{"limit": 1}, nil)
		assert.Nil(t, err.Errors)
		assert.Len(t, first, 1)
		assert.Equal(t, "5147403464212480", first[0].Id)
		assert.Equal(t, "next-page", cursor)
		assert.Contains(t, (*recorded)[0].Url, "limit=1")
		_, cursor, err = AiTranscript.Page(map[string]interface{}{"cursor": cursor}, nil)
		assert.Nil(t, err.Errors)
		assert.Equal(t, "", cursor)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})

	t.Run("test_go_query_follows_the_cursor", func(t *testing.T) {
		recorded := answerInOrder(t, pageAnswer("transcripts", "next-page", []string{transcriptAnswer}), pageAnswer("transcripts", "", []string{transcriptAnswer}))
		found := collectTranscripts(t, nil)
		assert.Len(t, found, 2)
		assert.Contains(t, (*recorded)[1].Url, "cursor=next-page")
	})
}
