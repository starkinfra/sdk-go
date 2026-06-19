package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	IssuingTokenRequest "github.com/starkinfra/sdk-go/starkinfra/issuingtokenrequest"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingTokenRequestPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	cardId := firstIssuingCardId(t)
	if cardId == "" {
		return
	}

	request := IssuingTokenRequest.IssuingTokenRequest{
		CardId:     cardId,
		WalletId:   "google",
		MethodCode: "app",
	}

	tokenRequest, err := IssuingTokenRequest.Create(request, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
		return
	}

	assert.NotNil(t, tokenRequest.Content)
	assert.IsType(t, "", tokenRequest.Content)
	assert.IsType(t, "", tokenRequest.Signature)
	assert.IsType(t, map[string]interface{}{}, tokenRequest.Metadata)
}

func firstIssuingCardId(t *testing.T) string {

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	cards, errorChannel := IssuingCard.Query(params, nil)
loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			return card.Id
		}
	}
	return ""
}
