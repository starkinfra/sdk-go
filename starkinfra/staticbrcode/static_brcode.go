package staticbrcode

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	StaticBrcode struct
//
//	A StaticBrcode stores account information in the form of a PixKey and can be used to create
//  Pix transactions easily.
//  When you initialize a StaticBrcode, the entity will not be automatically
//  created in the Stark Infra API. The 'create' function sends the structs
//
//	Parameters (required):
//	- Name [string]: Receiver's name. ex: "Tony Stark"
//  - KeyId [string]: Receiver's Pix key id. ex: "+5541999999999"
//  - City [string, default São Paulo]: Receiver's city name. ex: "Rio de Janeiro"
//
//	Parameters (optional):
//	- Amount [int, default 0]: Positive int that represents the amount in cents of the resulting Pix transaction. If the amount is zero, the sender can choose any amount in the moment of payment. ex: 1234 (= R$ 12.34)
//	- ReconciliationId [string, default nil]: Id to be used for conciliation of the resulting Pix transaction. This id must have up to 25 alphanumeric digits ex: "ah27s53agj6493hjds6836v49"
//	- CashierBankCode [string, default nil]: Cashier's bank code. ex: "20018183".
//	- Description [string, default nil]: Optional description to override default description to be shown in the bank statement. ex: "Payment for service #1234"
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: Id returned on creation, this is the BR code. ex: "00020126360014br.gov.bcb.pix0114+552840092118152040000530398654040.095802BR5915Jamie Lannister6009Sao Paulo620705038566304FC6C"
//  - Uuid [string]: Unique uuid returned when a StaticBrcode is created. ex: "97756273400d42ce9086404fe10ea0d6"
//  - Url [string]: Url link to the BR Code image. ex: "https://brcode-h.sandbox.starkinfra.com/static-qrcode/97756273400d42ce9086404fe10ea0d6.png"
//  - Updated [time.Time]: Latest update datetime for the StaticBrcode. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Created [time.Time]: Creation datetime for the StaticBrcode. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type StaticBrcode struct {
	Name             string     `json:",omitempty"`
	KeyId            string     `json:",omitempty"`
	City             string     `json:",omitempty"`
	Amount           int        `json:",omitempty"`
	ReconciliationId string     `json:",omitempty"`
	CashierBankCode  string     `json:",omitempty"`
	Description      string     `json:",omitempty"`
	Tags             []string   `json:",omitempty"`
	Id               string     `json:",omitempty"`
	Uuid             string     `json:",omitempty"`
	Url              string     `json:",omitempty"`
	Updated          *time.Time `json:",omitempty"`
	Created          *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "StaticBrcode"}

func Create(brcodes []StaticBrcode, user user.User) ([]StaticBrcode, Error.StarkErrors) {
	//	Create StaticBrcodes
	//
	//	Send a slice of StaticBrcode structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- brcodes [slice of StaticBrcode structs]: Slice of StaticBrcode structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of StaticBrcode structs with updated attributes
	create, err := utils.Multi(resource, brcodes, nil, user)
	unmarshalError := json.Unmarshal(create, &brcodes)
	if unmarshalError != nil {
		return brcodes, err
	}
	return brcodes, err
}

func Get(uuid string, user user.User) (StaticBrcode, Error.StarkErrors) {
	//	Retrieve a specific StaticBrcode by its uuid
	//
	//	Receive a single StaticBrcode struct previously created in the Stark Infra API by its uuid
	//
	//	Parameters (required):
	//	- uuid [string]: Struct's unique uuid. ex: "97756273400d42ce9086404fe10ea0d6"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- staticBrcode struct that corresponds to the given uuid.
	var staticBrcode StaticBrcode
	get, err := utils.Get(resource, uuid, nil, user)
	unmarshalError := json.Unmarshal(get, &staticBrcode)
	if unmarshalError != nil {
		return staticBrcode, err
	}
	return staticBrcode, err
}

func Query(params map[string]interface{}, user user.User) chan StaticBrcode {
	//	Retrieve StaticBrcode structs
	//
	//	Receive a channel of StaticBrcode structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- uuids [slice of strings, default nil]: Slice of uuids to filter retrieved structs. ex: []string{"97756273400d42ce9086404fe10ea0d6", "e3da0b6d56fa4045b9b295b2be82436e"}
	//  	- tags [slice of strings, default nil]: Slice of tags to filter retrieved structs. ex: []string{"travel", "food"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of StaticBrcode structs with updated attributes
	var staticBrcode StaticBrcode
	brcodes := make(chan StaticBrcode)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &staticBrcode)
			if err != nil {
				print(err)
			}
			brcodes <- staticBrcode
		}
		close(brcodes)
	}()
	return brcodes
}

func Page(params map[string]interface{}, user user.User) ([]StaticBrcode, string, Error.StarkErrors) {
	//	Retrieve paged PixKey structs
	//
	//	Receive a slice of up to 100 PixKey structs previously created in the Stark Infra API and the cursor to the next page.
	//  Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//  	- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- uuids [slice of strings, default nil]: Slice of uuids to filter retrieved structs. ex: []string{"97756273400d42ce9086404fe10ea0d6", "e3da0b6d56fa4045b9b295b2be82436e"}
	//  	- tags [slice of strings, default nil]: Slice of tags to filter retrieved structs. ex: []string{"travel", "food"}
	//  - user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of StaticBrcode structs with updated attributes
	//  - cursor to retrieve the next page of StaticBrcode structs
	var staticBrcodes []StaticBrcode
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &staticBrcodes)
	if unmarshalError != nil {
		return staticBrcodes, cursor, err
	}
	return staticBrcodes, cursor, err
}
