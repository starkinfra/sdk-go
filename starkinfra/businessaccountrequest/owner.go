package businessaccountrequest

//	BusinessAccountRequest.Owner struct
//
//	The Owner struct represents a company owner referenced by a BusinessAccountRequest. Each owner
//	completes its own identity verification through an independent webview. It is embedded on the
//	parent's Owners field and has no endpoints of its own.
//
//	Parameters (required):
//	- TaxId [string]: owner's tax ID (CPF). ex: "012.345.678-90"
//	- Name [string]: owner's full name (minimum 5 characters). ex: "Jamie Lannister"
//	- Role [string]: owner's role in the company. Options: "partner", "representative"
//
//	Attributes (return-only):
//	- IdentityId [string]: unique id of the identity verification linked to this owner. ex: "5709594221805568"
//	- ValidatorLink [string]: webview link to be delivered to the owner to complete biometrics and document capture. Treat it as a credential: deliver it through a secure channel, never log it or write it to disk.
//	- Status [string]: current status of the owner verification. Options: "created", "approved", "denied"

type Owner struct {
	TaxId         string `json:",omitempty"`
	Name          string `json:",omitempty"`
	Role          string `json:",omitempty"`
	IdentityId    string `json:",omitempty"`
	ValidatorLink string `json:",omitempty"`
	Status        string `json:",omitempty"`
}
