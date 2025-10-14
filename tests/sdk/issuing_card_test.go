package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingCardPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	cards, err := IssuingCard.Create(Example.IssuingCard(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, card := range cards {
		assert.NotNil(t, card.Id)
	}
}

func TestIssuingCardQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	cards, errorChannel := IssuingCard.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			assert.NotNil(t, card.Id)
		}
	}
}

func TestIssuingCardPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var params = map[string]interface{}{}
	params["limit"] = limit

	cards, cursor, err := IssuingCard.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, card := range cards {
		assert.NotNil(t, card.Id)
	}

	assert.NotNil(t, cursor)
}

func TestIssuingCardGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	
	var cardList []IssuingCard.IssuingCard

	cards, errorChannel := IssuingCard.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			cardList = append(cardList, card)
		}
	}

	for _, card := range cardList {
		getCard, err := IssuingCard.Get(card.Id, nil, nil)
		if err.Errors != nil {
			for _, e := range err.Errors {
				t.Errorf("code: %s, message: %s", e.Code, e.Message)
			}
		}
		assert.NotNil(t, getCard.Id)
	}
	assert.Equal(t, limit, len(cardList))
}

func TestIssuingCardDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "active"

	var cardList []IssuingCard.IssuingCard

	cards, errorChannel := IssuingCard.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			cardList = append(cardList, card)
		}
	}

	card, err := IssuingCard.Cancel(cardList[0].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, card.Id)
}

func TestIssuingCardUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 1
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = limit
	paramsQuery["status"] = "active"
	
	var cardList []IssuingCard.IssuingCard

	cards, errorChannel := IssuingCard.Query(paramsQuery, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case card, ok := <-cards:
			if !ok {
				break loop
			}
			cardList = append(cardList, card)
		}
	}

	var patchData = map[string]interface{}{}
	patchData["displayName"] = "ANTHONY EDWARD"

	card, err := IssuingCard.Update(cardList[0].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotNil(t, card.Id)
	assert.Equal(t, "ANTHONY EDWARD", card.DisplayName)
}
