package pixrequest

//	PixRequest object
//	PixRequests are used to receive or send instant payments to accounts
//	hosted in any Pix participant.
//	When you initialize a PixRequest, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the list of created objects.
//
//	Parameters (required):
//	- amount [integer]: amount in cents to be transferred. ex: 11234 (= R$ 112.34)
//	- external_id [string]: string that must be unique among all your PixRequests. Duplicated external IDs will cause failures. By default, this parameter will block any PixRequests that repeats amount and receiver information on the same date. ex: "my-internal-id-123456"
//	- sender_name [string]: sender's full name. ex: "Edward Stark"
//	- sender_tax_id [string]: sender's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- sender_branch_code [string]: sender's bank account branch code. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- sender_account_number [string]: sender's bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- sender_account_type [string]: sender's bank account type. ex: "checking", "savings", "salary" or "payment"
//	- receiver_name [string]: receiver's full name. ex: "Edward Stark"
//	- receiver_tax_id [string]: receiver's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- receiver_bank_code [string]: receiver's bank institution code in Brazil. ex: "20018183"
//	- receiver_account_number [string]: receiver's bank account number. Use '-' before the verifier digit. ex: "876543-2"
//	- receiver_branch_code [string]: receiver's bank account branch code. Use '-' in case there is a verifier digit. ex: "1357-9"
//	- receiver_account_type [string]: receiver's bank account type. ex: "checking", "savings", "salary" or "payment"
//	- end_to_end_id [string]: central bank's unique transaction ID. ex: "E79457883202101262140HHX553UPqeq"
//
//	Parameters (optional):
//	- receiver_key_id [string, default None]: receiver's dict key. ex: "20.018.183/0001-80"
//	- description [string, default None]: optional description to override default description to be shown in the bank statement. ex: "Payment for service #1234"
//	- reconciliation_id [string, default None]: Reconciliation ID linked to this payment. ex: "b77f5236-7ab9-4487-9f95-66ee6eaf1781"
//	- initiator_tax_id [string, default None]: Payment initiator's tax id (CPF/CNPJ). ex: "01234567890" or "20.018.183/0001-80"
//	- cash_amount [integer, default None]: Amount to be withdrawal from the cashier in cents. ex: 1000 (= R$ 10.00)
//	- cashier_bank_code [string, default None]: Cashier's bank code. ex: "00000000"
//	- cashier_type [string, default None]: Cashier's type. ex: [merchant, other, participant]
//	- tags [list of strings, default None]: list of strings for reference when searching for PixRequests. ex: ["employees", "monthly"]
//	- method [string, default None]: execution  method for thr creation of the Pix. ex: "manual", "payerQrcode", "dynamicQrcode".
//
//	Attributes (return-only):
//	- id [string]: unique id returned when the PixRequest is created. ex: "5656565656565656"
//	- fee [integer]: fee charged when PixRequest is paid. ex: 200 (= R$ 2.00)
//	- status [string]: current PixRequest status. ex: "created", "processing", "success", "failed"
//	- flow [string]: direction of money flow. ex: "in" or "out"
//	- sender_bank_code [string]: sender's bank institution code in Brazil. ex: "20018183"
//	- created [datetime.datetime]: creation datetime for the PixRequest. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: latest update datetime for the PixRequest. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type PixRequest struct {
	Amount                int
	ExternalId            string
	SenderName            string
	SenderTaxId           string
	SenderBranchCode      string
	SenderAccountNumber   string
	SenderAccountType     string
	ReceiverName          string
	ReceiverTaxId         string
	ReceiverBankCode      string
	ReceiverAccountNumber string
	ReceiverBranchCode    string
	ReceiverAccountType   string
	EndToEndId            string
	ReceiverKeyId         string
	Description           string
	ReconciliationId      string
	InitiatorTaxId        string
	CashAmount            int
	CashierBankCode       string
	CashierType           string
	Tags                  []string
	Method                string
	Id                    string
	Fee                   int
	Status                string
	Flow                  string
	SenderBankCode        string
	Created               string
	Updated               string
}

var resource = map[string]string{"class": PixRequest{}, "name": "PixRequest"}

func Create() {
	//	Create PixRequests
	//	Send a list of PixRequest objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- requests [list of PixRequest objects]: list of PixRequest objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixRequest objects with updated attributes
}

func Get() {
	//	Retrieve a specific PixRequest
	//	Receive a single PixRequest object previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixRequest object with updated attributes
}

func Query() {
	//	Retrieve PixRequests
	//	Receive a generator of PixRequest objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "processing", "success", "failed"]
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- end_to_end_ids [list of strings, default None]: central bank's unique transaction IDs. ex: ["E79457883202101262140HHX553UPqeq", "E79457883202101262140HHX553UPxzx"]
	//	- external_ids [list of strings, default None]: url safe strings that must be unique among all your PixRequests. Duplicated external IDs will cause failures. By default, this parameter will block any PixRequests that repeats amount and receiver information on the same date. ex: ["my-internal-id-123456", "my-internal-id-654321"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixRequest objects with updated attributes
}

func Page() {
	//	Retrieve paged PixRequests
	//	Receive a list of up to 100 PixRequest objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Max = 100. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "processing", "success", "failed"]
	//	- tags [list of strings, default None]: tags to filter retrieved objects. ex: ["tony", "stark"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- end_to_end_ids [list of strings, default None]: central bank's unique transaction IDs. ex: ["E79457883202101262140HHX553UPqeq", "E79457883202101262140HHX553UPxzx"]
	//	- external_ids [list of strings, default None]: url safe strings that must be unique among all your PixRequests. Duplicated external IDs will cause failures. By default, this parameter will block any PixRequests that repeats amount and receiver information on the same date. ex: ["my-internal-id-123456", "my-internal-id-654321"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixRequest objects with updated attributes
	//	- cursor to retrieve the next page of PixRequest objects
}

func Parse() {
	//	Create single verified PixRequest object from a content string
	//	Create a single PixRequest object from a content string received from a handler listening at the request url.
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
	//	- Parsed PixRequest object
}

func Response() {
	//	Helps you respond PixRequests
	//
	//	Parameters (required):
	//	- status [string]: response to the authorization. ex: "approved" or "denied"
	//
	//	Parameters (conditionally required):
	//	- reason [string]: denial reason. Options: "invalidAccountNumber", "blockedAccount", "accountClosed", "invalidAccountType", "invalidTransactionType", "taxIdMismatch", "invalidTaxId", "orderRejected", "reversalTimeExpired", "settlementFailed"
	//
	//	Return:
	//	- Dumped JSON string that must be returned to us on the PixRequest
}
