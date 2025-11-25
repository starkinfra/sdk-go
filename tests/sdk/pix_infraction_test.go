package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixInfraction "github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixInfractionPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	infractions, err := PixInfraction.Create(Example.PixInfraction(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, infraction := range infractions {
		assert.NotNil(t, infraction.Id)
	}
}

func TestPixInfractionQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	infractions, errorChannel := PixInfraction.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case infraction, ok := <-infractions:
			if !ok {
				break loop
			}
			assert.NotNil(t, infraction.Id)
		}
	}
}

func TestPixInfractionPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	infractions, cursor, err := PixInfraction.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, infraction := range infractions {
		assert.NotNil(t, infraction.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixInfractionInfoGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var infractionList []PixInfraction.PixInfraction

	infractions, errorChannel := PixInfraction.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case infraction, ok := <-infractions:
			if !ok {
				break loop
			}
			infractionList = append(infractionList, infraction)
		}
	}

	for _, infraction := range infractionList {
		getInfraction, err := PixInfraction.Get(infraction.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getInfraction.Id)
	}

	assert.Equal(t, limit, len(infractionList))
}

func TestPixInfractionInfoDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "delivered"

	var infractionList []PixInfraction.PixInfraction

	infractions, errorChannel := PixInfraction.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case infraction, ok := <-infractions:
			if !ok {
				break loop
			}
			infractionList = append(infractionList, infraction)
		}
	}

	infraction, err := PixInfraction.Cancel(infractionList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, infraction.Id)
}

func TestPixInfractionInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "delivered"
	
	var infractionList []PixInfraction.PixInfraction

	infractions, errorChannel := PixInfraction.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case infraction, ok := <-infractions:
			if !ok {
				break loop
			}
			infractionList = append(infractionList, infraction)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["result"] = "agreed"
	patchData["analysis"] = "Upon investigation fraud was confirmed."
	patchData["fraudType"] = "scam"

	infraction, err := PixInfraction.Update(infractionList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, infraction.Id)
	assert.Equal(t, infraction.Result, "agreed")
}
