package issuingstock

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingStock struct
//
//	The IssuingStock object represents the current stock of a certain IssuingDesign linked to an Embosser available to your workspace.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when IssuingStock is created. ex: "5656565656565656"
//	- Balance [int]: [EXPANDABLE] current stock balance. ex: 1000
//	- DesignId [string]: IssuingDesign unique id. ex: "5656565656565656"
//	- EmbosserId [string]: Embosser unique id. ex: "5656565656565656"
//	- Updated [time.Time]: Latest update datetime for the IssuingStock. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingStock. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingStock struct {
	Id         string     `json:",omitempty"`
	Balance    int        `json:",omitempty"`
	DesignId   string     `json:",omitempty"`
	EmbosserId string     `json:",omitempty"`
	Updated    *time.Time `json:",omitempty"`
	Created    *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingStock"}

func Get(id string, expand map[string]interface{}, user user.User) (IssuingStock, Error.StarkErrors) {
	//	Retrieve a specific IssuingStock by its id
	//
	//	Receive a single IssuingStock struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//	- expand [slice of strings, default nil]: Fields to expand information. ex: []string{"balance"}
	//
	//	Return:
	//	- issuingStock struct that corresponds to the given id.
	var issuingStock IssuingStock
	get, err := utils.Get(resource, id, expand, user)
	unmarshalError := json.Unmarshal(get, &issuingStock)
	if unmarshalError != nil {
		return issuingStock, err
	}
	return issuingStock, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingStock {
	//	Retrieve IssuingStock structs
	//
	//	Receive a channel of IssuingStock structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- designIds [slice of strings, default nil]: IssuingDesign unique ids. ex: []string{"5656565656565656", "4545454545454545"}
	//		- embosserIds [slice of strings, default nil]: Embosser unique ids. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//		- expand [slice of strings, default nil]: fields to expand information. ex: []string{"balance"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingStock structs with updated attributes
	var issuingStock IssuingStock
	stocks := make(chan IssuingStock)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingStock)
			if err != nil {
				print(err)
			}
			stocks <- issuingStock
		}
		close(stocks)
	}()
	return stocks
}

func Page(params map[string]interface{}, user user.User) ([]IssuingStock, string, Error.StarkErrors) {
	//	Retrieve paged IssuingStock structs
	//
	//	Receive a slice of up to 100 IssuingStock structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- designIds [slice of strings, default nil]: IssuingDesign unique ids. ex: []string{"5656565656565656", "4545454545454545"}
	//		- embosserIds [slice of strings, default nil]: Embosser unique ids. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved issuingStocks. ex: []string{"5656565656565656", "4545454545454545"}
	//		- expand [slice of strings, default nil]: fields to expand information. ex: []string{"balance"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingStock structs with updated attributes
	//	- cursor to retrieve the next page of IssuingStock structs
	var issuingStocks []IssuingStock
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingStocks)
	if unmarshalError != nil {
		return issuingStocks, cursor, err
	}
	return issuingStocks, cursor, err
}
