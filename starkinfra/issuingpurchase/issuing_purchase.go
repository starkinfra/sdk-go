package issuingpurchase

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingPurchase struct
//
//	Displays the IssuingPurchase structs created in your Workspace.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingPurchase is created. ex: "5656565656565656"
//	- HolderName [string]: Card holder name. ex: "Tony Stark"
//	- CardId [string]: Unique id returned when IssuingCard is created. ex: "5656565656565656"
//	- CardEnding [string]: Last 4 digits of the card number. ex: "1234"
//	- Purpose [string]: Purchase purpose. ex: "purchase"
//  - InstallmentCount [int]: quantity of installments to be confirmed. Minimum = 1. ex: 12
//	- Amount [int]: IssuingPurchase value in cents. Minimum = 0. ex: 1234 (= R$ 12.34)
//	- Tax [int]: Iof amount taxed for international purchases. ex: 1234 (= R$ 12.34)
//	- IssuerAmount [int]: Issuer amount. ex: 1234 (= R$ 12.34)
//	- IssuerCurrencyCode [string]: Issuer currency code. ex: "USD"
//	- IssuerCurrencySymbol [string]: Issuer currency symbol. ex: "$"
//	- MerchantAmount [int]: Merchant amount. ex: 1234 (= R$ 12.34)
//	- MerchantCurrencyCode [string]: Merchant currency code. ex: "USD"
//	- MerchantCurrencySymbol [string]: Merchant currency symbol. ex: "$"
//	- MerchantCategoryCode [string]: Merchant category code. ex: "fastFoodRestaurants"
//	- MerchantCountryCode [string]: Merchant country code. ex: "USA"
//	- AcquirerId [string]: Acquirer ID. ex: "5656565656565656"
//	- MerchantId [string]: Merchant ID. ex: "5656565656565656"
//	- MerchantName [string]: Merchant name. ex: "Google Cloud Platform"
//	- MerchantFee [int]: Fee charged by the merchant to cover specific costs, such as ATM withdrawal logistics, etc. ex: 200 (= R$ 2.00)
//	- WalletId [string]: Virtual wallet ID. ex: "5656565656565656"
//	- MethodCode [string]: Method code. ex: "chip", "token", "server", "manual", "magstripe" or "contactless"
//	- Score [float64]: Internal score calculated for the authenticity of the purchase. nil in case of insufficient data. ex: 7.6
//	- EndToEndId [string]: Unique id used to identify the transaction through all of its life cycle, even before the purchase is denied or accepted and gets its usual id. Example: endToEndId="679cd385-642b-49d0-96b7-89491e1249a5"
//	- Tags [string]: Slice of strings for tagging returned by the sub-issuer during the authorization. ex: []string{"travel", "food"}
//
//	Attributes (IssuingPurchase only):
//	- IssuingTransactionIds [slice of string]: Ledger transaction ids linked to this Purchase
//	- Status [string]: Current IssuingCard status. ex: "approved", "canceled", "denied", "confirmed", "voided"
//	- Updated [time.Time]: Latest update datetime for the IssuingPurchase. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: Creation datetime for the IssuingPurchase. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//
//	Attributes (Authorization Request only):
//	- IsPartialAllowed [bool]: True i the merchant allows partial purchases. ex: False
//	- CardTags [slice of strings]: Tags of the IssuingCard responsible for this purchase. ex: []string{"travel", "food"}
//	- HolderTags [slice of strings]: Tags of the IssuingHolder responsible for this purchase. ex: []string{"technology", "john snow"]

