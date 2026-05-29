package individualaccountrequest

//	IndividualAccountRequest.Address struct
//
//	Embedded value object describing the individual's residential address. It has no endpoints —
//	it is exposed only as the Address field on the parent IndividualAccountRequest and is serialized
//	as a nested JSON object on the wire, never flattened into addressStreet/addressCity/etc.
//
//	Parameters (required):
//	- Street [string]: Street name. ex: "Rua do Estilo Barroco"
//	- Number [string]: Street number. String, not integer (may contain non-digit chars). ex: "648"
//	- Neighborhood [string]: Neighborhood / district. ex: "Santo Amaro"
//	- City [string]: City. ex: "Sao Paulo"
//	- State [string]: State (BR 2-letter code). ex: "SP"
//	- ZipCode [string]: ZIP code (BR CEP). Accepts formatted or digit-only. ex: "05724005"

type Address struct {
	Street       string `json:",omitempty"`
	Number       string `json:",omitempty"`
	Neighborhood string `json:",omitempty"`
	City         string `json:",omitempty"`
	State        string `json:",omitempty"`
	ZipCode      string `json:",omitempty"`
}
