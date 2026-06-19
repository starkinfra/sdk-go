package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingBillingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuingbillinginvoice"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIssuingBillingInvoiceQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	invoices, errorChannel := IssuingBillingInvoice.Query(params, nil)
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

func TestIssuingBillingInvoicePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	invoices, cursor, err := IssuingBillingInvoice.Page(params, nil)
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

func TestIssuingBillingInvoiceGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var invoiceList []IssuingBillingInvoice.IssuingBillingInvoice

	invoices, errorChannel := IssuingBillingInvoice.Query(paramsQuery, nil)
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
		getInvoice, err := IssuingBillingInvoice.Get(invoice.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getInvoice.Id)
		assert.IsType(t, "", getInvoice.TaxId)
		assert.IsType(t, "", getInvoice.Name)
		assert.IsType(t, "", getInvoice.Status)
		assert.IsType(t, "", getInvoice.Brcode)
		assert.IsType(t, "", getInvoice.Link)
		assert.IsType(t, float64(0), getInvoice.Fine)
		assert.IsType(t, float64(0), getInvoice.Interest)
		assert.IsType(t, 0, getInvoice.Amount)
		assert.IsType(t, 0, getInvoice.NominalAmount)
		assert.IsType(t, &time.Time{}, getInvoice.Due)
		assert.IsType(t, &time.Time{}, getInvoice.Start)
		assert.IsType(t, &time.Time{}, getInvoice.End)
		assert.IsType(t, &time.Time{}, getInvoice.Created)
		assert.IsType(t, &time.Time{}, getInvoice.Updated)
	}

	if len(invoiceList) > 0 {
		var idsParams = map[string]interface{}{}
		idsParams["ids"] = []string{invoiceList[0].Id}

		idInvoices, idErrorChannel := IssuingBillingInvoice.Query(idsParams, nil)
	idLoop:
		for {
			select {
			case err := <-idErrorChannel:
				if err.Errors != nil {
					for _, e := range err.Errors {
						t.Errorf("code: %s, message: %s", e.Code, e.Message)
					}
				}
			case invoice, ok := <-idInvoices:
				if !ok {
					break idLoop
				}
				assert.Equal(t, invoiceList[0].Id, invoice.Id)
			}
		}
	}
}

func TestIssuingBillingInvoiceQueryParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = []string{"created"}
	params["after"] = time.Now().AddDate(0, -1, 0)
	params["before"] = time.Now()
	params["tags"] = []string{"travel", "food"}

	invoices, errorChannel := IssuingBillingInvoice.Query(params, nil)
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

func TestIssuingBillingInvoicePageParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["status"] = []string{"created"}
	params["after"] = time.Now().AddDate(0, -1, 0)
	params["before"] = time.Now()
	params["tags"] = []string{"travel", "food"}

	invoices, cursor, err := IssuingBillingInvoice.Page(params, nil)
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
