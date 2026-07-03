package issuingtokendesign

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingTokenDesign struct
//
//	The IssuingTokenDesign object displays the information of the token designs created in your Workspace.
//	This resource represents the existent designs for the cards which will be tokenized.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when IssuingTokenDesign is created. ex: "5656565656565656"
//	- Name [string]: design name. ex: "Stark Bank - White Metal"
//	- Created [time.Time]: creation datetime for the IssuingTokenDesign. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the IssuingTokenDesign. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC),

type IssuingTokenDesign struct {
	Id      string     `json:",omitempty"`
	Name    string     `json:",omitempty"`
	Created *time.Time `json:",omitempty"`
	Updated *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingTokenDesign"}

func Get(id string, user user.User) (IssuingTokenDesign, Error.StarkErrors) {
	//	Retrieve a specific IssuingTokenDesign by its id
	//
	//	Receive a single IssuingTokenDesign struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingTokenDesign struct that corresponds to the given id.
	var issuingTokenDesign IssuingTokenDesign
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingTokenDesign)
	if unmarshalError != nil {
		return issuingTokenDesign, err
	}
	return issuingTokenDesign, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingTokenDesign, chan Error.StarkErrors) {
	//	Retrieve IssuingTokenDesigns
	//
	//	Receive a channel of IssuingTokenDesign structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingTokenDesign structs with updated attributes
	var issuingTokenDesign IssuingTokenDesign
	designs := make(chan IssuingTokenDesign)
	designsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingTokenDesign)
			if err != nil {
				designsError <- Error.UnknownError(err.Error())
				continue
			}
			designs <- issuingTokenDesign
		}
		for err := range errorChannel {
			designsError <- err
		}
		close(designs)
		close(designsError)
	}()
	return designs, designsError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingTokenDesign, string, Error.StarkErrors) {
	//	Retrieve paged IssuingTokenDesign structs
	//
	//	Receive a slice of up to 100 IssuingTokenDesign structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingTokenDesign structs with updated attributes
	//	- cursor to retrieve the next page of IssuingTokenDesign structs
	var issuingTokenDesigns []IssuingTokenDesign
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingTokenDesigns)
	if unmarshalError != nil {
		return issuingTokenDesigns, cursor, err
	}
	return issuingTokenDesigns, cursor, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific IssuingTokenDesign pdf file
	//
	//	Receive a single IssuingTokenDesign pdf file generated in the Stark Infra API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- IssuingTokenDesign pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}
