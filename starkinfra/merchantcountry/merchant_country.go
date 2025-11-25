package merchantcountry

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	MerchantCountry struct
//
//	MerchantCountry's codes are used to define country filters in IssuingRules.
//
//	Parameters (conditionally required):
//	- Code [string]: Country's code. ex: "BRA"
//
//	Attributes (return-only):
//	- Name [string]: Country's name. ex: "Brazil"
//  - Number [string]: Country's number. ex: "076"
//  - ShortCode [string]: Country's short code. ex: "BR"

type MerchantCountry struct {
	Code      string `json:",omitempty"`
	Name      string `json:",omitempty"`
	Number    string `json:",omitempty"`
	ShortCode string `json:",omitempty"`
}

var resource = map[string]string{"name": "MerchantCountry"}

func Query(params map[string]interface{}, user user.User) (chan MerchantCountry, chan Error.StarkErrors) {
	//	Retrieve MerchantCountry structs
	//
	//	Receive a channel of MerchantCountry structs available in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- search [string, default nil]: Keyword to search for code, name, number or shortCode
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of MerchantCountry structs with updated attributes
	var merchantCountry MerchantCountry
	countries := make(chan MerchantCountry)
	countriesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &merchantCountry)
			if err != nil {
				countriesError <- Error.UnknownError(err.Error())
				continue
			}
			countries <- merchantCountry
		}
		for err := range errorChannel {
			countriesError <- err
		}
		close(countries)
		close(countriesError)
	}()
	return countries, countriesError
}
