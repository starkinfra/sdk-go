package issuingembossingrequest

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingEmbossingRequest struct
//
//  The IssuingEmbossingRequest object displays the information of embossing requests in your Workspace.
//
//	Parameters (required):
//	- CardId [string]: Id of the IssuingCard to be embossed. ex "5656565656565656"
//	- KitId [string]: Card embossing kit id. ex "5656565656565656"
//	- DisplayName1 [string]: Card displayed name. ex: "ANTHONY STARK"
//	- ShippingCity [string]: Shipping city. ex: "NEW YORK"
//	- ShippingCountryCode [string]: Shipping country code. ex: "US"
//	- ShippingDistrict [string]: Shipping district. ex: "NY"
//	- ShippingStateCode [string]: Shipping state code. ex: "NY"
//	- ShippingStreetLine1 [string]: Shipping main address. ex: "AVENUE OF THE AMERICAS"
//	- ShippingStreetLine2 [string]: Shipping address complement. ex: "Apt. 6"
//	- ShippingService [string]: Shipping service. ex: "loggi"
//	- ShippingTrackingNumber [string]: Shipping tracking number. ex: "5656565656565656"
//	- ShippingZipCode [string]: Shipping zip code. ex: "12345-678"
//
//	Parameters (optional):
//	- EmbosserId [string, default nil]: Id of the card embosser. ex: "5656565656565656"
//	- DisplayName2 [string, default nil]: Card displayed name. ex: "IT Services"
//	- DisplayName3 [string, default nil]: Card displayed name. ex: "StarkBank S.A."
//	- ShippingPhone [string, default nil]: Shipping phone. ex: "+5511999999999"
//	- Tags [slice of string, default nil]: Slice of strings for tagging. ex: []string{"card", "corporate"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingEmbossingRequest is created. ex: "5656565656565656"
//	- Fee [string]: Fee charged when IssuingEmbossingRequest is created. ex: 1000
//	- Status [string]: Status of the IssuingEmbossingRequest. ex: "created", "processing", "success", "failed"
//  - Updated [time.Time]: Latest update datetime for the IssuingEmbossingRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingEmbossingRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingEmbossingRequest struct {
	CardId                 string     `json:",omitempty"`
	KitId                  string     `json:",omitempty"`
	DisplayName1           string     `json:",omitempty"`
	ShippingCity           string     `json:",omitempty"`
	ShippingCountryCode    string     `json:",omitempty"`
	ShippingDistrict       string     `json:",omitempty"`
	ShippingStateCode      string     `json:",omitempty"`
	ShippingStreetLine1    string     `json:",omitempty"`
	ShippingStreetLine2    string     `json:",omitempty"`
	ShippingService        string     `json:",omitempty"`
	ShippingTrackingNumber string     `json:",omitempty"`
	ShippingZipCode        string     `json:",omitempty"`
	EmbosserId             string     `json:",omitempty"`
	DisplayName2           string     `json:",omitempty"`
	DisplayName3           string     `json:",omitempty"`
	ShippingPhone          string     `json:",omitempty"`
	Tags                   []string   `json:",omitempty"`
	Id                     string     `json:",omitempty"`
	Fee                    int        `json:",omitempty"`
	Status                 string     `json:",omitempty"`
	Updated                *time.Time `json:",omitempty"`
	Created                *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingEmbossingRequest"}

func Create(requests []IssuingEmbossingRequest, user user.User) ([]IssuingEmbossingRequest, Error.StarkErrors) {
	//	Create IssuingEmbossingRequests
	//
	//	Send a slice of IssuingEmbossingRequest structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- requests [slice of IssuingEmbossingRequest structs]: Slice of IssuingEmbossingRequest structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- slice of IssuingEmbossingRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return requests, err
	}
	return requests, err
}

func Get(id string, user user.User) (IssuingEmbossingRequest, Error.StarkErrors) {
	//	Retrieve a specific IssuingEmbossingRequest by its id
	//
	//	Receive a single IssuingEmbossingRequest struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingEmbossingRequest struct that corresponds to the given id.
	var issuingEmbossingRequest IssuingEmbossingRequest
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingEmbossingRequest)
	if unmarshalError != nil {
		return issuingEmbossingRequest, err
	}
	return issuingEmbossingRequest, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingEmbossingRequest {
	//	Retrieve IssuingEmbossingRequests
	//
	//	Receive a channel of IssuingEmbossingRequest structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed"}
	//  	- cardIds [slice of strings, default nil]: Slice of cardIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//  	- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//  	- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingEmbossingRequest structs with updated attributes
	var issuingEmbossingRequest IssuingEmbossingRequest
	requests := make(chan IssuingEmbossingRequest)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingEmbossingRequest)
			if err != nil {
				print(err)
			}
			requests <- issuingEmbossingRequest
		}
		close(requests)
	}()
	return requests
}

func Page(params map[string]interface{}, user user.User) ([]IssuingEmbossingRequest, string, Error.StarkErrors) {
	//	Retrieve paged IssuingEmbossingRequest structs
	//
	//	Receive a slice of up to 100 IssuingEmbossingRequest structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed"}
	//  	- cardIds [slice of strings, default nil]: Slice of cardIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//  	- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//  	- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingEmbossingRequest structs with updated attributes
	//	- cursor to retrieve the next page of IssuingEmbossingRequest structs
	var issuingEmbossingRequests []IssuingEmbossingRequest
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingEmbossingRequests)
	if unmarshalError != nil {
		return issuingEmbossingRequests, cursor, err
	}
	return issuingEmbossingRequests, cursor, err
}
