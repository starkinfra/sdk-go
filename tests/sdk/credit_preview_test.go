package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditPreview "github.com/starkinfra/sdk-go/starkinfra/creditpreview"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestCreditPreview(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	previews, err := CreditPreview.Create(Example.CreditPreview(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, preview := range previews {
		assert.NotNil(t, preview.Type)
	}
}

func TestCreditPreviewWithNominalAmountOrAmount(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	value := 100000 + rand.Intn(900000)
	types := []string{"sac", "price", "american", "bullet"}

	var previews []CreditPreview.CreditPreview
	var sentAmountOnly []bool
	for _, amountOnly := range []bool{true, false} {
		for _, previewType := range types {
			previews = append(previews, CreditPreview.CreditPreview{
				Type:   "credit-note",
				Credit: amountPreview(previewType, amountOnly, value),
			})
			sentAmountOnly = append(sentAmountOnly, amountOnly)
		}
	}
	sentTypes := make([]string, len(previews))
	for i, preview := range previews {
		sentTypes[i] = preview.Credit.Type
	}

	// Create decodes the response into the slice it receives, so what was sent is recorded above
	received, err := CreditPreview.Create(previews, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	assert.Equal(t, len(sentTypes), len(received))
	for i, preview := range received {
		t.Logf("%+v", preview.Credit)
		assert.Equal(t, sentTypes[i], preview.Credit.Type)
		if sentAmountOnly[i] {
			assert.Equal(t, value, preview.Credit.Amount)
			assert.Greater(t, preview.Credit.NominalAmount, 0)
			continue
		}
		assert.Equal(t, value, preview.Credit.NominalAmount)
		assert.Greater(t, preview.Credit.Amount, 0)
	}
}

func amountPreview(previewType string, amountOnly bool, value int) CreditPreview.CreditNotePreview {
	today := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC)
	scheduled := today.AddDate(0, 0, 11+rand.Intn(9))
	initialDue := today.AddDate(0, 0, 30+rand.Intn(10))

	preview := CreditPreview.CreditNotePreview{
		Type:            previewType,
		Scheduled:       &scheduled,
		TaxId:           Example.Cpf(),
		InitialDue:      &initialDue,
		NominalInterest: 12.5,
	}
	if amountOnly {
		preview.Amount = value
	}
	if !amountOnly {
		preview.NominalAmount = value
	}
	// bullet has a single installment, and the API rejects an interval for it
	if previewType != "bullet" {
		preview.Count = 12
		preview.Interval = "month"
	}
	return preview
}
