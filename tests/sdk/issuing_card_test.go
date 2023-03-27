package sdk

import (
	"fmt"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	"github.com/starkinfra/sdk-go/tests/utils"
	Example "github.com/starkinfra/sdk-go/tests/utils/generator"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIssuingCardPost(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	cards, err := IssuingCard.Create(Example.IssuingCard(), nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, card := range cards {
		assert.NotNil(t, card.Id)
		fmt.Printf("%+v\n", card)
	}
}

func TestIssuingCardQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 10

	cards := IssuingCard.Query(params, nil)
	for card := range cards {
		assert.NotNil(t, card.Id)
		fmt.Printf("%+v\n", card)
	}
}

func TestIssuingCardPage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var params = map[string]interface{}{}
	params["limit"] = 1

	cards, cursor, err := IssuingCard.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	for _, card := range cards {
		assert.NotNil(t, card.Id)
		fmt.Println(card.Id)
	}

	fmt.Println(cursor)
}

func TestIssuingCardGet(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var cardList []IssuingCard.IssuingCard
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	cards := IssuingCard.Query(paramsQuery, nil)
	for card := range cards {
		cardList = append(cardList, card)
	}

	card, err := IssuingCard.Get(cardList[rand.Intn(len(cardList))].Id, nil, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, card.Id)
	fmt.Println(card.Id)
}

func TestIssuingCardDelete(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var cardList []IssuingCard.IssuingCard
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	cards := IssuingCard.Query(paramsQuery, nil)
	for card := range cards {
		cardList = append(cardList, card)
	}

	card, err := IssuingCard.Cancel(cardList[rand.Intn(len(cardList))].Id, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, card.Id)
	fmt.Println(card.Id)
}

func TestIssuingCardUpdate(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var cardList []IssuingCard.IssuingCard
	var paramsQuery = map[string]interface{}{}
	paramsQuery["limit"] = rand.Intn(100)
	paramsQuery["status"] = "created"

	cards := IssuingCard.Query(paramsQuery, nil)
	for card := range cards {
		cardList = append(cardList, card)
	}

	var patchData = map[string]interface{}{}
	patchData["displayName"] = "ANTHONY EDWARD"

	card, err := IssuingCard.Update(cardList[rand.Intn(len(cardList))].Id, patchData, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			panic(fmt.Sprintf("code: %s, message: %s", e.Code, e.Message))
		}
	}

	assert.NotNil(t, card.Id)
	fmt.Println(card.Id)
	fmt.Println(card.DisplayName)
}
