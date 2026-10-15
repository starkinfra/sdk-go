package aitranscript

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	AiTranscript struct
//
//	An AiTranscript is the text of an audio file you upload, from any speaker, cloned or not.
//	When you initialize an AiTranscript, the entity will not be automatically
//	created in the Stark Infra API. The 'Create' function sends the struct
//	to the Stark Infra API and returns the created struct.
//
//	Parameters (required):
//	- Audio [string]: Base64-encoded audio to transcribe. Up to 10000000 characters. The format is read from the file's own header.
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the AiTranscript is created. ex: "5656565656565656"
//	- Text [string]: Transcribed text.
//	- Status [string]: Current status of the transcript. Options: "processing", "success", "failed"
//	- Errors [slice of strings]: Reasons the transcription failed. Empty when it worked.
//	- Created [time.Time]: Creation datetime for the AiTranscript. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)
//	- Updated [time.Time]: Latest update datetime for the AiTranscript. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type AiTranscript struct {
	Audio   string     `json:",omitempty"`
	Id      string     `json:",omitempty"`
	Text    string     `json:",omitempty"`
	Status  string     `json:",omitempty"`
	Errors  []string   `json:",omitempty"`
	Created *time.Time `json:",omitempty"`
	Updated *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "AiTranscript"}

func Create(transcript AiTranscript, user user.User) (AiTranscript, Error.StarkErrors) {
	//	Create an AiTranscript
	//
	//	Send an AiTranscript struct for creation at the Stark Infra API. The audio is transcribed during the call.
	//
	//	Parameters (required):
	//	- transcript [AiTranscript struct]: AiTranscript struct to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- AiTranscript struct with updated attributes
	var created AiTranscript
	create, err := utils.Single(resource, transcript, user)
	unmarshalError := json.Unmarshal(create, &created)
	if unmarshalError != nil {
		return created, err
	}
	return created, err
}

func Query(params map[string]interface{}, user user.User) (chan AiTranscript, chan Error.StarkErrors) {
	//	Retrieve AiTranscripts
	//
	//	Receive a channel of AiTranscript structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Channel of AiTranscript structs with updated attributes
	transcripts := make(chan AiTranscript)
	transcriptsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			var transcript AiTranscript
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &transcript)
			if err != nil {
				transcriptsError <- Error.UnknownError(err.Error())
				continue
			}
			transcripts <- transcript
		}
		for err := range errorChannel {
			transcriptsError <- err
		}
		close(transcripts)
		close(transcriptsError)
	}()
	return transcripts, transcriptsError
}

func Page(params map[string]interface{}, user user.User) ([]AiTranscript, string, Error.StarkErrors) {
	//	Retrieve paged AiTranscripts
	//
	//	Receive a slice of up to 100 AiTranscript structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of AiTranscript structs with updated attributes
	//	- cursor to retrieve the next page of AiTranscript structs, empty on the last page
	var transcripts []AiTranscript
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &transcripts)
	if unmarshalError != nil {
		return transcripts, cursor, err
	}
	return transcripts, cursor, err
}
