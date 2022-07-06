package pix_reversal

//	PixReversal object
//	PixReversals are instant payments used to revert PixRequests. You can only
//	revert inbound PixRequests.
//	When you initialize a PixReversal, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the list of created objects.
//
//	Parameters (required):
//	- amount [integer]: amount in cents to be reversed from the PixRequest. ex: 1234 (= R$ 12.34)
//	- external_id [string]: string that must be unique among all your PixReversals. Duplicated external IDs will cause failures. By default, this parameter will block any PixReversal that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//	- end_to_end_id [string]: central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
//	- reason [string]: reason why the PixRequest is being reversed. Options are "bankError", "fraud", "chashierError", "customerRequest"
//
//	Parameters (optional):
//	- tags [list of strings, default None]: list of strings for reference when searching for PixReversals. ex: ["employees", "monthly"]
//
//	Attributes (return-only):
//	- id [string]: unique id returned when the PixReversal is created. ex: "5656565656565656".
//	- return_id [string]: central bank's unique reversal transaction ID. ex: "D20018183202202030109X3OoBHG74wo".
//	- bank_code [string]: code of the bank institution in Brazil. ex: "20018183"
//	- fee [string]: fee charged by this PixReversal. ex: 200 (= R$ 2.00)
//	- status [string]: current PixReversal status. ex: "created", "processing", "success", "failed"
//	- flow [string]: direction of money flow. ex: "in" or "out"
//	- created [datetime.datetime]: creation datetime for the PixReversal. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: latest update datetime for the PixReversal. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type PixReversal struct {
	Amount     int      `json:"amount"`
	ExternalId string   `json:"externalId"`
	EndToEndId string   `json:"endToEndId"`
	Reason     string   `json:"reason"`
	Tags       []string `json:"tags"`
	Id         string   `json:"id"`
	ReturnId   string   `json:"returnId"`
	BankCode   string   `json:"bankCode"`
	Fee        string   `json:"fee"`
	Status     string   `json:"status"`
	Flow       string   `json:"flow"`
	Created    string   `json:"created"`
	Updated    string   `json:"updated"`
}

var resource = map[string]string{"class": PixReversal{}, "name": "PixReversal"}

func Create() {
	//	Create PixReversals
	//	Send a list of PixReversal objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- reversals [list of PixReversal objects]: list of PixReversal objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixReversal objects with updated attributes
}

func Get() {
	//	Retrieve a specific PixReversal
	//	Receive a single PixReversal object previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixReversal object with updated attributes
}

func Query() {
	//	Retrieve PixReversals
	//	Receive a generator of PixReversal objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "processing", "success", "failed"]
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- return_ids [list of strings, default None]: central bank's unique reversal transaction IDs. ex: ["D20018183202202030109X3OoBHG74wo", "D20018183202202030109X3OoBHG72rd"].
	//	- external_ids [list of strings, default None]: url safe strings that must be unique among all your PixReversals. Duplicated external IDs will cause failures. By default, this parameter will block any PixReversal that repeats amount and receiver information on the same date. ex: ["my-internal-id-123456", "my-internal-id-654321"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixReversal objects with updated attributes
}

func Page() {
	//	Retrieve paged PixReversals
	//	Receive a list of up to 100 PixReversal objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your reversals.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Max = 100. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "processing", "success", "failed"]
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- return_ids [list of strings, default None]: central bank's unique reversal transaction ID. ex: ["D20018183202202030109X3OoBHG74wo", "D20018183202202030109X3OoBHG72rd"].
	//	- external_ids [list of strings, default None]: url safe string that must be unique among all your PixReversals. Duplicated external IDs will cause failures. By default, this parameter will block any PixReversal that repeats amount and receiver information on the same date. ex: ["my-internal-id-123456", "my-internal-id-654321"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixReversal objects with updated attributes
	//	- cursor to retrieve the next page of PixReversal objects
}

func Parse() {
	//	Create single verified PixReversal object from a content string
	//	Create a single PixReversal object from a content string received from a handler listening at the reversal url.
	//	If the provided digital signature does not check out with the StarkInfra public key, a
	//	starkinfra.error.InvalidSignatureError will be raised.
	//
	//	Parameters (required):
	//	- content [string]: response content from request received at user endpoint (not parsed)
	//	- signature [string]: base-64 digital signature received at response header "Digital-Signature"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- Parsed PixReversal object
}

func Response() {
	//	Helps you respond PixReversal requests
	//
	//	Parameters (required):
	//	- status [string]: response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: denial reason. Options: "invalidAccountNumber", "blockedAccount", "accountClosed", "invalidAccountType", "invalidTransactionType", "taxIdMismatch", "invalidTaxId", "orderRejected", "reversalTimeExpired", "settlementFailed"
	//
	//	Return:
	//	- Dumped JSON string that must be returned to us on the PixReversal requests
}
