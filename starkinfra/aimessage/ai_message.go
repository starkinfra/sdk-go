package aimessage

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	AiMessage struct
//
//	An AiMessage is a single turn of an AiChat. You post what the user said and the same call returns the user's
//	message and the agent's answer.
//	When you initialize an AiMessage, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the user's message and the agent's answer.
//
//	Parameters (required):
//	- ChatId [string]: Id of the AiChat to post to. ex: "5656565656565656"
//	- Text [string]: Content of the user's message. Between 1 and 50000 characters. ex: "What is the status of my order?"
//
//	Parameters (optional):
//	- Model [string, default ""]: AI model to use for this turn only. Options: "bender-1.0", "prime-1.0". The API defaults to the agent's own model.
//
//	Attributes (return-only):
//	- Id [string]: Unique id of the AiMessage. ex: "5656565656565656"
//	- Sender [string]: Who wrote the message. Options: "user", "system". The agent's answers are sent by "system".
//	- Speech [string]: Version of the text written to be heard rather than read, ready to be sent to AiSpeech. Only filled when the agent has a voice.
//	- Metadata [map[string]interface{}]: Structured data the agent extracted, shaped by the agent's MetadataSchema. The keys are the agent's, exactly as it declared them.
//	- ChatName [string]: Title of the chat. Only present when Create is called with expand []string{"chatName"}.
//	- Created [time.Time]: Creation datetime for the AiMessage. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiMessage struct {
	ChatId   string                 `json:",omitempty"`
	Text     string                 `json:",omitempty"`
	Model    string                 `json:",omitempty"`
	Id       string                 `json:",omitempty"`
	Sender   string                 `json:",omitempty"`
	Speech   string                 `json:",omitempty"`
	Metadata map[string]interface{} `json:",omitempty"`
	ChatName string                 `json:",omitempty"`
	Created  *time.Time             `json:",omitempty"`
}

var resource = map[string]string{"name": "AiMessage"}

func Create(message AiMessage, expand []string, user user.User) ([]AiMessage, Error.StarkErrors) {
	//	Create an AiMessage
	//
	//	Post the user's message to an AiChat. The call waits for the agent, which takes a few seconds, and returns both messages.
	//	The API answers a list under "messages" plus the chat name, and takes expand in the query string, so this
	//	function builds the request and reads the response itself.
	//
	//	Parameters (required):
	//	- message [AiMessage struct]: AiMessage struct with ChatId and Text, to be created in the API
	//
	//	Parameters (optional):
	//	- expand [slice of strings, default nil]: Extra attributes to compute. Options: "chatName", which returns the chat title on every message, useful on the first turn, when the title is generated.
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice with the user's AiMessage and the agent's AiMessage
	payload := map[string]interface{}{"text": message.Text}
	if message.ChatId != "" {
		payload["chatId"] = message.ChatId
	}
	if message.Model != "" {
		payload["model"] = message.Model
	}
	var query map[string]interface{}
	if len(expand) > 0 {
		query = map[string]interface{}{"expand": expand}
	}
	response, err := utils.PostRaw(api.Endpoint(resource), payload, user, query, "", true)
	if err.Errors != nil {
		return nil, err
	}
	var answer struct {
		ChatName string
		Messages []AiMessage
	}
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return nil, Error.UnknownError(unmarshalError.Error())
	}
	for index := range answer.Messages {
		answer.Messages[index].ChatName = answer.ChatName
	}
	return answer.Messages, err
}

func Query(params map[string]interface{}, user user.User) (chan AiMessage, chan Error.StarkErrors) {
	//	Retrieve AiMessages
	//
	//	Receive a channel of AiMessage structs, newest first, following the cursor until the history ends.
	//	Without a chatId, the messages of every chat in the workspace are returned.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- chatId [string, default nil]: Id of the AiChat whose messages you want. An unknown or deleted chat is refused by the API. ex: "5656565656565656"
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiMessage structs with updated attributes
	messages := make(chan AiMessage)
	messagesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var message AiMessage
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &message)
			if err != nil {
				messagesError <- Error.UnknownError(err.Error())
				continue
			}
			messages <- message
		}
		for err := range errorChannel {
			messagesError <- err
		}
		close(messages)
		close(messagesError)
	}()
	return messages, messagesError
}

func Page(params map[string]interface{}, user user.User) ([]AiMessage, string, Error.StarkErrors) {
	//	Retrieve paged AiMessages
	//
	//	Receive a slice of up to 100 AiMessage structs and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//	Without a chatId, the messages of every chat in the workspace are returned.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- chatId [string, default nil]: Id of the AiChat whose messages you want. ex: "5656565656565656"
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiMessage structs with updated attributes
	//	- cursor to retrieve the next page of AiMessage structs, empty on the last page
	var messages []AiMessage
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &messages)
	if unmarshalError != nil {
		return messages, cursor, err
	}
	return messages, cursor, err
}
