package businessaccountrequest

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	BusinessAccountRequest struct
//
//	You can create a business account request to request an account for a specific company, opening the
//	account with identity verification by webview for each of its owners.
//
//	When you initialize a BusinessAccountRequest, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the slice of created objects.
//
//	Parameters (required):
//	- Address [businessaccountrequest.Address struct]: company's structured address. ex: businessaccountrequest.Address{Street: "Av. Faria Lima", Number: "2000", Neighborhood: "Itaim Bibi", City: "Sao Paulo", State: "SP", ZipCode: "04538-132"}
//	- Revenue [int]: company's annual revenue in cents. ex: 100000000 (= R$ 1,000,000.00)
//	- Name [string]: company's legal name (minimum 5 characters). ex: "Stark Bank S.A."
//	- TaxId [string]: company's tax ID (CNPJ). ex: "20.018.183/0001-80"
//	- Owners [slice of businessaccountrequest.Owner structs]: slice of 1 to 10 company owners. ex: []businessaccountrequest.Owner{{TaxId: "012.345.678-90", Name: "Jamie Lannister", Role: "partner"}}
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for BusinessAccountRequests. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the BusinessAccountRequest is created. ex: "5656565656565656"
//	- AccountType [string]: type of the account. ex: "business"
//	- Flags [slice of maps]: flags that motivated the decision, populated when the request is denied. Each flag has a code and a message. ex: []map[string]interface{}{{"code": "failedIdentityProof", "message": "O representante: 012.345.678-90 falhou na verificação de identidade."}}
//	- Status [string]: current status of the BusinessAccountRequest. Options: "created", "processing", "approved", "denied", "failed"
//	- Created [time.Time]: creation datetime for the BusinessAccountRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the BusinessAccountRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type BusinessAccountRequest struct {
	Address     Address                  `json:",omitempty"`
	Revenue     int                      `json:",omitempty"`
	Name        string                   `json:",omitempty"`
	TaxId       string                   `json:",omitempty"`
	Owners      []Owner                  `json:",omitempty"`
	Tags        []string                 `json:",omitempty"`
	Id          string                   `json:",omitempty"`
	AccountType string                   `json:",omitempty"`
	Flags       []map[string]interface{} `json:",omitempty"`
	Status      string                   `json:",omitempty"`
	Created     *time.Time               `json:",omitempty"`
	Updated     *time.Time               `json:",omitempty"`
}

var resource = map[string]string{"name": "BusinessAccountRequest"}

func Create(requests []BusinessAccountRequest, user user.User) ([]BusinessAccountRequest, Error.StarkErrors) {
	//	Create BusinessAccountRequests
	//
	//	Send a slice of BusinessAccountRequest structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- requests [slice of BusinessAccountRequest structs]: slice of BusinessAccountRequest structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of BusinessAccountRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return requests, err
	}
	return requests, err
}

func Get(id string, user user.User) (BusinessAccountRequest, Error.StarkErrors) {
	//	Retrieve a specific BusinessAccountRequest
	//
	//	Receive a single BusinessAccountRequest struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- BusinessAccountRequest struct that corresponds to the given id.
	var request BusinessAccountRequest
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &request)
	if unmarshalError != nil {
		return request, err
	}
	return request, err
}

func Query(params map[string]interface{}, user user.User) (chan BusinessAccountRequest, chan Error.StarkErrors) {
	//	Retrieve BusinessAccountRequests
	//
	//	Receive a channel of BusinessAccountRequest structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. Options: "created", "processing", "approved", "denied", "failed"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of BusinessAccountRequest structs with updated attributes
	requests := make(chan BusinessAccountRequest)
	requestsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var request BusinessAccountRequest
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &request)
			if err != nil {
				requestsError <- Error.UnknownError(err.Error())
				continue
			}
			requests <- request
		}
		for err := range errorChannel {
			requestsError <- err
		}
		close(requests)
		close(requestsError)
	}()
	return requests, requestsError
}

func Page(params map[string]interface{}, user user.User) ([]BusinessAccountRequest, string, Error.StarkErrors) {
	//	Retrieve paged BusinessAccountRequests
	//
	//	Receive a slice of up to 100 BusinessAccountRequest structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: maximum number of structs to be retrieved. It must be an integer between 1 and 100. ex: 50
	//		- after [string, default nil]: date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. Options: "created", "processing", "approved", "denied", "failed"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of BusinessAccountRequest structs with updated attributes
	//	- cursor to retrieve the next page of BusinessAccountRequest structs
	var requests []BusinessAccountRequest
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &requests)
	if unmarshalError != nil {
		return requests, cursor, err
	}
	return requests, cursor, err
}
