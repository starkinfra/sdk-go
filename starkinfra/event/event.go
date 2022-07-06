package event

var ResourceBySubscription = map[string]string{
	"pix-key":          "",
	"pix-claim":        "",
	"pix-chargeback":   "",
	"pix-infraction":   "",
	"pix-request.in":   "",
	"pix-request.out":  "",
	"pix-reversal.in":  "",
	"pix-reversal.out": "",
	"issuing-card":     "",
	"issuing-invoice":  "",
	"issuing-purchase": "",
	"credit-note":      "",
}

//	Webhook Event object
//	An Event is the notification received from the subscription to the Webhook.
//	Events cannot be created, but may be retrieved from the Stark Infra API to
//	list all generated updates on entities.
//
//	Attributes:
//	- id [string]: unique id returned when the Event is created. ex: "5656565656565656"
//	- log [Log]: a Log object from one of the subscribed services (PixRequestLog, PixReversalLog)
//	- created [datetime.datetime]: creation datetime for the notification Event. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- is_delivered [bool]: true if the Event has been successfully delivered to the user url. ex: False
//	- subscription [string]: service that triggered this Event. ex: "pix-request.in", "pix-request.out"
//	- workspace_id [string]: ID of the Workspace that generated this Event. Mostly used when multiple Workspaces have Webhooks registered to the same endpoint. ex: "4545454545454545"

type Event struct {
	Id           string `json:"id"`
	Log          Log    `json:"log"`
	Created      string `json:"created"`
	IsDelivered  bool   `json:"isDelivered"`
	Subscription string `json:"subscription"`
	Workspace    string `json:"workspace"`
}

func Get() {
	//	Retrieve a specific notification Event
	//	Receive a single notification Event object previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- Event object with updated attributes

}

func Query() {
	//	Retrieve notification Events
	//	Receive a generator of notification Event objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- is_delivered [bool, default None]: bool to filter successfully delivered events. ex: True or False
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of Event objects with updated attributes

}

func Page() {
	//	Retrieve paged Events
	//	Receive a list of up to 100 Event objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. It must be an integer between 1 and 100. ex: 50
	//	- after [datetime.date or string, default None]: date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- is_delivered [bool, default None]: bool to filter successfully delivered events. ex: True or False
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of Event objects with updated attributes
	//	- cursor to retrieve the next page of Event objects

}

func Delete() {
	//	Delete a Webhook Event entity
	//	Delete a notification Event entity previously created in the Stark Infra API by its ID
	//
	//	Parameters (required):
	//	- id [string]: Event unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- deleted Event object

}

func Update() {
	//	Update notification Event entity
	//	Update notification Event by passing id.
	//	If is_delivered is True, the event will no longer be returned on queries with is_delivered=False.
	//
	//	Parameters (required):
	//	- id [list of strings]: Event unique ids. ex: "5656565656565656"
	//	- is_delivered [bool]: If True and event hasn't been delivered already, event will be set as delivered. ex: True
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- target Event with updated attributes

}

func Parse() {
	//	Create single notification Event from a content string
	//	Create a single Event object received from Event listening at subscribed user endpoint.
	//	If the provided digital signature does not check out with the StarkInfra public key, a
	//	starkinfra.error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: response content from request received at user endpoint (not parsed)
	//	- signature [string]: base-64 digital signature received at response header "Digital-Signature"
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- Parsed Event object

}
