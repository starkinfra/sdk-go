package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixDirector "github.com/starkinfra/sdk-go/starkinfra/pixdirector"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixDirectorPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	director, err := PixDirector.Create(Example.PixDirector(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, director.Name)
}
