package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IndividualIdentity "github.com/starkinfra/sdk-go/starkinfra/individualidentity"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIndividualIdentityPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	identities, err := IndividualIdentity.Create(Example.IndividualIdentity(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, identity := range identities {
		assert.NotNil(t, identity.Id)
	}
}

func TestIndividualIdentityGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var identityList []IndividualIdentity.IndividualIdentity

	identities, errorChannel := IndividualIdentity.Query(paramsQuery, nil)
	
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case identity, ok := <-identities:
			if !ok {
				break loop
			}
			identityList = append(identityList, identity)
		}
	}

	for _, identity := range identityList {
		getIdentity, err := IndividualIdentity.Get(identity.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getIdentity.Id)
	}
	assert.Equal(t, limit, len(identityList))
}

func TestIndividualIdentityQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	identities, errorChannel := IndividualIdentity.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case identity, ok := <-identities:
			if !ok {
				break loop
			}
			assert.NotNil(t, identity.Id)
		}
	}
}

func TestIndividualIdentityPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	identities, cursor, err := IndividualIdentity.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, identity := range identities {
		assert.NotNil(t, identity.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIndividualIdentityUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []IndividualIdentity.IndividualIdentity
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	identities, errorChannel := IndividualIdentity.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case identity, ok := <-identities:
			if !ok {
				break loop
			}
			identityList = append(identityList, identity)
		}
	}


	identity, err := IndividualIdentity.Update(identityList[0].Id, "processing", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, identity.Id)
	assert.Equal(t, "processing", identity.Status)
}

func TestIndividualIdentityCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "created"

	var identityList []IndividualIdentity.IndividualIdentity

	identities, errorChannel := IndividualIdentity.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case identity, ok := <-identities:
			if !ok {
				break loop
			}
			identityList = append(identityList, identity)
		}
	}

	identity, err := IndividualIdentity.Cancel(identityList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, identity.Id)
	assert.Equal(t, "canceled", identity.Status)
}
