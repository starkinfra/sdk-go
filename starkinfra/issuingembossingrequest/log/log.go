package log

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	IssuingEmbossingRequest "github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingEmbossingRequest.Log struct
//
//	Every time a IssuingEmbossingRequest entity is updated, a corresponding IssuingEmbossingRequest.Log
//	is generated for the entity. This log is never generated by the
//	user, but it can be retrieved to check additional information
//	on the IssuingEmbossingRequest.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Request [IssuingEmbossingRequest struct]: IssuingEmbossingRequest entity to which the log refers to.
//	- Errors [slice of strings]: Slice of errors linked to this IssuingEmbossingRequest event
//	- Type [string]: Type of the IssuingEmbossingRequest event which triggered the log creation. ex: "registered" or "paid"
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 0, 0, time.UTC),

type Log struct {
	Id      string                                          `json:",omitempty"`
	Request IssuingEmbossingRequest.IssuingEmbossingRequest `json:",omitempty"`
	Errors  []string                                        `json:",omitempty"`
	Type    string                                          `json:",omitempty"`
	Created *time.Time                                      `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingEmbossingRequestLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific IssuingEmbossingRequest.Log by its id
	//
	//	Receive a single IssuingEmbossingRequest.Log struct previously created by the Stark Bank API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingEmbossingRequest.Log struct that corresponds to the given id.
	var issuingEmbossingRequestLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingEmbossingRequestLog)
	if unmarshalError != nil {
		return issuingEmbossingRequestLog, err
	}
	return issuingEmbossingRequestLog, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve IssuingEmbossingRequest.Log structs
	//
	//	Receive a channel of IssuingEmbossingRequest.Log structs previously created in the Stark Bank API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"created", "sending", "sent", "processing", "success", "failed"}
	//		- requestIds [slice of strings, default nil]: List of IssuingEmbossingRequest ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of note.Log structs with updated attributes
	var issuingEmbossingRequestLog Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingEmbossingRequestLog)
			if err != nil {
				print(err)
			}
			logs <- issuingEmbossingRequestLog
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged IssuingEmbossingRequest.Log structs
	//
	//	Receive a slice of up to 100 IssuingEmbossingRequest.Log structs previously created in the Stark Bank API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. It must be an int between 1 and 100. ex: 50
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"created", "sending", "sent", "processing", "success", "failed"}
	//		- requestIds [slice of strings, default nil]: List of IssuingEmbossingRequest ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingEmbossingRequest.Log structs with updated attributes
	//	- cursor to retrieve the next page of IssuingEmbossingRequest.Log structs
	var issuingEmbossingRequestLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingEmbossingRequestLogs)
	if unmarshalError != nil {
		return issuingEmbossingRequestLogs, cursor, err
	}
	return issuingEmbossingRequestLogs, cursor, err
}
