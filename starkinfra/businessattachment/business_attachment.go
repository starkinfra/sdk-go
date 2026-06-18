package businessattachment

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	BusinessAttachment struct
//
//	A BusinessAttachment represents a file attached to a BusinessIdentity, used to validate
//	the identity of the company. You must reference the desired BusinessIdentity by its id.
//
//	When you initialize a BusinessAttachment, the entity will not be automatically
//	created in the Stark Infra API. The 'create' function sends the objects
//	to the Stark Infra API and returns the slice of created objects.
//
//	Parameters (required):
//	- Name [string]: name of the BusinessAttachment. ex: "articles-of-incorporation.pdf"
//  - Content [string]: Base64 data url of the file. ex: data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAA...
//  - BusinessIdentityId [string]: Unique id of BusinessIdentity. ex: "5656565656565656"
//
//	Parameters (optional):
//  - ContentType [string]: content MIME type. This parameter is required as input only. ex: "image/png" or "application/pdf"
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for BusinessAttachments. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the attachment is created. ex: "5656565656565656"
//	- AttachmentId [string]: Unique id of the attachment file. ex: "5656565656565656"
//	- Status [string]: current status of the BusinessAttachment. Options: "created", "canceled", "approved", "denied"
//	- Created [time.Time]: creation datetime for the BusinessAttachment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- Updated [time.Time]: latest update datetime for the BusinessAttachment. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type BusinessAttachment struct {
	Name               string     `json:",omitempty"`
	Content            string     `json:",omitempty"`
	ContentType        string     `json:",omitempty"`
	BusinessIdentityId string     `json:",omitempty"`
	Tags               []string   `json:",omitempty"`
	Id                 string     `json:",omitempty"`
	AttachmentId       string     `json:",omitempty"`
	Status             string     `json:",omitempty"`
	Created            *time.Time `json:",omitempty"`
	Updated            *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "BusinessAttachment"}

func Create(attachments []BusinessAttachment, user user.User) ([]BusinessAttachment, Error.StarkErrors) {
	//	Create BusinessAttachments
	//
	//	Send a slice of BusinessAttachment objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- attachments [slice of BusinessAttachment structs]: slice of BusinessAttachment objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of BusinessAttachment struct with updated attributes
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

func Get(id string, expand map[string]interface{}, user user.User) (BusinessAttachment, Error.StarkErrors) {
	//	Retrieve a specific BusinessAttachment by its id
	//
	//	Receive a single BusinessAttachment struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- expand [slice of strings, default nil]: fields to expand information. ex: []string{"content"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- businessAttachment struct that corresponds to the given id.
	var businessAttachment BusinessAttachment
	get, err := utils.Get(resource, id, expand, user)
	unmarshalError := json.Unmarshal(get, &businessAttachment)
	if unmarshalError != nil {
		return businessAttachment, err
	}
	return businessAttachment, err
}

func Query(params map[string]interface{}, user user.User) (chan BusinessAttachment, chan Error.StarkErrors) {
	//	Retrieve BusinessAttachments
	//
	//	Receive a channel of BusinessAttachment structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of objects to be retrieved. Unlimited if nil. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- status [slice of strings, default nil]: filter for status of retrieved objects. Options: "created", "canceled", "approved" and "denied"
	//  	- tags [slice of strings, default nil]: tags to filter retrieved objects. ex: []string{"tony", "stark"}
	//  	- ids [slice of strings, default nil]: slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of BusinessAttachment structs with updated attributes
	var businessAttachment BusinessAttachment
	attachments := make(chan BusinessAttachment)
	attachmentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &businessAttachment)
			if err != nil {
				attachmentsError <- Error.UnknownError(err.Error())
				continue
			}
			attachments <- businessAttachment
		}
		for err := range errorChannel {
			attachmentsError <- err
		}
		close(attachments)
		close(attachmentsError)
	}()
	return attachments, attachmentsError
}

func Page(params map[string]interface{}, user user.User) ([]BusinessAttachment, string, Error.StarkErrors) {
	//	Retrieve paged BusinessAttachment structs
	//
	//	Receive a slice of up to 100 BusinessAttachment structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: maximum number of businessAttachments to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved businessAttachments. Options: ["created", "canceled", "approved", "denied"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved businessAttachments. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved businessAttachments. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of BusinessAttachment structs with updated attributes
	//	- cursor to retrieve the next page of BusinessAttachment structs
	var businessAttachments []BusinessAttachment
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &businessAttachments)
	if unmarshalError != nil {
		return businessAttachments, cursor, err
	}
	return businessAttachments, cursor, err
}

func Cancel(id string, user user.User) (BusinessAttachment, Error.StarkErrors) {
	//	Cancel a BusinessAttachment entity
	//
	//	Cancel a BusinessAttachment by passing id.
	//
	//	Parameters (required):
	//	- id [string]: BusinessAttachment unique id. ex: "6306109539221504"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call.
	//
	//	Return:
	//	- canceled BusinessAttachment struct
	var businessAttachment BusinessAttachment
	cancel, err := utils.Delete(resource, id, user)
	unmarshalError := json.Unmarshal(cancel, &businessAttachment)
	if unmarshalError != nil {
		return businessAttachment, err
	}
	return businessAttachment, err
}
