package issuingembossingkit

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	IssuingDesign "github.com/starkinfra/sdk-go/starkinfra/issuingdesign"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingEmbossingKit struct
//
//	The IssuingEmbossingKit object displays information on the embossing kits available to your Workspace.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingEmbossingKit is created. ex: "5656565656565656"
//	- Name [string]: Embossing kit name. ex: "stark-plastic-dark-001"
//	- Designs [slice of IssuingDesigns]: slice of IssuingDesign objects. ex: "created", "processing", "success", "failed"
//	- Updated [time.Time]: Latest update datetime for the IssuingEmbossingKit. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingEmbossingKit. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingEmbossingKit struct {
	Id      string                        `json:",omitempty"`
	Name    string                        `json:",omitempty"`
	Designs []IssuingDesign.IssuingDesign `json:",omitempty"`
	Updated *time.Time                    `json:",omitempty"`
	Created *time.Time                    `json:",omitempty"`
}

var object IssuingEmbossingKit
var objects []IssuingEmbossingKit
var resource = map[string]string{"name": "IssuingEmbossingKit"}

func Get(id string, user user.User) (IssuingEmbossingKit, Error.StarkErrors) {
	//	Retrieve a specific IssuingEmbossingKit by its id
	//
	//	Receive a single IssuingEmbossingKit struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingEmbossingKit struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingEmbossingKit {
	//	Retrieve IssuingEmbossingKits
	//
	//	Receive a channel of IssuingEmbossingKit structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of objects to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved objects. ex: []string{"created", "processing", "success", "failed"}
	//		- designIds [slice of strings, default nil]: Slice of designIds to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingEmbossingKit structs with updated attributes
	kits := make(chan IssuingEmbossingKit)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			kits <- object
		}
		close(kits)
	}()
	return kits
}

func Page(params map[string]interface{}, user user.User) ([]IssuingEmbossingKit, string, Error.StarkErrors) {
	//	Retrieve paged IssuingEmbossingKit structs
	//
	//	Receive a slice of up to 100 IssuingEmbossingKit structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: Maximum number of objects to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved objects. ex: []string{"created", "processing", "success", "failed"}
	//		- designIds [slice of strings, default nil]: Slice of designIds to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingEmbossingKit structs with updated attributes
	//	- cursor to retrieve the next page of IssuingEmbossingKit structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
