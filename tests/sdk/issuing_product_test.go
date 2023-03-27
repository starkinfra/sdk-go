package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingProduct "github.com/starkinfra/sdk-go/starkinfra/issuingproduct"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingProductQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	products := IssuingProduct.Query(params, nil)
	for product := range products {
		assert.NotNil(t, product.Id)
		fmt.Println(product.Id)
	}
}

func TestIssuingProductPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	products, cursor, err := IssuingProduct.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, product := range products {
		assert.NotNil(t, product.Id)
		fmt.Println(product.Id)
	}

	fmt.Println(cursor)
}
