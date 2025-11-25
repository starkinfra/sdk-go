package issuingtransaction

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingTransaction struct
//
//	The IssuingTransaction structs created in your Workspace to represent each balance shift.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingTransaction is created. ex: "5656565656565656"
//	- Amount [int]: IssuingTransaction value in cents. ex: 1234 (= R$ 12.34)
//	- Balance [int]: Balance amount of the Workspace at the instant of the Transaction in cents. ex: 200 (= R$ 2.00)
//	- Description [string]: IssuingTransaction description. ex: "Buying food"
//	- Source [string]: Source of the transaction. ex: "issuing-purchase/5656565656565656"
//	- Tags [string]: Slice of strings inherited from the source resource. ex: []string{"tony", "stark"}
//	- Created [time.Time]: Creation datetime for the IssuingTransaction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingTransaction struct {
	Id          string     `json:",omitempty"`
	Amount      int        `json:",omitempty"`
	Balance     int        `json:",omitempty"`
	Description string     `json:",omitempty"`
	Source      string     `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingTransaction"}

func Get(id string, user user.User) (IssuingTransaction, Error.StarkErrors) {
	//	Retrieve a specific IssuingTransaction by its id
	//
	//	Receive a single IssuingTransaction struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingTransaction struct that corresponds to the given id.
	var issuingTransaction IssuingTransaction
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingTransaction)
	if unmarshalError != nil {
		return issuingTransaction, err
	}
	return issuingTransaction, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingTransaction, chan Error.StarkErrors) {
	//	Retrieve IssuingTransaction structs
	//
	//	Receive a channel of IssuingTransaction structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "approved", "canceled", "denied", "confirmed" or "voided"
	//		- ids [slice of strings, default nil]: Purchase IDs
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingTransaction structs with updated attributes
	var issuingTransaction IssuingTransaction
	transactions := make(chan IssuingTransaction)
	transactionsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingTransaction)
			if err != nil {
				transactionsError <- Error.UnknownError(err.Error())
				continue
			}
			transactions <- issuingTransaction
		}
		for err := range errorChannel {
			transactionsError <- err
		}
		close(transactions)
		close(transactionsError)
	}()
	return transactions, transactionsError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingTransaction, string, Error.StarkErrors) {
	//	Retrieve paged IssuingTransaction structs
	//
	//	Receive a slice of up to 100 IssuingTransaction structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [string, default nil]: Filter for status of retrieved structs. ex: "approved", "canceled", "denied", "confirmed" or "voided"
	//		- ids [slice of strings, default nil, default nil]: Purchase IDs
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingTransaction structs with updated attributes
	//	- cursor to retrieve the next page of IssuingPurchase structs
	var issuingTransactions []IssuingTransaction
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingTransactions)
	if unmarshalError != nil {
		return issuingTransactions, cursor, err
	}
	return issuingTransactions, cursor, err
}
