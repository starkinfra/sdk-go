package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	BrcodePreview "github.com/starkinfra/sdk-go/starkinfra/brcodepreview"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccess(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	previews, err := BrcodePreview.Create(Example.BrcodePreview(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, preview := range previews {
		assert.NotNil(t, preview.Id)
	}
}
