package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingToken "github.com/starkinfra/sdk-go/starkinfra/issuingtoken"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

func TestIssuingTokenQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	tokens, errorChannel := IssuingToken.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case token, ok := <-tokens:
			if !ok {
				break loop
			}
			assert.NotNil(t, token.Id)
		}
	}
}

func TestIssuingTokenPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	tokens, cursor, err := IssuingToken.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, token := range tokens {
		assert.NotNil(t, token.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingTokenQueryParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["after"] = time.Now().AddDate(0, -1, 0)
	params["before"] = time.Now()
	params["status"] = []string{"active"}
	params["tags"] = []string{"travel", "food"}

	tokens, errorChannel := IssuingToken.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case token, ok := <-tokens:
			if !ok {
				break loop
			}
			assert.NotNil(t, token.Id)
		}
	}
}

func TestIssuingTokenPageParams(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit
	params["after"] = time.Now().AddDate(0, -1, 0)
	params["before"] = time.Now()
	params["status"] = []string{"active"}
	params["tags"] = []string{"travel", "food"}

	tokens, cursor, err := IssuingToken.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, token := range tokens {
		assert.NotNil(t, token.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingTokenGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	tokenId := firstIssuingTokenId(t)
	if tokenId == "" {
		return
	}

	token, err := IssuingToken.Get(tokenId, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, token.Id)
	assert.IsType(t, "", token.CardId)
	assert.IsType(t, "", token.WalletId)
	assert.IsType(t, "", token.WalletName)
	assert.IsType(t, "", token.MerchantId)
	assert.IsType(t, "", token.ExternalId)
	assert.IsType(t, []string{}, token.Tags)
	assert.IsType(t, "", token.Status)
	assert.IsType(t, &time.Time{}, token.Updated)
	assert.IsType(t, &time.Time{}, token.Created)
}

func TestIssuingTokenUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	tokenId := firstIssuingTokenId(t)
	if tokenId == "" {
		return
	}

	var patchData = map[string]interface{}{}
	patchData["status"] = "blocked"

	token, err := IssuingToken.Update(tokenId, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, token.Id)
}

func TestIssuingTokenCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	tokenId := firstIssuingTokenId(t)
	if tokenId == "" {
		return
	}

	token, err := IssuingToken.Cancel(tokenId, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, token.Id)
}

func TestIssuingTokenParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"cardId\": \"5656565656565656\", \"walletId\": \"google\", \"id\": \"5656565656565656\", \"status\": \"pending\", \"methodCode\": \"app\"}"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	_, err := IssuingToken.Parse(content, invalidSignature, nil)
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestIssuingTokenResponseAuthorizationApproved(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var approved = map[string]interface{}{}
	approved["status"] = "approved"

	response := IssuingToken.ResponseAuthorization(approved)
	assert.True(t, strings.Contains(response, "approved"))
}

func TestIssuingTokenResponseAuthorizationDenied(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var denied = map[string]interface{}{}
	denied["status"] = "denied"
	denied["reason"] = "lostCard"

	response := IssuingToken.ResponseAuthorization(denied)
	assert.True(t, strings.Contains(response, "denied"))
}

func TestIssuingTokenResponseActivationApproved(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var approved = map[string]interface{}{}
	approved["status"] = "approved"

	response := IssuingToken.ResponseActivation(approved)
	assert.True(t, strings.Contains(response, "approved"))
}

func TestIssuingTokenResponseActivationDenied(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var denied = map[string]interface{}{}
	denied["status"] = "denied"
	denied["reason"] = "lostCard"

	response := IssuingToken.ResponseActivation(denied)
	assert.True(t, strings.Contains(response, "denied"))
}

func firstIssuingTokenId(t *testing.T) string {

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	tokens, errorChannel := IssuingToken.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case token, ok := <-tokens:
			if !ok {
				break loop
			}
			return token.Id
		}
	}
	return ""
}
