package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixKeyHolmesLog "github.com/starkinfra/sdk-go/starkinfra/pixkeyholmes/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixKeyHolmesLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	var logsList []PixKeyHolmesLog.Log

	logs, errorChannel := PixKeyHolmesLog.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			assert.NotEmpty(t, log.Id)
			assert.NotEmpty(t, log.Type)
			assert.NotEmpty(t, log.Holmes.Id)
			logsList = append(logsList, log)
		}
	}
	assert.NotEmpty(t, logsList)
}

func TestPixKeyHolmesLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 2

	logs, cursor, err := PixKeyHolmesLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	ids := map[string]bool{}
	for _, log := range logs {
		assert.False(t, ids[log.Id])
		ids[log.Id] = true
	}

	assert.NotEmpty(t, ids)
	assert.NotNil(t, cursor)
}

func TestPixKeyHolmesLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	logs, _, err := PixKeyHolmesLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	if len(logs) == 0 {
		t.Fatal("no PixKeyHolmes logs available")
	}

	log, err := PixKeyHolmesLog.Get(logs[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.Equal(t, logs[0].Id, log.Id)
	assert.NotEmpty(t, log.Holmes.Id)
}
