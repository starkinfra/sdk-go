package creditpreview

import (
	Invoice "github.com/starkinfra/sdk-go/starkinfra/creditnote/invoice"
	"time"
)

//	CreditNotePreview struct
//
//	A CreditNotePreview is used to preview a CCB contract between the borrower and lender with a specific table type.
//
//	When you initialize a CreditNotePreview, the entity will be automatically sent to the Stark Infra API.
//	The 'create' function sends the structs to the Stark Infra API and returns the slice of preview data.
//
//	Parameters (required):
//	- Type [string]: Table type that defines the amortization system. Options: "sac", "price", "american", "bullet", "custom"
//	- Scheduled [time.Time]: Date of transfer execution. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- TaxId [string]: Credit receiver's tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//
//	Parameters (conditionally required):
//	Each `Type` value requires its own subset of these fields: "sac", "price", "american" and "bullet" require NominalInterest and InitialDue; "sac" and "price" require exactly one of Count or InitialAmount (Count is also required for "american") plus Interval; only "custom" requires Invoices, leaving the others unused. For every type, including "custom", exactly one of NominalAmount or Amount must also be provided.
//	- Invoices [slice of Invoice structs]: Slice of Invoice structs to be created and sent to the credit receiver. ex: []Invoice{Invoice{}, Invoice{}}
//	- NominalAmount [int]: Amount in cents transferred to the credit receiver, before deductions, for every type including "custom". Provide exactly one of NominalAmount or Amount; the other value, along with TaxAmount and the interest rates, is computed from the invoice schedule. ex: 11234 (= R$ 112.34)
//	- Amount [int]: Net amount in cents to be disbursed to the credit receiver, for every type including "custom". Provide exactly one of NominalAmount or Amount; the other value, along with TaxAmount and the interest rates, is computed from the invoice schedule.
//	- NominalInterest [float64]: Yearly nominal interest rate of the credit note, in percentage. ex: 12.5
//	- InitialDue [time.Time]: Date of the first invoice. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- Count [int]: Quantity of invoices for payment. ex: 12
//	- InitialAmount [int]: Value of the first invoice in cents. ex: 1234 (= R$12.34)
//	- Interval [string]: Interval between invoices. ex: "year", "month"
//
//	Parameters (optional):
//	- RebateAmount [int, default nil]: Credit analysis fee deducted from lent amount. ex: 11234 (= R$ 112.34)
//
//	Attributes (return-only):
//	- Interest [float64]: Yearly effective interest rate of the credit note, in percentage. ex: 12.5
//	- TaxAmount [int]: Tax amount included in the credit note. ex: 100

type CreditNotePreview struct {
	Type            string            `json:",omitempty"`
	NominalAmount   int               `json:",omitempty"`
	Scheduled       *time.Time        `json:",omitempty"`
	TaxId           string            `json:",omitempty"`
	Invoices        []Invoice.Invoice `json:",omitempty"`
	NominalInterest float64           `json:",omitempty"`
	InitialDue      *time.Time        `json:",omitempty"`
	Count           int               `json:",omitempty"`
	InitialAmount   int               `json:",omitempty"`
	Interval        string            `json:",omitempty"`
	RebateAmount    int               `json:",omitempty"`
	Amount          int               `json:",omitempty"`
	Interest        float64           `json:",omitempty"`
	TaxAmount       int               `json:",omitempty"`
}
