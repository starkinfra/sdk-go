package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixReversal "github.com/starkinfra/sdk-go/starkinfra/pixreversal"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestPixReversalPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	reversals, err := PixReversal.Create(Example.PixReversal(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, reversal := range reversals {
		assert.NotNil(t, reversal.Id)
	}
}

func TestPixReversalQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	reversals, errorChannel := PixReversal.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case reversal, ok := <-reversals:
			if !ok {
				break loop
			}
			assert.NotNil(t, reversal.Id)
		}
	}
}

func TestPixReversalPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	reversals, cursor, err := PixReversal.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, reversal := range reversals {
		assert.NotNil(t, reversal.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixReversalInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var reversalList []PixReversal.PixReversal

	reversals, errorChannel := PixReversal.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case reversal, ok := <-reversals:
			if !ok {
				break loop
			}
			reversalList = append(reversalList, reversal)
		}
	}

	for _, reversal := range reversalList {
		getReversal, err := PixReversal.Get(reversal.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getReversal.Id)
	}

	assert.Equal(t, limit, len(reversalList))
}

func TestPixReversalParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"status\": \"processing\", \"returnId\": \"D34052649202212081809BSc6b12oLsF\", \"amount\": 10, \"updated\": \"2022-12-08T18:09:38.344943+00:00\", \"tags\": [\"lannister\", \"chargeback\"], \"reason\": \"fraud\", \"created\": \"2022-12-08T18:09:38.344936+00:00\", \"flow\": \"in\", \"id\": \"5685338043318272\", \"endToEndId\": \"E35547753202201201450oo8srGorhf1\"}"
	validSignature := "MEQCIFiONlW6TV4+U3XWfACP2IttNrxPi8E++FCuXEsf1NjuAiAD2wktgT1tTzxcz+MMJWDPuw3PZjp2kJG+Wf9yF1lcGg=="

	parsed, err := PixReversal.Parse(content, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	assert.NotNil(t, parsed.Id)
}

func TestPixReversalParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"amount\": \"10\", \"external_id\": \"82635892395\", \"end_to_end_id\": \"E20018183202201201450u34sDGd19lz\", \"reason\": \"bankError\", \"tags\": [\"teste\",\"sdk\"], \"senderAccountType\": \"payment\", \"fee\": 0, \"receiverName\": \"Cora\", \"cashierType\": \"\", \"externalId\": \"\", \"method\": \"manual\", \"status\": \"processing\", \"updated\": \"2022-02-16T17:23:53.980250+00:00\", \"description\": \"\", \"tags\": [], \"receiverKeyId\": \"\", \"cashAmount\": 0, \"senderBankCode\": \"20018183\", \"senderBranchCode\": \"0001\", \"bankCode\": \"34052649\", \"senderAccountNumber\": \"5647143184367616\", \"receiverAccountNumber\": \"5692908409716736\", \"initiatorTaxId\": \"\", \"receiverTaxId\": \"34.052.649/0001-78\", \"created\": \"2022-02-16T17:23:53.980238+00:00\", \"flow\": \"in\", \"endToEndId\": \"E20018183202202161723Y4cqxlfLFcm\", \"amount\": 1, \"receiverAccountType\": \"checking\", \"reconciliationId\": \"\", \"receiverBankCode\": \"34052649\"}"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	_, err := PixReversal.Parse(content, invalidSignature, nil)
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestPixReversalResponseApproved(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var approved = map[string]interface{}{}
	approved["status"] = "approved"

	response := PixReversal.Response(approved)
	assert.True(t, strings.Contains(response, "approved"))
}

func TestPixReversalResponseDenied(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var denied = map[string]interface{}{}
	denied["status"] = "denied"
	denied["reason"] = "taxIdMismatch"

	response := PixReversal.Response(denied)
	assert.True(t, strings.Contains(response, "denied"))
}
