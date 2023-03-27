package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixRequest "github.com/starkinfra/sdk-go/starkinfra/pixrequest"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixRequestPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	requests, err := PixRequest.Create(Example.PixRequest(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
		fmt.Println(request.Id)
	}
}

func TestPixRequestQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 50

	requests := PixRequest.Query(params, nil)
	for request := range requests {
		assert.NotNil(t, request.Id)
		fmt.Println(request.Tags)
	}
}

func TestPixRequestPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	requests, cursor, err := PixRequest.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, request := range requests {
		assert.NotNil(t, request.Id)
		fmt.Println(request.Id)
	}
	fmt.Println(cursor)
}

func TestPixRequestInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var requestList []PixRequest.PixRequest
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	requests := PixRequest.Query(paramsQuery, nil)
	for request := range requests {
		requestList = append(requestList, request)
	}

	request, err := PixRequest.Get(requestList[rand.Intn(len(requestList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, request.Id)
	fmt.Println(request.EndToEndId)
}

func TestPixRequestParseRight(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"receiverBranchCode\": \"0001\", \"cashierBankCode\": \"\", \"senderTaxId\": \"20.018.183/0001-80\", \"senderName\": \"Stark Bank S.A. - Instituicao de Pagamento\", \"id\": \"4508348862955520\", \"senderAccountType\": \"payment\", \"fee\": 0, \"receiverName\": \"Cora\", \"cashierType\": \"\", \"externalId\": \"\", \"method\": \"manual\", \"status\": \"processing\", \"updated\": \"2022-02-16T17:23:53.980250+00:00\", \"description\": \"\", \"tags\": [], \"receiverKeyId\": \"\", \"cashAmount\": 0, \"senderBankCode\": \"20018183\", \"senderBranchCode\": \"0001\", \"bankCode\": \"34052649\", \"senderAccountNumber\": \"5647143184367616\", \"receiverAccountNumber\": \"5692908409716736\", \"initiatorTaxId\": \"\", \"receiverTaxId\": \"34.052.649/0001-78\", \"created\": \"2022-02-16T17:23:53.980238+00:00\", \"flow\": \"in\", \"endToEndId\": \"E20018183202202161723Y4cqxlfLFcm\", \"amount\": 1, \"receiverAccountType\": \"checking\", \"reconciliationId\": \"\", \"receiverBankCode\": \"34052649\"}"
	validSignature := "MEUCIQC7FVhXdripx/aXg5yNLxmNoZlehpyvX3QYDXJ8o02X2QIgVwKfJKuIS5RDq50NC/+55h/7VccDkV1vm8Q/7jNu0VM="

	parsed := PixRequest.Parse(content, validSignature, nil)
	fmt.Println(parsed)
}

func TestPixRequestParseWrong(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	content := "{\"receiverBranchCode\": \"0001\", \"cashierBankCode\": \"\", \"senderTaxId\": \"20.018.183/0001-80\", \"senderName\": \"Stark Bank S.A. - Instituicao de Pagamento\", \"id\": \"4508348862955520\", \"senderAccountType\": \"payment\", \"fee\": 0, \"receiverName\": \"Cora\", \"cashierType\": \"\", \"externalId\": \"\", \"method\": \"manual\", \"status\": \"processing\", \"updated\": \"2022-02-16T17:23:53.980250+00:00\", \"description\": \"\", \"tags\": [], \"receiverKeyId\": \"\", \"cashAmount\": 0, \"senderBankCode\": \"20018183\", \"senderBranchCode\": \"0001\", \"bankCode\": \"34052649\", \"senderAccountNumber\": \"5647143184367616\", \"receiverAccountNumber\": \"5692908409716736\", \"initiatorTaxId\": \"\", \"receiverTaxId\": \"34.052.649/0001-78\", \"created\": \"2022-02-16T17:23:53.980238+00:00\", \"flow\": \"in\", \"endToEndId\": \"E20018183202202161723Y4cqxlfLFcm\", \"amount\": 1, \"receiverAccountType\": \"checking\", \"reconciliationId\": \"\", \"receiverBankCode\": \"34052649\"}"
	invalidSignature := "MEUCIQDOpo1j+V40DNZK2URL2786UQK/8mDXon9ayEd8U0/l7AIgYXtIZJBTs8zCRR3vmted6Ehz/qfw1GRut/eYyvf1yOk="

	parsed := PixRequest.Parse(content, invalidSignature, nil)
	fmt.Println(parsed)
}

func TestPixRequestResponseApproved(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var approved = map[string]interface{}{}
	approved["status"] = "approved"

	response := PixRequest.Response(approved)
	fmt.Println(response)
}

func TestPixRequestResponseDenied(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var denied = map[string]interface{}{}
	denied["status"] = "denied"
	denied["reason"] = "taxIdMismatch"

	response := PixRequest.Response(denied)
	fmt.Println(response)
}
