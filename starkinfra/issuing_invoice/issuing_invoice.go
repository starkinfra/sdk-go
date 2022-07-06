package issuing_invoice

//	IssuingInvoice object
//	The IssuingInvoice objects created in your Workspace load your Issuing balance when paid.
//
//	Parameters (required):
//	- amount [integer]: IssuingInvoice value in cents. ex: 1234 (= R$ 12.34)
//
//	Parameters (optional):
//	- tax_id [string, default sub-issuer tax ID]: payer tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- name [string, default sub-issuer name]: payer name. ex: "Iron Bank S.A."
//	- tags [list of strings, default None]: list of strings for tagging. ex: ["travel", "food"]
//
//	Attributes (return-only):
//	- id [string]: unique id returned when IssuingInvoice is created. ex: "5656565656565656"
//	- status [string]: current IssuingInvoice status. ex: "created", "expired", "overdue", "paid"
//	- issuing_transaction_id [string]: ledger transaction ids linked to this IssuingInvoice. ex: "issuing-invoice/5656565656565656"
//	- updated [datetime.datetime]: latest update datetime for the IssuingInvoice. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- created [datetime.datetime]: creation datetime for the IssuingInvoice. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type IssuingInvoice struct {
	Amount               int      `json:"amount"`
	TaxId                string   `json:"taxId"`
	Name                 string   `json:"name"`
	Tags                 []string `json:"tags"`
	Id                   string   `json:"id"`
	Status               string   `json:"status"`
	IssuingTransactionId string   `json:"issuingTransactionId"`
	Updated              string   `json:"updated"`
	Created              string   `json:"created"`
}

var resource = map[string]string{"class": IssuingInvoice{}, "name": "IssuingInvoice"}

func Create() {
	//	Create IssuingInvoices
	//	Send an IssuingInvoice object for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- invoice [IssuingInvoice object]: IssuingInvoice object to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- IssuingInvoice object with updated attributes
}

func Get() {
	//	Retrieve a specific IssuingInvoice
	//	Receive a single IssuingInvoice object previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- IssuingInvoice object with updated attributes

}

func Query() {
	//	Retrieve IssuingInvoices
	//	Receive a generator of IssuingInvoices objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "expired", "overdue", "paid"]
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of IssuingInvoices objects with updated attributes
}

func Page() {
	//	Retrieve IssuingInvoices
	//	Receive a list of IssuingInvoices objects previously created in the Stark Infra API and the cursor to the next page.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None] date filter for objects created only after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None] date filter for objects created only before specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "expired", "overdue", "paid"]
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of IssuingInvoices objects with updated attributes
	//	- cursor to retrieve the next page of IssuingInvoices objects
}
