package issuingdesign

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingDesign struct
//
//  The IssuingDesign object displays information on the card and card package designs available to your Workspace.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when IssuingDesign is created. ex: "5656565656565656"
//	- Name [string]: card or package design name. ex: "stark-plastic-dark-001"
//	- EmbosserIds [slice of string]: slice of embosser unique ids. ex: []string{"5136459887542272", "5136459887542273"}
//	- Type [string]: card or package design type. Options: "card", "envelope"
//  - Updated [time.Time]: updated datetime for the IssuingDesign. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the IssuingDesign. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingDesign struct {
	Id          string     `json:",omitempty"`
	Name        string     `json:",omitempty"`
	EmbosserIds []string   `json:",omitempty"`
	Type        string     `json:",omitempty"`
	Updated     *time.Time `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingDesign"}

func Get(id string, user user.User) (IssuingDesign, Error.StarkErrors) {
	//	Retrieve a specific IssuingDesign by its id
	//
	//	Receive a single IssuingDesign struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingDesign struct that corresponds to the given id.
	var issuingDesign IssuingDesign
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingDesign)
	if unmarshalError != nil {
		return issuingDesign, err
	}
	return issuingDesign, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingDesign {
	//	Retrieve IssuingDesigns
	//
	//	Receive a channel of IssuingDesign structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//  	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingDesign structs with updated attributes
	var issuingDesign IssuingDesign
	designs := make(chan IssuingDesign)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingDesign)
			if err != nil {
				print(err)
			}
			designs <- issuingDesign
		}
		close(designs)
	}()
	return designs
}

func Page(params map[string]interface{}, user user.User) ([]IssuingDesign, string, Error.StarkErrors) {
	//	Retrieve paged IssuingDesign structs
	//
	//	Receive a slice of up to 100 IssuingDesign structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingDesign structs with updated attributes
	//	- cursor to retrieve the next page of IssuingDesign structs
	var issuingDesigns []IssuingDesign
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingDesigns)
	if unmarshalError != nil {
		return issuingDesigns, cursor, err
	}
	return issuingDesigns, cursor, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific IssuingDesign pdf file
	//
	//	Receive a single IssuingDesign pdf file generated in the Stark Infra API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- IssuingDesign pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}
