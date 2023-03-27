package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditPreview "github.com/starkinfra/sdk-go/starkinfra/creditpreview"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreditPreview(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	previews, err := CreditPreview.Create(Example.CreditPreview(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, preview := range previews {
		assert.NotNil(t, preview.Type)
	}
}
