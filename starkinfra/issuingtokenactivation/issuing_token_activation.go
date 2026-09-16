package issuingtokenactivation

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	IssuingTokenActivation struct
//
//	The IssuingTokenActivation struct displays the necessary information to proceed with the card tokenization.
//	You will receive this struct at your registered URL to notify you which method your user want to receive the activation code.
//
//	Attributes (return-only):
//	- CardId [string]: Card ID which the token is bounded to. ex: "5656565656565656"
//	- TokenId [string]: Token unique id. ex: "5656565656565656"
//	- Tags [slice of strings]: Tags to filter retrieved struct. ex: []string{"tony", "stark"}
//	- ActivationMethod [map[string]interface{}]: Dictionary object with "type":string and "value":string pairs

type IssuingTokenActivation struct {
	CardId           string                 `json:",omitempty"`
	TokenId          string                 `json:",omitempty"`
	Tags             []string               `json:",omitempty"`
	ActivationMethod map[string]interface{} `json:",omitempty"`
}

func Parse(content string, signature string, user user.User) (IssuingTokenActivation, Error.StarkErrors) {
	//	Create a single verified IssuingTokenActivation request from a content string
	//
	//	Use this method to parse and verify the authenticity of the request received at the informed endpoint.
	//	Activation requests are posted to your registered endpoint whenever IssuingTokenActivations are received.
	//	If the provided digital signature does not check out with the StarkInfra public key, a stark.exception.InvalidSignatureException will be raised.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- parsed IssuingTokenActivation struct
	var issuingTokenActivation IssuingTokenActivation
	parsed, err := utils.ParseAndVerify(content, signature, "", user)
	if err.Errors != nil {
		return issuingTokenActivation, err
	}

	unmarshalError := json.Unmarshal([]byte(parsed), &issuingTokenActivation)
	if unmarshalError != nil {
		return issuingTokenActivation, Error.UnknownError(unmarshalError.Error())
	}

	return issuingTokenActivation, Error.StarkErrors{}
}
