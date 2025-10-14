package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingDesign "github.com/starkinfra/sdk-go/starkinfra/issuingdesign"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingDesignQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	designs, errorChannel := IssuingDesign.Query(nil, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case design, ok := <-designs:
			if !ok {
				break loop
			}
			assert.NotNil(t, design.Id)
		}
	}
}

func TestIssuingDesignPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 3
	var params = map[string]interface{}{}
	params["limit"] = limit

	designs, cursor, err := IssuingDesign.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, design := range designs {
		assert.NotNil(t, design.Id)
	}
	assert.NotNil(t, cursor)
}

func TestIssuingDesignInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 2
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var designList []IssuingDesign.IssuingDesign

	designs, errorChannel := IssuingDesign.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case design, ok := <-designs:
			if !ok {
				break loop
			}
			designList = append(designList, design)
		}
	}

	for _, design := range designList {
		getDesign, err := IssuingDesign.Get(design.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getDesign.Id)
	}
	assert.Equal(t, limit, len(designList))
}

func TestIssuingDesignPdf(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 2
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var designList []IssuingDesign.IssuingDesign

	designs, errorChannel := IssuingDesign.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case design, ok := <-designs:
			if !ok {
				break loop
			}
			designList = append(designList, design)
		}
	}
	

	design, err := IssuingDesign.Pdf(designList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, design)
}
