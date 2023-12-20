package pixreversal

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixReversal struct
//
//	PixReversals are instant payments used to revert PixRequests. You can only
//	revert inbound PixRequests.
//	When you initialize a PixReversal, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the slice of created structs.
//
//	Parameters (required):
//	- Amount [int]: Amount in cents to be reversed from the PixRequest. ex: 1234 (= R$ 12.34)
//	- ExternalId [string]: String that must be unique among all your PixReversals. Duplicated external IDs will cause failures. By default, this parameter will block any PixReversal that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//	- EndToEndId [string]: Central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
//	- Reason [string]: Reason why the PixRequest is being reversed. Options are "bankError", "fraud", "chashierError", "customerRequest"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for PixReversals. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixReversal is created. ex: "5656565656565656".
//	- ReturnId [string]: Central bank's unique reversal transaction ID. ex: "D20018183202202030109X3OoBHG74wo".
//	- Fee [string]: Fee charged by this PixReversal. ex: 200 (= R$ 2.00)
//	- Status [string]: Current PixReversal status. ex: "created", "processing", "success", "failed"
//	- Flow [string]: Direction of money flow. ex: "in" or "out"
//	- Created [time.Time]: Creation datetime for the PixReversal. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the PixReversal. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixReversal struct {
	Amount     int        `json:",omitempty"`
	ExternalId string     `json:",omitempty"`
	EndToEndId string     `json:",omitempty"`
	Reason     string     `json:",omitempty"`
	Tags       []string   `json:",omitempty"`
	Id         string     `json:",omitempty"`
	ReturnId   string     `json:",omitempty"`
	Fee        int        `json:",omitempty"`
	Status     string     `json:",omitempty"`
	Flow       string     `json:",omitempty"`
	Created    *time.Time `json:",omitempty"`
	Updated    *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixReversal"}

func Create(reversals []PixReversal, user user.User) ([]PixReversal, Error.StarkErrors) {
	//	Create PixReversals
	//
	//	Send a slice of PixReversal structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- reversals [slice of PixReversal structs]: Slice of PixReversal structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixReversal structs with updated attributes
	create, err := utils.Multi(resource, reversals, nil, user)
	unmarshalError := json.Unmarshal(create, &reversals)
	if unmarshalError != nil {
		return reversals, err
	}
	return reversals, err
}

func Get(id string, user user.User) (PixReversal, Error.StarkErrors) {
	//	Retrieve a specific PixReversal by its id
	//
	//	Receive a single PixReversal struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixReversal struct that corresponds to the given id.
	var pixReversal PixReversal
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixReversal)
	if unmarshalError != nil {
		return pixReversal, err
	}
	return pixReversal, err
}

func Query(params map[string]interface{}, user user.User) chan PixReversal {
	//	Retrieve PixReversals
	//
	//	Receive a channel of PixReversal structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- returnIds [slice of strings, default nil]: Central bank's unique reversal transaction IDs. ex: []string{"D20018183202202030109X3OoBHG74wo", "D20018183202202030109X3OoBHG72rd"].
	//		- externalIds [slice of strings, default nil]: Url safe strings that must be unique among all your PixReversals. Duplicated external IDs will cause failures. By default, this parameter will block any PixReversal that repeats amount and receiver information on the same date. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixReversal structs with updated attributes
	var pixReversal PixReversal
	reversals := make(chan PixReversal)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixReversal)
			if err != nil {
				print(err)
			}
			reversals <- pixReversal
		}
		close(reversals)
	}()
	return reversals
}

func Page(params map[string]interface{}, user user.User) ([]PixReversal, string, Error.StarkErrors) {
	//	Retrieve paged PixReversals
	//
	//	Receive a slice of up to 100 PixReversal structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your reversals.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- returnIds [slice of strings, default nil]: Central bank's unique reversal transaction ID. ex: []string{"D20018183202202030109X3OoBHG74wo", "D20018183202202030109X3OoBHG72rd"].
	//		- externalIds [slice of strings, default nil]: Url safe string that must be unique among all your PixReversals. Duplicated external IDs will cause failures. By default, this parameter will block any PixReversal that repeats amount and receiver information on the same date. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixReversal structs with updated attributes
	//	- cursor to retrieve the next page of PixReversal structs
	var pixReversals []PixReversal
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixReversals)
	if unmarshalError != nil {
		return pixReversals, cursor, err
	}
	return pixReversals, cursor, err
}

func Parse(content string, signature string, user user.User) PixReversal {
	//	Create single verified PixReversal struct from a content string
	//
	//	Create a single PixReversal struct from a content string received from a handler listening at the reversal url.
	//	If the provided digital signature does not check out with the StarkInfra public key, a
	//	starkinfra.error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- parsed PixReversal struct
	var pixReversal PixReversal
	unmarshalError := json.Unmarshal([]byte(utils.ParseAndVerify(content, signature, "", user)), &pixReversal)
	if unmarshalError != nil {
		return pixReversal
	}
	return pixReversal
}

func Response(authorization map[string]interface{}) string {
	//	Helps you respond PixReversal authorization
	//
	//	Parameters (required):
	//	- status [string]: Response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: Denial reason. Options: "invalidAccountNumber", "blockedAccount", "accountClosed", "invalidAccountType", "invalidTransactionType", "taxIdMismatch", "invalidTaxId", "orderRejected", "reversalTimeExpired", "settlementFailed"
	//
	//	Return:
	//	- dumped JSON string that must be returned to us on the PixReversal requests
	params := map[string]map[string]interface{}{
		"authorization": authorization,
	}
	response, _ := json.MarshalIndent(params, "", "  ")
	return string(response)
}
