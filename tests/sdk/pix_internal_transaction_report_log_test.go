package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixInternalTransactionReport "github.com/starkinfra/sdk-go/starkinfra/pixinternaltransactionreport"
	PixInternalTransactionReportLog "github.com/starkinfra/sdk-go/starkinfra/pixinternaltransactionreport/log"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixInternalTransactionReportLogQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	created, createErr := PixInternalTransactionReport.Create(Example.PixInternalTransactionReport(), nil)
	if createErr.Errors != nil {
		for _, e := range createErr.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	var reportIds []string
	for _, report := range created {
		reportIds = append(reportIds, report.Id)
	}

	var params = map[string]interface{}{}
	params["reportIds"] = reportIds

	logs, errorChannel := PixInternalTransactionReportLog.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			assert.NotNil(t, log.Id)
		}
	}
}

func TestPixInternalTransactionReportLogPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	logs, cursor, err := PixInternalTransactionReportLog.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, log := range logs {
		assert.NotNil(t, log.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixInternalTransactionReportLogGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	created, createErr := PixInternalTransactionReport.Create(Example.PixInternalTransactionReport(), nil)
	if createErr.Errors != nil {
		for _, e := range createErr.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	var reportIds []string
	for _, report := range created {
		reportIds = append(reportIds, report.Id)
	}

	var paramsQuery = map[string]interface{}{}
	paramsQuery["reportIds"] = reportIds

	var logList []PixInternalTransactionReportLog.Log

	logs, errorChannel := PixInternalTransactionReportLog.Query(paramsQuery, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			logList = append(logList, log)
		}
	}

	assert.NotEmpty(t, logList)
	for _, log := range logList {
		getLog, err := PixInternalTransactionReportLog.Get(log.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getLog.Id)

		assert.NotNil(t, getLog.Report.Id)
	}
}
