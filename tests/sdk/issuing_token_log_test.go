package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingTokenLog "github.com/starkinfra/sdk-go/starkinfra/issuingtoken/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingTokenLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, errorChannel := IssuingTokenLog.Query(params, nil)
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
			assert.NotNil(t, log.Id)
		}
	}
}

func TestIssuingTokenLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, cursor, err := IssuingTokenLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}

	assert.NotNil(t, cursor)
}
