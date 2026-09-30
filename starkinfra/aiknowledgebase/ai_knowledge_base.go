package aiknowledgebase

import (
	"encoding/json"
	"fmt"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"strconv"
	"time"
)

//	AiKnowledgeBase struct
//
//	An AiKnowledgeBase turns a website into material an AiAgent can read. You give it a root URL;
//	Stark Infra crawls the page, follows its links, converts everything to Markdown and indexes it for retrieval.
//	When you initialize an AiKnowledgeBase, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- Name [string]: Name of the knowledge base. Between 1 and 100 characters. ex: "Product Documentation"
//	- RootUrl [string]: Absolute http or https URL the crawl starts from. ex: "https://docs.starkinfra.com"
//
//	Parameters (optional):
//	- IsRecursive [*bool, default nil]: Whether the crawl may follow links into other subdomains of the root URL's registered domain. The API defaults to true. ex: &isRecursive
//	- Tags [slice of strings, default nil]: Slice of up to 100 strings for reference when searching for AiKnowledgeBases. ex: []string{"support", "public"}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the AiKnowledgeBase is created. ex: "5656565656565656"
//	- Status [string]: Current status of the knowledge base. Options: "processing", "success", "failed". An agent retrieves from a base only once it reaches "success".
//	- Created [time.Time]: Creation datetime for the AiKnowledgeBase. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the AiKnowledgeBase. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiKnowledgeBase struct {
	Name        string     `json:",omitempty"`
	RootUrl     string     `json:",omitempty"`
	IsRecursive *bool      `json:",omitempty"`
	Tags        []string   `json:",omitempty"`
	Id          string     `json:",omitempty"`
	Status      string     `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
	Updated     *time.Time `json:",omitempty"`
}

//	HostPage struct
//
//	A page the crawler has seen, as listed by the Hosts function.
//
//	Attributes (return-only):
//	- OriginalUrl [string]: URL the page was crawled from. ex: "https://docs.starkinfra.com/get-started"
//	- StorageUrl [string]: URL of the page converted to Markdown. ex: "https://storage.googleapis.com/ai-knowledge/6767676767676767/get-started.md"
//	- Status [string]: Current status of the page. Options: "pending", "success", "failed"

type HostPage struct {
	OriginalUrl string `json:",omitempty"`
	StorageUrl  string `json:",omitempty"`
	Status      string `json:",omitempty"`
}

var resource = map[string]string{"name": "AiKnowledgeBase"}

const maxPageSize = 100

func Create(knowledgeBase AiKnowledgeBase, user user.User) (AiKnowledgeBase, Error.StarkErrors) {
	//	Create an AiKnowledgeBase
	//
	//	Send an AiKnowledgeBase struct for creation at the Stark Infra API and start crawling it.
	//	The API answers under "knowledgeBase" and "knowledgeBases", while the core derives the last word of the
	//	resource name ("base"), so the AiKnowledgeBase functions read the responses themselves.
	//	The call returns immediately with the knowledge base in "processing" status.
	//
	//	Parameters (required):
	//	- knowledgeBase [AiKnowledgeBase struct]: AiKnowledgeBase struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiKnowledgeBase struct with updated attributes
	var created AiKnowledgeBase
	payload := map[string]interface{}{"name": knowledgeBase.Name, "rootUrl": knowledgeBase.RootUrl}
	if knowledgeBase.IsRecursive != nil {
		payload["isRecursive"] = *knowledgeBase.IsRecursive
	}
	if knowledgeBase.Tags != nil {
		payload["tags"] = knowledgeBase.Tags
	}
	response, err := utils.PostRaw(api.Endpoint(resource), payload, user, nil, "", true)
	if err.Errors != nil {
		return created, err
	}
	var answer struct{ KnowledgeBase AiKnowledgeBase }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return created, Error.UnknownError(unmarshalError.Error())
	}
	return answer.KnowledgeBase, err
}

func Get(id string, user user.User) (AiKnowledgeBase, Error.StarkErrors) {
	//	Retrieve a specific AiKnowledgeBase
	//
	//	Receive a single AiKnowledgeBase struct previously created in the Stark Infra API by its id.
	//	This is the call to poll while the crawl runs.
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiKnowledgeBase struct that corresponds to the given id
	var knowledgeBase AiKnowledgeBase
	response, err := utils.GetRaw(api.Endpoint(resource)+"/"+id, nil, user, "", true)
	if err.Errors != nil {
		return knowledgeBase, err
	}
	var answer struct{ KnowledgeBase AiKnowledgeBase }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return knowledgeBase, Error.UnknownError(unmarshalError.Error())
	}
	return answer.KnowledgeBase, err
}

func Query(params map[string]interface{}, user user.User) (chan AiKnowledgeBase, chan Error.StarkErrors) {
	//	Retrieve AiKnowledgeBases
	//
	//	Receive a channel of AiKnowledgeBase structs previously created in the Stark Infra API, following the cursor
	//	until the list ends or the limit is reached.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- name [string, default nil]: Case-insensitive substring of the name to filter retrieved structs. ex: "docs"
	//		- status [string, default nil]: Filter for status of retrieved structs. Options: "processing", "success", "failed"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiKnowledgeBase structs with updated attributes
	knowledgeBases := make(chan AiKnowledgeBase)
	knowledgeBasesError := make(chan Error.StarkErrors)
	go func() {
		defer close(knowledgeBases)
		defer close(knowledgeBasesError)
		query := map[string]interface{}{}
		for key, value := range params {
			query[key] = value
		}
		hasLimit := params["limit"] != nil
		remaining, _ := strconv.Atoi(fmt.Sprintf("%v", params["limit"]))
		for {
			if hasLimit {
				query["limit"] = minInt(remaining, maxPageSize)
			}
			entities, cursor, err := Page(query, user)
			if err.Errors != nil {
				knowledgeBasesError <- err
				return
			}
			for _, entity := range entities {
				knowledgeBases <- entity
			}
			remaining -= len(entities)
			if cursor == "" || (hasLimit && remaining <= 0) {
				return
			}
			query["cursor"] = cursor
		}
	}()
	return knowledgeBases, knowledgeBasesError
}

func Page(params map[string]interface{}, user user.User) ([]AiKnowledgeBase, string, Error.StarkErrors) {
	//	Retrieve paged AiKnowledgeBases
	//
	//	Receive a slice of up to 100 AiKnowledgeBase structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//		- name [string, default nil]: Case-insensitive substring of the name to filter retrieved structs. ex: "docs"
	//		- status [string, default nil]: Filter for status of retrieved structs. Options: "processing", "success", "failed"
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiKnowledgeBase structs with updated attributes
	//	- cursor to retrieve the next page of AiKnowledgeBase structs, empty on the last page
	response, err := utils.GetRaw(api.Endpoint(resource), params, user, "", true)
	if err.Errors != nil {
		return nil, "", err
	}
	var answer struct {
		Cursor         string
		KnowledgeBases []AiKnowledgeBase
	}
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return nil, "", Error.UnknownError(unmarshalError.Error())
	}
	return answer.KnowledgeBases, answer.Cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (AiKnowledgeBase, Error.StarkErrors) {
	//	Update AiKnowledgeBase entity
	//
	//	Rename a knowledge base, retag it or change whether its crawl is recursive. The root URL cannot be changed.
	//	Only the parameters you give are sent.
	//
	//	Parameters (required):
	//	- id [string]: AiKnowledgeBase unique id. ex: "5656565656565656"
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated.
	//		Parameters (optional):
	//		- name [string]: New name of the knowledge base. Between 1 and 100 characters.
	//		- isRecursive [bool]: Whether the next crawl may follow links into other subdomains of the root URL's registered domain.
	//		- tags [slice of strings]: New slice of up to 100 strings. Replaces the current list.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Target AiKnowledgeBase with updated attributes
	var knowledgeBase AiKnowledgeBase
	response, err := utils.PatchRaw(api.Endpoint(resource)+"/"+id, patchData, user, nil, "", true)
	if err.Errors != nil {
		return knowledgeBase, err
	}
	var answer struct{ KnowledgeBase AiKnowledgeBase }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return knowledgeBase, Error.UnknownError(unmarshalError.Error())
	}
	return answer.KnowledgeBase, err
}

func Hosts(id string, user user.User) (map[string][]HostPage, Error.StarkErrors) {
	//	List the pages of an AiKnowledgeBase
	//
	//	Receive every page the crawler has seen, grouped by host. While a crawl is running this is the live picture,
	//	merged with the last finished one.
	//
	//	Parameters (required):
	//	- id [string]: AiKnowledgeBase unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- map of each host to its slice of AiKnowledgeBase.HostPage structs
	var hosts map[string][]HostPage
	response, err := utils.GetRaw(api.Endpoint(resource)+"/"+id+"/hosts", nil, user, "", true)
	if err.Errors != nil {
		return hosts, err
	}
	var answer struct{ Hosts map[string][]HostPage }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return hosts, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Hosts, err
}

func Delete(ids []string, user user.User) ([]AiKnowledgeBase, Error.StarkErrors) {
	//	Delete AiKnowledgeBases
	//
	//	Delete up to 100 AiKnowledgeBases at once. Agents still referencing a deleted base simply retrieve nothing from it.
	//
	//	Parameters (required):
	//	- ids [slice of strings]: Ids of the AiKnowledgeBases to be deleted. Up to 100 ids. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of deleted AiKnowledgeBase structs
	var deleted []AiKnowledgeBase
	response, err := utils.DeleteQuery(api.Endpoint(resource), map[string]interface{}{"ids": ids}, user, "", true)
	if err.Errors != nil {
		return deleted, err
	}
	var answer struct{ KnowledgeBases []AiKnowledgeBase }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return deleted, Error.UnknownError(unmarshalError.Error())
	}
	return answer.KnowledgeBases, err
}

func minInt(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
