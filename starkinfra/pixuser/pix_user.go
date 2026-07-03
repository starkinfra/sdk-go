package pixuser

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	PixUser struct
//
//	Pix Users are used to get fraud statistics of a user.
//
//	Parameters (required):
//	- Id [string]: user tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//
//	Attributes (return-only):
//	- Statistics [slice of Statistics structs]: Slice of PixUser.Statistics structs with fraud statistics of the user. ex: []pixuser.Statistics{{Value: 3, Type: "infractions", Source: "keyManagement"}}

type PixUser struct {
	Statistics []Statistics `json:",omitempty"`
	Id         string       `json:",omitempty"`
}

var resource = map[string]string{"name": "PixUser"}

func Get(id string, user user.User) (PixUser, Error.StarkErrors) {
	//	Retrieve a PixUser struct
	//
	//	Retrieve a PixUser object information by passing its taxId
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656".
	//
	//	Parameters (optional):
	//	- keyId [string]: marked PixKey id. ex: "+5511989898989"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- PixUser struct that corresponds to the given id.
	var pixUser PixUser
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixUser)
	if unmarshalError != nil {
		return pixUser, err
	}
	return pixUser, err
}
