package invoice

import "time"

//	CreditNote.Invoice.Discount struct
//
//	Invoice discount information.
//
//	Parameters (required):
//	- Percentage [float64]: Percentage of discount applied until specified due date. ex: 2.0
//	- Due [time.Time]: Due datetime for the discount. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Discount struct {
	Percentage float64    `json:",omitempty"`
	Due        *time.Time `json:",omitempty"`
}
