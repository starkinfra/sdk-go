package businessaccountrequest

//	BusinessAccountRequest.Address struct
//
//	The Address struct is the structured address of the company referenced by a
//	BusinessAccountRequest. It is embedded on the parent's Address field and has no endpoints of its own.
//
//	Parameters (required):
//	- Street [string]: street name. ex: "Av. Faria Lima"
//	- Number [string]: street number. ex: "2000"
//	- Neighborhood [string]: neighborhood / district. ex: "Itaim Bibi"
//	- City [string]: city. ex: "Sao Paulo"
//	- State [string]: state (BR 2-letter code). ex: "SP"
//	- ZipCode [string]: ZIP code (BR CEP), formatted or digit-only. ex: "04538-132"
//
//	Parameters (optional):
//	- Complement [string, default nil]: address complement. ex: "Sala 42"

type Address struct {
	Street       string `json:",omitempty"`
	Number       string `json:",omitempty"`
	Neighborhood string `json:",omitempty"`
	City         string `json:",omitempty"`
	State        string `json:",omitempty"`
	ZipCode      string `json:",omitempty"`
	Complement   string `json:",omitempty"`
}
