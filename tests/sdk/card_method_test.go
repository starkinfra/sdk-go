package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	CardMethod "github.com/starkinfra/sdk-go/starkinfra/cardmethod"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCardMethodQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var methodsList []CardMethod.CardMethod

	methods, errorChannel := CardMethod.Query(nil, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case method, ok := <-methods:
			if !ok {
				break loop
			}
			methodsList = append(methodsList, method)
		}
	}
	for _, method := range methodsList {
		assert.NotNil(t, method.Code)
	}
}
