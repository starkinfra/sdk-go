package issuinginvoice

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingInvoice struct
//
//	The IssuingInvoice structs created in your Workspace load your Issuing balance when paid.
//
//	Parameters (required):
//	- Amount [int]: IssuingInvoice value in cents. ex: 1234 (= R$ 12.34)
//
//	Parameters (optional):
//	- TaxId [string, default sub-issuer tax ID]: payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- Name [string, default sub-issuer name]: payer name. ex: "Iron Bank S.A."
//	- Tags [slice of strings, default nil]: slice of strings for tagging. ex: []string{"travel", "food"}
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when IssuingInvoice is created. ex: "5656565656565656"
//  - Brcode [string]: BR Code for the Invoice payment. ex: "00020101021226930014br.gov.bcb.pix2571brcode-h.development.starkinfra.com/v2/d7f6546e194d4c64a153e8f79f1c41ac5204000053039865802BR5925Stark Bank S.A. - Institu6009Sao Paulo62070503***63042109"
//  - Due [time.Time]: Invoice due and expiration date in UTC ISO format. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC),
//  - Link [string]: public Invoice webpage URL. ex: "https://starkbank-card-issuer.development.starkbank.com/invoicelink/d7f6546e194d4c64a153e8f79f1c41ac"
//	- Status [string]: current IssuingInvoice status. ex: "created", "expired", "overdue", "paid"
//	- IssuingTransactionId [string]: ledger transaction ids linked to this IssuingInvoice. ex: "issuing-invoice/5656565656565656"
//	- Updated [time.Time]: latest update datetime for the IssuingInvoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Created [time.Time]: creation datetime for the IssuingInvoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingInvoice struct {
	Id                   string     `json:",omitempty"`
	Amount               int        `json:",omitempty"`
	TaxId                string     `json:",omitempty"`
	Name                 string     `json:",omitempty"`
	Tags                 []string   `json:",omitempty"`
	Brcode               string     `json:",omitempty"`
	Due                  *time.Time `json:",omitempty"`
	Link                 string     `json:",omitempty"`
	Status               string     `json:",omitempty"`
	IssuingTransactionId string     `json:",omitempty"`
	Updated              *time.Time `json:",omitempty"`
	Created              *time.Time `json:",omitempty"`
}

var object IssuingInvoice
var objects []IssuingInvoice
var resource = map[string]string{"name": "IssuingInvoice"}

func Create(invoice IssuingInvoice, user user.User) (IssuingInvoice, Error.StarkErrors) {
	//	Create an IssuingInvoice
	//
	//	Send an IssuingInvoice struct for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- invoice [IssuingInvoice struct]: IssuingInvoice struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingInvoice struct with updated attributes
	create, err := utils.Single(resource, invoice, user)
	unmarshalError := json.Unmarshal(create, &invoice)
	if unmarshalError != nil {
		return invoice, err
	}
	return invoice, err
}

func Get(id string, user user.User) (IssuingInvoice, Error.StarkErrors) {
	//	Retrieve a specific IssuingInvoice by its id
	//
	//	Receive a single IssuingInvoice struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingInvoice struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan IssuingInvoice {
	//	Retrieve IssuingInvoice
	//
	//	Receive a channel of IssuingInvoices structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "expired", "overdue", "paid"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingInvoices structs with updated attributes
	invoices := make(chan IssuingInvoice)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			invoices <- object
		}
		close(invoices)
	}()
	return invoices
}

func Page(params map[string]interface{}, user user.User) ([]IssuingInvoice, string, Error.StarkErrors) {
	//	Retrieve IssuingInvoices
	//
	//	Receive a slice of up to 100 IssuingInvoice structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "expired", "overdue", "paid"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingInvoices structs with updated attributes
	//	- cursor to retrieve the next page of IssuingInvoices structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}
