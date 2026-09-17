package creditnote

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	Invoice "github.com/starkinfra/sdk-go/starkinfra/creditnote/invoice"
	Signer "github.com/starkinfra/sdk-go/starkinfra/creditsigner"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	CreditNote struct
//
//	CreditNotes are used to generate CCB contracts between you and your customers.
//
//	When you initialize a CreditNote, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the structs
//	to the Stark Infra API and returns the slice of created structs.
//
//	Parameters (required):
//	- TemplateId [string]: ID of the contract template on which the credit note will be based. ex: "0123456789101112"
//	- Name [string]: Credit receiver's full name. ex: "Edward Stark"
//	- TaxId [string]: Credit receiver's tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//	- Scheduled [time.Time]: Date of transfer execution. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- Invoices [slice of Invoice structs]: Slice of up to 100 Invoice structs representing the installments to be paid by the borrower. All invoices in the same CreditNote must share the same Fine and Interest. ex: []invoice.Invoice{{Amount: 1000}, {Amount: 1000}}
//	- Payment [creditNote.Transfer struct]: Payment entity to be created and sent to the credit receiver. ex: creditnote.Transfer()
//	- Signers [slice of CreditSigner structs]: slice of up to 10 CreditSigner structs, one per person or entity that must sign the contract. Method options: "link", "token", "server", "organization". Signers registered in your credit profile and the SCD signature are appended automatically. ex: []creditsigner.CreditSigner{creditnote.Signer(), creditnote.Signer()}
//	- ExternalId [string]: A string that must be unique among all your CreditNotes, used to avoid resource duplication. ex: "my-internal-id-123456"
//	- StreetLine1 [string]: Credit receiver main address. ex: "Av. Paulista, 200"
//	- StreetLine2 [string]: Credit receiver address complement. ex: "Apto. 123"
//	- District [string]: Credit receiver address district/neighbourhood. ex: "Bela Vista"
//	- City [string]: Credit receiver address city. ex: "Rio de Janeiro"
//	- StateCode [string]: Credit receiver address state. ex: "GO"
//	- ZipCode [string]: Credit receiver address zip code. ex: "01311-200"
//
//	Parameters (conditionally required):
//	- PaymentType [string]: Payment type, inferred from the payment parameter if it is not a map. ex: "transfer"
//	- NominalAmount [int]: CreditNote value in cents. The nominalAmount parameter is required when amount is not sent. ex: 1234 (= R$ 12.34)
//	- Amount [int]: Amount in cents transferred to the credit receiver, before deductions. The amount parameter is required when nominalAmount is not sent. ex: 1234 (= R$ 12.34)
//
//	Parameters (optional):
//	- RebateAmount [int, default 0]: Credit analysis fee deducted from lent amount. ex: 1234 (= R$ 112.34)
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for CreditNotes. ex: []string{"employees", "monthly"}
//	- Expiration [int, default default 604800 (7 days)]: Time interval in seconds between due date and expiration date. ex: 123456789
//	- Rules [slice of CreditNote.Rule structs, default nil]: Slice of CreditNote.Rule structs that modify the CreditNote behavior. ex: []creditnote.Rule{{Key: "invoiceCreationMode", Value: "scheduled"}}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the CreditNote is created. ex: "5656565656565656"
//	- DocumentId [string]: ID of the signed document to execute this CreditNote. ex: "4545454545454545"
//	- Status [string]: Current status of the CreditNote. ex: "canceled", "created", "expired", "failed", "processing", "signed", "success"
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this CreditNote. ex: []string{"19827356981273"}
//	- WorkspaceId [string]: ID of the Workspace that generated this CreditNote. ex: "4545454545454545"
//	- DebtorWorkspaceId [string]: ID of the debtor's Workspace, when it differs from the Workspace that generated this CreditNote. ex: "4545454545454545"
//	- TaxAmount [int]: Tax amount included in the CreditNote. ex: 100
//	- NominalInterest [float64]: Yearly nominal interest rate of the CreditNote, in percentage. ex: 11.5
//	- Interest [float64]: Yearly effective interest rate of the credit note, in percentage. ex: 12.5
//	- Created [time.Time]: Creation datetime for the CreditNote. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the CreditNote. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type CreditNote struct {
	Id                string                `json:",omitempty"`
	TemplateId        string                `json:",omitempty"`
	Name              string                `json:",omitempty"`
	TaxId             string                `json:",omitempty"`
	Scheduled         *time.Time            `json:",omitempty"`
	Payment           Transfer              `json:",omitempty"`
	Invoices          []Invoice.Invoice     `json:",omitempty"`
	Signers           []Signer.CreditSigner `json:",omitempty"`
	ExternalId        string                `json:",omitempty"`
	StreetLine1       string                `json:",omitempty"`
	StreetLine2       string                `json:",omitempty"`
	District          string                `json:",omitempty"`
	City              string                `json:",omitempty"`
	StateCode         string                `json:",omitempty"`
	ZipCode           string                `json:",omitempty"`
	NominalAmount     int                   `json:",omitempty"`
	PaymentType       string                `json:",omitempty"`
	Amount            int                   `json:",omitempty"`
	RebateAmount      int                   `json:",omitempty"`
	Tags              []string              `json:",omitempty"`
	Expiration        int                   `json:",omitempty"`
	Rules             []Rule                `json:",omitempty"`
	DocumentId        string                `json:",omitempty"`
	Status            string                `json:",omitempty"`
	TransactionIds    []string              `json:",omitempty"`
	WorkspaceId       string                `json:",omitempty"`
	DebtorWorkspaceId string                `json:",omitempty"`
	TaxAmount         int                   `json:",omitempty"`
	NominalInterest   float64               `json:",omitempty"`
	Interest          float64               `json:",omitempty"`
	Created           *time.Time            `json:",omitempty"`
	Updated           *time.Time            `json:",omitempty"`
}

