package pixpullrequest

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"time"
)

//	PixPullRequest struct
//
//	A Pix Pull Request is a command sent to the payer's bank to trigger the automatic
//	debit linked to an active Pix Pull Subscription. It confirms the receiver’s intent
//	to collect the agreed amount within the current billing cycle and initiates the
//	settlement process through the Pix infrastructure.
//	When you initialize a PixPullRequest, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- Amount [int]: Amount to be charged.
//	- Due [time.Time]: Due date for answering with an approval or denial. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- EndToEndId [string]: Central Bank's unique transaction id. Example: endToEndId="E00002649202201172211u34srod19le"
//	- ReceiverAccountNumber [string]: Receiver's bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- ReceiverAccountType [string]: Receiver's account type. Options: "checking", "savings", "salary" and "payment"
//	- ReceiverBankCode [string]: Receiver's bank code.
//	- ReconciliationId [string]: Id to be used for conciliation of the resulting Pix transaction. Up to 25 alphanumeric characters. Example: reconciliationId = "123456"
//	- SubscriptionId [string]: Unique ID of the Pix Pull Subscription.
//
//	Parameters (optional):
//	- AttemptType [string]: Defines the type of attempt for the Pix Pull Request. Options: "default", "instantRetry" and "scheduledRetry".
//	- Description [string]: Additional information to be delivered to the sender.
//	- ReceiverBranchCode [string]: Receiver's branch code.
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for PixPullRequests. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixPullRequest is created. ex: "5656565656565656"
//	- Status [string]: Current PixPullRequest status.
//	- Created [time.Time]: Creation datetime for the PixPullRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the PixPullRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Flow [string]: Direction of money flow. ex: "in" or "out"
//	- ReceiverTaxId [string]: Receiver's tax ID (CPF or CNPJ).
//	- SenderBankCode [string]: Sender's bank institution code in Brazil.
//	- SenderFinalName [string]: Sender's final name.
//	- SenderTaxId [string]: Sender's tax ID (CPF or CNPJ).

type PixPullRequest struct {
	Amount                int        `json:",omitempty"`
	Due                   *time.Time `json:",omitempty"`
	EndToEndId            string     `json:",omitempty"`
	ReceiverAccountNumber string     `json:",omitempty"`
	ReceiverAccountType   string     `json:",omitempty"`
	ReceiverBankCode      string     `json:",omitempty"`
	ReconciliationId      string     `json:",omitempty"`
	SubscriptionId        string     `json:",omitempty"`
	AttemptType           string     `json:",omitempty"`
	Description           string     `json:",omitempty"`
	ReceiverBranchCode    string     `json:",omitempty"`
	Tags                  []string   `json:",omitempty"`
	Id                    string     `json:",omitempty"`
	Status                string     `json:",omitempty"`
	Created               *time.Time `json:",omitempty"`
	Updated               *time.Time `json:",omitempty"`
	Flow                  string     `json:",omitempty"`
	ReceiverTaxId         string     `json:",omitempty"`
	SenderBankCode        string     `json:",omitempty"`
	SenderFinalName       string     `json:",omitempty"`
	SenderTaxId           string     `json:",omitempty"`
}

var resource = map[string]string{"name": "PixPullRequest"}

func Create(requests []PixPullRequest, user user.User) ([]PixPullRequest, Error.StarkErrors) {
	//	Create PixPullRequests
	//
	//	Send a slice of PixPullRequest structs for creation in the Stark Infra API
	//
	//	Parameters (required):
	//	- requests [slice of PixPullRequest structs]: Slice of PixPullRequest structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixPullRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return requests, err
	}
	return requests, err
}

func Get(id string, user user.User) (PixPullRequest, Error.StarkErrors) {
	//	Retrieve a specific PixPullRequest
	//
	//	Receive a single PixPullRequest struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- PixPullRequest struct with updated attributes
	var request PixPullRequest
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &request)
	if unmarshalError != nil {
		return request, err
	}
	return request, err
}

