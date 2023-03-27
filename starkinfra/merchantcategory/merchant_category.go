package merchantcategory

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	MerchantCategory struct
//
//	MerchantCategory's codes and types are used to define categories filters in IssuingRules.
//  A MerchantCategory filter must define exactly one parameter between code and type.
//  A type, such as "food", "services", etc., defines an entire group of merchant codes,
//  whereas a code only specifies a specific MCC.
//
//	Parameters (conditionally required):
//	- Code [string, default nil]: Category's code. ex: "veterinaryServices", "fastFoodRestaurants"
//  - Type [string, default nil]: Category's type. ex: "pets", "food"
//
//	Attributes (return-only):
//	- Name [string]: Category's name. ex: "Veterinary services", "Fast food restaurants"
//  - Number [string]: Category's number. ex: "742", "5814"

type MerchantCategory struct {
	Code   string `json:",omitempty"`
	Type   string `json:",omitempty"`
	Name   string `json:",omitempty"`
	Number string `json:",omitempty"`
}

var object MerchantCategory
var resource = map[string]string{"name": "MerchantCategory"}

func Query(params map[string]interface{}, user user.User) chan MerchantCategory {
	//	Retrieve MerchantCategory structs
	//
	//	Receive a channel of MerchantCategory structs available in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- search [string, default nil]: Keyword to search for code, type, name or number
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of MerchantCategory structs with updated attributes
	categories := make(chan MerchantCategory)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			categories <- object
		}
		close(categories)
	}()
	return categories
}
