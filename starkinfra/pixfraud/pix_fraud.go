package pixfraud

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixInfraction struct
//
//	PixFrauds are used to report a PixKey or taxId when a fraud
//	has been confirmed.
//	When you initialize a PixFraud, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- ExternalId [string]: EndToEndId or ReturnId of the transaction being reported. ex: "my_external_id"
//	- Type [string]: Type of PixFraud. Options: "identity", "mule", "scam", "other"
//	- TaxId [string]: User tax Id (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//
//	Parameters (optional):
// 	- KeyId [string]: Marked PixKey id. ex: "+5511989898989"
// 	- Tags [slice of strings, default nil]: List of strings for tagging. ex: ["fraudulent"]
//
//	Attributes (return-only):
// 	- Id [string]: Unique id returned when the PixFraud is created. ex: "5656565656565656"
// 	- BacenId [string]: Unique transaction id returned from Central Bank. ex: "ccf9bd9c-e99d-999e-bab9-b999ca999f99"
// 	- Status [string]: Current PixFraud status. Options: "created", "failed", "registered", "canceled".
// 	- Created [string]: Creation datetime for the PixFraud. ex: "2020-03-10 10:30:00.000000+00:00"
// 	- Updated [string]: Latest update datetime for the PixFraud. ex: "2020-03-10 10:30:00.000000+00:00"

type PixFraud struct {
	ExternalId     	 string     `json:",omitempty"`
	Type             string     `json:",omitempty"`
	TaxId 	         string     `json:",omitempty"`
	KeyId		     string     `json:",omitempty"`
	Tags             []string   `json:",omitempty"`
	Id               string     `json:",omitempty"`
	BacenId			 string		`json:",omitempty"`
	Status           string     `json:",omitempty"`
	Created          *time.Time `json:",omitempty"`
	Updated          *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixFraud"}

func Create(frauds []PixFraud, user user.User) ([]PixFraud, Error.StarkErrors) {
	//	Create PixFraud structs
	//
	//	Create PixFrauds in the Stark Infra API
	//
	//	Parameters (required):
	//	- frauds [slice of PixFraud structs]: Slice of PixFraud structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixFraud structs with updated attributes
	create, err := utils.Multi(resource, frauds, nil, user)
	unmarshalError := json.Unmarshal(create, &frauds)
	if unmarshalError != nil {
		return frauds, err
	}
	return frauds, err
}

func Get(id string, user user.User) (PixFraud, Error.StarkErrors) {
	//	Retrieve a PixFraud struct
	//
	//	Retrieve the PixFraud struct linked to your Workspace in the Stark Infra API using its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656".
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixFraud struct that corresponds to the given id.
	var pixFraud PixFraud
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixFraud)
	if unmarshalError != nil {
		return pixFraud, err
	}
	return pixFraud, err
}

func Query(params map[string]interface{}, user user.User) (chan PixFraud, chan Error.StarkErrors) {
	//	Retrieve PixFraud structs
	//
	//	Receive a channel of PixFraud structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "registered", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- bacenId [string, default nil]: Unique transaction id returned from Central Bank. ex: "ccf9bd9c-e99d-999e-bab9-b999ca999f99"
	//		- type [slice of strings, default nil]: Filter for the type of retrieved PixFrauds. Options: "reversal", "reversalChargeback"
	//  	- flow [string, default nil]: Direction of the PixFraud flow. Options: "out" if you created the PixFraud, "in" if you received the PixFraud.
	//  	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"fraudulent"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixFraud structs with updated attributes
	var pixFraud PixFraud
	frauds := make(chan PixFraud)
	fraudsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixFraud)
			if err != nil {
				fraudsError <- Error.UnknownError(err.Error())
				continue
			}
			frauds <- pixFraud
		}
		for err := range errorChannel {
			fraudsError <- err
		}
		close(frauds)
		close(fraudsError)
	}()
	return frauds, fraudsError
}

func Page(params map[string]interface{}, user user.User) ([]PixFraud, string, Error.StarkErrors) {
	//	Retrieve paged PixFraud structs.
	//
	//	Receive a slice of up to 100 PixFraud structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call.
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "registered", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- bacenId [string, default nil]: Unique transaction id returned from Central Bank. ex: "ccf9bd9c-e99d-999e-bab9-b999ca999f99"
	//		- type [slice of strings, default nil]: Filter for the type of retrieved PixFrauds. Options: "reversal", "reversalChargeback"
	//  	- tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"fraudulent"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixFraud structs with updated attributes
	//  - Cursor to retrieve the next page of PixFraud structs
	var pixFrauds []PixFraud
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixFrauds)
	if unmarshalError != nil {
		return pixFrauds, cursor, err
	}
	return pixFrauds, cursor, err
}

func Cancel(id string, user user.User) (PixFraud, Error.StarkErrors) {
	//	Cancel a PixFraud entity
	//
	//	Cancel a PixFraud entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled PixFraud struct
	var pixFraud PixFraud
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &pixFraud)
	if unmarshalError != nil {
		return pixFraud, err
	}
	return pixFraud, err
}
