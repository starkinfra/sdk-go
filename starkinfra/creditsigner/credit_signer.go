package creditsigner

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	CreditSigner struct
//
//	CreditNote signer's information.
//
//	Parameters (required):
//	- Name [string]: Signer's name. ex: "Tony Stark"
//	- Contact [string]: Signer's contact information. ex: "tony@starkindustries.com"
//	- Method [string]: Delivery method for the contract. ex: "link"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the CreditSigner is created. ex: "5656565656565656"

type CreditSigner struct {
	Name    string `json:",omitempty"`
	Contact string `json:",omitempty"`
	Method  string `json:",omitempty"`
	Id      string `json:",omitempty"`
}

var resource = map[string]string{"name": "CreditSigner"}

func ResendToken(signerId string, user user.User) (CreditSigner, Error.StarkErrors) {
	//	Resend token to signer
	//
	//	Resend token to a specific signer.
	//
	//	Parameters (required):
	//	- signerId [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- CreditSigner struct with updated attributes
	var signer CreditSigner
	payload := map[string]interface{}{"isSent": false}
	update, err := utils.Patch(resource, signerId, payload, user)
	unmarshalError := json.Unmarshal(update, &signer)
	if unmarshalError != nil {
		return signer, err
	}
	return signer, err
}
