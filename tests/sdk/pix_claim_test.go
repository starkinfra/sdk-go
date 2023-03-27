package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestPixClaimPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	claim, err := PixClaim.Create(Example.PixClaim(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, claim.Id)
	fmt.Println(claim.Id)
	fmt.Println(claim.Status)
}

func TestPixClaimQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	claims := PixClaim.Query(params, nil)
	for claim := range claims {
		assert.NotNil(t, claim.Id)
		fmt.Println(claim.Id)
	}
}

func TestPixClaimPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "delivered"

	claims, cursor, err := PixClaim.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, claim := range claims {
		assert.NotNil(t, claim.Id)
		fmt.Println(claim.Id)
		fmt.Println(claim.Status)
	}
	fmt.Println(cursor)
}

func TestPixClaimGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var claimList []PixClaim.PixClaim
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	claims := PixClaim.Query(paramsQuery, nil)
	for claim := range claims {
		claimList = append(claimList, claim)
	}

	claim, err := PixClaim.Get(claimList[rand.Intn(len(claimList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, claim.Id)
	fmt.Println(claim.Id)
}

func TestPixClaimInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var claimList []PixClaim.PixClaim
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)

	claims := PixClaim.Query(paramsQuery, nil)
	for claim := range claims {
		claimList = append(claimList, claim)
	}

	var patchData = map[string]interface{}{}
	patchData["status"] = "canceled"
	patchData["reason"] = "userRequested"

	claim, err := PixClaim.Update(claimList[rand.Intn(len(claimList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, claim.Id)
	fmt.Println(claim.Id)
}
