package pixuser

import (
	"time"
)

//	PixUser.Statistics struct
//
//	The PixUser.Statistics struct stores fraud statistics of a Pix user.
//
//	Attributes (return-only):
//	- Value [int]: Value of the statistic. ex: 3
//	- Type [string]: Type of the statistic. ex: "infractions"
//	- Source [string]: Source of the statistic. ex: "keyManagement"
//	- After [time.Time]: Start datetime considered for the statistic. ex: time.Date(2020, 4, 23, 23, 0, 0, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the statistic. ex: time.Date(2020, 4, 23, 23, 0, 0, 0, time.UTC),

type Statistics struct {
	Value   int        `json:",omitempty"`
	Type    string     `json:",omitempty"`
	Source  string     `json:",omitempty"`
	After   *time.Time `json:",omitempty"`
	Updated *time.Time `json:",omitempty"`
}
