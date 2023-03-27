package invoice

//	CreditNote.Invoice.Description struct
//
//	Invoice description information.
//
//	Parameters (required):
//	- Key [string]: Description for the value. ex: "Taxes"
//
//	Parameters (optional):
//	- Value [string, default nil]: Amount related to the described key. ex: "R$100,00"

type Description struct {
	Key   string `json:",omitempty"`
	Value string `json:",omitempty"`
}
