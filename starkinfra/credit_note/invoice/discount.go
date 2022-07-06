package invoice

//	credit_note.invoice.Discount object
//	Invoice discount information.
//
//	Parameters (required):
//	- percentage [float]: percentage of discount applied until specified due date
//	- due [datetime.datetime or string]: due datetime for the discount

type Discount struct {
	Percentage float64 `json:"percentage"`
	Due        string  `json:"due"`
}

var resource = map[string]string{"class": Description{}, "name": "Description"}
