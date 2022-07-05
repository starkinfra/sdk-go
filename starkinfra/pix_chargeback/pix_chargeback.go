package pix_chargeback

//	PixChargeback object
//	A Pix chargeback can be created when fraud is detected on a transaction or a system malfunction
//	results in an erroneous transaction.
//	It notifies another participant of your request to reverse the payment they have received.
//	When you initialize a PixChargeback, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the created object.
//
//	Parameters (required):
//	- amount [integer]: amount in cents to be reversed. ex: 11234 (= R$ 112.34)
//	- reference_id [string]: end_to_end_id or return_id of the transaction to be reversed. ex: "E20018183202201201450u34sDGd19lz"
//	- reason [string]: reason why the reversal was requested. Options: "fraud", "flaw", "reversalChargeback"
//
//	Parameters (optional):
//	- description [string, default None]: description for the PixChargeback.
//
//	Attributes (return-only):
//	- id [string]: unique id returned when the PixChargeback is created. ex: "5656565656565656"
//	- analysis [string]: analysis that led to the result.
//	- bacen_id [string]: central bank's unique UUID that identifies the PixChargeback.
//	- sender_bank_code [string]: bank_code of the Pix participant that created the PixChargeback. ex: "20018183"
//	- receiver_bank_code [string]: bank_code of the Pix participant that received the PixChargeback. ex: "20018183"
//	- rejection_reason [string]: reason for the rejection of the Pix chargeback. Options: "noBalance", "accountClosed", "unableToReverse"
//	- reversal_reference_id [string]: return id of the reversal transaction. ex: "D20018183202202030109X3OoBHG74wo".
//	- result [string]: result after the analysis of the PixChargeback by the receiving party. Options: "rejected", "accepted", "partiallyAccepted"
//	- status [string]: current PixChargeback status. Options: "created", "failed", "delivered", "closed", "canceled".
//	- created [datetime.datetime]: creation datetime for the PixChargeback. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: latest update datetime for the PixChargeback. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type PixChargeback struct {
	Amount              int
	ReferenceId         string
	Reason              string
	Description         string
	Id                  string
	Analysis            string
	BacenId             string
	SenderBankCode      string
	RecieverBankCode    string
	RejectionReason     string
	ReversalReferenceId string
	Result              string
	Status              string
	Created             string
	Updated             string
}

var resource = map[string]string{"class": PixChargeback{}, "name": "PixChargeback"}

func Create() {
	//	Create PixChargeback objects
	//	Create PixChargebacks in the Stark Infra API
	//
	//	Parameters (optional):
	//	- chargebacks [list of PixChargeback]: list of PixChargeback objects to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixChargeback objects with updated attributes

}

func Get() {
	//	Retrieve a PixChargeback object
	//	Retrieve the PixChargeback object linked to your Workspace in the Stark Infra API using its id.
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656".
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixChargeback object that corresponds to the given id.
}

func Query() {
	//	Retrieve PixChargebacks
	//	Receive a generator of PixChargebacks objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "failed", "delivered", "closed", "canceled"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixChargeback objects with updated attributes
}

func Page() {
	//	Retrieve PixChargebacks
	//	Receive a generator of PixChargebacks objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call.
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Max = 100. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "failed", "delivered", "closed", "canceled"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- cursor to retrieve the next page of PixChargeback objects
	//	- generator of PixChargeback objects with updated attributes
}

func Update() {
	//	Update PixChargeback entity
	//	Respond to a received PixChargeback.
	//
	//	Parameters (required):
	//	- id [string]: PixChargeback id. ex: '5656565656565656'
	//	- result [string]: result after the analysis of the PixChargeback. Options: "rejected", "accepted", "partiallyAccepted".
	//
	//	Parameters (conditionally required):
	//	- rejection_reason [string, default None]: if the PixChargeback is rejected a reason is required. Options: "noBalance", "accountClosed", "unableToReverse",
	//	- reversal_reference_id [string, default None]: return_id of the reversal transaction. ex: "D20018183202201201450u34sDGd19lz"
	//
	//	Parameters (optional):
	//	- analysis [string, default None]: description of the analysis that led to the result.
	//
	//	Return:
	//	- PixChargeback with updated attributes
}

func Cancel() {
	//	Cancel a PixChargeback entity
	//	Cancel a PixChargeback entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- canceled PixChargeback object
}
