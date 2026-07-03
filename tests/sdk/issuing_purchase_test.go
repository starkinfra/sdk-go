package sdk

import (
	"strings"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingPurchase "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingPurchaseQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	purchases, errorChannel := IssuingPurchase.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case purchase, ok := <-purchases:
			if !ok {
				break loop
			}
			assert.NotNil(t, purchase.Id)
		}
	}
}

func TestIssuingPurchasePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	purchases, cursor, err := IssuingPurchase.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, purchase := range purchases {
		assert.NotNil(t, purchase.Id)
	}
	assert.NotNil(t, cursor)
}

func TestIssuingPurchaseGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var purchaseList []IssuingPurchase.IssuingPurchase

	purchases, errorChannel := IssuingPurchase.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case purchase, ok := <-purchases:
			if !ok {
				break loop
			}
			purchaseList = append(purchaseList, purchase)
		}
	}

	for _, purchase := range purchaseList {
		getPurchase, err := IssuingPurchase.Get(purchase.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getPurchase.Id)
	}

	assert.Equal(t, limit, len(purchaseList))
}

func TestIssuingPurchaseParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"acquirerId\": \"236090\", \"amount\": 100, \"cardId\": \"5671893688385536\", \"cardTags\": [], \"endToEndId\": \"2fa7ef9f-b889-4bae-ac02-16749c04a3b6\", \"holderId\": \"5917814565109760\", \"holderTags\": [], \"isPartialAllowed\": false, \"issuerAmount\": 100, \"issuerCurrencyCode\": \"BRL\", \"merchantAmount\": 100, \"merchantCategoryCode\": \"bookStores\", \"merchantCountryCode\": \"BRA\", \"merchantCurrencyCode\": \"BRL\", \"merchantFee\": 0, \"merchantId\": \"204933612653639\", \"merchantName\": \"COMPANY 123\", \"methodCode\": \"token\", \"purpose\": \"purchase\", \"score\": null, \"tax\": 0, \"walletId\": \"\"}"
	validSignature := "MEUCIBxymWEpit50lDqFKFHYOgyyqvE5kiHERi0ZM6cJpcvmAiEA2wwIkxcsuexh9BjcyAbZxprpRUyjcZJ2vBAjdd7o28Q="

	_, err := IssuingPurchase.Parse(content, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
}

func TestIssuingPurchaseParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"acquirerId\": \"236090\", \"amount\": 100, \"cardId\": \"5671893688385536\", \"cardTags\": [], \"endToEndId\": \"2fa7ef9f-b889-4bae-ac02-16749c04a3b6\", \"holderId\": \"5917814565109760\", \"holderTags\": [], \"isPartialAllowed\": false, \"issuerAmount\": 100, \"issuerCurrencyCode\": \"BRL\", \"merchantAmount\": 100, \"merchantCategoryCode\": \"bookStores\", \"merchantCountryCode\": \"BRA\", \"merchantCurrencyCode\": \"BRL\", \"merchantFee\": 0, \"merchantId\": \"204933612653639\", \"merchantName\": \"COMPANY 123\", \"methodCode\": \"token\", \"purpose\": \"purchase\", \"score\": null, \"tax\": 0, \"walletId\": \"\"}"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	_, err := IssuingPurchase.Parse(content, invalidSignature, nil)
	if err.Errors == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestIssuingPurchaseResponseApproved(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var approved = map[string]interface{}{}
	approved["status"] = "approved"
	approved["amount"] = 10000
	approved["tags"] = []string{"tony", "stark"}

	response := IssuingPurchase.Response(approved)
	assert.True(t, strings.Contains(response, "approved"))
}

func TestIssuingPurchaseResponseDenied(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var denied = map[string]interface{}{}
	denied["status"] = "denied"
	denied["reason"] = "other"
	denied["tags"] = []string{"tony", "stark"}

	response := IssuingPurchase.Response(denied)
	assert.True(t, strings.Contains(response, "denied"))
}

func TestIssuingPurchaseNewFields(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	purchase := IssuingPurchase.IssuingPurchase{
		ProductId:            "53810200",
		MerchantCategoryType: "food",
		Description:          "Office Supplies",
		HolderId:             "5917814565109760",
		ZipCode:              "02101234",
		Metadata:             map[string]interface{}{"authorizationId": "OjZAqj"},
		AcquirerId:           "236090",
	}

	assert.Equal(t, "53810200", purchase.ProductId)
	assert.Equal(t, "food", purchase.MerchantCategoryType)
	assert.Equal(t, "Office Supplies", purchase.Description)
	assert.Equal(t, "5917814565109760", purchase.HolderId)
	assert.Equal(t, "02101234", purchase.ZipCode)
	assert.Equal(t, "236090", purchase.AcquirerId)
	assert.Equal(t, "OjZAqj", purchase.Metadata["authorizationId"])
}

func TestIssuingPurchaseParseNewFields(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"acquirerId\": \"236090\", \"amount\": 100, \"cardId\": \"5671893688385536\", \"cardTags\": [], \"description\": \"Office Supplies\", \"endToEndId\": \"2fa7ef9f-b889-4bae-ac02-16749c04a3b6\", \"holderId\": \"5917814565109760\", \"holderTags\": [], \"isPartialAllowed\": false, \"issuerAmount\": 100, \"issuerCurrencyCode\": \"BRL\", \"merchantAmount\": 100, \"merchantCategoryCode\": \"bookStores\", \"merchantCategoryType\": \"food\", \"merchantCountryCode\": \"BRA\", \"merchantCurrencyCode\": \"BRL\", \"merchantFee\": 0, \"merchantId\": \"204933612653639\", \"merchantName\": \"COMPANY 123\", \"metadata\": {\"authorizationId\": \"OjZAqj\"}, \"methodCode\": \"token\", \"productId\": \"53810200\", \"purpose\": \"purchase\", \"score\": null, \"tax\": 0, \"walletId\": \"\", \"zipCode\": \"02101234\"}"
	validSignature := "MEUCIBxymWEpit50lDqFKFHYOgyyqvE5kiHERi0ZM6cJpcvmAiEA2wwIkxcsuexh9BjcyAbZxprpRUyjcZJ2vBAjdd7o28Q="

	parsed, err := IssuingPurchase.Parse(content, validSignature, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.Equal(t, "236090", parsed.AcquirerId)
	assert.Equal(t, "53810200", parsed.ProductId)
	assert.Equal(t, "food", parsed.MerchantCategoryType)
	assert.Equal(t, "Office Supplies", parsed.Description)
	assert.Equal(t, "5917814565109760", parsed.HolderId)
	assert.Equal(t, "02101234", parsed.ZipCode)
	assert.Equal(t, "OjZAqj", parsed.Metadata["authorizationId"])
}
