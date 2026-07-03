package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixInternalTransactionReport "github.com/starkinfra/sdk-go/starkinfra/pixinternaltransactionreport"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixInternalTransactionReportPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	reports, err := PixInternalTransactionReport.Create(Example.PixInternalTransactionReport(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, report := range reports {
		assert.NotNil(t, report.Id)
		assert.NotNil(t, report.Status)
		assert.NotNil(t, report.Updated)
	}
}

func TestPixInternalTransactionReportQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	created, createErr := PixInternalTransactionReport.Create(Example.PixInternalTransactionReport(), nil)
	if createErr.Errors != nil {
		for _, e := range createErr.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	var ids []string
	for _, report := range created {
		ids = append(ids, report.Id)
	}

	var params = map[string]interface{}{}
	params["limit"] = len(ids)
	params["ids"] = ids

	reports, errorChannel := PixInternalTransactionReport.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case report, ok := <-reports:
			if !ok {
				break loop
			}
			assert.Contains(t, ids, report.Id)
		}
	}
}

func TestPixInternalTransactionReportPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	reports, cursor, err := PixInternalTransactionReport.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, report := range reports {
		assert.NotNil(t, report.Id)
	}

	assert.NotNil(t, cursor)
}

func TestPixInternalTransactionReportInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	created, createErr := PixInternalTransactionReport.Create(Example.PixInternalTransactionReport(), nil)
	if createErr.Errors != nil {
		for _, e := range createErr.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, report := range created {
		getReport, err := PixInternalTransactionReport.Get(report.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.Equal(t, report.Id, getReport.Id)

		assert.NotNil(t, getReport.Created)
		assert.NotNil(t, getReport.Updated)

		assert.NotEmpty(t, getReport.Status)
	}
}
