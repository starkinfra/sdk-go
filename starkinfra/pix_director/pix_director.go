package pix_director

//	PixDirector object
//	Mandatory data that must be registered within the Central Bank for emergency contact purposes.
//	When you initialize a PixDirector, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the list of created objects.
//
//	Parameters (required):
//	- name [string]: name of the PixDirector. ex: "Edward Stark".
//	- tax_id [string]: tax ID (CPF) of the PixDirector. ex: "012.345.678-90"
//	- phone [string]: phone of the PixDirector. ex: "+551198989898"
//	- email [string]: email of the PixDirector. ex: "ned.stark@starkbank.com"
//	- password [string]: password of the PixDirector. ex: "12345678"
//	- team_email [string]: team email. ex: "pix.team@company.com"
//	- team_phones [list of strings]: list of phones of the team. ex: ["+5511988889999", "+5511988889998"]
//
//	Attributes (return-only):
//	- id [string]: unique id returned when the PixDirector is created. ex: "5656565656565656"
//	- status [string]: current PixDirector status. ex: "success"

type PixDirector struct {
	Name       string   `json:"name"`
	TaxId      string   `json:"taxId"`
	Phone      string   `json:"phone"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	TeamEmail  string   `json:"teamEmail"`
	TeamPhones []string `json:"teamPhones"`
	Id         string   `json:"id"`
	Status     string   `json:"status"`
}

var resource = map[string]string{"class": PixDirector{}, "name": "PixDirector"}

func Create() {
	//	Create a PixDirector Object
	//	Send a PixDirector object for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- director [list of PixDirector Object]: list of PixDirector objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project object, default None]: Organization or Project object. Not necessary if starkinfra.user was set before function call
	//
	//	Return:
	//	- PixDirector object with updated attributes
}
