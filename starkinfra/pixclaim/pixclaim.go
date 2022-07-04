package pixclaim

//	PixClaim object
//	PixClaims intend to transfer a PixKey from one account to another.
//	When you initialize a PixClaim, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the created object.
//
//	Parameters (required):
//	- account_created [datetime.date, datetime.datetime or string]: opening Date or DateTime for the account claiming the PixKey. ex: "2022-01-01".
//	- account_number [string]: number of the account claiming the PixKey. ex: "76543".
//	- account_type [string]: type of the account claiming the PixKey. Options: "checking", "savings", "salary" or "payment".
//	- branch_code [string]: branch code of the account claiming the PixKey. ex: 1234".
//	- name [string]: holder's name of the account claiming the PixKey. ex: "Jamie Lannister".
//	- tax_id [string]: holder's taxId of the account claiming the PixKey (CPF/CNPJ). ex: "012.345.678-90".
//	- key_id [string]: id of the registered Pix Key to be claimed. Allowed keyTypes are CPF, CNPJ, phone number or email. ex: "+5511989898989".
//
//	Attributes (return-only):
//	- id [string]: unique id returned when the PixClaim is created. ex: "5656565656565656"
//	- status [string]: current PixClaim status. Options: "created", "failed", "delivered", "confirmed", "success", "canceled"
//	- type [string]: type of Pix Claim. Options: "ownership", "portability".
//	- key_type [string]: keyType of the claimed PixKey. Options: "CPF", "CNPJ", "phone" or "email"
//	- agent [string]: Options: "claimer" if you requested the PixClaim or "claimed" if you received a PixClaim request.
//	- bank_code [string]: bank_code of the account linked to the PixKey being claimed. ex: "20018183".
//	- claimed_bank_code [string]: bank_code of the account donating the PixKey. ex: "20018183".
//	- created [datetime.datetime]: creation datetime for the PixClaim. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)
//	- updated [datetime.datetime]: update datetime for the PixClaim. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type PixClaim struct {
	AccountCreated  string
	AccountNumber   string
	AccountType     string
	BranchCode      string
	Name            string
	TaxId           string
	KeyId           string
	Id              string
	Status          string
	Type            string
	KeyType         string
	Agente          string
	BankCode        string
	ClaimedBankCode string
	Created         string
	Updates         string
}

var resource = map[string]string{"class": PixClaim{}, "name": "PixClaim"}

func Create() {
	//	Create a PixClaim object
	//	Create a PixClaim to request the transfer of a PixKey to an account
	//	hosted at other Pix participants in the Stark Infra API.
	//
	//	Parameters (required):
	//	- claim [PixClaim object]: PixClaim object to be created in the API.
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixClaim object with updated attributes.
}

func Get() {
	//	Retrieve a PixClaim object
	//	Retrieve a PixClaim object linked to your Workspace in the Stark Infra API by its id.
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixClaim object that corresponds to the given id.
}

func Query() {
	//	Retrieve PixClaims
	//	Receive a generator of PixClaims objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "failed", "delivered", "confirmed", "success", "canceled"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- type [strings, default None]: filter for the type of retrieved PixClaims. Options: "ownership" or "portability".
	//	- agent [string, default None]: filter for the agent of retrieved PixClaims. Options: "claimer" or "claimed".
	//	- key_type [string, default None]: filter for the PixKey type of retrieved PixClaims. Options: "cpf", "cnpj", "phone", "email" and "evp",
	//	- key_id [string, default None]: filter PixClaims linked to a specific PixKey id. Example: "+5511989898989".
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixClaim objects with updated attributes
}

func Page() {
	//	Retrieve paged PixClaims
	//	Receive a list of up to 100 PixClaims objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call.
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Max = 100. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- status [list of strings, default None]: filter for status of retrieved objects. ex: ["created", "failed", "delivered", "confirmed", "success", "canceled"]
	//	- ids [list of strings, default None]: list of ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- type [strings, default None]: filter for the type of retrieved PixClaims. Options: "ownership" or "portability".
	//	- agent [string, default None]: filter for the agent of retrieved PixClaims. Options: "claimer" or "claimed".
	//	- key_type [string, default None]: filter for the PixKey type of retrieved PixClaims. Options: "cpf", "cnpj", "phone", "email" and "evp",
	//	- key_id [string, default None]: filter PixClaims linked to a specific PixKey id. Example: "+5511989898989".
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixClaim objects with updated attributes and cursor to retrieve the next page of PixClaim objects
}

func Update() {
	//	Update PixClaim entity
	//	Update a PixClaim parameters by passing id.
	//
	//	Parameters (required):
	//	- id [string]: PixClaim id. ex: '5656565656565656'
	//	- status [string]: patched status for Pix Claim. Options: "confirmed" and "canceled"
	//
	//	Parameters (optional):
	//	- reason [string, default: "userRequested"]: reason why the PixClaim is being patched. Options: "fraud", "userRequested", "accountClosure".
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixClaim with updated attributes
}
