package aichat

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	AiChat struct
//
//	An AiChat is one conversation thread with an AiAgent and holds the history. Each turn is an AiMessage.
//	When you initialize an AiChat, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- AgentId [string]: Id of the AiAgent that will answer in this chat. ex: "5656565656565656"
//
//	Parameters (optional):
//	- Title [string, default ""]: Title of the conversation. Up to 100 characters. When omitted, the first message posted to the chat generates one.
//	- Tags [slice of strings, default nil]: Up to 100 strings, each up to 100 characters and stored in lowercase, to find the chat later. ex: []string{"customer-123", "whatsapp"}
//	- Context [map[string]interface{}, default nil]: Data about the person on the other side of the chat that the agent reads before every reply. Up to 16384 bytes, treated as reference data and never as instructions. The keys are yours and are sent exactly as written. ex: map[string]interface{}{"name": "Ana", "balance": 1520.33}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the AiChat is created. ex: "5656565656565656"
//	- AgentName [string]: Name of the agent. Only present when requested with expand []string{"agentName"}.
//	- Updated [time.Time]: Latest update datetime for the AiChat. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiChat struct {
	AgentId   string                 `json:",omitempty"`
	Title     string                 `json:",omitempty"`
	Tags      []string               `json:",omitempty"`
	Context   map[string]interface{} `json:",omitempty"`
	Id        string                 `json:",omitempty"`
	AgentName string                 `json:",omitempty"`
	Updated   *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "AiChat"}

func Create(chat AiChat, user user.User) (AiChat, Error.StarkErrors) {
	//	Create an AiChat
	//
	//	Send an AiChat struct for creation at the Stark Infra API. Only the parameters you filled are sent.
	//	The Context keys are yours and reach the API exactly as written. An empty Title is not sent, which the API
	//	treats like an empty value on create; Update does send "", [] and {}.
	//
	//	Parameters (required):
	//	- chat [AiChat struct]: AiChat struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiChat struct with updated attributes
	var created AiChat
	payload := map[string]interface{}{"agentId": chat.AgentId}
	if chat.Title != "" {
		payload["title"] = chat.Title
	}
	if chat.Tags != nil {
		payload["tags"] = chat.Tags
	}
	if chat.Context != nil {
		payload["context"] = chat.Context
	}
	response, err := utils.PostRaw(api.Endpoint(resource), payload, user, nil, "", true)
	if err.Errors != nil {
		return created, err
	}
	var answer struct{ Chat AiChat }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return created, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Chat, err
}

func Get(id string, params map[string]interface{}, user user.User) (AiChat, Error.StarkErrors) {
	//	Retrieve a specific AiChat
	//
	//	Receive a single AiChat struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the request
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "agentName".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiChat struct that corresponds to the given id
	var chat AiChat
	get, err := utils.Get(resource, id, params, user)
	unmarshalError := json.Unmarshal(get, &chat)
	if unmarshalError != nil {
		return chat, err
	}
	return chat, err
}

func Query(params map[string]interface{}, user user.User) (chan AiChat, chan Error.StarkErrors) {
	//	Retrieve AiChats
	//
	//	Receive a channel of AiChat structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "agentName".
	//		- tags [slice of strings, default nil]: Up to 30 tags. Retrieves the chats that have any of them. ex: []string{"customer-123"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiChat structs with updated attributes
	chats := make(chan AiChat)
	chatsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var chat AiChat
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &chat)
			if err != nil {
				chatsError <- Error.UnknownError(err.Error())
				continue
			}
			chats <- chat
		}
		for err := range errorChannel {
			chatsError <- err
		}
		close(chats)
		close(chatsError)
	}()
	return chats, chatsError
}

func Page(params map[string]interface{}, user user.User) ([]AiChat, string, Error.StarkErrors) {
	//	Retrieve paged AiChats
	//
	//	Receive a slice of up to 100 AiChat structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "agentName".
	//		- tags [slice of strings, default nil]: Up to 30 tags. Retrieves the chats that have any of them. ex: []string{"customer-123"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiChat structs with updated attributes
	//	- cursor to retrieve the next page of AiChat structs, empty on the last page
	var chats []AiChat
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &chats)
	if unmarshalError != nil {
		return chats, cursor, err
	}
	return chats, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (AiChat, Error.StarkErrors) {
	//	Update AiChat entity
	//
	//	Update an AiChat's parameters by passing its id. All four parameters are named in the request: the ones you
	//	leave out (or give as nil) go as null and the API keeps what you do not send. Clear a field on purpose with
	//	"" for the title, an empty slice for tags or an empty map for context.
	//	The context keys are yours and reach the API exactly as written.
	//
	//	Parameters (required):
	//	- id [string]: AiChat unique id. ex: "5656565656565656"
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated.
	//		Parameters (optional):
	//		- title [string]: New title for the conversation. Up to 100 characters.
	//		- agentId [string]: Id of the AiAgent that should answer from now on.
	//		- tags [slice of strings]: New slice of up to 100 strings. Replaces the current list as a whole; an empty slice removes them.
	//		- context [map[string]interface{}]: New data about the person on the other side of the chat. Replaces the current map as a whole and is used from the next message on; an empty map removes it.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Target AiChat with updated attributes
	var chat AiChat
	payload := map[string]interface{}{}
	for _, key := range []string{"title", "agentId", "tags", "context"} {
		payload[key] = patchData[key]
	}
	response, err := utils.PatchRaw(api.Endpoint(resource)+"/"+id, payload, user, nil, "", true)
	if err.Errors != nil {
		return chat, err
	}
	var answer struct{ Chat AiChat }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return chat, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Chat, err
}

func Delete(ids []string, user user.User) ([]AiChat, Error.StarkErrors) {
	//	Delete AiChats
	//
	//	Delete up to 100 AiChats at once, with their messages.
	//
	//	Parameters (required):
	//	- ids [slice of strings]: Ids of the AiChats to be deleted. Up to 100 ids. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of deleted AiChat structs
	var deleted []AiChat
	response, err := utils.DeleteQuery(api.Endpoint(resource), map[string]interface{}{"ids": ids}, user, "", true)
	if err.Errors != nil {
		return deleted, err
	}
	var answer struct{ Chats []AiChat }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return deleted, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Chats, err
}
