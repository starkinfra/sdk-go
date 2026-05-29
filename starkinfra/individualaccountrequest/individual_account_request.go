package individualaccountrequest

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IndividualAccountRequest struct
//
//	Request to open a Stark Infra account for an individual. The caller submits the individual's
//	identifying data and income, and the API runs the approval flow asynchronously — moving the
//	request through "created" -> "processing" -> ("success" | "failed" | "canceled").
//	Supporting documents are uploaded as IndividualAccountAttachment and reference this request
//	by its id.
//
//	Parameters (required):
//	- Name [string]: Full legal name of the individual. ex: "Tony Stark"
//	- TaxId [string]: Brazilian CPF. Accepts formatted or digit-only. ex: "012.345.678-90" or "01234567890"
//	- Address [IndividualAccountRequest.Address]: Structured residential address, serialized as a nested object.
//	- Income [int]: Monthly income in cents. Must be > 0. ex: 1000000 (= R$ 10,000.00)
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the IndividualAccountRequest is created. ex: "5189530608992256"
//	- Status [string]: Current IndividualAccountRequest status. ex: "created", "processing", "success", "failed", "canceled"
//	- AccountType [string]: Type of the requested account. Always "individual" for this resource.
//	- Flags [slice of strings]: Server-side review flags. Empty unless the request triggered a manual-review condition.
//	- Created [time.Time]: Creation datetime for the IndividualAccountRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: Latest update datetime for the IndividualAccountRequest. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IndividualAccountRequest struct {
	Name        string     `json:",omitempty"`
	TaxId       string     `json:",omitempty"`
	Address     Address    `json:",omitempty"`
	Income      int        `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Id          string     `json:",omitempty"`
	Status      string     `json:",omitempty"`
	AccountType string     `json:",omitempty"`
	Flags       []string   `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
	Updated     *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IndividualAccountRequest"}

func Create(requests []IndividualAccountRequest, user user.User) ([]IndividualAccountRequest, Error.StarkErrors) {
	//	Create IndividualAccountRequests
	//
	//	Send a slice of IndividualAccountRequest structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- requests [slice of IndividualAccountRequest structs]: slice of IndividualAccountRequest structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualAccountRequest structs with updated attributes
	create, err := utils.Multi(resource, requests, nil, user)
	unmarshalError := json.Unmarshal(create, &requests)
	if unmarshalError != nil {
		return requests, err
	}
	return requests, err
}

func Get(id string, user user.User) (IndividualAccountRequest, Error.StarkErrors) {
	//	Retrieve a specific IndividualAccountRequest by its id
	//
	//	Receive a single IndividualAccountRequest struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5189530608992256"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- individualAccountRequest struct that corresponds to the given id.
	var individualAccountRequest IndividualAccountRequest
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &individualAccountRequest)
	if unmarshalError != nil {
		return individualAccountRequest, err
	}
	return individualAccountRequest, err
}

func Query(params map[string]interface{}, user user.User) (chan IndividualAccountRequest, chan Error.StarkErrors) {
	//	Retrieve IndividualAccountRequests
	//
	//	Receive a channel of IndividualAccountRequest structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed", "canceled"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"employees", "monthly"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5189530608992256", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IndividualAccountRequest structs with updated attributes
	var individualAccountRequest IndividualAccountRequest
	requests := make(chan IndividualAccountRequest)
	requestsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &individualAccountRequest)
			if err != nil {
				requestsError <- Error.UnknownError(err.Error())
				continue
			}
			requests <- individualAccountRequest
		}
		for err := range errorChannel {
			requestsError <- err
		}
		close(requests)
		close(requestsError)
	}()
	return requests, requestsError
}

func Page(params map[string]interface{}, user user.User) ([]IndividualAccountRequest, string, Error.StarkErrors) {
	//	Retrieve paged IndividualAccountRequest structs
	//
	//	Receive a slice of up to 100 IndividualAccountRequest structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "processing", "success", "failed", "canceled"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"employees", "monthly"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5189530608992256", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualAccountRequest structs with updated attributes
	//	- cursor to retrieve the next page of IndividualAccountRequest structs
	var individualAccountRequests []IndividualAccountRequest
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &individualAccountRequests)
	if unmarshalError != nil {
		return individualAccountRequests, cursor, err
	}
	return individualAccountRequests, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (IndividualAccountRequest, Error.StarkErrors) {
	//	Update IndividualAccountRequest entity
	//
	//	Update an IndividualAccountRequest by passing id.
	//
	//	Parameters (required):
	//	- id [string]: IndividualAccountRequest id. ex: "5189530608992256"
	//  - patchData [map[string]interface{}]: map containing the attributes to be updated. ex: map[string]interface{}{"status": "processing"}
	//		Parameters (optional):
	//		- name [string]: Replace the legal name.
	//		- taxId [string]: Replace the CPF.
	//		- address [map[string]interface{}]: Replace the address as a whole object (no partial address PATCH).
	//		- income [int]: Replace monthly income in cents.
	//		- status [string]: Manual state transition. ex: "processing"
	//		- tags [slice of strings]: Replace tag list.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- target IndividualAccountRequest with updated attributes
	var individualAccountRequest IndividualAccountRequest
	update, err := utils.Patch(resource, id, patchData, user)
	unmarshalError := json.Unmarshal(update, &individualAccountRequest)
	if unmarshalError != nil {
		return individualAccountRequest, err
	}
	return individualAccountRequest, err
}
