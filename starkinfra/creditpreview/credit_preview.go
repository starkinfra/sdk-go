package creditpreview

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	CreditPreview struct
//
//	A CreditPreview is used to get information from a credit before taking it.
//	This resource can be used to preview credit notes
//
//	Parameters (required):
//	- Credit [CreditNotePreview struct]: Information preview of the informed credit.
//
//	Parameters (conditionally required):
//	- Type [string]: Credit type. ex: "credit-note"

type CreditPreview struct {
	Credit CreditNotePreview `json:",omitempty"`
	Type   string            `json:",omitempty"`
}

var subResource = map[string]string{"name": "CreditPreview"}

func Create(previews []CreditPreview, user user.User) ([]CreditPreview, Error.StarkErrors) {
	//	Create CreditPreviews
	//
	//	Send a slice of CreditPreview structs for processing in the Stark Infra API
	//
	//	Parameters (required):
	//	- previews [slice of CreditPreview structs]: Slice of CreditPreview structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of CreditPreview structs with updated attributes
	create, err := utils.Multi(subResource, previews, nil, user)
	unmarshalError := json.Unmarshal(create, &previews)
	if unmarshalError != nil {
		return previews, err
	}
	return previews, err
}
