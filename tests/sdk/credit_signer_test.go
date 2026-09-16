package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
	CreditSigner "github.com/starkinfra/sdk-go/starkinfra/creditsigner"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCreditSignerResendToken(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	notes, err := CreditNote.Create(Example.CreditNote(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	note := notes[0]
	assert.NotEmpty(t, note.Signers)

	// the SCD signer is appended automatically alongside the requested one, and
	// order isn't guaranteed, so pick the requested signer by name instead of index.
	var signerId string
	for _, s := range note.Signers {
		if !strings.HasPrefix(strings.ToLower(strings.TrimSpace(s.Name)), "stark") {
			signerId = s.Id
			break
		}
	}
	if signerId == "" {
		t.Fatal("No signer found")
	}

	signer, err := CreditSigner.ResendToken(signerId, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, signer.Id)
}
