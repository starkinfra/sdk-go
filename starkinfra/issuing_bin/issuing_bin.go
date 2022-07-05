package issuing_bin

//	IssuingBin struct
//	The IssuingBin object displays information of registered BINs to your Workspace.
//	They represent a group of cards that begin with the same numbers (BIN) and offer the same product to end customers.
//
//	Attributes (return-only):
//	- id [string]: unique BIN number registered within the card network. ex: "53810200"
//	- network [string]: card network flag. ex: "mastercard"
//	- settlement [string]: settlement type. ex: "credit"
//	- category [string]: purchase category. ex: "prepaid"
//	- client [string]: client type. ex: "business"
//	- created [datetime.datetime]: creation datetime for the Bin. ex: datetime.datetime(2020, 3, 10, 10, 30, 0, 0)

type IssuingBin struct {
	Id         string
	Network    string
	Settlement string
	Category   string
	Client     string
	Created    string
}

var resource = map[string]string{"class": IssuingBin{}, "name": "IssuingBin"}

func Query() {
	//	Retrieve IssuingBins
	//	Receive a generator of IssuingBin objects previously registered in the Stark Infra API
	//
	//	Parameters (optional):
	//	- limit [integer, default None]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- generator of IssuingBin objects with updated attributes

}

func Page() {
	//	Retrieve paged IssuingBins
	//	Receive a list of up to 100 IssuingBin objects previously registered in the Stark Infra API and the cursor to the next page.
	//
	//	Parameters (optional):
	//	- cursor [string, default None]: cursor returned on the previous page function call
	//	- limit [integer, default 100]: maximum number of objects to be retrieved. Unlimited if None. ex: 35
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- list of IssuingBin objects with updated attributes
	//	- cursor to retrieve the next page of IssuingBin objects

}
