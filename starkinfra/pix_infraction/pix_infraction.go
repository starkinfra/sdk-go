package pix_infraction

//	PixInfraction object
//	PixInfractions are used to report transactions that are suspected of
//	fraud, to request a refund or to reverse a refund.
//	When you initialize a PixInfraction, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the created object.
//
//	Parameters (required):
//	- reference_id [string]: end_to_end_id or return_id of the transaction being reported. ex: "E20018183202201201450u34sDGd19lz"
//	- type [string]: type of infraction report. Options: "fraud", "reversal", "reversalChargeback"
//
//	Parameters (optional):
//	- description [string, default None]: description for any details that can help with the infraction investigation.
//
//	Attributes (return-only):
//	- id [string]: unique id returned when the PixInfraction is created. ex: "5656565656565656"
//	- credited_bank_code [string]: bank_code of the credited Pix participant in the reported transaction. ex: "20018183"
//	- debited_bank_code [string]: bank_code of the debited Pix participant in the reported transaction. ex: "20018183"
//	- agent [string]: Options: "reporter" if you created the PixInfraction, "reported" if you received the PixInfraction.
//	- analysis [string]: analysis that led to the result.
//	- bacen_id [string]: central bank's unique UUID that identifies the infraction report.
//	- reported_by [string]: agent that reported the PixInfraction. Options: "debited", "credited".
//	- result [string]: result after the analysis of the PixInfraction by the receiving party. Options: "agreed", "disagreed"
//	- status [string]: current PixInfraction status. Options: "created", "failed", "delivered", "closed", "canceled".
//	- created [datetime.datetime]: creation datetime for the PixInfraction. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: latest update datetime for the PixInfraction. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type PixInfraction struct {
	ReferenceId      string
	Type             string
	Description      string
	Id               string
	CreditedBankCode string
	DebitedBankCode  string
	Agent            string
	Analysis         string
	BacenId          string
	ReportedBy       string
	Result           string
	Status           string
	Created          string
	Updates          string
}

var resource = map[string]string{"class": PixInfraction{}, "name": "PixInfraction"}

func Create() {
	//	Create PixInfraction objects
	//	Create PixInfractions in the Stark Infra API
	//
	//	Parameters (required):
	//	- infractions [list of PixInfractions]: list of PixInfraction objects to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixInfraction objects with updated attributes
}

func Get() {
	//	Retrieve a PixInfraction object
	//	Retrieve the PixInfraction object linked to your Workspace in the Stark Infra API using its id.
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656".
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixInfraction object that corresponds to the given id.
}

func Query() {
	//	Retrieve PixInfractions
	//	Receive a generator of PixInfractions objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "failed", "delivered", "closed", "canceled"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- type [list of strings, default None]: filter for the type of retrieved PixInfractions. Options: "fraud", "reversal", "reversalChargeback"
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixInfraction objects with updated attributes
}

func Page() {
	//	Retrieve paged PixInfractions
	//	Receive a list of up to 100 PixInfractions objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call.
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Max = 100. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "failed", "delivered", "closed", "canceled"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- type [list of strings, default None]: filter for the type of retrieved PixInfractions. Options: "fraud", "reversal", "reversalChargeback"
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixInfraction objects with updated attributes and cursor to retrieve the next page of PixInfraction objects

}

func Update() {
	//	Update PixInfraction entity
	//	Respond to a received PixInfraction.
	//
	//	Parameters (required):
	//	- id [string]: PixInfraction id. ex: '5656565656565656'
	//	- result [string]: result after the analysis of the PixInfraction. Options: "agreed", "disagreed"
	//
	//	Parameters (optional):
	//	- analysis [string, default None]: analysis that led to the result.
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixInfraction with updated attributes
}

func Cancel() {
	//	Cancel a PixInfraction entity
	//	Cancel a PixInfraction entity previously created in the Stark Infra API
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- canceled PixInfraction object
}
