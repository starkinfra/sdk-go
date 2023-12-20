package log

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixInfraction.Log struct
//
//	Every time a PixInfraction entity is modified, a corresponding PixInfraction.Log
//	is generated for the entity. This log is never generated by the user.
//
//	Attributes:
//	- Id [string]: Unique id returned when the log is created. ex: "5656565656565656"
//	- Infraction [PixInfraction struct]: PixInfraction entity to which the log refers to.
//	- Type [string]: Type of the PixInfraction event which triggered the log creation. ex: "created", "failed", "delivering", "delivered", "closed", "canceled"
//	- Errors [slice of strings]: Slice of errors linked to this PixInfraction event
//	- Created [time.Time]: Creation datetime for the log. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Log struct {
	Id         string                      `json:",omitempty"`
	Infraction PixInfraction.PixInfraction `json:",omitempty"`
	Type       string                      `json:",omitempty"`
	Errors     []string                    `json:",omitempty"`
	Created    *time.Time                  `json:",omitempty"`
}

var resource = map[string]string{"name": "PixInfractionLog"}

func Get(id string, user user.User) (Log, Error.StarkErrors) {
	//	Retrieve a specific PixInfraction.Log by its id
	//
	//	Receive a single PixInfraction.Log struct previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- PixInfraction.Log struct that corresponds to the given id.
	var pixInfractionLog Log
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixInfractionLog)
	if unmarshalError != nil {
		return pixInfractionLog, err
	}
	return pixInfractionLog, err
}

func Query(params map[string]interface{}, user user.User) chan Log {
	//	Retrieve PixInfraction.Log structs
	//
	//	Receive a channel of PixInfraction.Log structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- ids [slice of strings, default nil]: Log ids to filter PixInfraction Logs. ex: []string{"5656565656565656"}
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- types [slice of strings, default nil]: Filter retrieved structs by types. ex: []string{"created", "failed", "delivering", "delivered", "closed", "canceled"}
	//		- infractionIds [slice of strings, default nil]: Slice of PixInfraction IDs to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel  of PixInfraction.Log structs with updated attributes
	var pixInfractionLog Log
	logs := make(chan Log)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixInfractionLog)
			if err != nil {
				print(err)
			}
			logs <- pixInfractionLog
		}
		close(logs)
	}()
	return logs
}

func Page(params map[string]interface{}, user user.User) ([]Log, string, Error.StarkErrors) {
	//	Retrieve paged PixInfraction.Log structs
	//
	//	Receive a slice of up to 100 PixInfraction.Log structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your infractions.
	//
	//	Parameters (optional):
	//	- cursor [string, default nil]: Cursor returned on the previous page function call
	//	- ids [slice of strings, default nil]: Log ids to filter PixInfraction Logs. ex: []string{"5656565656565656"}
	//	- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//	- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//	- types [slice of strings, default nil]: Filter retrieved structs by types. ex: []string{"created", "failed", "delivering", "delivered", "closed", "canceled"}
	//	- infractionIds [slice of strings, default nil]: Slice of PixInfraction ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of PixInfraction.Log structs with updated attributes
	//	- Cursor to retrieve the next page of PixInfraction.Log structs
	var pixInfractionLogs []Log
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixInfractionLogs)
	if unmarshalError != nil {
		return pixInfractionLogs, cursor, err
	}
	return pixInfractionLogs, cursor, err
}
