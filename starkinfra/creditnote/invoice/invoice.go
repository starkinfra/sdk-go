package invoice

//	CreditNote.Invoice struct
//
//	Invoice issued after the contract is signed, to be paid by the credit receiver.
//
//	Parameters (required):
//	- Amount [int]: Invoice value in cents. Minimum = 1 (any value will be accepted). ex: 1234 (= R$ 12.34)
//
//	Parameters (optional):
//	- Due [string, default now + 2 days]: Invoice due date in UTC ISO format. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC), for immediate invoices and time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC), for scheduled invoices
//	- Expiration [int, default 5097600 (59 days)]: Time interval in seconds between due date and expiration date. ex 123456789
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"tony", "stark"}
//	- Descriptions [slice of invoice.Description structs, default nil]: List Description structs
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when Invoice is created. ex: "5656565656565656"
//	- Name [string]: Payer name. ex: "Iron Bank S.A."
//	- TaxId [string]: Payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- Pdf [string]: Public Invoice PDF URL. ex: "https://invoice.starkbank.com/pdf/d454fa4e524441c1b0c1a729457ed9d8"
//	- Link [string]: Public Invoice webpage URL. ex: "https://my-workspace.sandbox.starkbank.com/invoicelink/d454fa4e524441c1b0c1a729457ed9d8"
//	- Fine [float64]: Invoice fine for overdue payment in %. ex: 2.5
//	- Interest [float64]: Invoice monthly interest for overdue payment in %. ex: 5.2
//	- NominalAmount [int]: Invoice emission value in cents (will change if invoice is updated, but not if it's paid). ex: 400000
//	- FineAmount [int]: Invoice fine value calculated over nominalAmount. ex: 20000
//	- InterestAmount [int]: Invoice interest value calculated over nominalAmount. ex: 10000
//	- DiscountAmount [int]: Invoice discount value calculated over nominalAmount. ex: 3000
//	- Discounts [slice of invoice.Discount structs]: Slice of Discount structs. ex: []string{Discount()]
//	- Brcode [string]: BR Code for the Invoice payment. ex: "00020101021226800014br.gov.bcb.pix2558invoice.starkbank.com/f5333103-3279-4db2-8389-5efe335ba93d5204000053039865802BR5913Arya Stark6009Sao Paulo6220051656565656565656566304A9A0"
//	- Status [string]: Current Invoice status. ex: "registered" or "paid"
//	- Fee [int]: Fee charged by this Invoice. ex: 200 (= R$ 2.00)
//	- TransactionIds [slice of strings]: Ledger transaction ids linked to this Invoice (if there are more than one, all but the first are reversals or failed reversal chargebacks). ex: []string{"19827356981273"}
//	- Created [time.Time]: Creation datetime for the Invoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the Invoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Invoice struct {
	Amount         int           `json:",omitempty"`
	Due            string        `json:",omitempty"`
	Expiration     int           `json:",omitempty"`
	Tags           []string      `json:",omitempty"`
	Description    []Description `json:",omitempty"`
	Fine           float64       `json:",omitempty"`
	Interest       float64       `json:",omitempty"`
	Id             string        `json:",omitempty"`
	Name           string        `json:",omitempty"`
	TaxId          string        `json:",omitempty"`
	Pdf            string        `json:",omitempty"`
	Link           string        `json:",omitempty"`
	NominalAmount  int           `json:",omitempty"`
	FineAmount     int           `json:",omitempty"`
	InterestAmount int           `json:",omitempty"`
	DiscountAmount int           `json:",omitempty"`
	Discounts      []Discount    `json:",omitempty"`
	BrCode         string        `json:",omitempty"`
	Status         string        `json:",omitempty"`
	Fee            int           `json:",omitempty"`
	TransactionIds []string      `json:",omitempty"`
	Created        string        `json:",omitempty"`
	Updated        string        `json:",omitempty"`
}
