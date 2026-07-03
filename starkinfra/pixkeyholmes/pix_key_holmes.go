package pixkeyholmes

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixKeyHolmes struct
//
//	PixKeyHolmes are used to investigate the registration status of a Pix Key
//	in the Central Bank's DICT.
//
//	When you initialize a PixKeyHolmes, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the slice of created objects.
//
//	Parameters (required):
//	- KeyId [string]: Pix Key to be investigated. ex: "+5511989898989", "11.222.333/0001-00", "valid@sandbox.com"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for PixKeyHolmes. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixKeyHolmes is created. ex: "5656565656565656"
//	- Result [string]: Result of the investigation after the case is solved. ex: "registered", "unregistered"
//	- Status [string]: Current status of the PixKeyHolmes. ex: "created", "solving", "solved", "failed"
//	- Created [time.Time]: Creation datetime for the PixKeyHolmes. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the PixKeyHolmes. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixKeyHolmes struct {
	KeyId   string     `json:",omitempty"`
	Tags    []string   `json:",omitempty"`
	Id      string     `json:",omitempty"`
	Result  string     `json:",omitempty"`
	Status  string     `json:",omitempty"`
	Created *time.Time `json:",omitempty"`
	Updated *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixKeyHolmes"}

func Create(holmes []PixKeyHolmes, user user.User) ([]PixKeyHolmes, Error.StarkErrors) {
	//	Create PixKeyHolmes
	//
	//	Send a slice of PixKeyHolmes structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- holmes [slice of PixKeyHolmes structs]: Slice of PixKeyHolmes structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Slice of PixKeyHolmes structs with updated attributes
	create, err := utils.Multi(resource, holmes, nil, user)
	unmarshalError := json.Unmarshal(create, &holmes)
	if unmarshalError != nil {
		return holmes, err
	}
	return holmes, err
}

func Query(params map[string]interface{}, user user.User) (chan PixKeyHolmes, chan Error.StarkErrors) {
	//	Retrieve PixKeyHolmes
	//
	//	Receive a channel of PixKeyHolmes structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. The live API accepts only "solved" or "solving" as filter values. ex: []string{"solved", "solving"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Channel of PixKeyHolmes structs with updated attributes
	var pixKeyHolmes PixKeyHolmes
	holmes := make(chan PixKeyHolmes)
	holmesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixKeyHolmes)
			if err != nil {
				holmesError <- Error.UnknownError(err.Error())
				continue
			}
			holmes <- pixKeyHolmes
		}
		for err := range errorChannel {
			holmesError <- err
		}
		close(holmes)
		close(holmesError)
	}()
	return holmes, holmesError
}

func Page(params map[string]interface{}, user user.User) ([]PixKeyHolmes, string, Error.StarkErrors) {
	//	Retrieve paged PixKeyHolmes structs
	//
	//	Receive a slice of up to 100 PixKeyHolmes structs previously created in the Stark Infra API and the cursor to the next page
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. The live API accepts only "solved" or "solving" as filter values. ex: []string{"solved", "solving"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- Slice of PixKeyHolmes structs with updated attributes
	//	- Cursor to retrieve the next page of PixKeyHolmes structs
	var pixKeyHolmes []PixKeyHolmes
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixKeyHolmes)
	if unmarshalError != nil {
		return pixKeyHolmes, cursor, err
	}
	return pixKeyHolmes, cursor, err
}
