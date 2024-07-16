package request

import (
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/request"
)

func Get(path string, query map[string]interface{}, user user.User) (request.Response, Error.StarkErrors) {
	//	Retrieve any StarkInfra resource
    //	Receive a json of resources previously created in StarkInfra's API
	// 
    //	Parameters (required):
    //  - path [string]: StarkInfra resource's route. ex: "/invoice/"
    //  
    //  Parameters (optional):
    //  - user [Organization/Project object, default nil]: Organization or Project object. Not necessary if starkinfra.user was set before function call
    // 	- query [map[string]interface{}, default nil]: Query parameters. ex: {"limit": 1, "status": paid}
	// 
	//  Return:
    //  Retrieve paged resources
	content, err := utils.GetRaw(path, query, user, "Joker", false)
	return content, err
}

func Post(path string, body map[string][]map[string]interface{}, query map[string]interface{}, user user.User) (request.Response, Error.StarkErrors) {
	//	Create any StarkInfra resource
    //	Send a map of string to interface and create any StarkInfra resource objects
	// 
    //	Parameters (required):
    // 	- path [string]: StarkInfra resource's route. ex: "/invoice/"
    // 	- body [map[string][]map[string]interface{}]: request parameters. ex: {"invoices": {{"amount": 100, "name": "Iron Infra S.A.", "taxId": "20.018.183/0001-80"}}}
    // 
	//	Parameters (optional):
    // 	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
    // 	- query [map[string]interface{}, default nil]: Query parameters. ex: {"limit": 1, "status": paid}
	// 
	//	Return:
    //	- Retrieve created resources
	//
	content, err := utils.PostRaw(path, body, user, query, "Joker", false)
	return content, err
}

func Patch(path string, body map[string]interface{}, query map[string]interface{}, user user.User) (request.Response, Error.StarkErrors) {
	//	Update any StarkInfra resource
    //	Send a json with parameters of a single StarkInfra resource object and update it
	// 
    //	Parameters (required):
    // 	- path [string]: StarkInfra resource's route. ex: "/invoice/5699165527090460"
    // 	- body [map[string]interface{}]: request parameters. ex: {"amount": 100}
    //	
	// 	Parameters (optional):
    // 	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
    //	
	// 	Return:
    // 	- Retrieve updated resource
	content, err := utils.PatchRaw(path, body, user, query, "Joker", false)
	return content, err
}

func Put(path string, body map[string][]map[string]interface{}, query map[string]interface{}, user user.User) (request.Response, Error.StarkErrors) {
	//	Put any StarkInfra resource
    //  Send a json with parameters of a single StarkInfra resource object and create it, if the resource alredy exists, you will update it.
    // 
	//  Parameters (required):
    //  - path [string]: StarkInfra resource's route. ex: "/invoice"
    //  - body [[string]interface{}]: request parameters. ex: {"amount": 100}
    // 
	//  Parameters (optional):
    //  - user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
    // 
	//  Return:
    //  - json of the resource with updated attributes
	content, err := utils.PutRaw(path, body, user, query, "Joker", false)
	return content, err
}

func Delete(path string, user user.User) (request.Response, Error.StarkErrors) {
	//	Delete any StarkInfra resource
	//	Send a json with parameters of a single StarkInfra resource object and delete it
	// 
	//	Parameters (required):
	//	- path [string]: StarkInfra resource's route. ex: "/invoice/5699165527090460"
	// 
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	// 
	//	Return:
	//	- json of the resource with updated attributes
	content, err := utils.DeleteRaw(path, user, "Joker", false)
	return content, err
}