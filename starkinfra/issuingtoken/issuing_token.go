package issuingtoken

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingToken struct
//
//	The IssuingToken struct displays the information of the tokens created in your Workspace.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingToken is created. ex: "5656565656565656"
//	- CardId [string]: Card ID which the token is bounded to. ex: "5656565656565656"
//	- WalletId [string]: Wallet provider which the token is bounded to. ex: "google"
//	- WalletName [string]: Wallet name. ex: "GOOGLE"
//	- MerchantId [string]: Merchant unique id. ex: "5656565656565656"
//	- Url [string]: Token URL. ex: "https://token.starkinfra.com/5656565656565656"
//	- Updated [time.Time]: Latest update datetime for the IssuingToken. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingToken. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//
//	Parameters (optional):
//	- ExternalId [string]: A unique string among all your IssuingTokens, used to avoid resource duplication. ex: "DSHRMC00002626944b0e3b539d4d459281bdba90c2588791"
//	- Tags [slice of strings]: Slice of strings for reference when searching for IssuingToken. ex: []string{"employees", "monthly"}
//	- Status [string]: Current IssuingToken status. ex: "active", "blocked", "canceled", "frozen" or "pending"
//	- ActivationCode [string]: Activation code received through the bank app or SMS. ex: "481632"
//	- MethodCode [string]: Provisioning method. Options: "app", "token", "manual", "server" or "browser"
//	- DeviceType [string]: Device type used for tokenization. ex: "Phone"
//	- DeviceName [string]: Device name used for tokenization. ex: "My phone"
//	- DeviceSerialNumber [string]: Device serial number used for tokenization. ex: "2F6D63"
//	- DeviceOsName [string]: Device operational system name used for tokenization. ex: "Android"
//	- DeviceOsVersion [string]: Device operational system version used for tokenization. ex: "4.4.4"
//	- DeviceImei [string]: Device imei used for tokenization. ex: "352099001761481"
//	- WalletInstanceId [string]: Unique id referred to the wallet app in the current device. ex: "71583be4777eb89aaf0345eebeb82594f096615ed17862d0"

type IssuingToken struct {
	Id                 string     `json:",omitempty"`
	CardId             string     `json:",omitempty"`
	WalletId           string     `json:",omitempty"`
	WalletName         string     `json:",omitempty"`
	MerchantId         string     `json:",omitempty"`
	Url                string     `json:",omitempty"`
	ExternalId         string     `json:",omitempty"`
	Tags               []string   `json:",omitempty"`
	Status             string     `json:",omitempty"`
	ActivationCode     string     `json:",omitempty"`
	MethodCode         string     `json:",omitempty"`
	DeviceType         string     `json:",omitempty"`
	DeviceName         string     `json:",omitempty"`
	DeviceSerialNumber string     `json:",omitempty"`
	DeviceOsName       string     `json:",omitempty"`
	DeviceOsVersion    string     `json:",omitempty"`
	DeviceImei         string     `json:",omitempty"`
	WalletInstanceId   string     `json:",omitempty"`
	Updated            *time.Time `json:",omitempty"`
	Created            *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingToken"}

func Get(id string, user user.User) (IssuingToken, Error.StarkErrors) {
	//	Retrieve a specific IssuingToken by its id
	//
	//	Receive a single IssuingToken struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingToken struct that corresponds to the given id.
	var issuingToken IssuingToken
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingToken)
	if unmarshalError != nil {
		return issuingToken, err
	}
	return issuingToken, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingToken, chan Error.StarkErrors) {
	//	Retrieve IssuingTokens
	//
	//	Receive a channel of IssuingToken structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled", "frozen", "pending"}
	//		- cardIds [slice of strings, default nil]: Slice of cardIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"travel", "food"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"DSHRMC00002626944b0e3b539d4d459281bdba90c2588791", "DSHRMC00002626941c531164a0b14c66ad9602ee716f1e85"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingToken structs with updated attributes
	var issuingToken IssuingToken
	tokens := make(chan IssuingToken)
	tokensError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingToken)
			if err != nil {
				tokensError <- Error.UnknownError(err.Error())
				continue
			}
			tokens <- issuingToken
		}
		for err := range errorChannel {
			tokensError <- err
		}
		close(tokens)
		close(tokensError)
	}()
	return tokens, tokensError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingToken, string, Error.StarkErrors) {
	//	Retrieve paged IssuingTokens
	//
	//	Receive a slice of up to 100 IssuingToken structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"active", "blocked", "canceled", "frozen", "pending"}
	//		- cardIds [slice of strings, default nil]: Slice of cardIds to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"travel", "food"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- externalIds [slice of strings, default nil]: External IDs. ex: []string{"DSHRMC00002626944b0e3b539d4d459281bdba90c2588791", "DSHRMC00002626941c531164a0b14c66ad9602ee716f1e85"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingToken structs with updated attributes
	//	- cursor to retrieve the next page of IssuingToken structs
	var issuingTokens []IssuingToken
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingTokens)
	if unmarshalError != nil {
		return issuingTokens, cursor, err
	}
	return issuingTokens, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (IssuingToken, Error.StarkErrors) {
	//	Update IssuingToken entity
	//
	//	Update an IssuingToken by passing id.
	//
	//	Parameters (required):
	//	- id [string]: IssuingToken id. ex: "5656565656565656"
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"status": "blocked"}
	//		Parameters (optional):
	//		- status [string]: You may block the IssuingToken by passing "blocked" or activate by passing "active" in the status. ex: "active", "blocked"
	//		- tags [slice of strings]: Slice of strings for tagging. ex: []string{"travel", "food"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- target IssuingToken with updated attributes
	var issuingToken IssuingToken
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &issuingToken)
	if unmarshalError != nil {
		return issuingToken, err
	}
	return issuingToken, err
}

