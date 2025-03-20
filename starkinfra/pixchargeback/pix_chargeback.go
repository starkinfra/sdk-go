package pixchargeback

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixChargeback struct
//
//	A Pix chargeback can be created when fraud is detected on a transaction or a system malfunction
//	results in an erroneous transaction.
//	It notifies another participant of your request to reverse the payment they have received.
//	When you initialize a PixChargeback, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- Amount [int]: Amount in cents to be reversed. ex: 11234 (= R$ 112.34)
//	- ReferenceId [string]: EndToEndId or returnId of the transaction to be reversed. ex: "E20018183202201201450u34sDGd19lz"
//	- Reason [string]: Reason why the reversal was requested. Options: "fraud", "flaw", "reversalChargeback"
//
//	Parameters (optional):
//	- Description [string, default nil]: Description for the PixChargeback.
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixChargeback is created. ex: "5656565656565656"
//	- Analysis [string]: Analysis that led to the result.
//	- SenderBankCode [string]: BankCode of the Pix participant that created the PixChargeback. ex: "20018183"
//  - ReceiverBankCode [string]: BankCode of the Pix participant that received the PixChargeback. ex: "20018183"
//  - RejectionReason [string]: Reason for the rejection of the Pix chargeback. Options: "noBalance", "accountClosed", "invalidRequest", "unableToReverse"
//  - ReversalReferenceId [string]: ReturnId or endToEndId of the reversal transaction. ex: "D20018183202202030109X3OoBHG74wo"
//  - Result [string]: Result after the analysis of the PixChargeback by the receiving party. Options: "rejected", "accepted", "partiallyAccepted"
//  - Flow [string]: Direction of the Pix Chargeback. Options: "in" for received chargebacks, "out" for chargebacks you requested
//  - Status [string]: Current PixChargeback status. Options: "created", "failed", "delivered", "closed", "canceled"
//  - Created [time.Time]: Creation datetime for the PixChargeback. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Updated [time.Time]: Latest update datetime for the PixChargeback. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixChargeback struct {
	Id                  string     `json:",omitempty"`
	Amount              int        `json:",omitempty"`
	ReferenceId         string     `json:",omitempty"`
	Reason              string     `json:",omitempty"`
	Description         string     `json:",omitempty"`
	Tags                []string   `json:",omitempty"`
	Analysis            string     `json:",omitempty"`
	SenderBankCode      string     `json:",omitempty"`
	ReceiverBankCode    string     `json:",omitempty"`
	RejectionReason     string     `json:",omitempty"`
	ReversalReferenceId string     `json:",omitempty"`
	Result              string     `json:",omitempty"`
	Flow                string     `json:",omitempty"`
	Status              string     `json:",omitempty"`
	Created             *time.Time `json:",omitempty"`
	Updated             *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixChargeback"}

func Create(chargebacks []PixChargeback, user user.User) ([]PixChargeback, Error.StarkErrors) {
	//	Create PixChargeback structs
	//
	//	Create PixChargebacks in the Stark Infra API
	//
	//	Parameters (required):
	//	- chargebacks [slice of PixChargeback structs]: Slice of PixChargeback structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixChargeback structs with updated attributes
	create, err := utils.Multi(resource, chargebacks, nil, user)
	unmarshalError := json.Unmarshal(create, &chargebacks)
	if unmarshalError != nil {
		return chargebacks, err
	}
	return chargebacks, err
}

func Get(id string, user user.User) (PixChargeback, Error.StarkErrors) {
	//	Retrieve a PixChargeback struct
	//
	//	Retrieve the PixChargeback struct linked to your Workspace in the Stark Infra API using its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656".
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixChargeback struct that corresponds to the given id.
	var pixChargeback PixChargeback
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixChargeback)
	if unmarshalError != nil {
		return pixChargeback, err
	}
	return pixChargeback, err
}

func Query(params map[string]interface{}, user user.User) chan PixChargeback {
	//	Retrieve PixChargeback structs
	//
	//	Receive a channel of PixChargebacks structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "delivered", "closed", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- flow [string, default nil]: Direction of the Pix Chargeback. Options: "in" for received chargebacks, "out" for chargebacks you requested
	//		- tags [slice of strings, default nil]: Filter for tags of retrieved structs. ex: []string{"travel", "food"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixChargeback structs with updated attributes
	var pixChargeback PixChargeback
	chargebacks := make(chan PixChargeback)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixChargeback)
			if err != nil {
				print(err)
			}
			chargebacks <- pixChargeback
		}
		close(chargebacks)
	}()
	return chargebacks
}

func Page(params map[string]interface{}, user user.User) ([]PixChargeback, string, Error.StarkErrors) {
	//	Retrieve paged PixChargeback structs
	//
	//	Receive a slice of up to 100 PixChareback.Log structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call.
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "delivered", "closed", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- flow [string, default nil]: Direction of the Pix Chargeback. Options: "in" for received chargebacks, "out" for chargebacks you requested
	//		- tags [slice of strings, default nil]: Filter for tags of retrieved structs. ex: []string{"travel", "food"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixChargeback structs with updated attributes
	//	- cursor to retrieve the next page of PixChargeback structs
	var pixChargebacks []PixChargeback
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixChargebacks)
	if unmarshalError != nil {
		return pixChargebacks, cursor, err
	}
	return pixChargebacks, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (PixChargeback, Error.StarkErrors) {
	//	Update PixChargeback entity
	//
	//	Respond to a received PixChargeback.
	//
	//	Parameters (required):
	//	- id [string]: PixChargeback id. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//      Parameters (required):
	//		- result [string]: Result after the analysis of the PixChargeback. Options: "rejected", "accepted", "partiallyAccepted".
	//		Parameters (conditionally required):
	//		- rejectionReason [string, default nil]: If the PixChargeback is rejected a reason is required. Options: "noBalance", "accountClosed", "invalidRequest", "unableToReverse".
	//		- reversalReferenceId [string, default nil]: ReturnId of the reversal transaction. ex: "D20018183202201201450u34sDGd19lz".
	//		- analysis [string, default nil]: Description of the analysis that led to the result. Required if rejection_reason is "invalidRequest".
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixChargeback struct with updated attributes
	var pixChargeback PixChargeback
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &pixChargeback)
	if unmarshalError != nil {
		return pixChargeback, err
	}
	return pixChargeback, err
}

func Cancel(id string, user user.User) (PixChargeback, Error.StarkErrors) {
	//	Cancel a PixChargeback entity
	//
	//	Cancel a PixChargeback entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled PixChargeback struct
	var pixChargeback PixChargeback
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &pixChargeback)
	if unmarshalError != nil {
		return pixChargeback, err
	}
	return pixChargeback, err
}
