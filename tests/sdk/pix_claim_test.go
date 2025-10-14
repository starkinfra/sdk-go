package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	PixClaim "github.com/starkinfra/sdk-go/starkinfra/pixclaim"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPixClaimPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	claim, err := PixClaim.Create(Example.PixClaim(), nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, claim.Id)
}

func TestPixClaimQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	claims, errorChannel := PixClaim.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case claim, ok := <-claims:
			if !ok {
				break loop
			}
			assert.NotNil(t, claim.Id)
		}
	}
}

func TestPixClaimPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["status"] = "delivered"

	claims, cursor, err := PixClaim.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, claim := range claims {
		assert.NotNil(t, claim.Id)
	}
	assert.NotNil(t, cursor)
}

func TestPixClaimGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var claimList []PixClaim.PixClaim

	claims, errorChannel := PixClaim.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case claim, ok := <-claims:
			if !ok {
				break loop
			}
			claimList = append(claimList, claim)
		}
	}

	for _, claim := range claimList {
		getClaim, err := PixClaim.Get(claim.Id, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getClaim.Id)
	}

	assert.Equal(t, limit, len(claimList))
}

func TestPixClaimInfoPatch(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "delivered"
	
	var claimList []PixClaim.PixClaim
	
	claims, errorChannel := PixClaim.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case claim, ok := <-claims:
			if !ok {
				break loop
			}
			claimList = append(claimList, claim)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["status"] = "canceled"
	patchData["reason"] = "userRequested"

	claim, err := PixClaim.Update(claimList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, claim.Id)
}
