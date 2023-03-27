package log

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingHolder.Log struct
//
//	Every time an IssuingHolder entity is updated, a corresponding IssuingHolder.Log
//	is generated for the entity. This log is never generated by the
//	user, but it can be retrieved to check additional information
//	on the IssuingHolder.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Holder [IssuingHolder]: IssuingHolder entity to which the log refers to.
//	- Type [string]: Type of the IssuingHolder event which triggered the log creation. ex: "blocked", "canceled", "created", "unblocked", "updated"
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id      string                      `json:",omitempty"`
	Holder  IssuingHolder.IssuingHolder `json:",omitempty"`
	Type    string                      `json:",omitempty"`
	Created *time.Time                  `json:",omitempty"`
}

var object Log
var objects []Log
var resource = map[string]string{"name": "IssuingHolderLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific IssuingHolder.Log
	//
	//	Receive a single IssuingHolder.Log struct previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingHolder.Log struct with updated attributes
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve IssuingHolder.Log
	//
	//	Receive a channel of IssuingHolder.Log structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"created", "blocked"}
	//		- holderIds [slice of strings, default nil]: Slice of IssuingHolder ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingHolder.Log structs with updated attributes
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			logs <- object
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged IssuingHolder.Log
	//
	//	Receive a slice of up to 100 IssuingHolder.Log structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"created", "blocked"}
	//		- holderIds [slice of strings, default nil]: Slice of IssuingHolder ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingHolder.Log structs with updated attributes
	//	- cursor to retrieve the next page of IssuingHolder.Log structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
