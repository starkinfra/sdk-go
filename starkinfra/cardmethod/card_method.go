package cardmethod

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	CardMethod struct
//
//	CardMethod's codes are used to define methods filters in IssuingRules.
//
//	Parameters (required):
//	- Code [string]: Method's code. Options: "chip", "token", "server", "manual", "magstripe", "contactless"
//
//	Attributes (return-only):
//	- Name [string]: Method's name. ex: "token"
//	- Number [string]: Method's number. ex: "81"

type CardMethod struct {
	Code   string `json:",omitempty"`
	Name   string `json:",omitempty"`
	Number string `json:",omitempty"`
}

var object CardMethod
var resource = map[string]string{"name": "CardMethod"}

func Query(params map[string]interface{}, user user.User) chan CardMethod {
	//	Retrieve CardMethod structs
	//
	//	Receive a channel of CardMethod structs available in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- search [string, default nil]: Keyword to search for code, name, number or shortCode
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of CardMethod structs with updated attributes
	methods := make(chan CardMethod)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			methods <- object
		}
		close(methods)
	}()
	return methods
}
