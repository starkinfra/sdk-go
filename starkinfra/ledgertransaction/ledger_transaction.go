package ledgertransaction

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	Ledger "github.com/starkinfra/sdk-go/starkinfra/ledger"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	LedgerTransaction struct
//
//	LedgerTransactions are used to move amounts in and out of a Ledger, updating its balance.
//	They can represent a deposit, a withdrawal, a transfer, an adjustment, etc.
//
//	Parameters (required):
//	- Amount [int]: Amount of the transaction. ex: 11234
//	- LedgerId [string]: Id of the Ledger containing the transaction. ex: "5656565656565656"
//	- ExternalId [string]: String that must be unique among all your LedgerTransactions in a single Ledger. ex: "my-internal-id-123456"
//	- Source [string]: Source of the LedgerTransaction. ex: "bank-transfer/123"
//
//	Parameters (optional):
//	- Fee [int, default nil]: Fee applied to the LedgerTransaction. ex: 100
//	- Rules [slice of Ledger.Rule structs, default nil]: Slice of Rule structs linked to the LedgerTransaction. Rules are used to overwrite the Ledger's rules for this transaction. ex: []ledger.Rule{{Key: "minimumBalance", Value: 0}}
//	- Metadata [map of strings, default nil]: Map used to store additional information about the LedgerTransaction struct. ex: map[string]interface{}{"orderId": "123", "orderType": "purchase"}
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for LedgerTransactions. ex: []string{"transfer/123", "savings"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the LedgerTransaction is created. ex: "5656565656565656"
//	- Balance [int]: Ledger's balance after the transaction. ex: 11234
//	- Created [time.Time]: Creation datetime for the LedgerTransaction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type LedgerTransaction struct {
	Amount     int                    `json:",omitempty"`
	LedgerId   string                 `json:",omitempty"`
	ExternalId string                 `json:",omitempty"`
	Source     string                 `json:",omitempty"`
	Fee        int                    `json:",omitempty"`
	Rules      []Ledger.Rule          `json:",omitempty"`
	Metadata   map[string]interface{} `json:",omitempty"`
	Tags       []string               `json:",omitempty"`
	Id         string                 `json:",omitempty"`
	Balance    int                    `json:",omitempty"`
	Created    *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "LedgerTransaction"}

func Create(transactions []LedgerTransaction, user user.User) ([]LedgerTransaction, Error.StarkErrors) {
	//	Create LedgerTransactions
	//
	//	Send a slice of LedgerTransaction structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- transactions [slice of LedgerTransaction structs]: Slice of LedgerTransaction structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of LedgerTransaction structs with updated attributes
	create, err := utils.Multi(resource, transactions, nil, user)
	unmarshalError := json.Unmarshal(create, &transactions)
	if unmarshalError != nil {
		return transactions, err
	}
	return transactions, err
}

func Get(id string, user user.User) (LedgerTransaction, Error.StarkErrors) {
	//	Retrieve a specific LedgerTransaction by its id
	//
	//	Receive a single LedgerTransaction struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- LedgerTransaction struct that corresponds to the given id.
	var transaction LedgerTransaction
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &transaction)
	if unmarshalError != nil {
		return transaction, err
	}
	return transaction, err
}

func Query(params map[string]interface{}, user user.User) (chan LedgerTransaction, chan Error.StarkErrors) {
	//	Retrieve LedgerTransaction structs
	//
	//	Receive a channel of LedgerTransaction structs previously created in the Stark Infra API
	//
	//	Parameters (conditionally required):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- ledgerId [string, default nil]: Id of the Ledger containing the transaction. Either ledgerId or ids must be provided. If both are sent, the query will be filtered by both. ex: "5656565656565656"
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. Either ledgerId or ids must be provided. If both are sent, the query will be filtered by both. ex: []string{"5656565656565656", "4545454545454545"}
	//	Parameters (optional):
	//		- flow [string, default nil]: Direction of the transaction. ex: "in" or "out"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"transfer/123", "savings"}
	//		- externalIds [slice of strings, default nil]: Slice of external ids to filter retrieved structs. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- limit [int, default 100, maximum 1000]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of LedgerTransaction structs with updated attributes
	var transaction LedgerTransaction
	transactions := make(chan LedgerTransaction)
	transactionsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)			
			err := json.Unmarshal(contentByte, &transaction)
			if err != nil {
				transactionsError <- Error.UnknownError(err.Error())
				continue
			}
			transactions <- transaction
		}
		for err := range errorChannel {
			transactionsError <- err
		}
		close(transactions)
		close(transactionsError)
	}()
	return transactions, transactionsError
}

func Page(params map[string]interface{}, user user.User) ([]LedgerTransaction, string, Error.StarkErrors) {
	//	Retrieve paged LedgerTransactions
	//
	//	Receive a slice of LedgerTransaction structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (conditionally required):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- ledgerId [string, default nil]: Id of the Ledger containing the transaction. Either ledgerId or ids must be provided. If both are sent, the query will be filtered by both. ex: "5656565656565656"
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. Either ledgerId or ids must be provided. If both are sent, the query will be filtered by both. ex: []string{"5656565656565656", "4545454545454545"}
	//	Parameters (optional):
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- flow [string, default nil]: Direction of the transaction. ex: "in" or "out"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"transfer/123", "savings"}
	//		- externalIds [slice of strings, default nil]: Slice of external ids to filter retrieved structs. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- limit [int, default 100, maximum 1000]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of LedgerTransaction structs with updated attributes
	//	- cursor to retrieve the next page of LedgerTransaction structs
	var transactions []LedgerTransaction
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &transactions)
	if unmarshalError != nil {
		return transactions, cursor, err
	}
	return transactions, cursor, err
}
