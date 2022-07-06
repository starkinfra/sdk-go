package log

type Log struct {
	Id      string   `json:"id"`
	Claim   PixClaim `json:"claim"`
	Type    string   `json:"type"`
	Errors  []string `json:"errors"`
	Agent   string   `json:"agent"`
	Reason  string   `json:"reason"`
	Created string   `json:"created"`
}

var resource = map[string]string{"class": Log{}, "name": "PixClaimLog"}

func Get() {
	//	Retrieve a specific PixClaim.Log
	//	Receive a single PixClaim.Log object previously created by the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: object unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixClaim.Log object with updated attributes
}

func Query() {
	//	Retrieve PixClaim.Logs
	//	Receive a generator of PixClaim.Log objects previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- ids [list of strings, default None]: Log ids to filter PixClaim Logs. ex: ["5656565656565656"]
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- types [list of strings, default None]: filter retrieved objects by types. ex: ["created", "failed", "delivering", "delivered", "confirming", "confirmed", "success", "canceling", "canceled"]
	//	- claim_ids [list of strings, default None]: list of PixClaim ids to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of PixClaim.Log objects with updated attributes
}

func Page() {
	//	Retrieve paged PixClaim.Logs
	//	Receive a list of up to 100 PixClaim.Log objects previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your claims.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- ids [list of strings, default None]: Log ids to filter PixClaim Logs. ex: ["5656565656565656"]
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Max = 100. ex: 35
	//	- after [datetime.date or string, default None]: date filter for objects created after a specified date. ex: datetime.date(2020, 3, 10)
	//	- before [datetime.date or string, default None]: date filter for objects created before a specified date. ex: datetime.date(2020, 3, 10)
	//	- types [list of strings, default None]: filter retrieved objects by types. ex: ["created", "failed", "delivering", "delivered", "confirming", "confirmed", "success", "canceling", "canceled"]
	//	- claim_ids [list of strings, default None]: list of PixClaim IDs to filter retrieved objects. ex: ["5656565656565656", "4545454545454545"]
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of PixClaim.Log objects with updated attributes
	//	- cursor to retrieve the next page of PixClaim.Log objects
}