func Query(params map[string]interface{}, user user.User) (chan PixPullRequest, chan Error.StarkErrors) {
	//	Retrieve PixPullRequests
	//
	//	Receive a channel of PixPullRequest structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}]: map containing the attributes to be retrieved.
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [string, default nil]: Filter for structs with a given status. ex: "created", "active", "canceled", "failed"
	//		- tags [slice of strings, default nil]: Filter for structs with a given tag. ex: []string{"employees", "monthly"}
	//      - ids [slice of strings, default nil]: Filter for structs with a given id. ex: []string{"5656565656565656", "4545454545454545"}
	//		- flow [string, default nil]: String to filter Pix Pull Requests by the specific flow. Options: "in", "out"
	//		- subscriptionIds [slice of strings, default nil]: Strings to filter Pix Pull Requets by the subscriptionIds. ex: []string{"5656565656565656", "4545454545454545"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixPullRequest structs with updated attributes
	var request PixPullRequest
	requests := make(chan PixPullRequest)
	requestsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &request)
			if err != nil {
				requestsError <- Error.UnknownError(err.Error())
				continue
			}
			requests <- request
		}
		for err := range errorChannel {
			requestsError <- err
		}
		close(requests)
		close(requestsError)
	}()
	return requests, requestsError
}

func Page(params map[string]interface{}, user user.User) ([]PixPullRequest, string, Error.StarkErrors) {
	//	Retrieve paged PixPullRequest structs
	//
	//	Receive a slice of up to 100 PixPullRequest structs previously created in the Stark Infra API and the cursor to the next page.
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
	//		- flow [string, default nil]: String to filter Pix Pull Requests by the specific flow. Options: "in", "out"
	//		- subscriptionIds [slice of strings, default nil]: Strings to filter Pix Pull Requets by the subscriptionIds. ex: []string{"5656565656565656", "4545454545454545"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixPullRequest structs with updated attributes
	//	- cursor to retrieve the next page of PixPullRequest structs
	var requests []PixPullRequest
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &requests)
	if unmarshalError != nil {
		return requests, cursor, err
	}
	return requests, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (PixPullRequest, Error.StarkErrors) {
	//  Update PixPullRequests
	//
	//	A Pix Pull Request can be updated to change the status to "scheduled" or "denied"
	// 
	//  Parameters (required):
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"status": "approved", "senderCityCode": "3550308"}
	//		Parameters (required):
	//		- status [string]: New status of the Pix Pull Request.
	//		Parameters (conditionally required):
	//		- reason [string]: Reason why the Pix Pull Request is being denied. Options: "senderAccountClosed", "senderAccountBLocked", "amountNotAllowed"
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixPullRequest with updated attributes
	var subscription PixPullRequest
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &subscription)
	if unmarshalError != nil {
		return subscription, err
	}
	return subscription, err
}

func Cancel(id string, reason string, user user.User) (PixPullRequest, Error.StarkErrors) {
	//	Cancel a PixPullRequest
	//
	//	As the receiver, you can also cancel a delivered or confirmed request by providing a specific reason:
	//  "accountClosed", "receiverOrganizationClosed", "receiverInternalError", "fraud", "receiverUserRequested".
	//	As the sender, you can cancel a confirmed request. The allowed reasons for cancellation are:
	//  "accountClosed", "senderDeceased", "fraud", "senderUserRequested".
	//
	//	Parameters (required):
	//	- id [string]: PixPullRequest id. ex: '5656565656565656'
	//	- reason [string]: Reason why the Pix Pull Request is being canceled. Options: "accountClosed", "receiverOrganizationClosed", "receiverInternalError", "fraud", "receiverUserRequested", "accountClosed", "senderDeceased", "fraud", "senderUserRequested"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled pixPullRequest struct
	var request PixPullRequest
	data := map[string]interface{}{}
	path := fmt.Sprintf("%v/%v?reason=%v", api.Endpoint(resource), id, reason)
	deleted, err := utils.DeleteRaw(path, user, "", true)
	if err.Errors != nil {
		return PixPullRequest{}, err
	}

	unmarshalError := json.Unmarshal(deleted.Content, &data)
	if unmarshalError != nil {
		return PixPullRequest{}, Error.UnknownError(string(deleted.Content))
	}
	jsonBytes, _ := json.Marshal(data[api.LastName(resource)])

	unmarshalError = json.Unmarshal(jsonBytes, &request)
	if unmarshalError != nil {
		return PixPullRequest{}, Error.UnknownError(unmarshalError.Error())
	}
	return request, err
}
