package creditsigner

//	CreditSigner struct
//
//	CreditNote signer's information.
//
//	Parameters (required):
//	- Name [string]: Signer's name. ex: "Tony Stark"
//	- Contact [string]: Signer's contact information. ex: "tony@starkindustries.com"
//	- Method [string]: Delivery method for the contract. ex: "link"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the CreditSigner is created. ex: "5656565656565656"

type CreditSigner struct {
	Name    string `json:",omitempty"`
	Contact string `json:",omitempty"`
	Method  string `json:",omitempty"`
	Id      string `json:",omitempty"`
}
