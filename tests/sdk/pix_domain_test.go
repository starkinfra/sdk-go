package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixDomain "github.com/starkinfra/sdk-go/starkinfra/pixdomain"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixDomainQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	domains := PixDomain.Query(nil)
	for domain := range domains {
		assert.NotNil(t, domain.Name)
		fmt.Println(domain.Name)
	}
}
