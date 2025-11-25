package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingProduct "github.com/starkinfra/sdk-go/starkinfra/issuingproduct"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingProductQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	products, errorChannel := IssuingProduct.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case product, ok := <-products:
			if !ok {
				break loop
			}
			assert.NotNil(t, product.Id)
		}
	}
}

func TestIssuingProductPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	products, cursor, err := IssuingProduct.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, product := range products {
		assert.NotNil(t, product.Id)
	}

	assert.NotNil(t, cursor)
}
