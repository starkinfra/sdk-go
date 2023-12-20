package event

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	CreditNoteLog "github.com/starkinfra/sdk-go/starkinfra/creditnote/log"
	IssuingCardLog "github.com/starkinfra/sdk-go/starkinfra/issuingcard/log"
	IssuingInvoiceLog "github.com/starkinfra/sdk-go/starkinfra/issuinginvoice/log"
	IssuingPurchaseLog "github.com/starkinfra/sdk-go/starkinfra/issuingpurchase/log"
	PixChargebackLog "github.com/starkinfra/sdk-go/starkinfra/pixchargeback/log"
	PixClaimLog "github.com/starkinfra/sdk-go/starkinfra/pixclaim/log"
	PixInfractionLog "github.com/starkinfra/sdk-go/starkinfra/pixinfraction/log"
	PixKeyLog "github.com/starkinfra/sdk-go/starkinfra/pixkey/log"
	PixRequestLog "github.com/starkinfra/sdk-go/starkinfra/pixrequest/log"
	PixReversalLog "github.com/starkinfra/sdk-go/starkinfra/pixreversal/log"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	Webhook Event struct
//
//	An Event is the notification received from the subscription to the Webhook.
//	Events cannot be created, but may be retrieved from the Stark Infra API to
//	list all generated updates on entities.
//
//	Attributes:
//	- Id [string]: Unique id returned when the Event is created. ex: "5656565656565656"
//	- Log [Log]: A Log struct from one of the subscribed services (PixRequestLog, PixReversalLog)
//	- Created [time.Time]: Creation datetime for the notification Event. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- IsDelivered [bool]: True if the Event has been successfully delivered to the user url. ex: False
//	- Subscription [string]: Service that triggered this Event. ex: "pix-request.in", "pix-request.out"
//	- WorkspaceId [string]: ID of the Workspace that generated this Event. Mostly used when multiple Workspaces have Webhooks registered to the same endpoint. ex: "4545454545454545"

type Event struct {
	Id           string      `json:",omitempty"`
	Log          interface{} `json:",omitempty"`
	Created      *time.Time  `json:",omitempty"`
	IsDelivered  bool        `json:",omitempty"`
	Subscription string      `json:",omitempty"`
	WorkspaceId  string      `json:",omitempty"`
}

var resource = map[string]string{"name": "Event"}

func Get(id string, user user.User) (Event, Error.StarkErrors) {
	//	Retrieve a specific notification Event
	//
	//	Receive a single notification Event struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- event struct that corresponds to the given id.
	var event Event
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &event)
	if unmarshalError != nil {
		return event, err
	}
	return event, err
}

func Query(params map[string]interface{}, user user.User) chan Event {
	//	Retrieve notification Events
	//
	//	Receive a channel of notification Event structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- isDelivered [bool, default nil]: Bool to filter successfully delivered events. ex: True or False
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of Event structs with updated attributes
	var event Event
	events := make(chan Event)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &event)
			if err != nil {
				print(err)
			}
			events <- event
		}
		close(events)
	}()
	return events
}

func Page(params map[string]interface{}, user user.User) ([]Event, string, Error.StarkErrors) {
	//	Retrieve paged Events
	//
	//	Receive a slice of up to 100 Event structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- isDelivered [bool, default nil]: Bool to filter successfully delivered events. ex: True or False
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of Event structs with updated attributes
	//	- cursor to retrieve the next page of Event structs
	var events []Event
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &events)
	if unmarshalError != nil {
		return ParseEvents(events), cursor, err
	}
	return ParseEvents(events), cursor, err
}

func Delete(id string, user user.User) (Event, Error.StarkErrors) {
	//	Delete a Webhook Event entity
	//
	//	Delete a notification Event entity previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Event unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- deleted Event struct
	var event Event
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &event)
	if unmarshalError != nil {
		return event, err
	}
	return event, err
}

func Update(id string, isDelivered bool, user user.User) (Event, Error.StarkErrors) {
	//	Update notification Event entity
	//
	//	Update notification Event by passing id.
	//	If isDelivered is True, the event will no longer be returned on queries with isDelivered: False.
	//
	//	Parameters (required):
	//	- id [slice of strings]: Event unique ids. ex: "5656565656565656"
	//	- isDelivered [bool]: If True and event hasn't been delivered already, event will be set as delivered. ex: True
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- target Event with updated attributes
	var event Event
	var patchData map[string]interface{}
	patchData["isDelivered"] = isDelivered
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &event)
	if unmarshalError != nil {
		return event, err
	}
	return event, err
}

func Parse(content string, signature string, user user.User) Event {
	//	Create single notification Event from a content string
	//
	//	Create a single Event struct received from Event listening at subscribed user endpoint.
	//	If the provided digital signature does not check out with the StarkInfra public key, an
	//	error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: Response content from request received at user endpoint (not parsed)
	//	- signature [string]: Base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- parsed Event struct
	var event Event
	unmarshalError := json.Unmarshal([]byte(utils.ParseAndVerify(content, signature, "event", user)), &event)
	if unmarshalError != nil {
		return event
	}
	return event
}

func (e Event) ParseLog() Event {
	if e.Subscription == "pix-key" {
		var log PixKeyLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-claim" {
		var log PixClaimLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-chargeback" {
		var log PixChargebackLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-infraction" {
		var log PixInfractionLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-request.in" {
		var log PixRequestLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-request.out" {
		var log PixReversalLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-reversal.in" {
		var log PixReversalLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "pix-reversal.out" {
		var log PixReversalLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "issuing-card" {
		var log IssuingCardLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "issuing-invoice" {
		var log IssuingInvoiceLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "issuing-purchase" {
		var log IssuingPurchaseLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	if e.Subscription == "credit-note" {
		var log CreditNoteLog.Log
		marshal, _ := json.Marshal(e.Log)
		err := json.Unmarshal(marshal, &log)
		if err != nil {
			panic(err)
		}
		e.Log = log
		return e
	}
	return e
}

func ParseEvents(events []Event) []Event {
	for i := 0; i < len(events); i++ {
		events[i] = events[i].ParseLog()
	}
	return events
}
