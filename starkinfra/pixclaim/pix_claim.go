package pixclaim

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixClaim struct
//
//	A Pix Claim is a request to transfer a Pix Key from an account hosted at another
//  Pix participant to an account under your bank code. Pix Claims must always be requested by the claimer.
//	When you initialize a PixClaim, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- AccountCreated [string]: Opening Date for the account claiming the PixKey. ex: "2022-01-01".
//	- AccountNumber [string]: Number of the account claiming the PixKey. ex: "76543".
//	- AccountType [string]: Type of the account claiming the PixKey. Options: "checking", "savings", "salary" or "payment".
//	- BranchCode [string]: Branch code of the account claiming the PixKey. ex: 1234".
//	- Name [string]: Holder's name of the account claiming the PixKey. ex: "Jamie Lannister".
//	- TaxId [string]: Holder's taxId of the account claiming the PixKey (CPF/CNPJ). ex: "012.345.678-90".
//	- KeyId [string]: Id of the registered Pix Key to be claimed. Allowed keyTypes are CPF, CNPJ, phone number or email. ex: "+5511989898989".
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixClaim is created. ex: "5656565656565656"
//	- Status [string]: Current PixClaim status. Options: "created", "failed", "delivered", "confirmed", "success", "canceled"
//	- Type [string]: Type of Pix Claim. Options: "ownership", "portability".
//	- KeyType [string]: KeyType of the claimed PixKey. Options: "CPF", "CNPJ", "phone" or "email"
//	- Flow [string]: Direction of the Pix Claim. Options: "in" if you received the PixClaim or "out" if you created the PixClaim.
//  - ClaimerBankCode [string]: BankCode of the Pix participant that created the PixClaim. ex: "20018183"
//  - ClaimedBankCode [string]: BankCode of the account donating the PixKey. ex: "20018183"
//  - Created [time.Time]: Creation datetime for the PixClaim. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Updated [time.Time]: Update datetime for the PixClaim. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixClaim struct {
	AccountCreated  string     `json:",omitempty"`
	AccountNumber   string     `json:",omitempty"`
	AccountType     string     `json:",omitempty"`
	BranchCode      string     `json:",omitempty"`
	Name            string     `json:",omitempty"`
	TaxId           string     `json:",omitempty"`
	KeyId           string     `json:",omitempty"`
	Tags            []string   `json:",omitempty"`
	Id              string     `json:",omitempty"`
	Status          string     `json:",omitempty"`
	Type            string     `json:",omitempty"`
	KeyType         string     `json:",omitempty"`
	Flow            string     `json:",omitempty"`
	ClaimerBankCode string     `json:",omitempty"`
	ClaimedBankCode string     `json:",omitempty"`
	Created         *time.Time `json:",omitempty"`
	Updated         *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixClaim"}

func Create(claim PixClaim, user user.User) (PixClaim, Error.StarkErrors) {
	//	Create a PixClaim struct
	//
	//	Create a PixClaim to request the transfer of a PixKey to an account
	//	hosted at other Pix participants in the Stark Infra API.
	//
	//	Parameters (required):
	//	- claim [PixClaim struct]: PixClaim struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixClaim struct with updated attributes.
	create, err := utils.Single(resource, claim, user)
	unmarshalError := json.Unmarshal(create, &claim)
	if unmarshalError != nil {
		return claim, err
	}
	return claim, err
}

func Get(id string, user user.User) (PixClaim, Error.StarkErrors) {
	//	Retrieve a PixClaim struct
	//
	//	Retrieve a PixClaim struct linked to your Workspace in the Stark Infra API by its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixClaim struct that corresponds to the given id.
	var pixClaim PixClaim
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixClaim)
	if unmarshalError != nil {
		return pixClaim, err
	}
	return pixClaim, err
}

func Query(params map[string]interface{}, user user.User) (chan PixClaim, chan Error.StarkErrors) {
	//	Retrieve PixClaim structs
	//
	//	Receive a channel of PixClaim structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "delivered", "confirmed", "success", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- type [strings, default nil]: Filter for the type of retrieved PixClaims. Options: "ownership" or "portability".
	//		- keyType [string, default nil]: Filter for the PixKey type of retrieved PixClaims. Options: "cpf", "cnpj", "phone", "email" and "evp"
	//  	- keyId [string, default nil]: Filter PixClaims linked to a specific PixKey id. ex: "+5511989898989"
	//  	- flow [string, default nil]: Direction of the Pix Claim. Options: "in" if you received the PixClaim or "out" if you created the PixClaim.
	//  	- tags [slice of strings, default nil]: Slice of strings to filter retrieved structs. ex: []string{"travel", "food"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- channel of PixClaim structs with updated attributes
	var pixClaim PixClaim
	claims := make(chan PixClaim)
	claimsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixClaim)
			if err != nil {
				claimsError <- Error.UnknownError(err.Error())
				continue
			}
			claims <- pixClaim
		}
		for err := range errorChannel {
			claimsError <- err
		}
		close(claims)
		close(claimsError)
	}()
	return claims, claimsError
}

func Page(params map[string]interface{}, user user.User) ([]PixClaim, string, Error.StarkErrors) {
	//	Retrieve paged PixClaim structs
	//
	//	Receive a slice of up to 100 PixClaim structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call.
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "delivered", "confirmed", "success", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- type [strings, default nil]: Filter for the type of retrieved PixClaims. Options: "ownership" or "portability".
	//		- keyType [string, default nil]: Filter for the PixKey type of retrieved PixClaims. Options: "cpf", "cnpj", "phone", "email" and "evp"
	//  	- keyId [string, default nil]: Filter PixClaims linked to a specific PixKey id. ex: "+5511989898989"
	//  	- flow [string, default nil]: Direction of the Pix Claim. Options: "in" if you received the PixClaim or "out" if you created the PixClaim.
	//  	- tags [slice of strings, default nil]: Slice of strings to filter retrieved structs. ex: []string{"travel", "food"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- slice of PixClaim structs with updated attributes
	//  - Cursor to retrieve the next page of PixClaim structs
	var pixClaims []PixClaim
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixClaims)
	if unmarshalError != nil {
		return pixClaims, cursor, err
	}
	return pixClaims, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (PixClaim, Error.StarkErrors) {
	//	Update PixClaim entity
	//
	//	Update a PixClaim parameters by passing id.
	//
	//	Parameters (required):
	//	- id [string]: PixClaim id. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//		Parameters (required):
	//		- status [string]: Patched status for Pix Claim. Options: "confirmed" and "canceled"
	//		Parameters (optional):
	//		- reason [string, default: "userRequested"]: Reason why the PixClaim is being patched. Options: "fraud", "userRequested", "accountClosure".
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixClaim with updated attributes
	var pixClaim PixClaim
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &pixClaim)
	if unmarshalError != nil {
		return pixClaim, err
	}
	return pixClaim, err
}
