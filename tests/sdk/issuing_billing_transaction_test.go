package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingBillingInvoice "github.com/starkinfra/sdk-go/starkinfra/issuingbillinginvoice"
	IssuingBillingTransaction "github.com/starkinfra/sdk-go/starkinfra/issuingbillingtransaction"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIssuingBillingTransactionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	transactions, errorChannel := IssuingBillingTransaction.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			assert.NotNil(t, transaction.Id)
		}
	}
}

func TestIssuingBillingTransactionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	transactions, cursor, err := IssuingBillingTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingBillingTransactionQueryParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["after"] = time.Now().AddDate(0, -1, 0)
	params["before"] = time.Now()
	params["tags"] = []string{"travel", "food"}

	invoiceId := firstIssuingBillingInvoiceId(t)
	if invoiceId != "" {
		params["invoiceId"] = invoiceId
	}

	transactions, errorChannel := IssuingBillingTransaction.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case transaction, ok := <-transactions:
			if !ok {
				break loop
			}
			assert.NotNil(t, transaction.Id)
		}
	}
}

func TestIssuingBillingTransactionPageParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["after"] = time.Now().AddDate(0, -1, 0)
	params["before"] = time.Now()
	params["tags"] = []string{"travel", "food"}

	invoiceId := firstIssuingBillingInvoiceId(t)
	if invoiceId != "" {
		params["invoiceId"] = invoiceId
	}

	transactions, cursor, err := IssuingBillingTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingBillingTransactionFields(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	transactions, cursor, err := IssuingBillingTransaction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, cursor)

	for _, transaction := range transactions {
		assert.NotNil(t, transaction.Id)
		assert.IsType(t, 0, transaction.Amount)
		assert.IsType(t, "", transaction.InvoiceId)
		assert.IsType(t, 0, transaction.Installment)
		assert.IsType(t, 0, transaction.InstallmentCount)
		assert.IsType(t, 0, transaction.Balance)
		assert.IsType(t, "", transaction.HolderName)
		assert.IsType(t, "", transaction.Source)
		assert.IsType(t, "", transaction.ExternalId)
		assert.IsType(t, "", transaction.Description)
		assert.IsType(t, "", transaction.CardEnding)
		assert.IsType(t, 0, transaction.Tax)
		assert.IsType(t, float64(0), transaction.Rate)
		assert.IsType(t, 0, transaction.MerchantAmount)
		assert.IsType(t, "", transaction.MerchantCurrencyCode)
		assert.IsType(t, &time.Time{}, transaction.Created)
	}
}

func firstIssuingBillingInvoiceId(t *testing.T) string {

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
			return invoice.Id
		}
	}
	return ""
}
