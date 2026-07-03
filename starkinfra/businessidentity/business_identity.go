package businessidentity

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	BusinessIdentity struct
//
//	A BusinessIdentity represents a company to be validated. It can have several business attachments
//	attached to it, which are used to validate the identity of the company. Once a business identity is created,
//	business attachments must be attached to it using the created method of the business attachment resource. When all
//	the required business attachments are attached to a business identity it can be sent to validation by patching its
//	status to processing.
//
//	When you initialize a BusinessIdentity, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the slice of created objects.
//
//	Parameters (required):
//	- TaxId [string]: company's tax ID (CNPJ). ex: "20.018.183/0001-80"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for BusinessIdentities. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the identity is created. ex: "5656565656565656"
//	- Name [string]: company's name. ex: "Stark Bank S.A."
//	- TaxIdStatus [string]: tax ID status of the BusinessIdentity. ex: "active"
//	- InsightTaxId [string]: tax ID retrieved through insights. ex: "20.018.183/0001-80"
//	- InsightDocumentType [string]: document type retrieved through insights. ex: "cnpj"
//	- NumPages [int]: number of pages of the BusinessIdentity. ex: 3
//	- Representatives [string]: JSON string with the representatives of the BusinessIdentity.
//	- Attachments [slice of strings]: slice of BusinessAttachment ids attached to the BusinessIdentity. ex: []string{"5656565656565656", "4545454545454545"}
//	- Rules [string]: JSON string with the rules of the BusinessIdentity.
//	- Status [string]: current status of the BusinessIdentity. Options: "created", "pending", "canceled", "processing", "success", "failed"
//	- Created [time.Time]: creation datetime for the BusinessIdentity. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the BusinessIdentity. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type BusinessIdentity struct {
	TaxId               string     `json:",omitempty"`
	Tags                []string   `json:",omitempty"`
	Id                  string     `json:",omitempty"`
	Name                string     `json:",omitempty"`
	TaxIdStatus         string     `json:",omitempty"`
	InsightTaxId        string     `json:",omitempty"`
	InsightDocumentType string     `json:",omitempty"`
	NumPages            int        `json:",omitempty"`
	Representatives     string     `json:",omitempty"`
	Attachments         []string   `json:",omitempty"`
	Rules               string     `json:",omitempty"`
	Status              string     `json:",omitempty"`
	Created             *time.Time `json:",omitempty"`
	Updated             *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "BusinessIdentity"}

func Create(identities []BusinessIdentity, user user.User) ([]BusinessIdentity, Error.StarkErrors) {
	//	Create BusinessIdentities
	//
	//	Send a slice of BusinessIdentity objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- identities [slice of BusinessIdentity structs]: slice of BusinessIdentity objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of BusinessIdentity struct with updated attributes
	create, err := utils.Multi(resource, identities, nil, user)
	unmarshalError := json.Unmarshal(create, &identities)
	if unmarshalError != nil {
		return identities, err
	}
	return identities, err
}

func Get(id string, user user.User) (BusinessIdentity, Error.StarkErrors) {
	//	Retrieve a specific BusinessIdentity by its id
	//
	//	Receive a single BusinessIdentity struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- BusinessIdentity struct that corresponds to the given id.
	var businessIdentity BusinessIdentity
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &businessIdentity)
	if unmarshalError != nil {
		return businessIdentity, err
	}
	return businessIdentity, err
}

func Query(params map[string]interface{}, user user.User) (chan BusinessIdentity, chan Error.StarkErrors) {
	//	Retrieve BusinessIdentitys
	//
	//	Receive a channel of BusinessIdentity structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//  	- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- status [slice of strings, default nil]: filter for status of retrieved structs. Options: "created", "pending", "canceled", "processing", "success" and "failed"
	//  	- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//  	- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//  	- taxIds [slice of strings, default nil]: slice of tax ids to filter retrieved structs. ex: []string{"20.018.183/0001-80"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of BusinessIdentity structs with updated attributes
	var businessIdentity BusinessIdentity
	identities := make(chan BusinessIdentity)
	identitiesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &businessIdentity)
			if err != nil {
				identitiesError <- Error.UnknownError(err.Error())
				continue
			}
			identities <- businessIdentity
		}
		for err := range errorChannel {
			identitiesError <- err
		}
		close(identities)
		close(identitiesError)
	}()
	return identities, identitiesError
}

func Page(params map[string]interface{}, user user.User) ([]BusinessIdentity, string, Error.StarkErrors) {
	//	Retrieve paged BusinessIdentity structs
	//
	//	Receive a slice of up to 100 BusinessIdentity structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. Options: "created", "pending", "canceled", "processing", "success" and "failed"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- taxIds [slice of strings, default nil]: slice of tax ids to filter retrieved structs. ex: []string{"20.018.183/0001-80"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of BusinessIdentity structs with updated attributes
	//	- cursor to retrieve the next page of BusinessIdentity structs
	var businessIdentities []BusinessIdentity
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &businessIdentities)
	if unmarshalError != nil {
		return businessIdentities, cursor, err
	}
	return businessIdentities, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (BusinessIdentity, Error.StarkErrors) {
	//	Update a BusinessIdentity entity
	//
	//	Update a BusinessIdentity by passing id.
	//
	//	Parameters (required):
	//	- id [string]: BusinessIdentity unique id. ex: "6306109539221504"
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated
	//		- status [string]: You may send BusinessIdentities to validation by passing 'processing' in the status
	//		- tags [slice of strings]: slice of strings for reference when searching for BusinessIdentities. ex: []string{"employees", "monthly"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- target BusinessIdentity with updated attributes
	var businessIdentity BusinessIdentity
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &businessIdentity)
	if unmarshalError != nil {
		return businessIdentity, err
	}
	return businessIdentity, err
}

func Cancel(id string, user user.User) (BusinessIdentity, Error.StarkErrors) {
	//	Cancel a BusinessIdentity entity
	//
	//	Cancel a BusinessIdentity by passing id.
	//
	//	Parameters (required):
	//	- id [string]: BusinessIdentity unique id. ex: "6306109539221504"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- canceled BusinessIdentity struct
	var businessIdentity BusinessIdentity
	cancel, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(cancel, &businessIdentity)
	if unmarshalError != nil {
		return businessIdentity, err
	}
	return businessIdentity, err
}
