package webhook

//	Webhook subscription object
//	A Webhook is used to subscribe to notification events on a user-selected endpoint.
//	Currently, available services for subscription are contract, credit-note, signer, issuing-card, issuing-invoice, issuing-purchase, pix-request.in, pix-request.out, pix-reversal.in, pix-reversal.out, pix-claim, pix-key, pix-chargeback, pix-infraction,
//
//	Parameters (required):
//	- url [string]: Url that will be notified when an event occurs.
//	- subscriptions [list of strings]: list of any non-empty combination of the available services. ex: ["contract", "credit-note", "signer", "issuing-card", "issuing-invoice", "issuing-purchase", "pix-request.in", "pix-request.out", "pix-reversal.in", "pix-reversal.out", "pix-claim", "pix-key", "pix-chargeback", "pix-infraction"]
//
//	Attributes:
//	- id [string]: unique id returned when the webhook is created. ex: "5656565656565656"

type Webhook struct {
	Url          string   `json:"url"`
	Subscription []string `json:"subscription"`
	Id           string   `json:"id"`
}

var resource = map[string]string{"class": Webhook{}, "name": "Webhook"}

func Create() {
	//	Create Webhook subscription
	//	Send a single Webhook subscription for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- url [string]: url to which notification events will be sent to. ex: "https://webhook.site/60e9c18e-4b5c-4369-bda1-ab5fcd8e1b29"
	//	- subscriptions [list of strings]: list of any non-empty combination of the available services. ex: ["contract", "credit-note", "signer"]
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- Webhook object with updated attributes

}

func Get() {
	//	Retrieve a specific Webhook subscription
	//	Receive a single Webhook subscription object previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- Webhook object with updated attributes
}

func Query() {
	//	Retrieve Webhook subcriptions
	//	Receive a generator of Webhook subcription objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of Webhook objects with updated attributes
}

func Page() {
	//	Retrieve paged Webhooks
	//	Receive a list of up to 100 Webhook objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. It must be an integer between 1 and 100. ex: 50
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of Webhook objects with updated attributes
	//	- cursor to retrieve the next page of Webhook objects
}

func Delete() {
	//	Delete a Webhook subscription entity
	//	Delete a Webhook subscription entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: Webhook unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- deleted Webhook object
}
