package sdk

import (
	"encoding/json"
	"strconv"
	"time"
	"github.com/starkinfra/sdk-go/starkinfra"
	Request "github.com/starkinfra/sdk-go/starkinfra/request"
	Utils "github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRequestGet(t *testing.T) {

	starkinfra.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "pix-request/"
	query["limit"] = 2

	response, err := Request.Get(
		path,
		query,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("error: %s", unmarshalError.Error())
	}
	requestData, ok1 := data["requests"].([]interface{})
	if !ok1 {
        t.Errorf("Erro ao converter os tipos content")
        return
    }
	for _, request := range requestData{
		requestMap, ok2 := request.(map[string]interface{})
        if !ok2 {
            t.Errorf("Erro ao converter item de list 'requests' para map[string]interface{}")
            continue
        }
        id, ok3 := requestMap["id"].(string)
        if !ok3 {
            t.Errorf("Erro ao converter list 'id' para string")
            continue
        }
		path = "pix-request/" + id
		for k := range data {
			delete(data, k)
		}
		response, err := Request.Get(
			path,
			nil,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			t.Errorf("error: %s", unmarshalError.Error())
		}
		requestData, ok4 := data["request"].(map[string]interface{})
        if !ok4 {
            t.Errorf("Erro ao converter 'id' para string")
            continue
        }
		getId, ok5 := requestData["id"].(string)
		if !ok5 {
            t.Errorf("Erro ao converter 'id' para string")
            continue
        }
		assert.Equal(t, id, getId)
	}
}

func TestRequestPostAndDelete(t *testing.T) {
	starkinfra.User = Utils.ExampleProject
	data := map[string]interface{}{}
    
	now := time.Now()
    milliseconds := now.UnixNano() / int64(time.Millisecond)
    ext_id := strconv.FormatInt(milliseconds, 10)

	body := map[string][]map[string]interface{}{
        "holders": {
            {
				"name": "Jaime Lannister",
				"externalId": ext_id,
				"taxId": "012.345.678-90",
			},
        },
    }

	path := "issuing-holder/"

	response, err := Request.Post(
		path,
		body,
		nil,
		Utils.ExampleProject,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("error: %s", unmarshalError.Error())
	}
	holdersData, ok1 := data["holders"].([]interface{})
	if !ok1 {
        t.Errorf("Erro ao converter os tipos content")
        return
    }
	for _, holder := range holdersData{
		holderMap, ok2 := holder.(map[string]interface{})
        if !ok2 {
            t.Errorf("Erro ao converter item de list 'invoices' para map[string]interface{}")
            continue
        }
        id, ok3 := holderMap["id"].(string)
        if !ok3 {
            t.Errorf("Erro ao converter list 'id' para string")
            continue
        }
		path = "issuing-holder/" + id
		for k := range data {
			delete(data, k)
		}
		response, err := Request.Delete(
			path,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			t.Errorf("error: %s", unmarshalError.Error())
		}
		assert.NotNil(t, data)
	}
}

func TestRequestPatch(t *testing.T) {
	starkinfra.User = Utils.ExampleProject
	data := map[string]interface{}{}
	var path string
	var query = map[string]interface{}{}

	path = "issuing-holder/"
	query["limit"] = 2
	query["status"] = "active"

	response, err := Request.Get(
		path,
		query,
		nil,
	)

	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}
	unmarshalError := json.Unmarshal(response.Content, &data)
	if unmarshalError != nil {
		t.Errorf("error: %s", unmarshalError.Error())
	}
	holdersData, ok1 := data["holders"].([]interface{})
	if !ok1 {
        t.Errorf("Erro ao converter os tipos content")
        return
    }
	for _, invoice := range holdersData{
		holderMap, _ := invoice.(map[string]interface{})
        id, _ := holderMap["id"].(string)
		path = "issuing-holder/" + id

		now := time.Now()
		milliseconds := now.UnixNano() / int64(time.Millisecond)
		testAssertion := strconv.FormatInt(milliseconds, 10)

		for k := range data {
			delete(data, k)
		}
		body := map[string]interface{}{
			"tags": []string{testAssertion},
		}
		response, err := Request.Patch(
			path,
			body,
			nil,
			Utils.ExampleProject,
		)
	
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		unmarshalError := json.Unmarshal(response.Content, &data)
		if unmarshalError != nil {
			t.Errorf("error: %s", unmarshalError.Error())
		}
		holderData, _ := data["holder"].(map[string]interface{})
		tags, _ := holderData["tags"]
		assert.Equal(t, tags, []interface {}([]interface {}{testAssertion}))
	}
}
