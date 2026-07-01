package ledger

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	Ledger struct
//
//	Ledgers are used to track the balance of a given amount by inserting LedgerTransactions to them.
//	They can represent a bank account, a digital wallet, an inventory product, etc.
//
//	Parameters (required):
//	- ExternalId [string]: String that must be unique among all your Ledgers. ex: "my-internal-id-123456"
//
//	Parameters (optional):
//	- Rules [slice of Ledger.Rule structs, default nil]: Slice of Rule structs linked to the Ledger. Rules are used to limit the balance of the Ledger. ex: []ledger.Rule{{Key: "minimumBalance", Value: 0}}
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for Ledgers. ex: []string{"account/123", "savings"}
//	- Metadata [map of strings, default nil]: Map used to store additional information about the Ledger struct. ex: map[string]interface{}{"accountId": "123", "accountType": "savings"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the Ledger is created. ex: "5656565656565656"
//	- Created [time.Time]: Creation datetime for the Ledger. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the Ledger. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Ledger struct {
	ExternalId string                 `json:",omitempty"`
	Rules      []Rule                 `json:",omitempty"`
	Tags       []string               `json:",omitempty"`
	Metadata   map[string]interface{} `json:",omitempty"`
	Id         string                 `json:",omitempty"`
	Created    *time.Time             `json:",omitempty"`
	Updated    *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "Ledger"}

func Create(ledgers []Ledger, user user.User) ([]Ledger, Error.StarkErrors) {
	//	Create Ledgers
	//
	//	Send a slice of Ledger structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- ledgers [slice of Ledger structs]: Slice of Ledger structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of Ledger structs with updated attributes
	create, err := utils.Multi(resource, ledgers, nil, user)
	unmarshalError := json.Unmarshal(create, &ledgers)
	if unmarshalError != nil {
		return ledgers, err
	}
	return ledgers, err
}

func Get(id string, user user.User) (Ledger, Error.StarkErrors) {
	//	Retrieve a specific Ledger by its id
	//
	//	Receive a single Ledger struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Ledger struct that corresponds to the given id.
	var ledger Ledger
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &ledger)
	if unmarshalError != nil {
		return ledger, err
	}
	return ledger, err
}

func Query(params map[string]interface{}, user user.User) (chan Ledger, chan Error.StarkErrors) {
	//	Retrieve Ledger structs
	//
	//	Receive a channel of Ledger structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- externalIds [slice of strings, default nil]: Slice of external ids to filter retrieved structs. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"account/123", "savings"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of Ledger structs with updated attributes
	var ledger Ledger
	ledgers := make(chan Ledger)
	ledgersError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &ledger)
			if err != nil {
				ledgersError <- Error.UnknownError(err.Error())
				continue
			}
			ledgers <- ledger
		}
		for err := range errorChannel {
			ledgersError <- err
		}
		close(ledgers)
		close(ledgersError)
	}()
	return ledgers, ledgersError
}

func Page(params map[string]interface{}, user user.User) ([]Ledger, string, Error.StarkErrors) {
	//	Retrieve paged Ledgers
	//
	//	Receive a slice of up to 100 Ledger structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- externalIds [slice of strings, default nil]: Slice of external ids to filter retrieved structs. ex: []string{"my-internal-id-123456", "my-internal-id-654321"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"account/123", "savings"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of Ledger structs with updated attributes
	//	- cursor to retrieve the next page of Ledger structs
	var ledgers []Ledger
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &ledgers)
	if unmarshalError != nil {
		return ledgers, cursor, err
	}
	return ledgers, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (Ledger, Error.StarkErrors) {
	//	Update Ledger entity
	//
	//	Update a Ledger by passing its id.
	//
	//	Parameters (required):
	//	- id [string]: Ledger id. ex: "5656565656565656"
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"tags": []string{"account/123", "updated"}}
	//		Parameters (optional):
	//		- rules [slice of maps, default nil]: Slice of maps with "key": string, "value": int pairs. Rules are used to limit the balance of the Ledger.
	//		- tags [slice of strings, default nil]: Slice of strings for reference when searching for Ledgers. ex: []string{"account/123", "savings"}
	//		- metadata [map of strings, default nil]: Map used to store additional information about the Ledger struct. ex: map[string]interface{}{"accountId": "123", "accountType": "savings"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- target Ledger with updated attributes
	var ledger Ledger
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &ledger)
	if unmarshalError != nil {
		return ledger, err
	}
	return ledger, err
}
