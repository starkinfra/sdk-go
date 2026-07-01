package ledger

//	Ledger.Rule struct
//
//	The Ledger.Rule struct modifies the behavior of Ledger structs when passed as an argument upon their creation or update.
//
//	Parameters (required):
//	- Key [string]: Rule to be customized, describes what Ledger behavior will be altered. ex: "minimumBalance", "maximumBalance"
//	- Value [int]: Value of the rule. ex: 1000

type Rule struct {
	Key   string `json:",omitempty"`
	Value int    `json:"value"`
}
