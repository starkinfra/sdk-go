package attempt

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	Event.Attempt struct
//
//	When an Event delivery fails, an event attempt will be registered.
//	It carries information meant to help you debug event reception issues.
//
//	Attributes (return-only):
//	- Id [string]: Unique id that identifies the delivery attempt. ex: "5656565656565656"
//	- Code [string]: Delivery error code. ex: badHttpStatus, badConnection, timeout
//	- Message [string]: Delivery error full description. ex: "HTTP POST request returned status 404"
//	- EventId [string]: ID of the Event whose delivery failed. ex: "4848484848484848"
//	- WebhookId [string]: ID of the Webhook that triggered this event. ex: "5656565656565656"
//	- Created [time.Time]: Datetime representing the moment when the attempt was made. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type Attempt struct {
	Id        string     `json:",omitempty"`
	Code      string     `json:",omitempty"`
	Message   string     `json:",omitempty"`
	EventId   string     `json:",omitempty"`
	WebhookId string     `json:",omitempty"`
	Created   *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "EventAttempt"}

func Get(id string, user user.User) (Attempt, Error.StarkErrors) {
	//	Retrieve a specific Event.Attempt
	//
	//	Receive a single event.Attempt struct previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- event.Attempt struct that corresponds to the given id.
	var attempt Attempt
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &attempt)
	if unmarshalError != nil {
		return attempt, err
	}
	return attempt, err
}

func Query(params map[string]interface{}, user user.User) (chan Attempt, chan Error.StarkErrors) {
	//	Retrieve event.Attempt structs
	//
	//	Receive a channel of Event.Attempt structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- eventIds [slice of strings, default nil]: Slice of Event ids to filter attempts. ex: []string{"5656565656565656", "4545454545454545"}
	//		- webhookIds [slice of strings, default nil]: Slice of Webhook ids to filter attempts. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- channel of Event.Attempt structs with updated attributes
	var attempt Attempt
	attempts := make(chan Attempt)
	attemptsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &attempt)
			if err != nil {
				attemptsError <- Error.UnknownError(err.Error())
				continue
			}
			attempts <- attempt
		}
		for err := range errorChannel {
			attemptsError <- err
		}
		close(attempts)
		close(attemptsError)
	}()
	return attempts, attemptsError
}

func Page(params map[string]interface{}, user user.User) ([]Attempt, string, Error.StarkErrors) {
	//	Retrieve paged Event.Attempt structs
	//
	//	Receive a slice of up to 100 Event.Attempt structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- eventIds [slice of strings, default nil]: Slice of Event ids to filter attempts. ex: []string{"5656565656565656", "4545454545454545"}
	//		- webhookIds [slice of strings, default nil]: Slice of Webhook ids to filter attempts. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- slice of Event.Attempt structs with updated attributes
	//	- cursor to retrieve the next page of Event.Attempt structs
	var attempts []Attempt
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &attempts)
	if unmarshalError != nil {
		return attempts, cursor, err
	}
	return attempts, cursor, err
}
