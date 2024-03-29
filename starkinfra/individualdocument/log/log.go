package log

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	IndividualDocument "github.com/starkinfra/sdk-go/starkinfra/individualdocument"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IndividualDocument.Log struct
//
//	Every time an IndividualDocument entity is updated, a corresponding IndividualDocument.Log
//	is generated for the entity. This log is never generated by the
//	user, but it can be retrieved to check additional information
//	on the IndividualDocument.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Document [IndividualDocument struct]: IndividualDocument entity to which the log refers to.
//	- Errors [slice of strings]: Slice of errors linked to this CreditNote event
//	- Type [string]: Type of the IndividualDocument event which triggered the log creation. ex: "blocked", "canceled", "created", "expired", "unblocked", "updated"
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id       string                                `json:",omitempty"`
	Document IndividualDocument.IndividualDocument `json:",omitempty"`
	Errors   []string                              `json:",omitempty"`
	Type     string                                `json:",omitempty"`
	Created  *time.Time                            `json:",omitempty"`
}

var resource = map[string]string{"name": "IndividualDocumentLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific IndividualDocument.Log
	//
	//	Receive a single IndividualDocument.Log struct previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- individualDocument.Log struct that corresponds to the given id.
	var individualDocumentLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &individualDocumentLog)
	if unmarshalError != nil {
		return individualDocumentLog, err
	}
	return individualDocumentLog, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve IndividualDocument.Log
	//
	//	Receive a channel of IndividualDocument.Log structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"created", "canceled", "processing", "failed", "success"}
	//		- documentsIds [slice of strings, default nil]: list of IndividualDocument ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IndividualDocument.Log structs with updated attributes
	var individualDocumentLog Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &individualDocumentLog)
			if err != nil {
				print(err)
			}
			logs <- individualDocumentLog
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged IndividualDocument.Log
	//
	//	Receive a slice of up to 100 IndividualDocument.Log structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter for log event types. ex: []string{"created", "canceled", "processing", "failed", "success"}
	//		- documentsIds [slice of strings, default nil]: list of IndividualDocument ids to filter logs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualDocument.Log structs with updated attributes
	//	- cursor to retrieve the next page of IndividualDocument.Log structs
	var individualDocumentLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &individualDocumentLogs)
	if unmarshalError != nil {
		return individualDocumentLogs, cursor, err
	}
	return individualDocumentLogs, cursor, err
}
