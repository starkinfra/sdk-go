package credit_note

//	CreditNote object
//	CreditNotes are used to generate CCB contracts between you and your customers.
//	When you initialize a CreditNote, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the list of created objects.
//
//	Parameters (required):
//	- template_id [string]: ID of the contract template on which the credit note will be based. ex: template_id="0123456789101112"
//	- name [string]: credit receiver's full name. ex: name="Edward Stark"
//	- tax_id [string]: credit receiver's tax ID (CPF or CNPJ). ex: "20.018.183/0001-80"
//	- nominal_amount [integer]: amount in cents transferred to the credit receiver, before deductions. ex: nominal_amount=11234 (= R$ 112.34)
//	- scheduled [datetime.date, datetime.datetime or string]: date of transfer execution. ex: scheduled=datetime(2020, 3, 10)
//	- invoices [list of Invoice objects]: list of Invoice objects to be created and sent to the credit receiver. ex: invoices=[Invoice(), Invoice()]
//	- payment [credit_note.Transfer]: payment entity to be created and sent to the credit receiver. ex: payment=credit_note.Transfer()
//	- signers [list of credit_note.Signer objects]: signer's name, contact and delivery method for the signature request. ex: signers=[credit_note.Signer(), credit_note.Signer()]
//	- external_id [string]: a string that must be unique among all your CreditNotes, used to avoid resource duplication. ex: "my-internal-id-123456"
//	- street_line_1 [string]: credit receiver main address. ex: "Av. Paulista, 200"
//	- street_line_2 [string]: credit receiver address complement. ex: "Apto. 123"
//	- district [string]: credit receiver address district / neighbourhood. ex: "Bela Vista"
//	- city [string]: credit receiver address city. ex: "Rio de Janeiro"
//	- state_code [string]: credit receiver address state. ex: "GO"
//	- zip_code [string]: credit receiver address zip code. ex: "01311-200"
//
//	Parameters (conditionally required):
//	- payment_type [string]: payment type, inferred from the payment parameter if it is not a dictionary. ex: "transfer"
//	Parameters (optional):
//	- rebate_amount [integer, default None]: credit analysis fee deducted from lent amount. ex: rebate_amount=11234 (= R$ 112.34)
//	- tags [list of strings, default None]: list of strings for reference when searching for CreditNotes. ex: tags=["employees", "monthly"]
//	Attributes (return-only):
//	- id [string]: unique id returned when the CreditNote is created. ex: "5656565656565656"
//	- amount [integer]: CreditNote value in cents. ex: 1234 (= R$ 12.34)
//	- expiration [integer or datetime.timedelta]: time interval in seconds between due date and expiration date. ex 123456789
//	- document_id [string]: ID of the signed document to execute this CreditNote. ex: "4545454545454545"
//	- status [string]: current status of the CreditNote. ex: "canceled", "created", "expired", "failed", "processing", "signed", "success"
//	- transaction_ids [list of strings]: ledger transaction ids linked to this CreditNote. ex: ["19827356981273"]
//	- workspace_id [string]: ID of the Workspace that generated this CreditNote. ex: "4545454545454545"
//	- tax_amount [integer]: tax amount included in the CreditNote. ex: 100
//	- interest [float]: yearly effective interest rate of the credit note, in percentage. ex: 12.5
//	- created [datetime.datetime]: creation datetime for the CreditNote. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: latest update datetime for the CreditNote. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type Creditnote struct {
	TemplateId     string
	Name           string
	TaxId          string
	NominalAmount  string
	Scheduled      string
	Payment        Transfer
	Invoices       []Invoice
	Signers        []Signer
	ExternalId     string
	StreetLine1    string
	StreetLine2    string
	District       string
	City           string
	StateCode      string
	ZipCode        string
	PaymentType    string
	RebateAmount   int
	Tags           []string
	Amount         int
	DocumentId     string
	Status         string
	TransactionIds []string
	WorkspaceId    string
	TaxAmount      int
	Interest       int
	Created        string
	Update         string
}

func ParseSigners() {

}

func ParseInvoices() {

}

func ParsePayment() {

}

func Create() {

}

func Get() {
	return utils.Get()
}

func Query() {

}

func Page() {

}

func Cancel() {

}