type IssuingPurchase struct {
	Id                     string     `json:",omitempty"`
	HolderName             string     `json:",omitempty"`
	CardId                 string     `json:",omitempty"`
	CardEnding             string     `json:",omitempty"`
	Purpose                string     `json:",omitempty"`
	InstallmentCount       int        `json:",omitempty"`
	Amount                 int        `json:",omitempty"`
	Tax                    int        `json:",omitempty"`
	IssuerAmount           int        `json:",omitempty"`
	IssuerCurrencyCode     string     `json:",omitempty"`
	IssuerCurrencySymbol   string     `json:",omitempty"`
	MerchantAmount         int        `json:",omitempty"`
	MerchantCurrencyCode   string     `json:",omitempty"`
	MerchantCurrencySymbol string     `json:",omitempty"`
	MerchantCategoryCode   string     `json:",omitempty"`
	MerchantCountryCode    string     `json:",omitempty"`
	AcquireId              string     `json:",omitempty"`
	MerchantId             string     `json:",omitempty"`
	MerchantName           string     `json:",omitempty"`
	MerchantFee            int        `json:",omitempty"`
	WalletId               string     `json:",omitempty"`
	MethodCode             string     `json:",omitempty"`
	Score                  float64    `json:",omitempty"`
	EndToEndId             string     `json:",omitempty"`
	Tags                   []string   `json:",omitempty"`
	IssuingTransactionIds  []string   `json:",omitempty"`
	Status                 string     `json:",omitempty"`
	Updated                *time.Time `json:",omitempty"`
	Created                *time.Time `json:",omitempty"`
	IsPartialAllowed       bool       `json:",omitempty"`
	CardTags               []string   `json:",omitempty"`
	HolderTags             []string   `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingPurchase"}

func Get(id string, user user.User) (IssuingPurchase, Error.StarkErrors) {
	//	Retrieve a specific IssuingPurchase by its id
	//
	//	Receive a single IssuingPurchase struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingPurchase struct that corresponds to the given id.
	var issuingPurchase IssuingPurchase
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingPurchase)
	if unmarshalError != nil {
		return issuingPurchase, err
	}
	return issuingPurchase, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingPurchase, chan Error.StarkErrors) {
	//	Retrieve IssuingPurchase structs
	//
	//	Receive a channel of IssuingPurchase structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- ids [slice of strings, default nil, default nil]: Purchase IDs
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- endToEndIds [slice of strings, default nil]: Central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
	//		- holderIds [slice of strings, default nil]: Card holder IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- cardIds [slice of strings, default nil]: Card  IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"approved", "canceled", "denied", "confirmed", "voided"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingPurchase structs with updated attributes
	var issuingPurchase IssuingPurchase
	purchases := make(chan IssuingPurchase)
	purchasesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingPurchase)
			if err != nil {
				purchasesError <- Error.UnknownError(err.Error())
				continue
			}
			purchases <- issuingPurchase
		}
		for err := range errorChannel {
			purchasesError <- err
		}
		close(purchases)
		close(purchasesError)
	}()
	return purchases, purchasesError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingPurchase, string, Error.StarkErrors) {
	//	Retrieve paged IssuingPurchase structs
	//
	//	Receive a slice of up to 100 IssuingPurchase structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- ids [slice of strings, default nil, default nil]: Purchase IDs
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- endToEndIds [slice of strings, default nil]: Central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
	//		- holderIds [slice of strings, default nil]: Card holder IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- cardIds [slice of strings, default nil]: Card  IDs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"approved", "canceled", "denied", "confirmed", "voided"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingPurchase structs with updated attributes
	//	- cursor to retrieve the next page of IssuingPurchase structs
	var issuingPurchases []IssuingPurchase
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingPurchases)
	if unmarshalError != nil {
		return issuingPurchases, cursor, err
	}
	return issuingPurchases, cursor, err
}

func Parse(content string, signature string, user user.User) (IssuingPurchase, Error.StarkErrors) {
	//	Create single verified IssuingPurchase authorization request from a content string
	//
	//	Use this method to parse and verify the authenticity of the authorization request received at the informed endpoint.
	//	Authorization requests are posted to your registered endpoint whenever IssuingPurchases are received.
	//	They present IssuingPurchase data that must be analyzed and answered with approval or declination.
	//	If the provided digital signature does not check out with the StarkInfra public key, a stark.exception.InvalidSignatureException will be raised.
	//	If the authorization request is not answered within 2 seconds or is not answered with an HTTP status code 200 the IssuingPurchase will go through the pre-configured stand-in validation.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- parsed IssuingPurchase struct
	var issuingPurchase IssuingPurchase
	parsed, err := utils.ParseAndVerify(content, signature, "", user)
	if err.Errors != nil {
		return issuingPurchase, err
	}

	unmarshalError := json.Unmarshal([]byte(parsed), &issuingPurchase)
	if unmarshalError != nil {
		return issuingPurchase, Error.UnknownError(unmarshalError.Error())
	}
	
	return issuingPurchase, Error.StarkErrors{}
}

func Response(authorization map[string]interface{}) string {
	//	Helps you respond IssuingPurchase requests
	//
	//	Parameters (required):
	//	- status [string]: Sub-issuer response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: Denial reason. Options: "other", "blocked", "lostCard", "stolenCard", "invalidPin", "invalidCard", "cardExpired", "issuerError", "concurrency", "standInDenial", "subIssuerError", "invalidPurpose", "invalidZipCode", "invalidWalletId", "inconsistentCard", "settlementFailed", "cardRuleMismatch", "invalidExpiration", "prepaidInstallment", "holderRuleMismatch", "insufficientBalance", "tooManyTransactions", "invalidSecurityCode", "invalidPaymentMethod", "confirmationDeadline", "withdrawalAmountLimit", "insufficientCardLimit", "insufficientHolderLimit"
	//
	//	Parameters (optional):
	//	- amount [int, default nil]: Amount in cents that was authorized. ex: 1234 (= R$ 12.34)
	//	- tags [slice of strings, default nil]: Tags to filter retrieved struct. ex: []string{"tony", "stark"}
	//
	//	Return:
	//	- dumped JSON string that must be returned to us on the IssuingPurchase request
	params := map[string]map[string]interface{}{
		"authorization": authorization,
	}
	response, _ := json.MarshalIndent(params, "", "  ")
	return string(response)
}
