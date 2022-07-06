package invoice

//	credit_note.Invoice object
//	Invoice issued after the contract is signed, to be paid by the credit receiver.
//
//	Parameters (required):
//	- amount [integer]: Invoice value in cents. Minimum = 1 (any value will be accepted). ex: 1234 (= R$ 12.34)
//
//	Parameters (optional):
//	- due [datetime.datetime or datetime.date or string, default now + 2 days]: Invoice due date in UTC ISO format. ex: "2020-10-28T17:59:26.249976+00:00" for immediate invoices and "2020-10-28" for scheduled invoices
//	- expiration [integer or datetime.timedelta, default 5097600 (59 days)]: time interval in seconds between due date and expiration date. ex 123456789
//	- tags [list of strings, default None]: list of strings for tagging
//	- descriptions [list of credit_note.invoice.Description objects or dictionaries, default nil]: list Description objects
//
//	Attributes (return-only):
//	- id [string]: unique id returned when Invoice is created. ex: "5656565656565656"
//	- name [string]: payer name. ex: "Iron Bank S.A."
//	- tax_id [string]: payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- pdf [string]: public Invoice PDF URL. ex: "https://invoice.starkbank.com/pdf/d454fa4e524441c1b0c1a729457ed9d8"
//	- link [string]: public Invoice webpage URL. ex: "https://my-workspace.sandbox.starkbank.com/invoicelink/d454fa4e524441c1b0c1a729457ed9d8"
//	- fine [float]: Invoice fine for overdue payment in %. ex: 2.5
//	- interest [float]: Invoice monthly interest for overdue payment in %. ex: 5.2
//	- nominal_amount [integer]: Invoice emission value in cents (will change if invoice is updated, but not if it's paid). ex: 400000
//	- fine_amount [integer]: Invoice fine value calculated over nominal_amount. ex: 20000
//	- interest_amount [integer]: Invoice interest value calculated over nominal_amount. ex: 10000
//	- discount_amount [integer]: Invoice discount value calculated over nominal_amount. ex: 3000
//	- discounts [list of credit_note.invoice.Discount objects]: list of Discount objects. ex: [Discount()]
//	- brcode [string]: BR Code for the Invoice payment. ex: "00020101021226800014br.gov.bcb.pix2558invoice.starkbank.com/f5333103-3279-4db2-8389-5efe335ba93d5204000053039865802BR5913Arya Stark6009Sao Paulo6220051656565656565656566304A9A0"
//	- status [string]: current Invoice status. ex: "registered" or "paid"
//	- fee [integer]: fee charged by this Invoice. ex: 200 (= R$ 2.00)
//	- transaction_ids [list of strings]: ledger transaction ids linked to this Invoice (if there are more than one, all but the first are reversals or failed reversal chargebacks). ex: ["19827356981273"]
//	- created [datetime.datetime]: creation datetime for the Invoice. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: latest update datetime for the Invoice. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type Invoice struct {
	Amount         int           `json:"amount"`
	Due            string        `json:"due"`
	Expiration     int           `json:"expiration"`
	Tags           []string      `json:"tags"`
	Description    []Description `json:"description"`
	Id             string        `json:"id"`
	Name           string        `json:"name"`
	TaxId          string        `json:"taxId"`
	Pdf            string        `json:"pdf"`
	Link           string        `json:"link"`
	Fine           float64       `json:"fine"`
	Interest       float64       `json:"interest"`
	NominalAmount  int           `json:"nominalAmount"`
	FineAmount     int           `json:"fineAmount"`
	InterestAmount int           `json:"interestAmount"`
	DiscountAmount int           `json:"discountAmount"`
	Discounts      []Discount    `json:"discounts"`
	BrCode         string        `json:"brCode"`
	Status         string        `json:"status"`
	Fee            int           `json:"fee"`
	TransactionIds []string      `json:"transactionIds"`
	Created        string        `json:"created"`
	Updated        string        `json:"updated"`
}

func ParseDiscounts() {

}

func ParseDescriptions() {

}
