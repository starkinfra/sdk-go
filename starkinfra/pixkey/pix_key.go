package pixkey

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixKey struct
//
//	PixKeys link bank account information to key ids.
//	Key ids are a convenient way to search and pass bank account information.
//	When you initialize a Pix Key, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- AccountCreated [time.Time]: Opening DateTime for the linked account. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- AccountNumber [string]: Number of the linked account. ex: "76543".
//	- AccountType [string]: Type of the linked account. Options: "checking", "savings", "salary" or "payment".
//	- BranchCode [string]: Branch code of the linked account. ex: 1234".
//	- Name [string]: Holder's name of the linked account. ex: "Jamie Lannister".
//	- TaxId [string]: Holder's taxId (CPF/CNPJ) of the linked account. ex: "012.345.678-90".
//
//	Parameters (optional):
//	- Id [string, default nil]: Id of the registered PixKey. Allowed types are: CPF, CNPJ, phone number or email. If this parameter is not passed, an EVP will be created. ex: "+5511989898989";
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for PixKeys. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Owned [time.Time]: Datetime when the key was owned by the holder. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- OwnerType [string]: Type of the owner of the PixKey. Options: "business" or "individual".
//	- Status [string]: Current PixKey status. Options: "created", "registered", "canceled", "failed"
//	- BankCode [string]: BankCode of the account linked to the Pix Key. ex: "20018183".
//	- BankName [string]: Name of the bank that holds the account linked to the PixKey. ex: "StarkBank"
//	- Type [string]: Type of the PixKey. Options: "cpf", "cnpj", "phone", "email" and "evp",
//	- Created [time.Time]: Creation datetime for the PixKey. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixKey struct {
	AccountCreated *time.Time `json:",omitempty"`
	AccountNumber  string     `json:",omitempty"`
	AccountType    string     `json:",omitempty"`
	BranchCode     string     `json:",omitempty"`
	Name           string     `json:",omitempty"`
	TaxId          string     `json:",omitempty"`
	Id             string     `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Owned          *time.Time `json:",omitempty"`
	OwnerType      string     `json:",omitempty"`
	Status         string     `json:",omitempty"`
	BankCode       string     `json:",omitempty"`
	BankName       string     `json:",omitempty"`
	Type           string     `json:",omitempty"`
	Created        *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixKey"}

func Create(key PixKey, user user.User) (PixKey, Error.StarkErrors) {
	//	Create a PixKey struct
	//
	//	Create a PixKey linked to a specific account in the Stark Infra API
	//
	//	Parameters (required):
	//	- key [PixKey struct]: PixKey struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixKey struct with updated attributes.
	create, err := utils.Single(resource, key, user)
	unmarshalError := json.Unmarshal(create, &key)
	if unmarshalError != nil {
		return key, err
	}
	return key, err
}
func Get(id string, query map[string]interface{}, user user.User) (PixKey, Error.StarkErrors) {
	//	Retrieve a PixKey struct
	//
	//	Retrieve the PixKey struct linked to your Workspace in the Stark Infra API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656".
	//  - query [map[string]interface{}]: map containing the attributes to be retrieved.
	//		Parameters (required):
	//		- payerId [string]: Tax id (CPF/CNPJ) of the individual or business requesting the PixKey information. This id is used by the Central Bank to limit request rates. ex: "20.018.183/0001-80".
	//		Parameters (optional):
	//		- endToEndId [string, default nil]: Central bank's unique transaction id. If the request results in the creation of a PixRequest, the same endToEndId should be used. If this parameter is not passed, one endToEndId will be automatically created. Example: "E00002649202201172211u34srod19le"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixKey struct that corresponds to the given id.
	var pixKey PixKey
	get, err := utils.Get(resource, id, query, user)
	unmarshalError := json.Unmarshal(get, &pixKey)
	if unmarshalError != nil {
		return pixKey, err
	}
	return pixKey, err
}

func Query(params map[string]interface{}, user user.User) chan PixKey {
	//	Retrieve PixKey structs
	//
	//	Receive a channel of PixKeys structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "registered", "canceled", "failed"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- type [string, default nil]: Filter for the type of retrieved PixKeys. Options: "cpf", "cnpj", "phone", "email" and "evp"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixKey structs with updated attributes
	var pixKey PixKey
	keys := make(chan PixKey)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixKey)
			if err != nil {
				print(err)
			}
			keys <- pixKey
		}
		close(keys)
	}()
	return keys
}

func Page(params map[string]interface{}, user user.User) ([]PixKey, string, Error.StarkErrors) {
	//	Retrieve paged PixKey structs
	//
	//	Receive a slice of up to 100 PixKey structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call.
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "registered", "canceled", "failed"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- type [string, default nil]: Filter for the type of retrieved PixKeys. Options: "cpf", "cnpj", "phone", "email" and "evp"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixKey structs with updated attributes
	//  - Cursor to retrieve the next page of PixKey structs
	var pixKeys []PixKey
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixKeys)
	if unmarshalError != nil {
		return pixKeys, cursor, err
	}
	return pixKeys, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (PixKey, Error.StarkErrors) {
	//	Update PixKey entity
	//
	//	Update a PixKey parameters by passing its id.
	//
	//	Parameters (required):
	//	- id [string]: PixKey id. Allowed types are: CPF, CNPJ, phone number or email. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//  	Parameters (required):
	//		- reason [string]: Reason why the PixKey is being patched. Options: "branchTransfer", "reconciliation" or "userRequested".
	//		Parameters (optional):
	//		- accountCreated [time.Time, default ""]: Opening DateTime for the account to be linked. ex: "2022-01-01.
	//		- accountNumber [string, default nil]: Number of the account to be linked. ex: "76543".
	//		- accountType [string, default nil]: Type of the account to be linked. Options: "checking", "savings", "salary" or "payment".
	//		- branchCode [string, default nil]: Branch code of the account to be linked. ex: 1234".
	//		- name [string, default nil]: Holder's name of the account to be linked. ex: "Jamie Lannister".
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixKey with updated attributes
	var pixKey PixKey
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &pixKey)
	if unmarshalError != nil {
		return pixKey, err
	}
	return pixKey, err
}

func Cancel(id string, user user.User) (PixKey, Error.StarkErrors) {
	//	Cancel a PixKey entity
	//
	//	Cancel a PixKey entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled pixKey struct
	var pixKey PixKey
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &pixKey)
	if unmarshalError != nil {
		return pixKey, err
	}
	return pixKey, err
}
