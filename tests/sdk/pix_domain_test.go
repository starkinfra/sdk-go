package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixDomain "github.com/starkinfra/sdk-go/starkinfra/pixdomain"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixDomainQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	domains, errorChannel := PixDomain.Query(nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case domain, ok := <-domains:
			if !ok {
				break loop
			}
			assert.NotNil(t, domain.Name)
		}
	}
}
