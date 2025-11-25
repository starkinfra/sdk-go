package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	MerchantCategory "github.com/starkinfra/sdk-go/starkinfra/merchantcategory"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantCategoryQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	categories, errorChannel := MerchantCategory.Query(nil, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case category, ok := <-categories:
			if !ok {
				break loop
			}
			assert.NotNil(t, category.Code)
		}
	}
}
