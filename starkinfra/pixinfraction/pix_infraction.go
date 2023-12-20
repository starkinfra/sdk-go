package pixinfraction

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixInfraction struct
//
//	PixInfractions are used to report transactions that are suspected of
//	fraud, to request a refund or to reverse a refund.
//	When you initialize a PixInfraction, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- ReferenceId [string]: EndToEndId or return_id of the transaction being reported. ex: "E20018183202201201450u34sDGd19lz"
//	- Type [string]: Type of infraction report. Options: "reversal", "reversalChargeback"
//	- Method [string]: Method of Pix Infraction. Options: "scam", "unauthorized", "coercion", "invasion", "other", "unknown"
//
//	Parameters (optional):
//	- Description [string, default nil]: Description for any details that can help with the infraction investigation.
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
//	- FraudType [string, default nil]: Type of Pix Fraud. Options: "identity", "mule", "scam", "unknown", "other"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixInfraction is created. ex: "5656565656565656"
//	- FraudId [string]: Id of the Pix Fraud. ex: "5741774970552320"
//	- CreditedBankCode [string]: BankCode of the credited Pix participant in the reported transaction. ex: "20018183"
//	- DebitedBankCode [string]: BankCode of the debited Pix participant in the reported transaction. ex: "20018183"
//	- Flow [string]: Direction of the PixInfraction flow. Options: "out" if you created the PixInfraction, "in" if you received the PixInfraction.
//  - Analysis [string]: Analysis that led to the result.
//  - ReportedBy [string]: Agent that reported the PixInfraction. Options: "debited", "credited"
//  - Result [string]: Result after the analysis of the PixInfraction by the receiving party. Options: "agreed", "disagreed"
//  - Status [string]: Current PixInfraction status. Options: "created", "failed", "delivered", "closed", "canceled"
//  - Created [time.Time]: Creation datetime for the PixInfraction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//  - Updated [time.Time]: Latest update datetime for the PixInfraction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixInfraction struct {
	ReferenceId      string     `json:",omitempty"`
	Type             string     `json:",omitempty"`
	Method           string     `json:",omitempty"`
	Description      string     `json:",omitempty"`
	Tags             []string   `json:",omitempty"`
	FraudType        string     `json:",omitempty"`
	Id               string     `json:",omitempty"`
	FraudId          string     `json:",omitempty"`
	CreditedBankCode string     `json:",omitempty"`
	DebitedBankCode  string     `json:",omitempty"`
	Flow             string     `json:",omitempty"`
	Analysis         string     `json:",omitempty"`
	ReportedBy       string     `json:",omitempty"`
	Result           string     `json:",omitempty"`
	Status           string     `json:",omitempty"`
	Created          *time.Time `json:",omitempty"`
	Updated          *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixInfraction"}

func Create(infractions []PixInfraction, user user.User) ([]PixInfraction, Error.StarkErrors) {
	//	Create PixInfraction structs
	//
	//	Create PixInfractions in the Stark Infra API
	//
	//	Parameters (required):
	//	- infractions [slice of PixInfraction structs]: Slice of PixInfraction structs to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixInfraction structs with updated attributes
	create, err := utils.Multi(resource, infractions, nil, user)
	unmarshalError := json.Unmarshal(create, &infractions)
	if unmarshalError != nil {
		return infractions, err
	}
	return infractions, err
}

func Get(id string, user user.User) (PixInfraction, Error.StarkErrors) {
	//	Retrieve a PixInfraction struct
	//
	//	Retrieve the PixInfraction struct linked to your Workspace in the Stark Infra API using its id.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656".
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixInfraction struct that corresponds to the given id.
	var pixInfraction PixInfraction
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixInfraction)
	if unmarshalError != nil {
		return pixInfraction, err
	}
	return pixInfraction, err
}

func Query(params map[string]interface{}, user user.User) chan PixInfraction {
	//	Retrieve PixInfraction structs
	//
	//	Receive a channel of PixInfraction structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "delivered", "closed", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- type [slice of strings, default nil]: Filter for the type of retrieved PixInfractions. Options: "reversal", "reversalChargeback"
	//  	- flow [string, default nil]: Direction of the PixInfraction flow. Options: "out" if you created the PixInfraction, "in" if you received the PixInfraction.
	//  	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixInfraction structs with updated attributes
	var pixInfraction PixInfraction
	infractions := make(chan PixInfraction)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixInfraction)
			if err != nil {
				print(err)
			}
			infractions <- pixInfraction
		}
		close(infractions)
	}()
	return infractions
}

func Page(params map[string]interface{}, user user.User) ([]PixInfraction, string, Error.StarkErrors) {
	//	Retrieve paged PixInfraction structs.
	//
	//	Receive a slice of up to 100 PixInfraction structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call.
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "delivered", "closed", "canceled"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- type [slice of strings, default nil]: Filter for the type of retrieved PixInfractions. Options: "reversal", "reversalChargeback"
	//  	- flow [string, default nil]: Direction of the PixInfraction flow. Options: "out" if you created the PixInfraction, "in" if you received the PixInfraction.
	//  	- tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"travel", "food"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixInfraction structs with updated attributes
	//  - Cursor to retrieve the next page of PixInfraction structs
	var pixInfractions []PixInfraction
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixInfractions)
	if unmarshalError != nil {
		return pixInfractions, cursor, err
	}
	return pixInfractions, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (PixInfraction, Error.StarkErrors) {
	//	Update PixInfraction entity
	//
	//	Update a PixInfraction by passing id.
	//
	//	Parameters (required):
	//	- id [string]: PixInfraction id. ex: '5656565656565656'
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"amount": 9090}
	//  	Parameters (required):
	//		- result [string]: Result after the analysis of the PixInfraction. Options: "agreed", "disagreed"
	//		Parameters (conditionally required):
	//		- fraudType [string, default nil]: Type of Pix Fraud. Options: "identity", "mule", "scam", "unknown", "other"
	//		Parameters (optional):
	//		- analysis [string, default nil]: Analysis that led to the result.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixInfraction with updated attributes
	var pixInfraction PixInfraction
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &pixInfraction)
	if unmarshalError != nil {
		return pixInfraction, err
	}
	return pixInfraction, err
}

func Cancel(id string, user user.User) (PixInfraction, Error.StarkErrors) {
	//	Cancel a PixInfraction entity
	//
	//	Cancel a PixInfraction entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- canceled PixInfraction struct
	var pixInfraction PixInfraction
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &pixInfraction)
	if unmarshalError != nil {
		return pixInfraction, err
	}
	return pixInfraction, err
}
