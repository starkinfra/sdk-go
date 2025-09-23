package pixrequest

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixRequest struct
//
//	PixRequests are used to receive or send instant payments to accounts
//	hosted in any Pix participant.
//	When you initialize a PixRequest, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the slice of created structs.
//
//	Parameters (required):
//	- Amount [int]: Amount in cents to be transferred. ex: 11234 (= R$ 112.34)
//	- ExternalId [string]: String that must be unique among all your PixRequests. Duplicated external IDs will cause failures. By default, this parameter will block any PixRequests that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//	- SenderName [string]: Sender's full name. ex: "Edward Stark"
//	- SenderTaxId [string]: Sender's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- SenderBranchCode [string]: Sender's bank account branch code. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- SenderAccountNumber [string]: Sender's bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- SenderAccountType [string]: Sender's bank account type. ex: "checking", "savings", "salary" or "payment"
//	- ReceiverName [string]: Receiver's full name. ex: "Edward Stark"
//	- ReceiverTaxId [string]: Receiver's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- ReceiverBankCode [string]: Receiver's bank institution code in Brazil. ex: "20018183"
//	- ReceiverAccountNumber [string]: Receiver's bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- ReceiverBranchCode [string]: Receiver's bank account branch code. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- ReceiverAccountType [string]: Receiver's bank account type. ex: "checking", "savings", "salary" or "payment"
//	- EndToEndId [string]: Central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
//
//	Parameters (optional):
//	- ReceiverKeyId [string, default nil]: Receiver's dict key. ex: "20.018.183/0001-80"
//	- Description [string, default nil]: Optional description to override default description to be shown in the bank statement. ex: "Payment for service #1234"
//	- ReconciliationId [string, default nil]: Reconciliation ID linked to this payment. ex: "b77f5236-7ab9-4487-9f95-66ee6eaf1781"
//	- InitiatorTaxId [string, default nil]: Payment initiator's tax id (CPF/CNPJ). ex: "01234567890" or "20.018.183/0001-80"
//	- CashAmount [int, default nil]: Amount to be withdrawal from the cashier in cents. ex: 1000 (= R$ 10.00)
//	- CashierBankCode [string, default nil]: Cashier's bank code. ex: "00000000"
//	- CashierType [string, default nil]: Cashier's type. ex: []string{merchant, other, participant]
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for PixRequests. ex: []string{"employees", "monthly"}
//	- Method [string, default nil]: Execution  method for the creation of the Pix. ex: "manual", "payerQrcode", "dynamicQrcode".
//	- Priority [string, default "high"]: Defines the channel through which the entities will be processed. Options: "low", "high"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixRequest is created. ex: "5656565656565656"
//	- Fee [int]: Fee charged when PixRequest is paid. ex: 200 (= R$ 2.00)
//	- Status [string]: Current PixRequest status. ex: "created", "processing", "success", "failed"
//	- Flow [string]: Direction of money flow. ex: "in" or "out"
//	- SenderBankCode [string]: Sender's bank institution code in Brazil. ex: "20018183"
//	- Created [time.Time]: Creation datetime for the PixRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the PixRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixRequest struct {
	Amount                int        `json:",omitempty"`
	ExternalId            string     `json:",omitempty"`
	SenderName            string     `json:",omitempty"`
	SenderTaxId           string     `json:",omitempty"`
	SenderBranchCode      string     `json:",omitempty"`
	SenderAccountNumber   string     `json:",omitempty"`
	SenderAccountType     string     `json:",omitempty"`
	ReceiverName          string     `json:",omitempty"`
	ReceiverTaxId         string     `json:",omitempty"`
	ReceiverBankCode      string     `json:",omitempty"`
	ReceiverAccountNumber string     `json:",omitempty"`
	ReceiverBranchCode    string     `json:",omitempty"`
	ReceiverAccountType   string     `json:",omitempty"`
	EndToEndId            string     `json:",omitempty"`
	ReceiverKeyId         string     `json:",omitempty"`
	Description           string     `json:",omitempty"`
	ReconciliationId      string     `json:",omitempty"`
	InitiatorTaxId        string     `json:",omitempty"`
	CashAmount            int        `json:",omitempty"`
	CashierBankCode       string     `json:",omitempty"`
	CashierType           string     `json:",omitempty"`
	Tags                  []string   `json:",omitempty"`
	Method                string     `json:",omitempty"`
	Priority              string     `json:",omitempty"`
	Id                    string     `json:",omitempty"`
	Fee                   int        `json:",omitempty"`
	Status                string     `json:",omitempty"`
	Flow                  string     `json:",omitempty"`
	SenderBankCode        string     `json:",omitempty"`
	Created               *time.Time `json:",omitempty"`
	Updated               *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixRequest"}

