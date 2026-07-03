package pixdispute

import "time"

//	PixDispute.Transaction struct
//
//	A PixDispute.Transaction is a return-only sub-object embedded in the
//	Transactions slice of a PixDispute. It represents a node of the dispute
//	graph created by the Central Bank and is never created by the user.
//
//	Attributes (return-only):
//	- EndToEndId [string]: Central Bank's unique transaction id. ex: "E20018183202201201450u34sDGd19lz"
//	- Amount [int]: Refundable amount in cents. ex: 11234 (= R$ 112.34)
//	- NominalAmount [int]: Transaction amount in cents. ex: 11234 (= R$ 112.34)
//	- ReceiverType [string]: Receiver's type. Options: "individual", "business"
//	- ReceiverTaxIdCreated [string]: Receiver's taxId creation date (business type only).
//	- ReceiverAccountCreated [string]: Receiver's account creation date.
//	- ReceiverBankCode [string]: Receiver's bank code. ex: "20018183"
//	- ReceiverId [string]: Identifier of the accountholder in the graph.
//	- SenderType [string]: Sender's type. Options: "individual", "business"
//	- SenderTaxIdCreated [string]: Sender's taxId creation date (business type only).
//	- SenderAccountCreated [string]: Sender's account creation date.
//	- SenderBankCode [string]: Sender's bank code. ex: "20018183"
//	- SenderId [string]: Identifier of the accountholder in the graph.
//	- Settled [time.Time]: Settled datetime of the transaction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Transaction struct {
	EndToEndId             string     `json:",omitempty"`
	Amount                 int        `json:",omitempty"`
	NominalAmount          int        `json:",omitempty"`
	ReceiverType           string     `json:",omitempty"`
	ReceiverTaxIdCreated   string     `json:",omitempty"`
	ReceiverAccountCreated string     `json:",omitempty"`
	ReceiverBankCode       string     `json:",omitempty"`
	ReceiverId             string     `json:",omitempty"`
	SenderType             string     `json:",omitempty"`
	SenderTaxIdCreated     string     `json:",omitempty"`
	SenderAccountCreated   string     `json:",omitempty"`
	SenderBankCode         string     `json:",omitempty"`
	SenderId               string     `json:",omitempty"`
	Settled                *time.Time `json:",omitempty"`
}
