package individualidentity

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IndividualIdentity struct
//
//	An IndividualIdentity represents an individual to be validated. It can have several individual identities attached
//	to it, which are used to validate the identity of the individual. Once an individual identity is created, individual
//	identities must be attached to it using the created method of the individual identity resource. When all the required
//	individual identities are attached to an individual identity it can be sent to validation by patching its status to
//	processing.
//
//	When you initialize a IndividualIdentity, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the slice of created objects.
//
//	Parameters (required):
//	- Name [string]: individual's full name. ex: "Edward Stark"
// 	- TaxId [string]: individual's tax ID (CPF). ex: "594.739.480-42"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for IndividualIdentities. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the identity is created. ex: "5656565656565656"
//	- Status [string]: current status of the IndividualIdentity. Options: "created", "canceled", "processing", "failed", "success"
//	- Created [time.Time]: creation datetime for the IndividualIdentity. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IndividualIdentity struct {
	Created *time.Time `json:",omitempty"`
	Id      string     `json:",omitempty"`
	Name    string     `json:",omitempty"`
	Status  string     `json:",omitempty"`
	Tags    []string   `json:",omitempty"`
	TaxId   string     `json:",omitempty"`
}

var resource = map[string]string{"name": "IndividualIdentity"}

func Create(identity []IndividualIdentity, user user.User) ([]IndividualIdentity, Error.StarkErrors) {
	//	Create IndividualIdentities
	//
	//	Send a slice of IndividualIdentity objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- identities [slice of IndividualIdentity structs]: slice of IndividualIdentity objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualIdentity struct with updated attributes
	create, err := utils.Multi(resource, identity, nil, user)
	fmt.Println(string(create))
	unmarshalError := json.Unmarshal(create, &identity)
	if unmarshalError != nil {
		return identity, err
	}
	return identity, err
}

func Get(id string, user user.User) (IndividualIdentity, Error.StarkErrors) {
	//	Retrieve a specific IndividualIdentity by its id
	//
	//	Receive a single IndividualIdentity struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- IndividualIdentity struct that corresponds to the given id.
	var individualIdentity IndividualIdentity
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &individualIdentity)
	if unmarshalError != nil {
		return individualIdentity, err
	}
	return individualIdentity, err
}

func Query(params map[string]interface{}, user user.User) (chan IndividualIdentity, chan Error.StarkErrors) {
	//	Retrieve IndividualIdentitys
	//
	//	Receive a channel of IndividualIdentity structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//  	- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- status [slice of strings, default nil]: filter for status of retrieved structs. Options: "created", "canceled", "processing", "failed" and "success"
	//  	- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//  	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IndividualIdentity structs with updated attributes
	var individualIdentity IndividualIdentity
	identities := make(chan IndividualIdentity)
	identitiesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &individualIdentity)
			if err != nil {
				identitiesError <- Error.UnknownError(err.Error())
			}
			identities <- individualIdentity
		}
		for err := range errorChannel {
			identitiesError <- err
		}
		close(identities)
		close(identitiesError)
	}()
	return identities, identitiesError
}

func Page(params map[string]interface{}, user user.User) ([]IndividualIdentity, string, Error.StarkErrors) {
	//	Retrieve paged IndividualIdentity structs
	//
	//	Receive a slice of up to 100 IndividualIdentity structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. Options: "created", "canceled", "processing", "failed" and "success"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualIdentity structs with updated attributes
	//	- cursor to retrieve the next page of IndividualIdentity structs
	var individualIdentities []IndividualIdentity
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &individualIdentities)
	if unmarshalError != nil {
		return individualIdentities, cursor, err
	}
	return individualIdentities, cursor, err
}

func Update(id string, status string, user user.User) (IndividualIdentity, Error.StarkErrors) {
	//	Update an IndividualIdentity entity
	//
	//	Update an IndividualIdentity by passing id.
	//
	//	Parameters (required):
	//	- id [string]: IndividualIdentity unique id. ex: "6306109539221504"
	//	- status [string]: You may send IndividualDocuments to validation by passing 'processing' in the status
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- target IndividualIdentity with updated attributes
	var individualIdentity IndividualIdentity
	patchData := map[string]interface{}{}
	patchData["status"] = status
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &individualIdentity)
	if unmarshalError != nil {
		return individualIdentity, err
	}
	return individualIdentity, err
}

func Cancel(id string, user user.User) (IndividualIdentity, Error.StarkErrors) {
	//	Cancel an IndividualIdentity entity
	//
	//	Cancel an IndividualIdentity by passing id.
	//
	//	Parameters (required):
	//	- id [string]: IndividualIdentity unique id. ex: "6306109539221504"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- canceled IndividualIdentity struct
	var individualIdentity IndividualIdentity
	cancel, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(cancel, &individualIdentity)
	if unmarshalError != nil {
		return individualIdentity, err
	}
	return individualIdentity, err
}
