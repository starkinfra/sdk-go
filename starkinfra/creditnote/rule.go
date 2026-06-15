package creditnote

//	CreditNote.Rule struct
//
//	The Rule struct modifies the behavior of CreditNotes when passed as an argument upon their creation.
//
//	Parameters (required):
//	- Key [string]: Rule to be customized, describes what CreditNote behavior will be altered. ex: "invoiceCreationMode"
//	- Value [string]: Value of the rule. ex: "scheduled", "instant", "never"

type Rule struct {
	Key   string `json:",omitempty"`
	Value string `json:",omitempty"`
}
