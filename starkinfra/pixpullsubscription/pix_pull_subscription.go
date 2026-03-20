package pixpullsubscription

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
	"fmt"
)

//	PixPullSubscription struct
//
//  Pix Pull Subscriptions allow you to set up and manage recurring Pix debits,
//	defining the frequency, amount, and required payer authorizations for each subscription.
//  When you initialize a PixPullSubscription, the entity will not be automatically
//  created in the Stark Infra API. The 'create' function sends the structs
//  to the Stark Infra API and returns the created struct.
//
//  Parameters (required):
//  - BacenId [string]: Central Bank's unique recurrency id.
//  - ExternalId [string]: Safe string that must be unique among all your Pix Pull Subscriptions.
//  - InstallmentStart [time.Time]: Start of settlements allowed for this Pix Pull Subscription.
//  - Interval [string]: Cycle definition of the Pix Pull Requests. ex: "week", "month", "quarter", "semester" or "year"
//  - ReceiverName [string]: Receiver's full name. ex: "Edward Stark"
//  - ReceiverTaxId [string]: Receiver's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//  - SenderAccountNumber [string]: Sender's bank account number. Use '-' before the verifier digit. ex: "876543-2"
//  - SenderBankCode [string]: Sender's bank institution code in Brazil. ex: "20018183"
//  - SenderBranchCode [string]: Sender's bank account branch code. Use '-' in case there is a verifier digit. ex: "1357-9"
//  - SenderTaxId [string]: Sender's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//  - Type [string]: Subscription journey type. #TODO
//
//  Parameters (optional):
//  - Amount [int]: Amount in cents to be transferred in every cycle of the subscription. Required if the subscription has a fixed value.
//  - AmountMinLimit [int]: The floor value for the maximum amount allowed to be set by the sender when approving the subscription.
//  - Description [string]: Additional information to be delivered to the sender.
//  - Due [time.Time]: Due date for answering with an approval or denial.
//  - InstallmentEnd [time.Time]: End of settlements allowed for this Pix Pull Subscription
//  - ReceiverBankCode [string]: Receiver's bank institution code in Brazil.
//  - ReferenceCode [string]: Represents the comercial relation. It can be a contract number, order identification or client code.
//  - PullRetryLimit [int]: Defines if the receiver is able to create Pix Pull Request for retries.
//  - SenderCityCode [string]:
//  - SenderFinalName [string]:
//  - SenderFinalTaxId [string]:
//  - Tags [slice of strings, default nil]: Slice of strings for reference when searching for PixPullSubscriptions. ex: []string{"employees", "monthly"}
//
//  Attributes (return-only):
//  - Id [string]: Unique id returned when the PixPullSubscription is created. ex: "5656565656565656"
//  - Status [string]: Current PixPullSubscription status.
//  - Created [time.Time]: Creation datetime for the PixPullSubscription. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Updated [time.Time]: Latest update datetime for the PixPullSubscription. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Flow [string]: Indicates the flow of the Pix Subscription. ex: "in" or "out"

type PixPullSubscription struct {
	BacenId             string     `json:",omitempty"`
	ExternalId          string     `json:",omitempty"`
	InstallmentStart    *time.Time `json:",omitempty"`
	Interval            string     `json:",omitempty"`
	ReceiverName        string     `json:",omitempty"`
	ReceiverTaxId       string     `json:",omitempty"`
	SenderAccountNumber string     `json:",omitempty"`
	SenderBankCode      string     `json:",omitempty"`
	SenderBranchCode    string     `json:",omitempty"`
	SenderTaxId         string     `json:",omitempty"`
	Type                string     `json:",omitempty"`
	Amount              int        `json:",omitempty"`
	AmountMinLimit      int        `json:",omitempty"`
	Description         string     `json:",omitempty"`
	Due                 *time.Time `json:",omitempty"`
	InstallmentEnd      *time.Time `json:",omitempty"`
	ReceiverBankCode    string     `json:",omitempty"`
	ReferenceCode       string     `json:",omitempty"`
	PullRetryLimit      int        `json:",omitempty"`
	SenderCityCode      string     `json:",omitempty"`
	SenderFinalName     string     `json:",omitempty"`
	SenderFinalTaxId    string     `json:",omitempty"`
	Tags                []string   `json:",omitempty"`
	Id                  string     `json:",omitempty"`
	Status              string     `json:",omitempty"`
	Created             *time.Time `json:",omitempty"`
	Updated             *time.Time `json:",omitempty"`
	Flow                string     `json:",omitempty"`
}

var resource = map[string]string{"name": "PixPullSubscription"}

func Create(subscriptions []PixPullSubscription, user user.User) ([]PixPullSubscription, Error.StarkErrors) {
	//	Create PixPullSubscriptions
	//
	//	Send a slice of PixPullSubscription structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- subscriptions [slice of PixPullSubscription structs]: Slice of PixPullSubscription structs to be created in the API.
	//
	//	Parameters (optional):
	// - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixPullSubscription structs with updated attributes
	create, err := utils.Multi(resource, subscriptions, nil, user)
	jsonStr := string(create)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"due":""`, `"due":null`)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"installmentEnd":""`, `"installmentEnd":null`)
	unmarshalError := json.Unmarshal([]byte(jsonStr), &subscriptions)
	if unmarshalError != nil {
		return subscriptions, err
	}
	return subscriptions, err
}

func Get(id string, user user.User) (PixPullSubscription, Error.StarkErrors) {
	//	Retrieve a specific PixPullSubscription
	//
	//	Receive a single PixPullSubscription struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- PixPullSubscription struct with updated attributes
	var subscription PixPullSubscription
	get, err := utils.Get(resource, id, nil, user)
	jsonStr := string(get)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"due":""`, `"due":null`)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"installmentEnd":""`, `"installmentEnd":null`)
	unmarshalError := json.Unmarshal([]byte(jsonStr), &subscription)
	if unmarshalError != nil {
		return subscription, err
	}
	return subscription, err
}

