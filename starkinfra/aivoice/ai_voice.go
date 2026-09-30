package aivoice

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/core-go/starkcore/utils/api"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	AiVoice struct
//
//	An AiVoice is a voice cloned from a recording you upload. Once cloned, it can read any text out loud through
//	an AiSpeech, and it can be attached to an AiAgent so every reply carries a speech ready to be synthesized.
//	Cloning is asynchronous: the voice is created in "processing" status and moves to "success" when it is ready
//	to speak, or to "failed" when the recording could not be cloned.
//	When you initialize an AiVoice, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- Audio [string]: Base64-encoded recording of the speaker. MP3, WAV, OGG, FLAC and WebM are accepted. Up to 10000000 characters.
//
//	Parameters (optional):
//	- Name [string, default ""]: Name of the voice. Up to 100 characters. Defaults to the voice's own id. ex: "Helena"
//	- Description [string, default ""]: Free-text description of the voice. Up to 1000 characters.
//	- Language [string, default ""]: Language the voice speaks. Options: "portuguese", "english". The API defaults to "portuguese".
//	- Gender [string, default ""]: Gender of the voice. Options: "male", "female", "neutral"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the AiVoice is created. This is the VoiceId you send to other AI resources. ex: "5656565656565656"
//	- Status [string]: Current status of the voice. Options: "processing", "success", "failed". Only a voice in "success" can speak.
//	- Errors [slice of strings]: Reasons the cloning failed. Empty while the voice is healthy.
//	- Created [time.Time]: Creation datetime for the AiVoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the AiVoice. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiVoice struct {
	Audio       string     `json:",omitempty"`
	Name        string     `json:",omitempty"`
	Description string     `json:",omitempty"`
	Language    string     `json:",omitempty"`
	Gender      string     `json:",omitempty"`
	Id          string     `json:",omitempty"`
	Status      string     `json:",omitempty"`
	Errors      []string   `json:",omitempty"`
	Created     *time.Time `json:",omitempty"`
	Updated     *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "AiVoice"}

func Create(voice AiVoice, user user.User) (AiVoice, Error.StarkErrors) {
	//	Create an AiVoice
	//
	//	Send an AiVoice struct for creation at the Stark Infra API and start cloning it.
	//	The call returns immediately with the voice in "processing" status.
	//
	//	Parameters (required):
	//	- voice [AiVoice struct]: AiVoice struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiVoice struct with updated attributes
	var created AiVoice
	create, err := utils.Single(resource, voice, user)
	unmarshalError := json.Unmarshal(create, &created)
	if unmarshalError != nil {
		return created, err
	}
	return created, err
}

func Query(params map[string]interface{}, user user.User) (chan AiVoice, chan Error.StarkErrors) {
	//	Retrieve AiVoices
	//
	//	Receive a channel of AiVoice structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiVoice structs with updated attributes
	voices := make(chan AiVoice)
	voicesError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var voice AiVoice
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &voice)
			if err != nil {
				voicesError <- Error.UnknownError(err.Error())
				continue
			}
			voices <- voice
		}
		for err := range errorChannel {
			voicesError <- err
		}
		close(voices)
		close(voicesError)
	}()
	return voices, voicesError
}

func Page(params map[string]interface{}, user user.User) ([]AiVoice, string, Error.StarkErrors) {
	//	Retrieve paged AiVoices
	//
	//	Receive a slice of up to 100 AiVoice structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiVoice structs with updated attributes
	//	- cursor to retrieve the next page of AiVoice structs, empty on the last page
	var voices []AiVoice
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &voices)
	if unmarshalError != nil {
		return voices, cursor, err
	}
	return voices, cursor, err
}

func Delete(ids []string, user user.User) ([]AiVoice, Error.StarkErrors) {
	//	Delete AiVoices
	//
	//	Delete up to 100 AiVoices at once.
	//
	//	Parameters (required):
	//	- ids [slice of strings]: Ids of the AiVoices to be deleted. Up to 100 ids. ex: []string{"5656565656565656", "4545454545454545"}
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of deleted AiVoice structs
	var deleted []AiVoice
	response, err := utils.DeleteQuery(api.Endpoint(resource), map[string]interface{}{"ids": ids}, user, "", true)
	if err.Errors != nil {
		return deleted, err
	}
	var answer struct{ Voices []AiVoice }
	unmarshalError := json.Unmarshal(response.Content, &answer)
	if unmarshalError != nil {
		return deleted, Error.UnknownError(unmarshalError.Error())
	}
	return answer.Voices, err
}
