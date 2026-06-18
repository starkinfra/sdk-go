package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	BusinessIdentity "github.com/starkinfra/sdk-go/starkinfra/businessidentity"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBusinessIdentityPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	identities, err := BusinessIdentity.Create(Example.BusinessIdentity(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, identity := range identities {
		assert.NotNil(t, identity.Id)
	}
}

func TestBusinessIdentityQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	identities, errorChannel := BusinessIdentity.Query(params, nil)
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

func TestBusinessIdentityPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var ids = map[string]bool{}
	var cursor string

	for i := 0; i < 2; i++ {
		var params = map[string]interface{}{}
		params["limit"] = 2
		params["cursor"] = cursor

		identities, nextCursor, err := BusinessIdentity.Page(params, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}

		for _, identity := range identities {
			assert.NotNil(t, identity.Id)
			ids[identity.Id] = true
		}

		cursor = nextCursor
		if cursor == "" {
			break
		}
	}

	assert.Equal(t, 4, len(ids))
}

func TestBusinessIdentityGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit

	var identityList []BusinessIdentity.BusinessIdentity

	identities, errorChannel := BusinessIdentity.Query(paramsQuery, nil)
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
		getIdentity, err := BusinessIdentity.Get(identity.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getIdentity.Id)
	}
	assert.Equal(t, limit, len(identityList))
}

func TestBusinessIdentityUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []BusinessIdentity.BusinessIdentity
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = 1
	paramsQuery["status"] = "created"

	identities, errorChannel := BusinessIdentity.Query(paramsQuery, nil)
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

	var patchData = map[string]interface{}{}
	patchData["tags"] = []string{"test", "testing"}

	identity, err := BusinessIdentity.Update(identityList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, identity.Id)
}

func TestBusinessIdentityCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "created"

	var identityList []BusinessIdentity.BusinessIdentity

	identities, errorChannel := BusinessIdentity.Query(paramsQuery, nil)
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

	identity, err := BusinessIdentity.Cancel(identityList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, identity.Id)
	assert.Equal(t, "canceled", identity.Status)
}