func Cancel(id string, user user.User) (IssuingToken, Error.StarkErrors) {
	//	Cancel an IssuingToken entity
	//
	//	Cancel an IssuingToken entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: IssuingToken unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled IssuingToken struct
	var issuingToken IssuingToken
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &issuingToken)
	if unmarshalError != nil {
		return issuingToken, err
	}
	return issuingToken, err
}

func Parse(content string, signature string, user user.User) (IssuingToken, Error.StarkErrors) {
	//	Create a single verified IssuingToken request from a content string
	//
	//	Use this method to parse and verify the authenticity of the request received at the informed endpoint.
	//	Token requests are posted to your registered endpoint whenever IssuingTokens are received.
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
	//	- parsed IssuingToken struct
	var issuingToken IssuingToken
	parsed, err := utils.ParseAndVerify(content, signature, "", user)
	if err.Errors != nil {
		return issuingToken, err
	}

	unmarshalError := json.Unmarshal([]byte(parsed), &issuingToken)
	if unmarshalError != nil {
		return issuingToken, Error.UnknownError(unmarshalError.Error())
	}

	return issuingToken, Error.StarkErrors{}
}

func ResponseAuthorization(authorization map[string]interface{}) string {
	//	Helps you respond IssuingToken authorization requests
	//
	//	When a new tokenization is triggered by your user, a POST request will be made to your registered URL to get your decision to complete the tokenization.
	//	The POST request must be answered in the following format, within 2 seconds, and with an HTTP status code 200.
	//
	//	Parameters (required):
	//	- status [string]: Sub-issuer response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: Denial reason. Options: "other", "bruteForce", "subIssuerError", "lostCard", "invalidCard", "invalidHolder", "expiredCard", "canceledCard", "blockedCard", "invalidExpiration", "invalidSecurityCode", "missingTokenAuthorizationUrl", "maxCardTriesExceeded", "maxWalletInstanceTriesExceeded"
	//	- activationMethods [slice of maps]: Slice of maps with "type": string and "value": string pairs
	//	- designId [string]: Design unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- tags [slice of strings]: Tags to filter retrieved struct. ex: []string{"tony", "stark"}
	//
	//	Return:
	//	- dumped JSON string that must be returned to us on the IssuingToken request
	params := map[string]map[string]interface{}{
		"authorization": authorization,
	}
	response, _ := json.MarshalIndent(params, "", "  ")
	return string(response)
}

func ResponseActivation(authorization map[string]interface{}) string {
	//	Helps you respond IssuingToken activation requests
	//
	//	When a new token activation is triggered by your user, a POST request will be made to your registered URL for you to confirm the activation code you informed to them.
	//	The POST request must be answered in the following format, within 2 seconds, and with an HTTP status code 200.
	//
	//	Parameters (required):
	//	- status [string]: Sub-issuer response to the activation. ex: "approved" or "denied"
	//
	//	Parameters (optional):
	//	- reason [string]: Denial reason. Options: "other", "bruteForce", "subIssuerError", "lostCard", "invalidCard", "invalidHolder", "expiredCard", "canceledCard", "blockedCard", "invalidExpiration", "invalidSecurityCode", "missingTokenAuthorizationUrl", "maxCardTriesExceeded", "maxWalletInstanceTriesExceeded"
	//	- tags [slice of strings]: Tags to filter retrieved struct. ex: []string{"tony", "stark"}
	//
	//	Return:
	//	- dumped JSON string that must be returned to us on the IssuingToken request
	params := map[string]map[string]interface{}{
		"authorization": authorization,
	}
	response, _ := json.MarshalIndent(params, "", "  ")
	return string(response)
}
