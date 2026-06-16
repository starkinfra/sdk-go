package pixdispute

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixDispute struct
//
//	A PixDispute is used to report a transaction that is suspected of fraud to the
//	Central Bank so that a graph analysis can be created to trace the funds.
//	When you initialize a PixDispute, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the slice of created structs.
//
//	Parameters (required):
//	- ReferenceId [string]: EndToEndId of the transaction being reported. ex: "E20018183202201201450u34sDGd19lz"
//	- Method [string]: Method of the dispute. Options: "scam", "unauthorized", "coercion", "invasion", "other"
//	- OperatorEmail [string]: Contact email of the operator responsible for the PixDispute.
//	- OperatorPhone [string]: Contact phone number of the operator responsible for the PixDispute.
//
//	Parameters (conditionally required):
//	- Description [string, default nil]: Details for any details that can help with the investigation. Required when Method is "other".
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//	- MinTransactionAmount [int, default nil]: Minimum transaction amount considered for the graph creation. ex: 11234 (= R$ 112.34)
//	- MaxTransactionCount [int, default nil]: Maximum number of transactions considered for the graph creation. ex: 100
//	- MaxHopInterval [int, default nil]: Mean time between transactions considered for the graph creation.
//	- MaxHopCount [int, default nil]: Depth considered for the graph creation.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixDispute is created. ex: "5656565656565656"
//	- BacenId [string]: Central Bank's unique dispute id. ex: "817fc523-9e9d-40ab-9e53-dacb71454a05"
//	- Flow [string]: Direction of the PixDispute flow. Options: "in" if you received the PixDispute, "out" if you created the PixDispute.
//	- Status [string]: Current PixDispute status. Options: "analysed", "canceled", "closed", "created", "delivered", "failed", "processing"
//	- Transactions [slice of PixDispute.Transaction structs]: Slice of transactions that make up the dispute graph.
//	- Created [time.Time]: Creation datetime for the PixDispute. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the PixDispute. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixDispute struct {
	ReferenceId          string        `json:",omitempty"`
	Method               string        `json:",omitempty"`
	OperatorEmail        string        `json:",omitempty"`
	OperatorPhone        string        `json:",omitempty"`
	Description          string        `json:",omitempty"`
	Tags                 []string      `json:",omitempty"`
	MinTransactionAmount int           `json:",omitempty"`
	MaxTransactionCount  int           `json:",omitempty"`
	MaxHopInterval       int           `json:",omitempty"`
	MaxHopCount          int           `json:",omitempty"`
	Id                   string        `json:",omitempty"`
	BacenId              string        `json:",omitempty"`
	Flow                 string        `json:",omitempty"`
	Status               string        `json:",omitempty"`
	Transactions         []Transaction `json:",omitempty"`
	Created              *time.Time    `json:",omitempty"`
	Updated              *time.Time    `json:",omitempty"`
}

var resource = map[string]string{"name": "PixDispute"}

func Create(disputes []PixDispute, user user.User) ([]PixDispute, Error.StarkErrors) {
	//	Create PixDispute structs
	//
	//	Send a slice of PixDispute structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- disputes [slice of PixDispute structs]: Slice of PixDispute structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixDispute structs with updated attributes
	create, err := utils.Multi(resource, disputes, nil, user)
	unmarshalError := json.Unmarshal(create, &disputes)
	if unmarshalError != nil {
		return disputes, err
	}
	return disputes, err
}

func Get(id string, user user.User) (PixDispute, Error.StarkErrors) {
	//	Retrieve a specific PixDispute
	//
	//	Receive a single PixDispute struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixDispute struct with updated attributes
	var pixDispute PixDispute
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixDispute)
	if unmarshalError != nil {
		return pixDispute, err
	}
	return pixDispute, err
}

func Query(params map[string]interface{}, user user.User) (chan PixDispute, chan Error.StarkErrors) {
	//	Retrieve PixDispute structs
	//
	//	Receive a channel of PixDispute structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "delivered", "analysed", "processing", "closed", "failed", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"travel", "food"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixDispute structs with updated attributes
	var pixDispute PixDispute
	disputes := make(chan PixDispute)
	disputesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixDispute)
			if err != nil {
				disputesError <- Error.UnknownError(err.Error())
				continue
			}
			disputes <- pixDispute
		}
		for err := range errorChannel {
			disputesError <- err
		}
		close(disputes)
		close(disputesError)
	}()
	return disputes, disputesError
}

func Page(params map[string]interface{}, user user.User) ([]PixDispute, string, Error.StarkErrors) {
	//	Retrieve paged PixDispute structs
	//
	//	Receive a slice of up to 100 PixDispute structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call.
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "delivered", "analysed", "processing", "closed", "failed", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"travel", "food"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixDispute structs with updated attributes
	//	- cursor to retrieve the next page of PixDispute structs
	var pixDisputes []PixDispute
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixDisputes)
	if unmarshalError != nil {
		return pixDisputes, cursor, err
	}
	return pixDisputes, cursor, err
}

func Cancel(id string, user user.User) (PixDispute, Error.StarkErrors) {
	//	Cancel a PixDispute entity
	//
	//	Cancel a PixDispute entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled PixDispute struct
	var pixDispute PixDispute
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &pixDispute)
	if unmarshalError != nil {
		return pixDispute, err
	}
	return pixDispute, err
}
