package log

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	PixReversal "github.com/starkinfra/sdk-go/starkinfra/pixreversal"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixReversal.Log struct
//
//	Every time a PixReversal entity is modified, a corresponding PixReversal.Log
//	is generated for the entity. This log is never generated by the user.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Reversal [PixReversal struct]: PixReversal entity to which the log refers to.
//	- Type [string]: Type of the PixReversal event which triggered the log creation. ex: "sent", "denied", "failed", "created", "success", "approved", "credited", "refunded", "processing"
//	- Errors [slice of strings]: Slice of errors linked to this PixReversal event
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id       string                  `json:",omitempty"`
	Reversal PixReversal.PixReversal `json:",omitempty"`
	Type     string                  `json:",omitempty"`
	Errors   interface{}             `json:",omitempty"`
	Created  *time.Time              `json:",omitempty"`
}

var resource = map[string]string{"name": "PixReversalLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific PixReversal.Log by its id
	//
	//	Receive a single PixReversal.Log struct previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixReversal.Log struct with updated attributes
	var pixReversalLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixReversalLog)
	if unmarshalError != nil {
		return pixReversalLog, err
	}
	return pixReversalLog, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve PixReversal.Log structs
	//
	//	Receive a channel of PixReversal.Log structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter retrieved structs by types. Options: ["sent", "denied", "failed", "created", "success", "approved", "credited", "refunded", "processing"}
	//		- reversalIds [slice of strings, default nil]: Slice of PixReversal IDs to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixReversal.Log structs with updated attributes
	var pixReversalLog Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixReversalLog)
			if err != nil {
				print(err)
			}
			logs <- pixReversalLog
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged PixReversal.Log structs
	//
	//	Receive a slice of up to 100 PixReversal.Log structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your reversals.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter retrieved structs by types. Options: ["sent", "denied", "failed", "created", "success", "approved", "credited", "refunded", "processing"}
	//		- reversalIds [slice of strings, default nil]: Slice of PixReversal IDs to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixReversal.Log structs with updated attributes
	//	- cursor to retrieve the next page of PixRequest.Log structs
	var pixReversalLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixReversalLogs)
	if unmarshalError != nil {
		return pixReversalLogs, cursor, err
	}
	return pixReversalLogs, cursor, err
}
