package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingTokenDesign "github.com/starkinfra/sdk-go/starkinfra/issuingtokendesign"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestIssuingTokenDesignQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	designs, errorChannel := IssuingTokenDesign.Query(params, nil)
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

func TestIssuingTokenDesignPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	designs, cursor, err := IssuingTokenDesign.Page(params, nil)
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

func TestIssuingTokenDesignQueryParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	designId := firstIssuingTokenDesignId(t)
	if designId != "" {
		params["ids"] = []string{designId}
	}

	designs, errorChannel := IssuingTokenDesign.Query(params, nil)
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

func TestIssuingTokenDesignPageParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	designId := firstIssuingTokenDesignId(t)
	if designId != "" {
		params["ids"] = []string{designId}
	}

	designs, cursor, err := IssuingTokenDesign.Page(params, nil)
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

func TestIssuingTokenDesignGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var designList []IssuingTokenDesign.IssuingTokenDesign

	designs, errorChannel := IssuingTokenDesign.Query(paramsQuery, nil)
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
		getDesign, err := IssuingTokenDesign.Get(design.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getDesign.Id)
		assert.IsType(t, "", getDesign.Name)
		assert.IsType(t, &time.Time{}, getDesign.Created)
		assert.IsType(t, &time.Time{}, getDesign.Updated)
	}
}

func TestIssuingTokenDesignPdf(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	designId := firstIssuingTokenDesignId(t)
	if designId == "" {
		return
	}

	pdf, err := IssuingTokenDesign.Pdf(designId, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	filename := fmt.Sprintf("%v%v.pdf", "issuingTokenDesign", designId)
	errFile := os.WriteFile(filename, pdf, 0666)
	if errFile != nil {
		t.Errorf("error: %s", errFile.Error())
	}
	assert.NotNil(t, pdf)
}

func firstIssuingTokenDesignId(t *testing.T) string {

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	designs, errorChannel := IssuingTokenDesign.Query(params, nil)
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
			return design.Id
		}
	}
	return ""
}
