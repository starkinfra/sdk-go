package issuingholder

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	IssuingRule "github.com/starkinfra/sdk-go/starkinfra/issuingrule"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingHolder struct
//
//	The IssuingHolder describes a card holder that may group several cards.
//
//	Parameters (required):
//	- Name [string]: Card holder's name. ex: Jannie Lanister
//	- TaxId [string]: Card holder's tax ID. ex: "012.345.678-90"
//	- ExternalId [string]: Card holder's external ID. "my_external_id1"
//
//	Parameters (optional):
//	- Rules [slice of IssuingRule, default nil]: [EXPANDABLE] Slice of holder spending rules
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingHolder is created. ex: "5656565656565656"
//	- Status [string]: Current IssuingHolder status. ex: "active", "blocked", "canceled"
//	- Updated [time.Time]: Latest update datetime for the IssuingHolder. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingHolder. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingHolder struct {
	Name       string                    `json:",omitempty"`
	TaxId      string                    `json:",omitempty"`
	ExternalId string                    `json:",omitempty"`
	Rules      []IssuingRule.IssuingRule `json:",omitempty"`
	Tags       []string                  `json:",omitempty"`
	Id         string                    `json:",omitempty"`
	Status     string                    `json:",omitempty"`
	Updated    *time.Time                `json:",omitempty"`
	Created    *time.Time                `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingHolder"}

func Create(holders []IssuingHolder, expand map[string]interface{}, user user.User) ([]IssuingHolder, Error.StarkErrors) {
	//	Create IssuingHolder
	//
	//	Send a slice of IssuingHolder structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- holders [slice of IssuingHolder structs]: Slice of IssuingHolder structs to be created in the API
	//
	//	Parameters (optional):
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingHolder structs with updated attributes
	var issuingHolders []IssuingHolder
	create, err := utils.Multi(resource, holders, expand, user)
	unmarshalError := json.Unmarshal(create, &issuingHolders)
	if unmarshalError != nil {
		return issuingHolders, err
	}
	return issuingHolders, err
}

func Get(id string, expand map[string]interface{}, user user.User) (IssuingHolder, Error.StarkErrors) {
	//	Retrieve a specific IssuingHolder by its id
	//
	//	Receive a single IssuingHolder struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"rules"}
	//
	//	Return:
	//	- issuingHolder struct that corresponds to the given id.
	var issuingHolder IssuingHolder
	get, err := utils.Get(resource, id, expand, user)
	unmarshalError := json.Unmarshal(get, &issuingHolder)
	if unmarshalError != nil {
		return issuingHolder, err
	}
	return issuingHolder, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingHolder {
	//	Retrieve IssuingHolders
	//
	//	Receive a channel of IssuingHolder structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- expand [string, default nil]: Fields to expand information. ex: "rules"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingHolder structs with updated attributes
	var issuingHolder IssuingHolder
	holders := make(chan IssuingHolder)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingHolder)
			if err != nil {
				print(err)
			}
			holders <- issuingHolder
		}
		close(holders)
	}()
	return holders
}

func Page(params map[string]interface{}, user user.User) ([]IssuingHolder, string, Error.StarkErrors) {
	//	Retrieve IssuingHolders
	//
	//	Receive a slice of up to 100 IssuingHolder structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- expand [string, default nil]: Fields to expand information. ex: "rules"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingHolder structs with updated attributes
	//	- cursor to retrieve the next page of IssuingHolder structs
	var issuingHolders []IssuingHolder
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingHolders)
	if unmarshalError != nil {
		return issuingHolders, cursor, err
	}
	return issuingHolders, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (IssuingHolder, Error.StarkErrors) {
	//	Update IssuingHolder entity
	//
	//	Update an IssuingHolder by passing id, if it hasn't been paid yet.
	//
	//	Parameters (required):
	//	- id [string]: IssuingHolder id. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//		Parameters (optional):
	//		- status [string]: You may block the IssuingHolder by passing 'blocked' in the status
	//		- name [string]: Card holder name.
	//		- tags [slice of strings]: Slice of strings for tagging
	//		- rules [slice of maps, default nil]: Slice of maps with "amount": int, "currencyCode": string, "id": string, "interval": string, "name": string pairs
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- target IssuingHolder with updated attributes
	var issuingHolder IssuingHolder
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &issuingHolder)
	if unmarshalError != nil {
		return issuingHolder, err
	}
	return issuingHolder, err
}

func Cancel(id string, user user.User) (IssuingHolder, Error.StarkErrors) {
	//	Cancel an IssuingHolder entity
	//
	//	Cancel an IssuingHolder entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: IssuingHolder unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled IssuingHolder struct
	var issuingHolder IssuingHolder
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &issuingHolder)
	if unmarshalError != nil {
		return issuingHolder, err
	}
	return issuingHolder, err
}
