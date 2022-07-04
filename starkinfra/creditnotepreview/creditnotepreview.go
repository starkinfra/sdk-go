package creditnotepreview

type CreditNotePreview struct {
	Type            string
	NominalAmount   int
	Scheduled       string
	TaxId           string
	Invoices        Invoice
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

var resource = map[string]string{"class": CreditNotePreview{}, "name": "CreditNotePreview"}

func ParseOptionalInvoices() {

}

func Create() {

}
