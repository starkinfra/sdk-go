package pixdirector

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	PixDirector struct
//
//	Mandatory data that must be registered within the Central Bank for emergency contact purposes.
//	When you initialize a PixDirector, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the slice of created structs.
//
//	Parameters (required):
//	- Name [string]: Name of the PixDirector. ex: "Edward Stark".
//	- TaxId [string]: Tax ID (CPF) of the PixDirector. ex: "012.345.678-90"
//	- Phone [string]: Phone of the PixDirector. ex: "+551198989898"
//	- Email [string]: Email of the PixDirector. ex: "ned.stark@starkbank.com"
//	- Password [string]: Password of the PixDirector. ex: "12345678"
//	- TeamEmail [string]: Team email. ex: "pix.team@company.com"
//	- TeamPhones [slice of strings]: Slice of phones of the team. ex: []string{"+5511988889999", "+5511988889998"}
//
//	Attributes (return-only):
//	- Status [string]: Current PixDirector status. ex: "success"

type PixDirector struct {
	Name       string   `json:",omitempty"`
	TaxId      string   `json:",omitempty"`
	Phone      string   `json:",omitempty"`
	Email      string   `json:",omitempty"`
	Password   string   `json:",omitempty"`
	TeamEmail  string   `json:",omitempty"`
	TeamPhones []string `json:",omitempty"`
	Status     string   `json:",omitempty"`
}

var subResource = map[string]string{"name": "PixDirector"}

func Create(director PixDirector, user user.User) (PixDirector, Error.StarkErrors) {
	//	Create a PixDirector struct
	//
	//	Send a PixDirector struct for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- director [PixDirector struct]: Slice of PixDirector structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixDirector struct with updated attributes
	create, err := utils.Single(subResource, director, user)
	unmarshalError := json.Unmarshal(create, &director)
	if unmarshalError != nil {
		return director, err
	}
	return director, err
}
