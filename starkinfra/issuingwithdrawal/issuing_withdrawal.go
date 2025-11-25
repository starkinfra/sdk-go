package issuingwithdrawal

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingWithdrawal struct
//
//	The IssuingWithdrawal structs created in your Workspace return cash from your Issuing balance to your
//	Banking balance.
//
//	Parameters (required):
//	- Amount [int]: IssuingWithdrawal value in cents. Minimum = 0 (any value will be accepted). ex: 1234 (= R$ 12.34)
//	- ExternalId [string] IssuingWithdrawal external ID. ex: "12345"
//	- Description [string]: IssuingWithdrawal description. ex: "sending money back"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"tony", "stark"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingWithdrawal is created. ex: "5656565656565656"
//	- TransactionId [string]: Stark Bank ledger transaction ids linked to this IssuingWithdrawal
//	- IssuingTransactionId [string]: Issuing ledger transaction ids linked to this IssuingWithdrawal
//	- Updated [time.Time]: Latest update datetime for the IssuingWithdrawal. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingWithdrawal. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingWithdrawal struct {
	Amount               int        `json:",omitempty"`
	ExternalId           string     `json:",omitempty"`
	Description          string     `json:",omitempty"`
	Tags                 []string   `json:",omitempty"`
	Id                   string     `json:",omitempty"`
	TransactionId        string     `json:",omitempty"`
	IssuingTransactionId string     `json:",omitempty"`
	Updated              *time.Time `json:",omitempty"`
	Created              *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingWithdrawal"}

func Create(withdrawal IssuingWithdrawal, user user.User) (IssuingWithdrawal, Error.StarkErrors) {
	//	Create an IssuingWithdrawal
	//
	//	Send a single IssuingWithdrawal struct for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- withdrawal [IssuingWithdrawal struct]: IssuingWithdrawal struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingWithdrawal struct with updated attributes
	create, err := utils.Single(resource, withdrawal, user)
	unmarshalError := json.Unmarshal(create, &withdrawal)
	if unmarshalError != nil {
		return withdrawal, err
	}
	return withdrawal, err
}

func Get(id string, user user.User) (IssuingWithdrawal, Error.StarkErrors) {
	//	Retrieve a specific IssuingWithdrawal by its id
	//
	//	Receive a single IssuingWithdrawal struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingWithdrawal struct that corresponds to the given id.
	var issuingWithdrawal IssuingWithdrawal
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingWithdrawal)
	if unmarshalError != nil {
		return issuingWithdrawal, err
	}
	return issuingWithdrawal, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingWithdrawal, chan Error.StarkErrors) {
	//	Retrieve IssuingWithdrawal structs
	//
	//	Receive a channel of IssuingWithdrawal structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingWithdrawal structs with updated attributes
	var issuingWithdrawal IssuingWithdrawal
	withdrawals := make(chan IssuingWithdrawal)
	withdrawalsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingWithdrawal)
			if err != nil {
				withdrawalsError <- Error.UnknownError(err.Error())
				continue
			}
			withdrawals <- issuingWithdrawal
		}
		for err := range errorChannel {
			withdrawalsError <- err
		}
		close(withdrawals)
		close(withdrawalsError)
	}()
	return withdrawals, withdrawalsError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingWithdrawal, string, Error.StarkErrors) {
	//	Retrieve paged IssuingWithdrawal structs
	//
	//	Receive a slice of up to 100 IssuingWithdrawal structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingWithdrawal structs with updated attributes
	//	- cursor to retrieve the next page of IssuingWithdrawal structs
	var issuingWithdrawals []IssuingWithdrawal
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingWithdrawals)
	if unmarshalError != nil {
		return issuingWithdrawals, cursor, err
	}
	return issuingWithdrawals, cursor, err
}
