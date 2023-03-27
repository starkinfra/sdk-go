package creditnote

import "time"

//	CreditNote.Transfer struct
//
//	Transfer sent to the credit receiver after the contract is signed.
//
//	Parameters (required):
//	- Name [string]: Receiver full name. ex: "Anthony Edward Stark"
//	- TaxId [string]: Receiver tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- BankCode [string]: Code of the receiver bank institution in Brazil. ex: "20018183" or "341"
//	- BranchCode [string]: Receiver bank account branch. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- AccountNumber [string]: Receiver bank account number. Use '-' before the verifier digit. ex: "876543-2"
//
//	Parameters (optional):
//	- AccountType [string, default "checking"]: Receiver bank account type. This parameter only has effect on Pix Transfers. ex: "checking", "savings", "salary" or "payment"
//	- Tags [slice of strings, default nil]: Slice of strings for reference when searching for transfers. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the transfer is created. ex: "5656565656565656"
//	- Amount [int]: Amount in cents to be transferred. ex: 1234 (= R$ 12.34)
//	- ExternalId [string]: URL safe string that must be unique among all your transfers. Duplicated external_ids will cause failures. By default, this parameter will block any transfer that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//  - Scheduled [time.Time]: Date when the transfer will be processed. May be pushed to next business day if necessary. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- Description [string]: Optional description to override default description to be shown in the bank statement. ex: "Payment for service #1234"
//	- Fee [int]: Fee charged when the Transfer is processed. ex: 200 (= R$ 2.00)
//	- Status [string]: Current transfer status. ex: "success" or "failed"
//	- TransactionIds [slice of strings]: Ledger Transaction IDs linked to this Transfer (if there are two, the second is the chargeback). ex: []string{"19827356981273"}
//	- Created [time.Time]: Creation datetime for the transfer. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the transfer. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Transfer struct {
	Name           string     `json:",omitempty"`
	TaxId          string     `json:",omitempty"`
	BankCode       string     `json:",omitempty"`
	BranchCode     string     `json:",omitempty"`
	AccountNumber  string     `json:",omitempty"`
	AccountType    string     `json:",omitempty"`
	Tags           []string   `json:",omitempty"`
	Id             string     `json:",omitempty"`
	Amount         int        `json:",omitempty"`
	ExternalId     string     `json:",omitempty"`
	Scheduled      *time.Time `json:",omitempty"`
	Description    string     `json:",omitempty"`
	Fee            int        `json:",omitempty"`
	Status         string     `json:",omitempty"`
	TransactionIds []string   `json:",omitempty"`
	Created        string     `json:",omitempty"`
	Updated        string     `json:",omitempty"`
}
