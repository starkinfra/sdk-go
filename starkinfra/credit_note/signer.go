package credit_note

//	credit_note.Signer object
//	CreditNote signer's information.
//
//	Parameters (required):
//	- name [string]: signer's name. ex: "Tony Stark"
//	- contact [string]: signer's contact information. ex: "tony@starkindustries.com"
//	- method [string]: delivery method for the contract. ex: "link"

type Signer struct {
	Name    string
	Contact string
	Method  string
}