func Query(params map[string]interface{}, user user.User) (chan PixPullSubscription, chan Error.StarkErrors) {
	//	Retrieve PixPullSubscriptions
	//
	//	Receive a channel of PixPullSubscription structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}]: map containing the attributes to be retrieved.
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [string, default nil]: Filter for structs with a given status. ex: "created", "active", "canceled", "failed"
	//		- tags [slice of strings, default nil]: Filter for structs with a given tag. ex: []string{"employees", "monthly"}
	//      - ids [slice of strings, default nil]: Filter for structs with a given id. ex: []string{"5656565656565656", "4545454545454545"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixPullSubscription structs with updated attributes
	var subscription PixPullSubscription
	subscriptions := make(chan PixPullSubscription)
	subscriptionsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			jsonStr := string(contentByte)
			jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"due":""`, `"due":null`)
			jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"installmentEnd":""`, `"installmentEnd":null`)
			err := json.Unmarshal([]byte(jsonStr), &subscription)
			if err != nil {
				subscriptionsError <- Error.UnknownError(err.Error())
				continue
			}
			subscriptions <- subscription
		}
		for err := range errorChannel {
			subscriptionsError <- err
		}
		close(subscriptions)
		close(subscriptionsError)
	}()
	return subscriptions, subscriptionsError
}

func Page(params map[string]interface{}, user user.User) ([]PixPullSubscription, string, Error.StarkErrors) {
	//	Retrieve paged PixPullSubscription structs
	//
	//	Receive a slice of up to 100 PixPullSubscription structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "active", "canceled", "failed"}
	//		- tags [slice of strings, default nil]: Filter for structs with a given tag. ex: []string{"employees", "monthly"}
	//      - ids [slice of strings, default nil]: Filter for structs with a given id. ex: []string{"5656565656565656", "4545454545454545"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixPullSubscription structs with updated attributes
	//	- cursor to retrieve the next page of PixPullSubscription structs
	var subscriptions []PixPullSubscription
	page, cursor, err := utils.Page(resource, params, user)
	jsonStr := string(page)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"due":""`, `"due":null`)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"installmentEnd":""`, `"installmentEnd":null`)
	unmarshalError := json.Unmarshal([]byte(jsonStr), &subscriptions)
	if unmarshalError != nil {
		return subscriptions, cursor, err
	}
	return subscriptions, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (PixPullSubscription, Error.StarkErrors) {
	//  Update PixPullSubscriptions
	//
	//  A Pix Subscription can be patched for three distinct purposes - to update, confirm or deny it.
	//  As the receiver, you can approve or deny the subscription if the subscription type is "subscritionAndPayment".
	//  As the sender, you can confirm or deny a delivered subscription.
	//
	//  Parameters (required):
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"status": "approved", "senderCityCode": "3550308"}
	//		Parameters (required):
	//		- status [string]: New status of the Pix Subscription.
	//		Parameters (conditionally required):
	//      - senderCityCode [string]: IBGE Code of the payer's city. Required if you are confirming the subscription.
	//		- reason [string]: Reason why the Pix Subscription is being patched. Options: "accountClosed", "accountBlocked", "invalidBranchCode", "notRecognizedBySender", "userRejected", "notOffered"
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixPullSubscription with updated attributes
	var subscription PixPullSubscription
	update, err := utils.Patch(resource, id, patchData, user)
	jsonStr := string(update)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"due":""`, `"due":null`)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"installmentEnd":""`, `"installmentEnd":null`)
	unmarshalError := json.Unmarshal([]byte(jsonStr), &subscription)
	if unmarshalError != nil {
		return subscription, err
	}
	return subscription, err
}

func Cancel(id string, reason string, user user.User) (PixPullSubscription, Error.StarkErrors) {
	//	Cancel a PixPullSubscription
	//
	//	As the receiver, you can also cancel a delivered or confirmed subscription by providing a specific reason:
	//  "accountClosed", "receiverOrganizationClosed", "receiverInternalError", "fraud", "receiverUserRequested".
	//	As the sender, you can cancel a confirmed subscription. The allowed reasons for cancellation are:
	//  "accountClosed", "senderDeceased", "fraud", "senderUserRequested".
	//
	//	Parameters (required):
	//	- id [string]: PixPullSubscription id. ex: '5656565656565656'
	//	- reason [string]: Reason why the Pix Pull Subscription is being canceled. Options: "accountClosed", "receiverOrganizationClosed", "receiverInternalError", "fraud", "receiverUserRequested", "accountClosed", "senderDeceased", "fraud", "senderUserRequested"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled pixPullSubscription struct
	var subscription PixPullSubscription
	data := map[string]interface{}{}
	path := fmt.Sprintf("%v/%v?reason=%v", api.Endpoint(resource), id, reason)
	deleted, err := utils.DeleteRaw(path, user, "", true)
	if err.Errors != nil {
		return PixPullSubscription{}, err
	}

	unmarshalError := json.Unmarshal(deleted.Content, &data)
	if unmarshalError != nil {
		return PixPullSubscription{}, Error.UnknownError(string(deleted.Content))
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])
	jsonStr := string(jsonBytes)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"due":""`, `"due":null`)
	jsonStr = utils.ReplaceEmptyStringField(jsonStr, `"installmentEnd":""`, `"installmentEnd":null`)

	unmarshalError = json.Unmarshal([]byte(jsonStr), &subscription)
	if unmarshalError != nil {
		return PixPullSubscription{}, Error.UnknownError(unmarshalError.Error())
	}
	return subscription, err
}
