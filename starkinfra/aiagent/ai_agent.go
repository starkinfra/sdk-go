package aiagent

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	AiKnowledgeBase "github.com/starkinfra/sdk-go/starkinfra/aiknowledgebase"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	AiAgent struct
//
//	An AiAgent is the configuration of an assistant: the model, the instructions, the knowledge it may consult and
//	the voice it speaks with. The agent never changes during a conversation; the conversation lives in an AiChat
//	and each turn is an AiMessage.
//	When you initialize an AiAgent, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- Name [string]: Name of the agent. Between 1 and 100 characters. ex: "Support assistant"
//	- Model [string]: AI model the agent runs on. Options: "bender-1.0" for everyday conversations, "prime-1.0" for harder reasoning.
//
//	Parameters (optional):
//	- SystemPrompt [string, default ""]: Instructions that define the agent's persona, tone and domain behavior. Up to 100000 characters. The API falls back to its default assistant prompt when omitted.
//	- VoiceId [string, default ""]: Id of the AiVoice the agent speaks with. When set, every reply also carries a speech string ready to be sent to AiSpeech. The API does not check that the voice exists.
//	- KnowledgeBaseIds [slice of strings, default nil]: Ids of up to 100 AiKnowledgeBases the agent retrieves from before answering. The API does not check that they exist. An empty, non-nil slice is sent as an empty list.
//	- MetadataSchema [map[string]interface{}, default nil]: Flat map whose keys are the fields the agent must extract on every reply. Each field takes a "type" (string, integer, number, boolean or array), an optional "description" of up to 2000 characters, an optional "enum" of up to 20 strings for string fields. The keys are yours and are sent exactly as written. ex: map[string]interface{}{"order_id": map[string]interface{}{"type": "string", "description": "Order the customer mentions"}}
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the AiAgent is created. ex: "5656565656565656"
//	- KnowledgeBases [slice of AiKnowledgeBase structs]: The knowledge bases themselves, with only their Id and Name filled. Only present when requested with expand []string{"knowledgeBases"}.
//	- Created [time.Time]: Creation datetime for the AiAgent. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the AiAgent. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiAgent struct {
	Name             string                            `json:",omitempty"`
	Model            string                            `json:",omitempty"`
	SystemPrompt     string                            `json:",omitempty"`
	VoiceId          string                            `json:",omitempty"`
	KnowledgeBaseIds []string                          `json:",omitempty"`
	MetadataSchema   map[string]interface{}            `json:",omitempty"`
	Id               string                            `json:",omitempty"`
	KnowledgeBases   []AiKnowledgeBase.AiKnowledgeBase `json:",omitempty"`
	Created          *time.Time                        `json:",omitempty"`
	Updated          *time.Time                        `json:",omitempty"`
}

var resource = map[string]string{"name": "AiAgent"}

func Create(agent AiAgent, user user.User) (AiAgent, Error.StarkErrors) {
	//	Create an AiAgent
	//
	//	Send an AiAgent struct for creation at the Stark Infra API. Only the parameters you filled are sent.
	//	The MetadataSchema keys are yours and reach the API exactly as written. An empty SystemPrompt or VoiceId is
	//	not sent, which the API treats like an empty value on create; Update does send "", [] and {}.
	//
	//	Parameters (required):
	//	- agent [AiAgent struct]: AiAgent struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiAgent struct with updated attributes
	var created AiAgent
	payload := map[string]interface{}{"name": agent.Name, "model": agent.Model}
	if agent.SystemPrompt != "" {
		payload["systemPrompt"] = agent.SystemPrompt
	}
	if agent.VoiceId != "" {
		payload["voiceId"] = agent.VoiceId
	}
	if agent.KnowledgeBaseIds != nil {
		payload["knowledgeBaseIds"] = agent.KnowledgeBaseIds
	}
	if agent.MetadataSchema != nil {
		payload["metadataSchema"] = agent.MetadataSchema
	}
	response, err := utils.PostRaw(api.Endpoint(resource), payload, user, nil, "", true)
	if err.Errors != nil {
		return created, err
	}
	var answer struct{ Agent AiAgent }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return created, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Agent, err
}

