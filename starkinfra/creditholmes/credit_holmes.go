package creditholmes

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	CreditHolmes struct
//
//	CreditHolmes are used to obtain debt information on your customers.
//	Before you create a CreditHolmes, make sure you have your customer's express
//	authorization to verify their information in the Central Bank's SCR.
//
//	When you initialize a CreditHolmes, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the slice of created objects.
//
//	Parameters (required):
//	- TaxId [string]: Customer's tax ID (CPF or CNPJ) for whom the credit operations will be verified. ex: "20.018.183/0001-80"
//
//	Parameters (optional):
//	- Competence [string, default 'two months before current date']: competence month of the operation verification. ex: "2006-01"
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for CreditHolmes. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the CreditHolmes is created. ex: "5656565656565656"
//	- Result [map]: result of the investigation after the case is solved.
//	- Status [string]: current status of the CreditHolmes. ex: "created", "failed", "success"
//	- Created [time.Time]: Creation datetime for the CreditHolmes. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the CreditHolmes. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type CreditHolmes struct {
	TaxId      string                 `json:",omitempty"`
	Competence string                 `json:",omitempty"`
	Tags       []string               `json:",omitempty"`
	Id         string                 `json:",omitempty"`
	Result     map[string]interface{} `json:",omitempty"`
	Status     string                 `json:",omitempty"`
	Created    *time.Time             `json:",omitempty"`
	Updated    *time.Time             `json:",omitempty"`
}

var object CreditHolmes
var objects []CreditHolmes
var resource = map[string]string{"name": "CreditHolmes"}

func Create(holmes []CreditHolmes, user user.User) ([]CreditHolmes, Error.StarkErrors) {
	//	Create CreditHolmes
	//
	//	Send a slice of CreditHolmes structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- holmes [slice of CreditHolmes structs]: Slice of CreditHolmes structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Slice of CreditHolmes structs with updated attributes
	create, err := utils.Multi(resource, holmes, nil, user)
	unmarshalError := json.Unmarshal(create, &holmes)
	if unmarshalError != nil {
		return holmes, err
	}
	return holmes, err
}

func Get(id string, user user.User) (CreditHolmes, Error.StarkErrors) {
	//	Retrieve a specific CreditHolmes
	//
	//	Receive a single CreditHolmes struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- CreditHolmes struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan CreditHolmes {
	//	Retrieve CreditHolmes
	//
	//	Receive a channel of CreditHolmes structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: "created", "failed", "success"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Channel  of CreditHolmes structs with updated attributes
	holmes := make(chan CreditHolmes)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			holmes <- object
		}
		close(holmes)
	}()
	return holmes
}

func Page(params map[string]interface{}, user user.User) ([]CreditHolmes, string, Error.StarkErrors) {
	//	Retrieve paged CreditHolmes structs
	//
	//	Receive a slice of up to 100 CreditHolmes structs previously created in the Stark Infra API and the cursor to the next page
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: "created", "failed", "success"
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Slice of CreditHolmes structs with updated attributes
	//	- Cursor to retrieve the next page of CreditHolmes structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
