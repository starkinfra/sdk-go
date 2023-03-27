package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	MerchantCategory "github.com/starkinfra/sdk-go/starkinfra/merchantcategory"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantCategoryQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	categories := MerchantCategory.Query(nil, nil)
	for category := range categories {
		assert.NotNil(t, category.Code)
		fmt.Println(category.Code)
	}
}