func Get(id string, params map[string]interface{}, user user.User) (AiAgent, Error.StarkErrors) {
	//	Retrieve a specific AiAgent
	//
	//	Receive a single AiAgent struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the request
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "knowledgeBases".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiAgent struct that corresponds to the given id
	var agent AiAgent
	get, err := utils.Get(resource, id, params, user)
	unmarshalError := json.Unmarshal(get, &agent)
	if unmarshalError != nil {
		return agent, err
	}
	return agent, err
}

func Query(params map[string]interface{}, user user.User) (chan AiAgent, chan Error.StarkErrors) {
	//	Retrieve AiAgents
	//
	//	Receive a channel of AiAgent structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "knowledgeBases".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiAgent structs with updated attributes
	agents := make(chan AiAgent)
	agentsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var agent AiAgent
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &agent)
			if err != nil {
				agentsError <- Error.UnknownError(err.Error())
				continue
			}
			agents <- agent
		}
		for err := range errorChannel {
			agentsError <- err
		}
		close(agents)
		close(agentsError)
	}()
	return agents, agentsError
}

func Page(params map[string]interface{}, user user.User) ([]AiAgent, string, Error.StarkErrors) {
	//	Retrieve paged AiAgents
	//
	//	Receive a slice of up to 100 AiAgent structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "knowledgeBases".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiAgent structs with updated attributes
	//	- cursor to retrieve the next page of AiAgent structs, empty on the last page
	var agents []AiAgent
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &agents)
	if unmarshalError != nil {
		return agents, cursor, err
	}
	return agents, cursor, err
}

func Update(id string, patchData map[string]interface{}, user user.User) (AiAgent, Error.StarkErrors) {
	//	Update AiAgent entity
	//
	//	Update an AiAgent's parameters by passing its id. All six parameters are named in the request: the ones you
	//	leave out (or give as nil) go as null and the API keeps what you do not send. Clear a field on purpose with
	//	"" for the strings, an empty slice for knowledgeBaseIds or an empty map for metadataSchema.
	//	The MetadataSchema keys are yours and reach the API exactly as written.
	//
	//	Parameters (required):
	//	- id [string]: AiAgent unique id. ex: "5656565656565656"
	//	- patchData [map[string]interface{}]: map containing the attributes to be updated.
	//		Parameters (optional):
	//		- name [string]: New name for the agent. Between 1 and 100 characters.
	//		- model [string]: New AI model. Options: "bender-1.0", "prime-1.0"
	//		- systemPrompt [string]: New instructions for the agent. Up to 100000 characters. "" removes them.
	//		- voiceId [string]: New AiVoice id. "" makes the agent text-only.
	//		- knowledgeBaseIds [slice of strings]: The AiKnowledgeBase ids the agent should end up with. Replaces the current list. An empty slice detaches every knowledge base.
	//		- metadataSchema [map[string]interface{}]: New schema of the structured data the agent must extract. An empty map removes it.
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Target AiAgent with updated attributes
	var agent AiAgent
	payload := map[string]interface{}{}
	for _, key := range []string{"name", "model", "systemPrompt", "voiceId", "knowledgeBaseIds", "metadataSchema"} {
		payload[key] = patchData[key]
	}
	response, err := utils.PatchRaw(api.Endpoint(resource)+"/"+id, payload, user, nil, "", true)
	if err.Errors != nil {
		return agent, err
	}
	var answer struct{ Agent AiAgent }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return agent, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Agent, err
}

func Delete(ids []string, user user.User) ([]AiAgent, Error.StarkErrors) {
	//	Delete AiAgents
	//
	//	Delete up to 100 AiAgents at once.
	//
	//	Parameters (required):
	//	- ids [slice of strings]: Ids of the AiAgents to be deleted. Up to 100 ids. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of deleted AiAgent structs
	var deleted []AiAgent
	response, err := utils.DeleteQuery(api.Endpoint(resource), map[string]interface{}{"ids": ids}, user, "", true)
	if err.Errors != nil {
		return deleted, err
	}
	var answer struct{ Agents []AiAgent }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return deleted, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Agents, err
}
