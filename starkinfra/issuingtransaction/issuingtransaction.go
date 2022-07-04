package issuingtransaction

//	IssuingTransaction object
//	The IssuingTransaction objects created in your Workspace to represent each balance shift.
//
//	Attributes (return-only):
//	- id [string]: unique id returned when IssuingTransaction is created. ex: "5656565656565656"
//	- amount [integer]: IssuingTransaction value in cents. ex: 1234 (= R$ 12.34)
//	- balance [integer]: balance amount of the Workspace at the instant of the Transaction in cents. ex: 200 (= R$ 2.00)
//	- description [string]: IssuingTransaction description. ex: "Buying food"
//	- source [string]: source of the transaction. ex: "issuing-purchase/5656565656565656"
//	- tags [string]: list of strings inherited from the source resource. ex: ["tony", "stark"]
//	- created [datetime.datetime]: creation datetime for the IssuingTransaction. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type IssuingTransaction struct {
	Id          string
	Amount      int
	Balance     int
	Description string
	Source      string
	Tags        []string
	Created     string
}

var resource = map[string]string{"class": IssuingInvoice{}, "name": "IssuingInvoice"}

func Get() {
	//	Retrieve a specific IssuingTransaction
	//	Receive a single IssuingTransaction object previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- IssuingTransaction object with updated attributes

}

func Query() {
	//	Retrieve IssuingTransaction
	//	Receive a generator of IssuingTransaction objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- external_ids [list of strings, default []]: external IDs. ex: ["5656565656565656", "4545454545454545"]
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- status [string, default None]: filter for status of retrieved objects. ex: "approved", "canceled", "denied", "confirmed" or "voided"
	//	- ids [list of strings, default [], default None]: purchase IDs
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of IssuingTransaction objects with updated attributes

}

func Page() {
	//	Retrieve paged IssuingTransaction
	//	Receive a list of IssuingTransaction objects previously created in the Stark Infra API and the cursor to the next page.
	//
	//	Parameters (optional):
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- external_ids [list of strings, default []]: external IDs. ex: ["5656565656565656", "4545454545454545"]
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- status [string, default None]: filter for status of retrieved objects. ex: "approved", "canceled", "denied", "confirmed" or "voided"
	//	- ids [list of strings, default [], default None]: purchase IDs
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of IssuingTransaction objects with updated attributes
	//	- cursor to retrieve the next page of IssuingPurchase objects
}
