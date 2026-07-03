package issuingstockrule

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingStockRule struct
//
//	The IssuingStockRule object defines a notification rule attached to an IssuingStock.
//	When the linked stock balance reaches MinimumBalance, the recipients listed in Emails
//	and Phones are notified.
//
//	Parameters (required):
//	- MinimumBalance [int]: Stock balance threshold that triggers a notification. ex: 10000
//	- StockId [string]: IssuingStock unique id the rule is linked to. ex: "5136459887542272"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"card", "corporate"}
//	- Emails [slice of strings, default nil]: Emails notified when the stock reaches the minimum balance. ex: []string{"john.doe@enterprise.com"}
//	- Phones [slice of strings, default nil]: Phones notified when the stock reaches the minimum balance. ex: []string{"+5511912345678"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingStockRule is created. ex: "5664445921492992"
//	- Status [string]: Current IssuingStockRule status. ex: "active", "canceled"
//	- Updated [time.Time]: Latest update datetime for the IssuingStockRule. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingStockRule. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingStockRule struct {
	MinimumBalance int        `json:",omitempty"`
	StockId        string     `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Emails         []string   `json:",omitempty"`
	Phones         []string   `json:",omitempty"`
	Id             string     `json:",omitempty"`
	Status         string     `json:",omitempty"`
	Updated        *time.Time `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingStockRule"}

func Create(rules []IssuingStockRule, user user.User) ([]IssuingStockRule, Error.StarkErrors) {
	//	Create IssuingStockRules
	//
	//	Send a slice of IssuingStockRule structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- rules [slice of IssuingStockRule structs]: Slice of IssuingStockRule structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingStockRule structs with updated attributes
	create, err := utils.Multi(resource, rules, nil, user)
	unmarshalError := json.Unmarshal(create, &rules)
	if unmarshalError != nil {
		return rules, err
	}
	return rules, err
}

func Get(id string, user user.User) (IssuingStockRule, Error.StarkErrors) {
	//	Retrieve a specific IssuingStockRule by its id
	//
	//	Receive a single IssuingStockRule struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingStockRule struct that corresponds to the given id.
	var issuingStockRule IssuingStockRule
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingStockRule)
	if unmarshalError != nil {
		return issuingStockRule, err
	}
	return issuingStockRule, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingStockRule, chan Error.StarkErrors) {
	//	Retrieve IssuingStockRule structs
	//
	//	Receive a channel of IssuingStockRule structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"active", "canceled"}
	//		- stockIds [slice of strings, default nil]: slice of stockIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"card", "corporate"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingStockRule structs with updated attributes
	var issuingStockRule IssuingStockRule
	rules := make(chan IssuingStockRule)
	rulesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingStockRule)
			if err != nil {
				rulesError <- Error.UnknownError(err.Error())
				continue
			}
			rules <- issuingStockRule
		}
		for err := range errorChannel {
			rulesError <- err
		}
		close(rules)
		close(rulesError)
	}()
	return rules, rulesError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingStockRule, string, Error.StarkErrors) {
	//	Retrieve paged IssuingStockRule structs
	//
	//	Receive a slice of up to 100 IssuingStockRule structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"active", "canceled"}
	//		- stockIds [slice of strings, default nil]: slice of stockIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"card", "corporate"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingStockRule structs with updated attributes
	//	- cursor to retrieve the next page of IssuingStockRule structs
	//
	var issuingStockRules []IssuingStockRule
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingStockRules)
	if unmarshalError != nil {
		return issuingStockRules, cursor, err
	}
	return issuingStockRules, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (IssuingStockRule, Error.StarkErrors) {
	//	Update IssuingStockRule entity
	//
	//	Update an IssuingStockRule by passing id.
	//
	//	Parameters (required):
	//	- id [string]: IssuingStockRule id. ex: "5656565656565656"
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"minimumBalance": 20000}
	//		Parameters (optional):
	//		- minimumBalance [int]: Stock balance threshold that triggers a notification. ex: 20000
	//		- tags [slice of strings]: Slice of strings for tagging
	//		- emails [slice of strings]: Emails notified when the stock reaches the minimum balance
	//		- phones [slice of strings]: Phones notified when the stock reaches the minimum balance
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- target IssuingStockRule with updated attributes
	var issuingStockRule IssuingStockRule
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &issuingStockRule)
	if unmarshalError != nil {
		return issuingStockRule, err
	}
	return issuingStockRule, err
}

func Cancel(id string, user user.User) (IssuingStockRule, Error.StarkErrors) {
	//	Cancel an IssuingStockRule entity
	//
	//	Cancel an IssuingStockRule entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: IssuingStockRule unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled IssuingStockRule struct
	var issuingStockRule IssuingStockRule
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &issuingStockRule)
	if unmarshalError != nil {
		return issuingStockRule, err
	}
	return issuingStockRule, err
}
