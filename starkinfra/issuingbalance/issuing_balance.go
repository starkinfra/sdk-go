package issuingbalance

import (
	"encoding/json"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingBalance struct
//
//	The IssuingBalance struct displays the current issuing balance of the Workspace,
//	which is the result of the sum of all transactions within this
//	Workspace. The balance is never generated by the user, but it
//	can be retrieved to see the available information.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when IssuingBalance is created. ex: "5656565656565656"
//	- Amount [int]: Current balance amount of the Workspace in cents. ex: 200 (= R$ 2.00)
//	- Currency [string]: Currency of the current Workspace. Expect others to be added eventually. ex: "BRL"
//	- Updated [string]: Latest update datetime for the IssuingBalance. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IssuingBalance struct {
	Id       string     `json:",omitempty"`
	Amount   int        `json:",omitempty"`
	Currency string     `json:",omitempty"`
	Updated  *time.Time `json:",omitempty"`
}

var object IssuingBalance
var resource = map[string]string{"name": "IssuingBalance"}

func Get(user user.User) IssuingBalance {
	//	Retrieve the IssuingBalance struct
	//
	//	Receive the IssuingBalance struct linked to your Workspace in the Stark Infra API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- IssuingBalance struct with updated attributes
	balance := make(chan IssuingBalance)
	query := utils.Query(resource, nil, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			balance <- object
		}
		close(balance)
	}()
	return <-balance
}