var resource = map[string]string{"name": "CreditNote"}

func Create(notes []CreditNote, user user.User) ([]CreditNote, Error.StarkErrors) {
	//	Create CreditNotes
	//
	//	Send a slice of CreditNote structs for creation at the Stark Infra API. You can create up to 100
	//	CreditNotes in a single request. Provide either NominalAmount (pre-tax) or Amount (net disbursed value) —
	//	the other one, along with TaxAmount (IOF) and the interest rates, is computed automatically from the invoice schedule.
	//
	//	Parameters (required):
	//	- notes [slice of CreditNote structs]: Slice of CreditNote structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of CreditNote structs with updated attributes
	var creditNote []CreditNote
	create, err := utils.Multi(resource, notes, nil, user)
	unmarshalError := json.Unmarshal(create, &creditNote)
	if unmarshalError != nil {
		return creditNote, err
	}
	return creditNote, err
}

func Get(id string, user user.User) (CreditNote, Error.StarkErrors) {
	//	Retrieve a specific CreditNote
	//
	//	Receive a single CreditNote struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- CreditNote struct that corresponds to the given id.
	var creditNote CreditNote
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &creditNote)
	if unmarshalError != nil {
		return creditNote, err
	}
	return creditNote, err
}

func Query(params map[string]interface{}, user user.User) (chan CreditNote, chan Error.StarkErrors) {
	//	Retrieve CreditNote structs
	//
	//	Receive a channel of CreditNote structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. Options: ["canceled", "created", "expired", "failed", "processing", "signed", "success"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- channel of CreditNote structs with updated attributes
	notes := make(chan CreditNote)
	notesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var creditNote CreditNote
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &creditNote)
			if err != nil {
				notesError <- Error.UnknownError(err.Error())
				continue
			}
			notes <- creditNote
		}
		for err := range errorChannel {
			notesError <- err
		}
		close(notes)
		close(notesError)
	}()
	return notes, notesError
}

func Page(params map[string]interface{}, user user.User) ([]CreditNote, string, Error.StarkErrors) {
	//	Retrieve paged CreditNote structs
	//
	//	Receive a slice of up to 100 CreditNote structs previously created in the Stark Infra API and the cursor to the next page
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"canceled", "created", "expired", "failed", "processing", "signed", "success"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- slice of CreditNote structs with updated attributes
	//	- cursor to retrieve the next page of CreditNote structs
	var creditNotes []CreditNote
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &creditNotes)
	if unmarshalError != nil {
		return creditNotes, cursor, err
	}
	return creditNotes, cursor, err
}

func Cancel(id string, user user.User) (CreditNote, Error.StarkErrors) {
	//	Cancel a CreditNote entity
	//
	//	Cancel a CreditNote entity previously created in the Stark Infra API. Only CreditNotes with status
	//	"created", "signed" or "processing" can be canceled, which also cancels the signing document. A CreditNote
	//	in "success", "failed", "expired" or already "canceled" status is returned unchanged.
	//
	//	Parameters (required):
	//	- id [string]: CreditNote unique id. ex: "6306109539221504"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- canceled CreditNote struct
	var creditNote CreditNote
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &creditNote)
	if unmarshalError != nil {
		return creditNote, err
	}
	return creditNote, err
}

func Pdf(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific CreditNote pdf file
	//
	//	Receive the CCB disbursement pdf file of a CreditNote previously created in the Stark Infra API by its id.
	//	Not available for canceled CreditNotes.
	//
	//	Parameters (required):
	//	- id [string]: CreditNote unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- CreditNote pdf file
	return utils.GetContent(resource, id, nil, user, "pdf")
}

func Payment(id string, user user.User) ([]byte, Error.StarkErrors) {
	//	Retrieve a specific CreditNote payment pdf file
	//
	//	Receive the CCB disbursement payment pdf file of a CreditNote previously created in the Stark Infra API by its id.
	//	Only CreditNotes with status "success" have a payment pdf.
	//
	//	Parameters (required):
	//	- id [string]: CreditNote unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- CreditNote payment pdf file
	return utils.GetContent(resource, id, nil, user, "payment/pdf")
}
