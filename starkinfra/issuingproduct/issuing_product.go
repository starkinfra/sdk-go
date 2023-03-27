package issuingproduct

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingProduct struct
//
//	The IssuingProduct struct displays information of registered card products to your Workspace.
//  They represent a group of cards that begin with the same numbers (id) and offer the same product to end customers.
//
//	Attributes (return-only):
//	- Id [string]: Unique card product number (BIN) registered within the card network. ex: "53810200"
//  - Network [string]: Card network flag. ex: "mastercard"
//  - FundingType [string]: Type of funding used for payment. ex: "credit", "debit"
//  - HolderType [string]: Holder type. ex: "business", "individual"
//  - Code [string]: Internal code from card flag informing the product. ex: "MRW", "MCO", "MWB", "MCS"
//  - Created [time.Time]: Creation datetime for the IssuingProduct. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingProduct struct {
	Id          string     `json:",omitempty"`
	Network     string     `json:",omitempty"`
	FundingType string     `json:",omitempty"`
	HolderType  string     `json:",omitempty"`
	Code        string     `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
}

var object IssuingProduct
var objects []IssuingProduct
var resource = map[string]string{"name": "IssuingProduct"}

func Query(params map[string]interface{}, user user.User) chan IssuingProduct {
	//	Retrieve IssuingProduct structs
	//
	//	Receive a channel of IssuingProduct structs previously registered in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingBin structs with updated attributes
	products := make(chan IssuingProduct)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			products <- object
		}
		close(products)
	}()
	return products
}

func Page(params map[string]interface{}, user user.User) ([]IssuingProduct, string, Error.StarkErrors) {
	//	Retrieve paged IssuingProduct structs
	//
	//	Receive a slice of up to 100 IssuingProduct structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingProduct structs with updated attributes
	//	- cursor to retrieve the next page of IssuingProduct structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
