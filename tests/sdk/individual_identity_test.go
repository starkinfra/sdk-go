package sdk

import (
	"fmt"
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
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, identity := range identities {
		assert.NotNil(t, identity.Id)
		fmt.Println(identity.Id)
	}
}

func TestIndividualIdentityGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []IndividualIdentity.IndividualIdentity
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = 10

	identities := IndividualIdentity.Query(paramsQuery, nil)
	for identity := range identities {
		fmt.Println(identity)
		identityList = append(identityList, identity)
	}

	identity, err := IndividualIdentity.Get(identityList[rand.Intn(len(identityList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	fmt.Println(identity.Id)
	assert.NotNil(t, identity.Id)
}

func TestIndividualIdentityQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	identities := IndividualIdentity.Query(params, nil)
	for identity := range identities {
		fmt.Println(identity)
		assert.NotNil(t, identity.Id)
	}
}

func TestIndividualIdentityPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 3

	identities, cursor, err := IndividualIdentity.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, identity := range identities {
		assert.NotNil(t, identity.Id)
		fmt.Println(identity.Id)
	}

	fmt.Println(cursor)
}

func TestIndividualIdentityUpdated(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []IndividualIdentity.IndividualIdentity
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	identities := IndividualIdentity.Query(paramsQuery, nil)
	for identity := range identities {
		identityList = append(identityList, identity)
	}

	fmt.Println(identityList[rand.Intn(len(identityList))].Id)

	identity, err := IndividualIdentity.Update(identityList[rand.Intn(len(identityList))].Id, "processing", nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, identity.Id)
	fmt.Println(identity.Id)
}

func TestIndividualIdentityCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var identityList []IndividualIdentity.IndividualIdentity
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	identities := IndividualIdentity.Query(paramsQuery, nil)
	for identity := range identities {
		identityList = append(identityList, identity)
	}

	identity, err := IndividualIdentity.Cancel(identityList[rand.Intn(len(identityList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, identity.Id)
	fmt.Println(identity.Id)
}
