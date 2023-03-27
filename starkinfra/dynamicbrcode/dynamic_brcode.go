package dynamicbrcode

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	DynamicBrcode struct
//
//	BR Codes store information represented by Pix QR Codes, which are used to
//  send or receive Pix transactions in a convenient way.
//
//  DynamicBrcodes represent charges with information that can change at any time,
//  since all data needed for the payment is requested dynamically to an URL stored
//  in the BR Code. Stark Infra will receive the GET request and forward it to your
//  registered endpoint with a GET request containing the UUID of the BR Code for
//  identification.
//
//  When you initialize a DynamicBrcode, the entity will not be automatically
//  created in the Stark Infra API. The 'create' function sends the structs
//  to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//  - Name [string]: Receiver's name. ex: "Tony Stark"
//  - City [string]: Receiver's city name. ex: "Rio de Janeiro"
//  - ExternalId [string]: String that must be unique among all your DynamicBrcodes. Duplicated external ids will cause failures. ex: "my-internal-id-123456"
//
//	Parameters (optional):
//  - Type [string, default "instant"]: Type of the DynamicBrcode. Options: "instant", "due"
//  - Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//  - Id [string]: Id returned on creation, this is the BR Code. ex: "00020126360014br.gov.bcb.pix0114+552840092118152040000530398654040.095802BR5915Jamie Lannister6009Sao Paulo620705038566304FC6C"
//  - Uuid [string]: Unique uuid returned when the DynamicBrcode is created. ex: "4e2eab725ddd495f9c98ffd97440702d"
//  - Url [string]: URL link to the BR Code image. ex: "https://brcode-h.sandbox.starkinfra.com/dynamic-qrcode/901e71f2447c43c886f58366a5432c4b.png"
//  - Updated [time.Time]: Latest update datetime for the DynamicBrcode. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Created [time.Time]: Creation datetime for the DynamicBrcode. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type DynamicBrcode struct {
	Name       string     `json:",omitempty"`
	City       string     `json:",omitempty"`
	ExternalId string     `json:",omitempty"`
	Type       string     `json:",omitempty"`
	Tags       []string   `json:",omitempty"`
	Id         string     `json:",omitempty"`
	Uuid       string     `json:",omitempty"`
	Url        string     `json:",omitempty"`
	Updated    *time.Time `json:",omitempty"`
	Created    *time.Time `json:",omitempty"`
}

var object DynamicBrcode
var objects []DynamicBrcode
var resource = map[string]string{"name": "DynamicBrcode"}

func Create(brcodes []DynamicBrcode, user user.User) ([]DynamicBrcode, Error.StarkErrors) {
	//	Create DynamicBrcodes
	//
	//	Send a slice of DynamicBrcode structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- brcodes [slice of DynamicBrcode structs]: Slice of DynamicBrcode structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of DynamicBrcode structs with updated attributes
	create, err := utils.Multi(resource, brcodes, nil, user)
	unmarshalError := json.Unmarshal(create, &brcodes)
	if unmarshalError != nil {
		return brcodes, err
	}
	return brcodes, err
}

func Get(uuid string, user user.User) (DynamicBrcode, Error.StarkErrors) {
	//	Retrieve a specific DynamicBrcode
	//
	//	Receive a single DynamicBrcode struct previously created in the Stark Infra API by its uuid
	//
	//	Parameters (required):
	//	- uuid [string]: Struct's unique uuid. ex: "901e71f2447c43c886f58366a5432c4b"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- dynamicBrcode struct that corresponds to the given id.
	get, err := utils.Get(resource, uuid, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan DynamicBrcode {
	//	Retrieve DynamicBrcode structs
	//
	//	Receive a channel of DynamicBrcode structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- externalIds [slice of strings, default nil]: Slice of externalIds to filter retrieved structs. ex: []string{"my_external_id1", "my_external_id2"}
	//  	- uuids [slice of strings, default nil]: Slice of uuids to filter retrieved structs. ex: []string{"901e71f2447c43c886f58366a5432c4b", "4e2eab725ddd495f9c98ffd97440702d"}
	//  	- tags [slice of strings, default nil]: Slice of tags to filter retrieved structs. ex: []string{"travel", "food"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of DynamicBrcode structs with updated attributes
	brcodes := make(chan DynamicBrcode)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			brcodes <- object
		}
		close(brcodes)
	}()
	return brcodes
}

