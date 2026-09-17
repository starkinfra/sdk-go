package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingTokenActivation "github.com/starkinfra/sdk-go/starkinfra/issuingtokenactivation"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingTokenActivationParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"activationMethod\": {\"type\": \"text\", \"value\": \"** *****-5678\"}, \"tokenId\": \"5585821789122165\", \"tags\": [\"token\", \"user/1234\"], \"cardId\": \"5189831499972623\"}"
	validSignature := "MEUCIAxn0FmsPWI4r3Y7Nq8xFNQHYZgo0QAGDQ4/7CajKoVuAiEA09kXWrPMhsw4JbgC3pmNccCWr+hidfop/KsSNqza0yE="

	activation, err := IssuingTokenActivation.Parse(content, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, activation.CardId)
	assert.NotNil(t, activation.TokenId)
}

func TestIssuingTokenActivationParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"activationMethod\": {\"type\": \"text\", \"value\": \"** *****-5678\"}, \"tokenId\": \"5585821789122165\", \"tags\": [\"token\", \"user/1234\"], \"cardId\": \"5189831499972623\"}"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	_, err := IssuingTokenActivation.Parse(content, invalidSignature, nil)
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}
