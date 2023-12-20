package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixUser "github.com/starkinfra/sdk-go/starkinfra/pixuser"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixUserInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	fraud, err := PixUser.Get("01234567890", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, fraud.Id)
	fmt.Println(fraud.Id)
}
