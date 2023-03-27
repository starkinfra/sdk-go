package webhook

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
)

//	Webhook struct
//
//	A Webhook is used to subscribe to notification events on a user-selected endpoint.
//	Currently, available services for subscription are contract, credit-note, signer, issuing-card, issuing-invoice, issuing-purchase, pix-request.in, pix-request.out, pix-reversal.in, pix-reversal.out, pix-claim, pix-key, pix-chargeback, pix-infraction,
//
//	Parameters (required):
//	- Url [string]: Url that will be notified when an event occurs.
//	- Subscriptions [slice of strings]: Slice of any non-empty combination of the available services. ex: []string{"contract", "credit-note", "signer", "issuing-card", "issuing-invoice", "issuing-purchase", "pix-request.in", "pix-request.out", "pix-reversal.in", "pix-reversal.out", "pix-claim", "pix-key", "pix-chargeback", "pix-infraction"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the webhook is created. ex: "5656565656565656"

type Webhook struct {
	Url           string   `json:",omitempty"`
	Subscriptions []string `json:",omitempty"`
	Id            string   `json:",omitempty"`
}

var object Webhook
var objects []Webhook
var resource = map[string]string{"name": "Webhook"}

func Create(webhook Webhook, user user.User) (Webhook, Error.StarkErrors) {
	//	Create Webhook
	//
	//	Send a single Webhook for creation at the Stark Infra API
	//
	//	Parameters (required):
	//  - webhook [Webhooks struct]: Webhook struct to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- webhook struct with updated attributes
	create, err := utils.Single(resource, webhook, user)
	unmarshalError := json.Unmarshal(create, &webhook)
	if unmarshalError != nil {
		return webhook, err
	}
	return webhook, err
}

func Get(id string, user user.User) (Webhook, Error.StarkErrors) {
	//	Retrieve a specific Webhook by its id
	//
	//	Receive a single Webhook struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- webhook struct that corresponds to the given id.
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}

func Query(params map[string]interface{}, user user.User) chan Webhook {
	//	Retrieve Webhook
	//
	//	Receive a channel of Webhook structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of Webhook structs with updated attributes
	webhooks := make(chan Webhook)
	query := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &object)
			if err != nil {
				print(err)
			}
			webhooks <- object
		}
		close(webhooks)
	}()
	return webhooks
}

func Page(params map[string]interface{}, user user.User) ([]Webhook, string, Error.StarkErrors) {
	//	Retrieve paged Webhook structs
	//
	//	Receive a slice of up to 100 Webhook structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of Webhook structs with updated attributes
	//	- cursor to retrieve the next page of Webhook structs
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &objects)
	if unmarshalError != nil {
		return objects, cursor, err
	}
	return objects, cursor, err
}

func Delete(id string, user user.User) (Webhook, Error.StarkErrors) {
	//	Delete a Webhook entity
	//
	//	Delete a Webhook entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Webhook unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- deleted Webhook struct
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &object)
	if unmarshalError != nil {
		return object, err
	}
	return object, err
}
