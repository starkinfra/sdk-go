package creditnote_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"testing"

	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/sdk-go/starkinfra"
	CreditNote "github.com/starkinfra/sdk-go/starkinfra/creditnote"
	"github.com/stretchr/testify/assert"
)

const testPrivateKey = "-----BEGIN EC PRIVATE KEY-----\nMHQCAQEEIBEcEJZLk/DyuXVsEjz0w4vrE7plPXhQxODvcG1Jc0WToAcGBSuBBAAK\noUQDQgAE6t4OGx1XYktOzH/7HV6FBukxq0Xs2As6oeN6re1Ttso2fwrh5BJXDq75\nmSYHeclthCRgU8zl6H1lFQ4BKZ5RCQ==\n-----END EC PRIVATE KEY-----"

type recordingTransport struct {
	requestBody []byte
	response    string
}

func (transport *recordingTransport) RoundTrip(request *http.Request) (*http.Response, error) {
	body, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	transport.requestBody = body
	return &http.Response{
		StatusCode: 200,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(bytes.NewReader([]byte(transport.response))),
		Request:    request,
	}, nil
}

// Fetch builds its own http.Client, which falls back to http.DefaultTransport,
// so swapping it intercepts the real SDK/core path without network.
func createWithTransport(t *testing.T, response string, note CreditNote.CreditNote) ([]byte, []CreditNote.CreditNote) {
	transport := &recordingTransport{response: response}
	originalTransport := http.DefaultTransport
	http.DefaultTransport = transport
	t.Cleanup(func() { http.DefaultTransport = originalTransport })

	user := project.Project{Id: "5656565656565656", PrivateKey: testPrivateKey, Environment: "sandbox"}
	notes, err := CreditNote.Create([]CreditNote.CreditNote{note}, user)
	assert.Nil(t, err.Errors)
	return transport.requestBody, notes
}

func sentNote(t *testing.T, body []byte) map[string]interface{} {
	var payload map[string][]map[string]interface{}
	assert.NoError(t, json.Unmarshal(body, &payload))
	assert.Len(t, payload["notes"], 1)
	return payload["notes"][0]
}

func TestCreditNoteTemplateReplacementsSentWithKeysIntact(t *testing.T) {
	starkinfra.User = nil
	replacements := map[string]interface{}{
		"guarantorName":    "Jon Snow",
		"guarantorTaxId":   "012.345.678-90",
		"guarantorAddress": "Av. Paulista, 200 - Bela Vista - Sao Paulo, SP",
	}

	body, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{
		Name:                 "Edward Stark",
		TemplateReplacements: replacements,
	})

	assert.Equal(t, replacements, sentNote(t, body)["templateReplacements"])
}

func TestCreditNoteTemplateReplacementsOmittedWhenAbsent(t *testing.T) {
	starkinfra.User = nil

	nilBody, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{Name: "Edward Stark"})
	emptyBody, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{
		Name:                 "Edward Stark",
		TemplateReplacements: map[string]interface{}{},
	})

	assert.NotContains(t, sentNote(t, nilBody), "templateReplacements")
	assert.NotContains(t, sentNote(t, emptyBody), "templateReplacements")
}

func TestCreditNoteTemplateReplacementsNilValueIsNotSent(t *testing.T) {
	starkinfra.User = nil

	body, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{
		Name:                 "Edward Stark",
		TemplateReplacements: map[string]interface{}{"guarantorName": "Jon Snow", "guarantorNickname": nil},
	})

	assert.Equal(
		t,
		map[string]interface{}{"guarantorName": "Jon Snow"},
		sentNote(t, body)["templateReplacements"],
	)
}

func TestCreditNoteTemplateReplacementsReadFromResponse(t *testing.T) {
	starkinfra.User = nil
	response := `{"notes": [{"id": "5656565656565656", "templateReplacements": {"guarantorName": "Jon Snow", "guarantorTaxId": "012.345.678-90"}}]}`

	_, notes := createWithTransport(t, response, CreditNote.CreditNote{Name: "Edward Stark"})

	assert.Len(t, notes, 1)
	assert.IsType(t, CreditNote.TemplateReplacements{}, notes[0].TemplateReplacements)
	assert.Equal(
		t,
		CreditNote.TemplateReplacements{"guarantorName": "Jon Snow", "guarantorTaxId": "012.345.678-90"},
		notes[0].TemplateReplacements,
	)
}

// A plain map keeps working: the named type has the same underlying type, so existing callers still compile
// and send the same payload.
func TestCreditNoteTemplateReplacementsAcceptsPlainMapAndNamedType(t *testing.T) {
	starkinfra.User = nil
	plain := map[string]interface{}{"guarantorName": "Jon Snow", "guarantorNickname": nil}
	named := CreditNote.TemplateReplacements{"guarantorName": "Jon Snow", "guarantorNickname": nil}

	plainBody, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{
		Name:                 "Edward Stark",
		TemplateReplacements: plain,
	})
	namedBody, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{
		Name:                 "Edward Stark",
		TemplateReplacements: named,
	})

	assert.Equal(
		t,
		map[string]interface{}{"guarantorName": "Jon Snow"},
		sentNote(t, plainBody)["templateReplacements"],
	)
	assert.Equal(t, sentNote(t, plainBody)["templateReplacements"], sentNote(t, namedBody)["templateReplacements"])
}

// Pins the key rewriting described in the field docs and the README, so they cannot drift from the core.
func TestCreditNoteTemplateReplacementsKeysAreRewrittenByCore(t *testing.T) {
	starkinfra.User = nil
	keys := map[string]string{
		"endereçoAvalista": "endereoAvalista",
		"avalista2nome":    "avalista2Nome",
		"CPFAvalista":      "cPFAvalista",
		"cpf/cnpj":         "cpfcnpj",
		"guarantor_name":   "guarantorName",
		"guarantor-tax.id": "guarantorTaxId",
		"guarantor name":   "guarantorName",
	}

	for key, expected := range keys {
		body, _ := createWithTransport(t, `{"notes": [{"id": "1"}]}`, CreditNote.CreditNote{
			TemplateReplacements: map[string]interface{}{key: "value"},
		})

		assert.Equal(
			t,
			map[string]interface{}{expected: "value"},
			sentNote(t, body)["templateReplacements"],
			key,
		)
	}
}
