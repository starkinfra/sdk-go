package invoice

//	credit_note.invoice.Description object
//	Invoice description information.
//
//	Parameters (required):
//	- key [string]: Description for the value. ex: "Taxes"
//
//	Parameters (optional):
//	- value [string, default ""]: amount related to the described key. ex: "R$100,00"

type Description struct {
	Key   string
	Value string
}

var resource = map[string]string{"class": Description{}, "name": "Description"}
