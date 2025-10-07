package individualdocument

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IndividualDocument struct
//
//	Parameters (required):
//	- Type [string]: type of the IndividualDocument. Options: "drivers-license-front", "drivers-license-back", "identity-front", "identity-back" or "selfie"
//  - Content [string]: Base64 data url of the picture. ex: data:image/jpeg;base64,/9j/4AAQSkZJRgABAQAASABIAAD...
//  - ContentType [string]: content MIME type. This parameter is required as input only. ex: "image/png" or "image/jpeg"
//  - IdentityId [string]: Unique id of IndividualIdentity. ex: "5656565656565656"
//
//	Parameters (optional):
//	- Tags [slice of strings, default nil]: slice of strings for reference when searching for IndividualDocuments. ex: []string{"employees", "monthly"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the document is created. ex: "5656565656565656"
//	- Status [string]: current status of the IndividualDocument. Options: "created", "canceled", "processing", "failed", "success"
//	- Created [time.Time]: creation datetime for the IndividualDocument. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type IndividualDocument struct {
	Type        string     `json:",omitempty"`
	Content     string     `json:",omitempty"`
	ContentType string     `json:",omitempty"`
	IdentityId  string     `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Id          string     `json:",omitempty"`
	Status      string     `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IndividualDocument"}

func Create(documents []IndividualDocument, user user.User) ([]IndividualDocument, Error.StarkErrors) {
	//	Create IndividualDocuments
	//
	//	Send a slice of IndividualDocument objects for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- documents [slice of IndividualDocument structs]: slice of IndividualDocument objects to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualDocument struct with updated attributes
	for i := 0; i < len(documents); i++ {
		if documents[i].ContentType != "" {
			documents[i].Content = fmt.Sprintf("data:%s;base64,%s", documents[i].ContentType, documents[i].Content)
			documents[i].ContentType = ""
		}
	}
	create, err := utils.Multi(resource, documents, nil, user)
	unmarshalError := json.Unmarshal(create, &documents)
	if unmarshalError != nil {
		return documents, err
	}
	return documents, err
}

func Get(id string, user user.User) (IndividualDocument, Error.StarkErrors) {
	//	Retrieve a specific IndividualDocument by its id
	//
	//	Receive a single IndividualDocument struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- individualDocument struct that corresponds to the given id.
	var individualDocument IndividualDocument
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &individualDocument)
	if unmarshalError != nil {
		return individualDocument, err
	}
	return individualDocument, err
}

func Query(params map[string]interface{}, user user.User) (chan IndividualDocument, chan Error.StarkErrors) {
	//	Retrieve IndividualDocuments
	//
	//	Receive a channel of IndividualDocument structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: maximum number of objects to be retrieved. Unlimited if nil. ex: 35
	//  	- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//  	- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//  	- status [slice of strings, default nil]: filter for status of retrieved objects. Options: "created", "canceled", "processing", "failed" and "success"
	//  	- tags [slice of strings, default nil]: tags to filter retrieved objects. ex: []string{"tony", "stark"}
	//  	- ids [slice of strings, default nil]: slice of ids to filter retrieved objects. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IndividualDocument structs with updated attributes
	var individualDocument IndividualDocument
	documents := make(chan IndividualDocument)
	documentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &individualDocument)
			if err != nil {
				documentsError <- Error.UnknownError(err.Error())
				continue
			}
			documents <- individualDocument
		}
		for err := range errorChannel {
			documentsError <- err
		}
		close(documents)
		close(documentsError)
	}()
	return documents, documentsError
}

func Page(params map[string]interface{}, user user.User) ([]IndividualDocument, string, Error.StarkErrors) {
	//	Retrieve paged IndividualDocument structs
	//
	//	Receive a slice of up to 100 IndividualDocument structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default nil]: maximum number of individualDocuments to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date.  ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: filter for status of retrieved individualDocuments. Options: ["created", "canceled", "processing", "failed", "success"}
	//		- tags [slice of strings, default nil]: tags to filter retrieved individualDocuments. ex: []string{"tony", "stark"}
	//		- ids [slice of strings, default nil]: slice of ids to filter retrieved individualDocuments. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IndividualDocument structs with updated attributes
	//	- cursor to retrieve the next page of IndividualDocument structs
	var individualDocuments []IndividualDocument
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &individualDocuments)
	if unmarshalError != nil {
		return individualDocuments, cursor, err
	}
	return individualDocuments, cursor, err
}