func Page(params map[string]interface{}, user user.User) ([]DynamicBrcode, string, Error.StarkErrors) {
	//	Retrieve paged DynamicBrcode structs
	//
	//	Receive a slice of up to 100 DynamicBrcode structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//  	- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- externalIds [slice of strings, default nil]: Slice of externalIds to filter retrieved structs. ex: []string{"my_external_id1", "my_external_id2"}
	//  	- uuids [slice of strings, default nil]: Slice of uuids to filter retrieved structs. ex: []string{"901e71f2447c43c886f58366a5432c4b", "4e2eab725ddd495f9c98ffd97440702d"}
	//  	- tags [slice of strings, default nil]: Slice of tags to filter retrieved structs. ex: []string{"travel", "food"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of DynamicBrcode structs with updated attributes
	//  - cursor to retrieve the next page of DynamicBrcode structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func ResponseDue(params map[string]interface{}) interface{} {
	//	Helps you respond to a due DynamicBrcode Read
	//
	//	When a Due DynamicBrcode is read by your user, a GET request containing the Brcode's
	//	UUID will be made to your registered URL to retrieve additional information needed
	//	to complete the transaction.
	//
	//  - params [map[string]interface{}, default nil]: map of parameters for the response
	//		Parameters (required):
	//		- version [int]: Int that represents how many times the BR Code was updated.
	//		- created [string]: Creation datetime in ISO format of the DynamicBrcode.
	//		- due [string]: Requested payment due datetime in ISO format.
	//		- expiration [int]: Time in seconds counted from the creation datetime until the DynamicBrcode expires. After expiration, the BR Code cannot be paid anymore.
	//		- keyId [string]: Receiver's PixKey id. Can be a tax_id (CPF/CNPJ), a phone number, an email or an alphanumeric sequence (EVP). ex: "+5511989898989"
	//		- status [string]: BR Code status. Options: "created", "overdue", "paid", "canceled" or "expired"
	//		- reconciliationId [string]: Id to be used for conciliation of the resulting Pix transaction. This id must have from to 26 to 35 alphanumeric characters ex: "cd65c78aeb6543eaaa0170f68bd741ee"
	//		- nominalAmount [int]: Positive int that represents the amount in cents of the resulting Pix transaction. ex: 1234 (= R$ 12.34)
	//  	- senderName [string]: Sender's full name. ex: "Anthony Edward Stark"
	//  	- receiverName [string]: Receiver's full name. ex: "Jamie Lannister"
	//  	- receiverStreetLine [string]: Receiver's main address. ex: "Av. Paulista, 200"
	//  	- receiverCity [string]: Receiver's address city name. ex: "Sao Paulo"
	//  	- receiverStateCode [string]: Receiver's address state code. ex: "SP"
	//  	- receiverZipCode [string]: Receiver's address zip code. ex: "01234-567"
	//
	//		Parameters (optional):
	//		- senderTaxId [string, default nil]: Sender's CPF (11 digits formatted or unformatted) or CNPJ (14 digits formatted or unformatted). ex: "01.001.001/0001-01"
	//		- receiverTaxId [string, default nil]: Receiver's CPF (11 digits formatted or unformatted) or CNPJ (14 digits formatted or unformatted). ex: "012.345.678-90"
	//		- fine [float64, default 2.0]: Percentage charged if the sender pays after the due datetime.
	//		- interest [float64, default 1.0]: Interest percentage charged if the sender pays after the due datetime.
	//		- discounts [slice of maps, default nil]: Slice of maps with "percentage":float64 and "due":date.datetime or string pairs.
	//		- description [string, default nil]: Additional information to be shown to the sender at the moment of payment.
	//
	//	Return:
	//	- dumped JSON string that must be returned to us
	due, _ := json.MarshalIndent(params, "", "  ")
	return string(due)
}

func ResponseInstant(params map[string]interface{}) string {
	//	Helps you respond to an instant DynamicBrcode Read
	//
	//	When an instant DynamicBrcode is read by your user, a GET request containing the BR Code's UUID will be made
	//  to your registered URL to retrieve additional information needed to complete the transaction.
	//  The get request must be answered in the following format within 5 seconds and with an HTTP status code 200.
	//
	//  - params [map[string]interface{}, default nil]: map of parameters for the response
	//		Parameters (required):
	//  	- created [string]: Creation datetime of the DynamicBrcode.
	//		- version [int]: Int that represents how many times the BR Code was updated.
	//  	- keyId [string]: Receiver's PixKey id. Can be a tax_id (CPF/CNPJ), a phone number, an email or an alphanumeric sequence (EVP). ex: "+5511989898989"
	//  	- status [string]: BR Code's status. Options: "created", "overdue", "paid", "canceled" or "expired"
	//  	- reconciliationId [string]: Id to be used for conciliation of the resulting Pix transaction. ex: "cd65c78aeb6543eaaa0170f68bd741ee"
	//  	- amount [int]: Positive int that represents the amount in cents of the resulting Pix transaction. ex: 1234 (= R$ 12.34)
	//
	//		Parameters (conditionally required):
	//		- cashierType [string, default nil]: Cashier's type. Required if the cashAmount is different from 0. Options: "merchant", "participant" and "other"
	//  	- cashierBankCode [string, default nil]: Cashier's bank code. Required if the cashAmount is different from 0. ex: "20018183"
	//
	//		Parameters (optional):
	//		- cashAmount [int, default nil]: Amount to be withdrawn from the cashier in cents. ex: 1000 (= R$ 10.00)
	//  	- expiration [int, default 86400 (1 day)]: Time in seconds counted from the creation datetime until the DynamicBrcode expires. After expiration, the BR Code cannot be paid anymore. Default value: 86400 (1 day)
	//  	- senderName [string, default nil]: Sender's full name. ex: "Anthony Edward Stark"
	//  	- senderTaxId [string, default nil]: Sender's CPF (11 digits formatted or unformatted) or CNPJ (14 digits formatted or unformatted). ex: "01.001.001/0001-01"
	//  	- amountType [string, default "fixed"]: Amount type of the Brcode. If the amount type is "custom" the Brcode's amount can be changed by the sender at the moment of payment. Options: "fixed" or "custom"
	//  	- description [string, default nil]: Additional information to be shown to the sender at the moment of payment.
	//
	//	Return:
	//	- dumped JSON string that must be returned to us
	instant, _ := json.MarshalIndent(params, "", "  ")
	return string(instant)
}

func Verify(uuid string, signature string, user user.User) string {
	//	Verify a DynamicBrcode Read
	//
	//	When a DynamicBrcode is read by your user, a GET request will be made to your registered URL to
	//  retrieve additional information needed to complete the transaction.
	//  Use this method to verify the authenticity of a GET request received at your registered endpoint.
	//  If the provided digital signature does not check out with the StarkInfra public key,
	//  a error.InvalidSignatureException will be raised.
	//
	//	Parameters (required):
	//	- uuid [string]: Unique uuid returned when a DynamicBrcode is created. ex: "4e2eab725ddd495f9c98ffd97440702d"
	//  - signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- verified Brcode's uuid.
	return utils.Verify(uuid, signature, user)
}
