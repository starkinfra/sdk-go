package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingInvoicePost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	invoice, err := IssuingInvoice.Create(Example.IssuingInvoice(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, invoice.Id)
	fmt.Println(invoice.Id)
}

func TestIssuingInvoiceQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices := IssuingInvoice.Query(params, nil)
	for invoice := range invoices {
		assert.NotNil(t, invoice.Id)
		fmt.Println(invoice.Id)
	}
}

func TestIssuingInvoicePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	invoices, cursor, err := IssuingInvoice.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, invoice := range invoices {
		assert.NotNil(t, invoice.Id)
		fmt.Println(invoice.Id)
	}
	fmt.Println(cursor)
}

func TestIssuingInvoiceGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var invoiceList []IssuingInvoice.IssuingInvoice
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	invoices := IssuingInvoice.Query(paramsQuery, nil)
	for invoice := range invoices {
		invoiceList = append(invoiceList, invoice)
	}

	invoice, err := IssuingInvoice.Get(invoiceList[rand.Intn(len(invoiceList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, invoice.Id)
	fmt.Println(invoice.Id)
}
