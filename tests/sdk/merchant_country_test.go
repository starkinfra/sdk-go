package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	MerchantCountry "github.com/starkinfra/sdk-go/starkinfra/merchantcountry"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMerchantCountryQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["search"] = "brazil"

	countries := MerchantCountry.Query(params, nil)
	for country := range countries {
		assert.NotNil(t, country.Code)
		fmt.Println(country.Code)
	}
}
