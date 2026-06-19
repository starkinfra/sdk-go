package issuingbillinginvoice

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingBillingInvoice struct
//
//	The IssuingBillingInvoice structs represent the invoices charged for your issuing usage.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the IssuingBillingInvoice is created. ex: "5656565656565656"
//	- TaxId [string]: payer tax ID (CPF or CNPJ). ex: "012.345.678-90"
//	- Name [string]: payer name. ex: "Tony Stark"
//	- Fine [float64]: Fine percentage applied when paid after the due date. ex: 2.0
//	- Interest [float64]: Monthly interest percentage applied when paid after the due date. ex: 1.0
//	- Status [string]: current IssuingBillingInvoice status. ex: "created", "paid"
//	- Amount [int]: invoice amount in cents. ex: 1234 (= R$ 12.34)
//	- NominalAmount [int]: nominal amount in cents. ex: 1234 (= R$ 12.34)
//	- Brcode [string]: BR Code for the invoice payment. ex: "00020101021226930014br.gov.bcb.pix..."
//	- Link [string]: public invoice webpage URL. ex: "https://starkinfra.com/invoicelink/d7f6546e194d4c64a153e8f79f1c41ac"
//	- Due [time.Time]: invoice due datetime in UTC ISO format. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC)
//	- Start [time.Time]: billing cycle start datetime. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- End [time.Time]: billing cycle end datetime. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Created [time.Time]: creation datetime for the IssuingBillingInvoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: latest update datetime for the IssuingBillingInvoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type IssuingBillingInvoice struct {
	Id            string     `json:",omitempty"`
	TaxId         string     `json:",omitempty"`
	Name          string     `json:",omitempty"`
	Fine          float64    `json:",omitempty"`
	Interest      float64    `json:",omitempty"`
	Status        string     `json:",omitempty"`
	Amount        int        `json:",omitempty"`
	NominalAmount int        `json:",omitempty"`
	Brcode        string     `json:",omitempty"`
	Link          string     `json:",omitempty"`
	Due           *time.Time `json:",omitempty"`
	Start         *time.Time `json:",omitempty"`
	End           *time.Time `json:",omitempty"`
	Created       *time.Time `json:",omitempty"`
	Updated       *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingBillingInvoice"}

func Get(id string, user user.User) (IssuingBillingInvoice, Error.StarkErrors) {
	//	Retrieve a specific IssuingBillingInvoice by its id
	//
	//	Receive a single IssuingBillingInvoice struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- issuingBillingInvoice struct that corresponds to the given id.
	var issuingBillingInvoice IssuingBillingInvoice
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &issuingBillingInvoice)
	if unmarshalError != nil {
		return issuingBillingInvoice, err
	}
	return issuingBillingInvoice, err
}

func Query(params map[string]interface{}, user user.User) (chan IssuingBillingInvoice, chan Error.StarkErrors) {
	//	Retrieve IssuingBillingInvoices
	//
	//	Receive a channel of IssuingBillingInvoice structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "paid"}
	//		- ids [slice of strings, default nil]: list of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingBillingInvoice structs with updated attributes
	var issuingBillingInvoice IssuingBillingInvoice
	invoices := make(chan IssuingBillingInvoice)
	invoicesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingBillingInvoice)
			if err != nil {
				invoicesError <- Error.UnknownError(err.Error())
				continue
			}
			invoices <- issuingBillingInvoice
		}
		for err := range errorChannel {
			invoicesError <- err
		}
		close(invoices)
		close(invoicesError)
	}()
	return invoices, invoicesError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingBillingInvoice, string, Error.StarkErrors) {
	//	Retrieve IssuingBillingInvoices
	//
	//	Receive a slice of up to 100 IssuingBillingInvoice structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved structs. ex: []string{"created", "paid"}
	//		- ids [slice of strings, default nil]: list of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingBillingInvoice structs with updated attributes
	//	- cursor to retrieve the next page of IssuingBillingInvoice structs
	var issuingBillingInvoices []IssuingBillingInvoice
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingBillingInvoices)
	if unmarshalError != nil {
		return issuingBillingInvoices, cursor, err
	}
	return issuingBillingInvoices, cursor, err
}