func Create(requests []PixRequest, user user.User) ([]PixRequest, Error.StarkErrors) {
	//	Create PixRequests
	//
	//	Send a slice of PixRequest structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- requests [slice of PixRequest structs]: Slice of PixRequest structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return requests, err
	}
	return requests, err
}

func Get(id string, user user.User) (PixRequest, Error.StarkErrors) {
	//	Retrieve a specific PixRequest
	//
	//	Receive a single PixRequest struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixRequest struct with updated attributes
	var pixRequest PixRequest
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixRequest)
	if unmarshalError != nil {
		return pixRequest, err
	}
	return pixRequest, err
}

func Query(params map[string]interface{}, user user.User) chan PixRequest {
	//	Retrieve PixRequests
	//
	//	Receive a channel of PixRequest structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- endToEndIds [slice of strings, default nil]: Central bank's unique transaction IDs. ex: []string{"E79457883202101262140HHX553UPqeq", "E79457883202101262140HHX553UPxzx"}
	//		- externalIds [slice of strings, default nil]: Url safe strings that must be unique among all your PixRequests. Duplicated external IDs will cause failures. By default, this parameter will block any PixRequests that repeats amount and receiver information on the same date. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixRequest structs with updated attributes
	var pixRequest PixRequest
	requests := make(chan PixRequest)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixRequest)
			if err != nil {
				print(err)
			}
			requests <- pixRequest
		}
		close(requests)
	}()
	return requests
}

func Page(params map[string]interface{}, user user.User) ([]PixRequest, string, Error.StarkErrors) {
	//	Retrieve paged PixRequest structs
	//
	//	Receive a slice of up to 100 PixRequest structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- endToEndIds [slice of strings, default nil]: Central bank's unique transaction IDs. ex: []string{"E79457883202101262140HHX553UPqeq", "E79457883202101262140HHX553UPxzx"}
	//		- externalIds [slice of strings, default nil]: Url safe strings that must be unique among all your PixRequests. Duplicated external IDs will cause failures. By default, this parameter will block any PixRequests that repeats amount and receiver information on the same date. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixRequest structs with updated attributes
	//	- cursor to retrieve the next page of PixRequest structs
	var pixRequests []PixRequest
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixRequests)
	if unmarshalError != nil {
		return pixRequests, cursor, err
	}
	return pixRequests, cursor, err
}

func Parse(content string, signature string, user user.User) PixRequest {
	//	Create single verified PixRequest struct from a content string
	//
	//	Create a single PixRequest struct from a content string received from a handler listening at the request url.
	//	If the provided digital signature does not check out with the StarkInfra public key, a
	//	error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- parsed PixRequest struct
	var pixRequest PixRequest
	unmarshalError := json.Unmarshal([]byte(utils.ParseAndVerify(content, signature, "", user)), &pixRequest)
	if unmarshalError != nil {
		return pixRequest
	}
	return pixRequest
}

func Response(authorization map[string]interface{}) string {
	//	Helps you respond PixRequests
	//
	//	Parameters (required):
	//	- status [string]: Response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: Denial reason. Options: "invalidAccountNumber", "blockedAccount", "accountClosed", "invalidAccountType", "invalidTransactionType", "taxIdMismatch", "invalidTaxId", "orderRejected", "reversalTimeExpired", "settlementFailed"
	//
	//	Return:
	//	- dumped JSON string that must be returned to us on the PixRequest
	params := map[string]map[string]interface{}{
		"authorization": authorization,
	}
	response, _ := json.MarshalIndent(params, "", "  ")
	return string(response)
}
