package individualaccountattachment

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IndividualAccountAttachment struct
//
//	Supporting document (identity, driver's license, selfie) attached to an IndividualAccountRequest
//	for the account-approval flow. The caller uploads the raw image bytes and a MIME content type;
//	Create encodes them into a data: URL client-side before sending. This resource references its
//	parent request by AccountRequestId.
//
//	Parameters (required):
//	- Type [string]: Type of the IndividualAccountAttachment. Options: "drivers-license-front", "drivers-license-back", "identity-front", "identity-back", "selfie"
//	- Content [string]: Raw image bytes (binary) at constructor time. After Create, becomes the data:<contentType>;base64,<payload> URL sent on the wire.
//	- ContentType [string]: MIME type of Content. ex: "image/png", "image/jpeg". Input-only — consumed client-side to build the data: URL; never sent as its own wire field.
//	- AccountRequestId [string]: Id of the parent IndividualAccountRequest. ex: "5189530608992256"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: Slice of strings for tagging. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the IndividualAccountAttachment is created. ex: "5656565656565656"
//	- Status [string]: Current IndividualAccountAttachment status. ex: "created", "success", "failed", "deleted"
//	- Created [time.Time]: Creation datetime for the IndividualAccountAttachment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IndividualAccountAttachment struct {
	Type             string     `json:",omitempty"`
	Content          string     `json:",omitempty"`
	ContentType      string     `json:",omitempty"`
	AccountRequestId string     `json:",omitempty"`
	Tags             []string   `json:",omitempty"`
	Id               string     `json:",omitempty"`
	Status           string     `json:",omitempty"`
	Created          *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IndividualAccountAttachment"}

func Create(attachments []IndividualAccountAttachment, user user.User) ([]IndividualAccountAttachment, Error.StarkErrors) {
	//	Create IndividualAccountAttachments
	//
	//	Send a slice of IndividualAccountAttachment structs for creation at the Stark Infra API.
	//	Each attachment's Content + ContentType are encoded into a data:<contentType>;base64,<payload>
	//	URL client-side before serialization, and ContentType is then blanked so it is never sent as its
	//	own wire field.
	//
	//	Parameters (required):
	//	- attachments [slice of IndividualAccountAttachment structs]: slice of IndividualAccountAttachment structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualAccountAttachment structs with updated attributes
	for i := 0; i < len(attachments); i++ {
		if attachments[i].ContentType != "" {
			attachments[i].Content = fmt.Sprintf("data:%s;base64,%s", attachments[i].ContentType, attachments[i].Content)
			attachments[i].ContentType = ""
		}
	}
	create, err := utils.Multi(resource, attachments, nil, user)
	unmarshalError := json.Unmarshal(create, &attachments)
	if unmarshalError != nil {
		return attachments, err
	}
	return attachments, err
}

func Get(id string, user user.User) (IndividualAccountAttachment, Error.StarkErrors) {
	//	Retrieve a specific IndividualAccountAttachment by its id
	//
	//	Receive a single IndividualAccountAttachment struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- individualAccountAttachment struct that corresponds to the given id.
	var individualAccountAttachment IndividualAccountAttachment
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &individualAccountAttachment)
	if unmarshalError != nil {
		return individualAccountAttachment, err
	}
	return individualAccountAttachment, err
}

func Query(params map[string]interface{}, user user.User) (chan IndividualAccountAttachment, chan Error.StarkErrors) {
	//	Retrieve IndividualAccountAttachments
	//
	//	Receive a channel of IndividualAccountAttachment structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "success", "failed", "deleted"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"employees", "monthly"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IndividualAccountAttachment structs with updated attributes
	var individualAccountAttachment IndividualAccountAttachment
	attachments := make(chan IndividualAccountAttachment)
	attachmentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &individualAccountAttachment)
			if err != nil {
				attachmentsError <- Error.UnknownError(err.Error())
				continue
			}
			attachments <- individualAccountAttachment
		}
		for err := range errorChannel {
			attachmentsError <- err
		}
		close(attachments)
		close(attachmentsError)
	}()
	return attachments, attachmentsError
}

func Page(params map[string]interface{}, user user.User) ([]IndividualAccountAttachment, string, Error.StarkErrors) {
	//	Retrieve paged IndividualAccountAttachment structs
	//
	//	Receive a slice of up to 100 IndividualAccountAttachment structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "success", "failed", "deleted"}
	//		- tags [slice of strings, default nil]: Tags to filter retrieved structs. ex: []string{"employees", "monthly"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualAccountAttachment structs with updated attributes
	//	- cursor to retrieve the next page of IndividualAccountAttachment structs
	var individualAccountAttachments []IndividualAccountAttachment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &individualAccountAttachments)
	if unmarshalError != nil {
		return individualAccountAttachments, cursor, err
	}
	return individualAccountAttachments, cursor, err
}

func Cancel(id string, user user.User) (IndividualAccountAttachment, Error.StarkErrors) {
	//	Cancel an IndividualAccountAttachment entity
	//
	//	Cancel an IndividualAccountAttachment entity previously created in the Stark Infra API.
	//	Maps to DELETE /individual-account-attachment/{id} and returns the deleted resource
	//	(with Status = "deleted").
	//
	//	Parameters (required):
	//	- id [string]: IndividualAccountAttachment unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- deleted IndividualAccountAttachment struct
	var individualAccountAttachment IndividualAccountAttachment
	deleted, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(deleted, &individualAccountAttachment)
	if unmarshalError != nil {
		return individualAccountAttachment, err
	}
	return individualAccountAttachment, err
}
