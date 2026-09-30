package sdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"
	"strings"
	"testing"

	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/stretchr/testify/assert"
)

type scriptedTransport struct {
	answers  []string
	recorded *[]recordedRequest
}

func (transport scriptedTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	body := ""
	if request.Body != nil {
		content, _ := io.ReadAll(request.Body)
		body = string(content)
	}
	index := len(*transport.recorded)
	*transport.recorded = append(*transport.recorded, recordedRequest{Method: request.Method, Url: request.URL.String(), Body: body})
	if index >= len(transport.answers) {
		index = len(transport.answers) - 1
	}
	return &http.Response{
		StatusCode: 200,
		Body:       io.NopCloser(strings.NewReader(transport.answers[index])),
		Header:     http.Header{},
		Request:    request,
	}, nil
}

type recordedRequest struct {
	Method string
	Url    string
	Body   string
}

func answerInOrder(t *testing.T, answers ...string) *[]recordedRequest {
	recorded := &[]recordedRequest{}
	original := http.DefaultTransport
	http.DefaultTransport = scriptedTransport{answers: answers, recorded: recorded}
	t.Cleanup(func() { http.DefaultTransport = original })
	return recorded
}

func bodyKeys(t *testing.T, body string) []string {
	sent := map[string]interface{}{}
	assert.Nil(t, json.Unmarshal([]byte(body), &sent))
	keys := []string{}
	for key := range sent {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func bodyOf(t *testing.T, body string) map[string]interface{} {
	sent := map[string]interface{}{}
	assert.Nil(t, json.Unmarshal([]byte(body), &sent))
	return sent
}

func assertInputError(t *testing.T, errs Error.StarkErrors) {
	if !assert.NotEmpty(t, errs.Errors) {
		return
	}
	assert.NotEqual(t, "internalServerError", errs.Errors[0].Code)
	assert.NotEqual(t, "unknownError", errs.Errors[0].Code)
}

func failOnErrors(t *testing.T, errs Error.StarkErrors) {
	for _, e := range errs.Errors {
		t.Errorf("code: %s, message: %s", e.Code, e.Message)
	}
}

func pageAnswer(key string, cursor string, items []string) string {
	cursorJson := "null"
	if cursor != "" {
		cursorJson = fmt.Sprintf("%q", cursor)
	}
	return fmt.Sprintf(`{"cursor":%s,"%s":[%s]}`, cursorJson, key, strings.Join(items, ","))
}

func repeated(item string, count int) []string {
	items := []string{}
	for index := 0; index < count; index++ {
		items = append(items, item)
	}
	return items
}
