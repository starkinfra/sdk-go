package credit_note_preview

//	CreditNotePreview object
//	A CreditNotePreview is used to preview a CCB contract between the borrower and lender with a specific table type.
//	When you initialize a CreditNotePreview, the entity will be automatically sent to the Stark Infra API.
//	The 'create' function sends the objects
//	to the Stark Infra API and returns the list of preview data.
//
//	Parameters (required):
//	- type [string]: table type that defines the amortization system. Options: "sac", "price", "american", "bullet", "custom"
//	- nominal_amount [integer]: amount in cents transferred to the credit receiver, before deductions. ex: nominal_amount=11234 (= R$ 112.34)
//	- scheduled [datetime.date, datetime.datetime or string]: date of transfer execution. ex: scheduled=datetime(2020, 3, 10)
//	- tax_id [string]: credit receiver's tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//
//	Parameters (conditionally required):
//	- invoices [list of Invoice objects]: list of Invoice objects to be created and sent to the credit receiver. ex: invoices=[Invoice(), Invoice()]
//	- nominal_interest [float]: yearly nominal interest rate of the credit note, in percentage. ex: 12.5
//	- initial_due [datetime.date, datetime.datetime or string]: date of the first invoice. ex: scheduled=datetime(2020, 3, 10)
//	- count [integer]: quantity of invoices for payment. ex: 12
//	- initial_amount [integer]: value of the first invoice in cents. ex: 1234 (= R$12.34)
//	- interval [string]: interval between invoices. ex: "year", "month"
//
//	Parameters (optional):
//	- rebate_amount [integer, default None]: credit analysis fee deducted from lent amount. ex: rebate_amount=11234 (= R$ 112.34)
//
//	Attributes (return-only):
//	- amount [integer]: CreditNote value in cents. ex: 1234 (= R$ 12.34)
//	- interest [float]: yearly effective interest rate of the credit note, in percentage. ex: 12.5
//	- tax_amount [integer]: tax amount included in the CreditNote. ex: 100

type CreditNotePreview struct {
	Type            string
	NominalAmount   int
	Scheduled       string
	TaxId           string
	Invoices        []Invoice
	NominalInterest float64
	InitialDue      string
	Count           int
	InitialAmount   int
	Interval        string
	RebateAmount    int
	Amount          int
	Interest        float64
	TaxAmount       int
}

var resource = map[string]any{"class": CreditNotePreview{}, "name": "CreditNotePreview"}

func ParseOptionalInvoices() {

}

func Create() {
	//	Create CreditNotePreview
	//	Send a list of CreditNotePreview objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- previews [list of CreditNotePreview objects]: list of CreditNotePreview objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of CreditNotePreview objects with updated attributes
}
