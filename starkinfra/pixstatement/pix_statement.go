package pixstatement

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixStatement struct
//
//	The PixStatement struct stores information about all the transactions that
//	happened on a specific day at your settlment account according to the Central Bank.
//	It must be created by the user before it can be accessed.
//	This feature is only available for direct participants.
//	When you initialize a PixStatement, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- After [time.Time]: Transactions that happened at this date are stored in the PixStatement, must be the same as before. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- Before [time.Time]: Transactions that happened at this date are stored in the PixStatement, must be the same as after. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- Type [string]: Types of entities to include in statement. Options: ["interchange", "interchangeTotal", "transaction"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixStatement is created. ex: "5656565656565656"
//	- Status [string]: Current PixStatement status. ex: []string{"success", "failed"}
//	- TransactionCount [int]: Number of transactions that happened during the day that the PixStatement was requested. ex: 11
//	- Created [time.Time]: Creation datetime for the PixStatement. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the PixStatement. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixStatement struct {
	After            *time.Time `json:",omitempty"`
	Before           *time.Time `json:",omitempty"`
	Types            string     `json:",omitempty"`
	Id               string     `json:",omitempty"`
	Status           string     `json:",omitempty"`
	TransactionCount int        `json:",omitempty"`
	Created          *time.Time `json:",omitempty"`
	Updated          *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixStatement"}

func Create(statement PixStatement, user user.User) (PixStatement, Error.StarkErrors) {
	//	Create a PixStatement struct
	//
	//	Create a PixStatement linked to your Workspace in the Stark Infra API
	//
	//	Parameters (required):
	//	- statement [PixStatement struct]: PixStatement struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- PixStatement struct with updated attributes.
	create, err := utils.Single(resource, statement, user)
	unmarshalError := json.Unmarshal(create, &statement)
	if unmarshalError != nil {
		return statement, err
	}
	return statement, err
}

func Get(id string, user user.User) (PixStatement, Error.StarkErrors) {
	//	Retrieve a specific PixStatement by its id
	//
	//	Retrieve the PixStatement struct linked to your Workspace in the Stark Infra API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixStatement struct that corresponds to the given id.
	var pixStatement PixStatement
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixStatement)
	if unmarshalError != nil {
		return pixStatement, err
	}
	return pixStatement, err
}

func Query(params map[string]interface{}, user user.User) chan PixStatement {
	//	Retrieve PixStatement structs
	//
	//	Receive a channel of PixStatement structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixStatement structs with updated attributes
	var pixStatement PixStatement
	statements := make(chan PixStatement)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixStatement)
			if err != nil {
				print(err)
			}
			statements <- pixStatement
		}
		close(statements)
	}()
	return statements
}

func Csv(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a .csv PixStatement
	//
	//	Retrieve a specific PixStatement by its ID in a .csv file.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- .zip file containing a PixStatement in .csv format
	return utils.GetContent(resource, id, nil, user, "csv")
}

func Page(params map[string]interface{}, user user.User) ([]PixStatement, string, Error.StarkErrors) {
	//	Retrieve paged PixStatement structs
	//
	//	Receive a slice of up to 100 PixStatement structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixStatement structs with updated attributes
	//	- cursor to retrieve the next page of PixStatement structs
	var pixStatements []PixStatement
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixStatements)
	if unmarshalError != nil {
		return pixStatements, cursor, err
	}
	return pixStatements, cursor, err
}
