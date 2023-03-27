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

	methods := CardMethod.Query(nil, nil)
	for method := range methods {
		assert.NotNil(t, method.Code)
	}
}
