package issuingtokenrequest

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	IssuingTokenRequest struct
//
//	The IssuingTokenRequest object displays the necessary information to proceed with the card tokenization.
//
//	Parameters (required):
//	- CardId [string]: card id to be tokenized. ex: "5734340247945216"
//	- WalletId [string]: desired wallet to be integrated. ex: "google"
//	- MethodCode [string]: method code. ex: "app" or "manual"
//
//	Attributes (return-only):
//	- Content [string]: token request content. ex: "eyJwdWJsaWNLZXlGaW5nZXJwcmludCI6ICJlNTNiZThjZTRhYWQxNWU2OWNmMjExOTA5Mjk4YzJkOTE0O..."
//	- Signature [string]: token request signature. ex: "eyJwdWJsaWNLZXlGaW5nZXJwcmludCI6ICJlNTNiZThjZTRhYWQxNWU2OWNmMjExOTA5Mjk4YzJkOTE0O..."
//	- Metadata [map[string]interface{}]: map used to store additional information about the IssuingTokenRequest object.

type IssuingTokenRequest struct {
	CardId     string                 `json:",omitempty"`
	WalletId   string                 `json:",omitempty"`
	MethodCode string                 `json:",omitempty"`
	Content    string                 `json:",omitempty"`
	Signature  string                 `json:",omitempty"`
	Metadata   map[string]interface{} `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingTokenRequest"}

func Create(request IssuingTokenRequest, user user.User) (IssuingTokenRequest, Error.StarkErrors) {
	//	Create an IssuingTokenRequest
	//
	//	Send a single IssuingTokenRequest to the Stark Infra API to generate the payload to proceed with the card tokenization
	//
	//	Parameters (required):
	//	- request [IssuingTokenRequest struct]: IssuingTokenRequest struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- IssuingTokenRequest struct with updated attributes
	create, err := utils.Single(resource, request, user)
	unmarshalError := json.Unmarshal(create, &request)
	if unmarshalError != nil {
		return request, err
	}
	return request, err
}
