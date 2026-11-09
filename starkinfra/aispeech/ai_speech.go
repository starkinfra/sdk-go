package aispeech

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	AiSpeech struct
//
//	An AiSpeech is one text read out loud by an AiVoice. The speech is synthesized when it is created and comes
//	back as a base64 MP3 in the Audio attribute.
//	When you initialize an AiSpeech, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- VoiceId [string]: Id of the AiVoice that should read the text. Only a voice in "success" can speak. ex: "5656565656565656"
//	- Text [string]: Text to read out loud. Between 1 and 100000 characters. ex: "Hello, how can I help you?"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the AiSpeech is created. ex: "5656565656565656"
//	- Status [string]: Current status of the speech. Options: "processing", "success", "failed"
//	- Audio [string]: Base64-encoded MP3 of the speech. Left out of Query and Page results; Get returns it.
//	- VoiceName [string]: Name of the voice. Only present when requested with expand []string{"voiceName"}.
//	- Errors [slice of strings]: Reasons the synthesis failed. Empty when it worked.
//	- Created [time.Time]: Creation datetime for the AiSpeech. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the AiSpeech. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiSpeech struct {
	VoiceId   string     `json:",omitempty"`
	Text      string     `json:",omitempty"`
	Id        string     `json:",omitempty"`
	Status    string     `json:",omitempty"`
	Audio     string     `json:",omitempty"`
	VoiceName string     `json:",omitempty"`
	Errors    []string   `json:",omitempty"`
	Created   *time.Time `json:",omitempty"`
	Updated   *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "AiSpeech"}

func Create(speech AiSpeech, user user.User) (AiSpeech, Error.StarkErrors) {
	//	Create an AiSpeech
	//
	//	Send an AiSpeech struct for creation at the Stark Infra API. The audio is synthesized during the call.
	//
	//	Parameters (required):
	//	- speech [AiSpeech struct]: AiSpeech struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiSpeech struct with updated attributes
	var created AiSpeech
	create, err := utils.Single(resource, speech, user)
	unmarshalError := json.Unmarshal(create, &created)
	if unmarshalError != nil {
		return created, err
	}
	return created, err
}

func Get(id string, params map[string]interface{}, user user.User) (AiSpeech, Error.StarkErrors) {
	//	Retrieve a specific AiSpeech
	//
	//	Receive a single AiSpeech struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the request
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "voiceName".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiSpeech struct that corresponds to the given id
	var speech AiSpeech
	get, err := utils.Get(resource, id, params, user)
	unmarshalError := json.Unmarshal(get, &speech)
	if unmarshalError != nil {
		return speech, err
	}
	return speech, err
}

func Query(params map[string]interface{}, user user.User) (chan AiSpeech, chan Error.StarkErrors) {
	//	Retrieve AiSpeeches
	//
	//	Receive a channel of AiSpeech structs previously created in the Stark Infra API. The audio is left out of the results.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "voiceName".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiSpeech structs with updated attributes
	speeches := make(chan AiSpeech)
	speechesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var speech AiSpeech
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &speech)
			if err != nil {
				speechesError <- Error.UnknownError(err.Error())
				continue
			}
			speeches <- speech
		}
		for err := range errorChannel {
			speechesError <- err
		}
		close(speeches)
		close(speechesError)
	}()
	return speeches, speechesError
}

func Page(params map[string]interface{}, user user.User) ([]AiSpeech, string, Error.StarkErrors) {
	//	Retrieve paged AiSpeeches
	//
	//	Receive a slice of up to 100 AiSpeech structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests. The audio is left out of the results.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- expand [slice of strings, default nil]: Extra attributes to compute. Options: "voiceName".
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiSpeech structs with updated attributes
	//	- cursor to retrieve the next page of AiSpeech structs, empty on the last page
	var speeches []AiSpeech
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &speeches)
	if unmarshalError != nil {
		return speeches, cursor, err
	}
	return speeches, cursor, err
}
