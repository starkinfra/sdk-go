package issuingrestock

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingRestock struct
//
//	The IssuingRestock object displays the information of the restock orders created in your Workspace.
//  This resource place a restock order for a specific IssuingStock object.
//
//	Attributes (required):
//	- Count [int]: number of restocks to be restocked. ex: 100
//	- StockId [string]: IssuingStock unique id ex: "5136459887542272"
//
//	Parameters (optional):
//	- Tags [slice of string]: Slice of strings for tagging returned by the sub-issuer during the authorization. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when IssuingRestock is created. ex: "5656565656565656"
//	- Status [string]: Current IssuingCard status. ex: "approved", "canceled", "denied", "confirmed", "voided"
//	- Updated [time.Time]: Latest update datetime for the IssuingRestock. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingRestock. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingRestock struct {
	Count   int        `json:",omitempty"`
	StockId string     `json:",omitempty"`
	Tags    []string   `json:",omitempty"`
	Id      string     `json:",omitempty"`
	Status  string     `json:",omitempty"`
	Updated *time.Time `json:",omitempty"`
	Created *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingRestock"}

func Create(restocks []IssuingRestock, user user.User) ([]IssuingRestock, Error.StarkErrors) {
	//	Create IssuingRestocks
	//
	//	Send a slice of IssuingRestock structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- restocks [slice of IssuingRestock structs]: Slice of IssuingRestock structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- slice of IssuingRestock structs with updated attributes
	create, err := utils.Multi(resource, restocks, nil, user)
	unmarshalError := json.Unmarshal(create, &restocks)
	if unmarshalError != nil {
		return restocks, err
	}
	return restocks, err
}

func Get(id string, user user.User) (IssuingRestock, Error.StarkErrors) {
	//	Retrieve a specific IssuingRestock by its id
	//
	//	Receive a single IssuingRestock struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingRestock struct that corresponds to the given id.
	var issuingRestock IssuingRestock
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingRestock)
	if unmarshalError != nil {
		return issuingRestock, err
	}
	return issuingRestock, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingRestock {
	//	Retrieve IssuingRestock structs
	//
	//	Receive a channel of IssuingRestock structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "processing", "confirmed"}
	//		- stockIds [slice of strings, default nil]: slice of stockIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"card", "corporate"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingRestock structs with updated attributes
	var issuingRestock IssuingRestock
	restocks := make(chan IssuingRestock)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingRestock)
			if err != nil {
				print(err)
			}
			restocks <- issuingRestock
		}
		close(restocks)
	}()
	return restocks
}

func Page(params map[string]interface{}, user user.User) ([]IssuingRestock, string, Error.StarkErrors) {
	//	Retrieve paged IssuingRestock structs
	//
	//	Receive a slice of up to 100 IssuingRestock structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "processing", "confirmed"}
	//		- stockIds [slice of strings, default nil]: slice of stock_ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"card", "corporate"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingRestock structs with updated attributes
	//	- cursor to retrieve the next page of IssuingRestock structs
	var issuingRestocks []IssuingRestock
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingRestocks)
	if unmarshalError != nil {
		return issuingRestocks, cursor, err
	}
	return issuingRestocks, cursor, err
}
