package log

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixClaim.Log struct
//
//	Every time a PixClaim entity is modified, a corresponding PixClaim.Log
//	is generated for the entity. This log is never generated by the user.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Claim [PixClaim struct]: PixClaim entity to which the log refers to.
//	- Type [string]: Type of the PixClaim event which triggered the log creation. ex: "created", "failed", "delivering", "delivered", "confirming", "confirmed", "success", "canceling", "canceled"
//	- Errors [slice of strings]: Slice of errors linked to this PixClaim event
//	- Reason [string]: Reason why the PixClaim was modified, resulting in the Log. Options: "fraud", "userRequested", "accountClosure", "defaultOperation", "reconciliation"
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id      string            `json:",omitempty"`
	Claim   PixClaim.PixClaim `json:",omitempty"`
	Type    string            `json:",omitempty"`
	Errors  []string          `json:",omitempty"`
	Reason  string            `json:",omitempty"`
	Created *time.Time        `json:",omitempty"`
}

var Object Log
var objects []Log
var resource = map[string]string{"name": "PixClaimLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific PixClaim.Log by its id
	//
	//	Receive a single PixClaim.Log struct previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixClaim.Log struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &Object)
	if unmarshalError != nil {
		return Object, err
	}
	return Object, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve PixClaim.Log structs
	//
	//	Receive a channel of PixClaim.Log structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- ids [slice of strings, default nil]: Log ids to filter PixClaim Logs. ex: []string{"5656565656565656"}
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter retrieved structs by types. ex: []string{"created", "failed", "delivering", "delivered", "confirming", "confirmed", "success", "canceling", "canceled"}
	//		- claimIds [slice of strings, default nil]: Slice of PixClaim ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixClaim.Log structs with updated attributes
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &Object)
			if err != nil {
				print(err)
			}
			logs <- Object
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged PixClaim.Logs
	//
	//	Receive a slice of up to 100 PixClaim.Log structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your claims.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- ids [slice of strings, default nil]: Log ids to filter PixClaim Logs. ex: []string{"5656565656565656"}
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter retrieved structs by types. ex: []string{"created", "failed", "delivering", "delivered", "confirming", "confirmed", "success", "canceling", "canceled"}
	//		- claimIds [slice of strings, default nil]: Slice of PixClaim IDs to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixClaim.Log structs with updated attributes
	//	- cursor to retrieve the next page of PixClaim.Log structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
