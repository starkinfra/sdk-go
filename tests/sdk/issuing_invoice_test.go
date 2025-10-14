package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingInvoicePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	invoice, err := IssuingInvoice.Create(Example.IssuingInvoice(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, invoice.Id)
}

func TestIssuingInvoiceQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	invoices, errorChannel := IssuingInvoice.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			assert.NotNil(t, invoice.Id)
		}
	}
}

func TestIssuingInvoicePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	invoices, cursor, err := IssuingInvoice.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, invoice := range invoices {
		assert.NotNil(t, invoice.Id)
	}
	assert.NotNil(t, cursor)
}

func TestIssuingInvoiceGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var invoiceList []IssuingInvoice.IssuingInvoice

	invoices, errorChannel := IssuingInvoice.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case invoice, ok := <-invoices:
			if !ok {
				break loop
			}
			invoiceList = append(invoiceList, invoice)
		}
	}

	for _, invoice := range invoiceList {
		getInvoice, err := IssuingInvoice.Get(invoice.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getInvoice.Id)
	}
	assert.Equal(t, limit, len(invoiceList))
}
