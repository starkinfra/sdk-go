package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingDesign "github.com/starkinfra/sdk-go/starkinfra/issuingdesign"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingDesignQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	designs := IssuingDesign.Query(nil, nil)
	for design := range designs {
		assert.NotNil(t, design.Id)
		fmt.Println(design.Id)
	}
}

func TestIssuingDesignPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	designs, cursor, err := IssuingDesign.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, design := range designs {
		assert.NotNil(t, design.Id)
		fmt.Println(design.Id)
	}
	fmt.Println(cursor)
}

func TestIssuingDesignInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var designList []IssuingDesign.IssuingDesign
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	designs := IssuingDesign.Query(paramsQuery, nil)
	for design := range designs {
		designList = append(designList, design)
	}

	design, err := IssuingDesign.Get(designList[rand.Intn(len(designList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, design.Id)
	fmt.Println(design.Id)
}

func TestIssuingDesignPdf(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var designList []IssuingDesign.IssuingDesign
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	designs := IssuingDesign.Query(paramsQuery, nil)
	for design := range designs {
		designList = append(designList, design)
	}

	design, err := IssuingDesign.Pdf(designList[rand.Intn(len(designList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}
	assert.NotNil(t, design)
}
