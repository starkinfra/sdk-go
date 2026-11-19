package sdk

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"

	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingCard "github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	IssuingHolder "github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	IssuingProduct "github.com/starkinfra/sdk-go/starkinfra/issuingproduct"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
)

type recordedRequest struct {
	Method string
	Url    string
	Body   string
}

// fakeTransport replaces only the outgoing HTTP call: payload building and signing run for real.
type fakeTransport struct {
	body     string
	recorded []recordedRequest
}

func (f *fakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	body := ""
	if req.Body != nil {
		raw, _ := io.ReadAll(req.Body)
		body = string(raw)
	}
	f.recorded = append(f.recorded, recordedRequest{Method: req.Method, Url: req.URL.String(), Body: body})
	return &http.Response{StatusCode: 200, Body: io.NopCloser(strings.NewReader(f.body)), Header: http.Header{}, Request: req}, nil
}

func useFakeApi(t *testing.T, body string) *fakeTransport {
	starkinfra.User = utils.ExampleProject
	fake := &fakeTransport{body: body}
	original := http.DefaultTransport
	http.DefaultTransport = fake
	t.Cleanup(func() { http.DefaultTransport = original })
	return fake
}

func sentBody(t *testing.T, fake *fakeTransport) map[string]interface{} {
	assert.Equal(t, 1, len(fake.recorded))
	var body map[string]interface{}
	assert.NoError(t, json.Unmarshal([]byte(fake.recorded[0].Body), &body))
	return body
}

func TestIssuingCardWithoutHolderPayloadOmitsHolderFields(t *testing.T) {
	fake := useFakeApi(t, `{"cards": [{"id": "1", "status": "pending", "holderId": "", "holderName": ""}]}`)

	cards, err := IssuingCard.Create([]IssuingCard.IssuingCard{{ProductId: "123", Type: "physical"}}, nil, nil)

	assert.Nil(t, err.Errors)
	assert.Equal(t, map[string]interface{}{"cards": []interface{}{map[string]interface{}{"productId": "123", "type": "physical"}}}, sentBody(t, fake))
	assert.Equal(t, "", cards[0].HolderId)
	assert.Equal(t, "", cards[0].HolderName)
}

func TestIssuingCardPayloadSendsHolderFieldsWhenInformed(t *testing.T) {
	fake := useFakeApi(t, `{"cards": [{"id": "1"}]}`)

	_, err := IssuingCard.Create([]IssuingCard.IssuingCard{{
		HolderName: "Tony Stark", HolderTaxId: "012.345.678-90", HolderExternalId: "my-entity/123", ProductId: "123",
	}}, nil, nil)

	assert.Nil(t, err.Errors)
	assert.Equal(t, map[string]interface{}{"cards": []interface{}{map[string]interface{}{
		"holderName": "Tony Stark", "holderTaxId": "012.345.678-90", "holderExternalId": "my-entity/123", "productId": "123",
	}}}, sentBody(t, fake))
}

func TestIssuingCardUpdatePayload(t *testing.T) {
	cases := map[string]map[string]interface{}{
		"holder fields are not sent when omitted": {"displayName": "ANTHONY STARK"},
		"holder fields are sent alone":            {"holderName": "Tony Stark", "holderTaxId": "012.345.678-90", "holderExternalId": "my-entity/123"},
		"holder fields are sent with status and pin": {
			"status": "active", "pin": "1234", "holderName": "Tony Stark", "holderTaxId": "012.345.678-90", "holderExternalId": "my-entity/123",
		},
	}
	for name, patchData := range cases {
		t.Run(name, func(t *testing.T) {
			fake := useFakeApi(t, `{"card": {"id": "5656565656565656", "status": "pending"}}`)

			_, err := IssuingCard.Update("5656565656565656", patchData, nil)

			assert.Nil(t, err.Errors)
			assert.Equal(t, "PATCH", fake.recorded[0].Method)
			assert.True(t, strings.HasSuffix(fake.recorded[0].Url, "/v2/issuing-card/5656565656565656"), fake.recorded[0].Url)
			assert.Equal(t, patchData, sentBody(t, fake))
		})
	}
}

func TestIssuingCardUpdateMultiPayload(t *testing.T) {
	fake := useFakeApi(t, `{"cards": [{"id": "1", "status": "pending"}, {"id": "2", "status": "pending"}]}`)

	cards, err := IssuingCard.UpdateMulti([]map[string]interface{}{
		{"id": "1", "holderName": "Tony Stark", "holderTaxId": "012.345.678-90", "holderExternalId": "my-entity/123"},
		{"id": "2", "tags": []string{"batch"}},
	}, nil)

	assert.Nil(t, err.Errors)
	assert.Equal(t, "PATCH", fake.recorded[0].Method)
	assert.True(t, strings.HasSuffix(fake.recorded[0].Url, "/v2/issuing-card"), fake.recorded[0].Url)
	assert.Equal(t, map[string]interface{}{"cards": []interface{}{
		map[string]interface{}{"id": "1", "holderName": "Tony Stark", "holderTaxId": "012.345.678-90", "holderExternalId": "my-entity/123"},
		map[string]interface{}{"id": "2", "tags": []interface{}{"batch"}},
	}}, sentBody(t, fake))
	assert.Equal(t, 2, len(cards))
	assert.Equal(t, "1", cards[0].Id)
	assert.Equal(t, "2", cards[1].Id)
}

// ---- sandbox tests: they run against the real API and cancel every card they create ----

func firstProduct(t *testing.T) IssuingProduct.IssuingProduct {
	starkinfra.User = utils.ExampleProject
	products, errorChannel := IssuingProduct.Query(map[string]interface{}{"limit": 1}, nil)
	var first IssuingProduct.IssuingProduct
	var queryError []Error.StarkError
	found := false
	for products != nil || errorChannel != nil {
		select {
		case err, ok := <-errorChannel:
			if !ok {
				errorChannel = nil
				continue
			}
			if err.Errors != nil && queryError == nil {
				queryError = err.Errors
			}
		case product, ok := <-products:
			if !ok {
				products = nil
				continue
			}
			if !found {
				first, found = product, true
			}
		}
	}
	// The channels are drained before failing so the query goroutine is never left blocked on a send.
	if queryError != nil {
		t.Fatalf("product query failed: %v", queryError)
	}
	if !found {
		t.Fatalf("no issuing product in the sandbox Project")
	}
	return first
}

// A holder can only be linked to a product of its own type: business holders have a CNPJ (14 digits) and individual holders a
// CPF (11 digits). Existing holders are used so no personal data is created by the test.
func findExistingHolderForProduct(t *testing.T, product IssuingProduct.IssuingProduct) (IssuingHolder.IssuingHolder, bool) {
	digits := map[string]int{"business": 14, "individual": 11}[product.HolderType]
	holders, errorChannel := IssuingHolder.Query(map[string]interface{}{"status": "active", "limit": 100}, nil)
	var found IssuingHolder.IssuingHolder
	var queryError []Error.StarkError
	ok := false
	for holders != nil || errorChannel != nil {
		select {
		case err, open := <-errorChannel:
			if !open {
				errorChannel = nil
				continue
			}
			if err.Errors != nil && queryError == nil {
				queryError = err.Errors
			}
		case holder, open := <-holders:
			if !open {
				holders = nil
				continue
			}
			if !ok && len(onlyDigits(holder.TaxId)) == digits {
				found, ok = holder, true
			}
		}
	}
	if queryError != nil {
		t.Fatalf("holder query failed: %v", queryError)
	}
	return found, ok
}

func onlyDigits(value string) string {
	var out strings.Builder
	for _, char := range value {
		if char >= '0' && char <= '9' {
			out.WriteRune(char)
		}
	}
	return out.String()
}

func cardsWithoutHolder(n int, cardType string, productId string) []IssuingCard.IssuingCard {
	cards := []IssuingCard.IssuingCard{}
	for i := 0; i < n; i++ {
		cards = append(cards, IssuingCard.IssuingCard{ProductId: productId, Type: cardType, Tags: []string{"sdk-test"}})
	}
	return cards
}

// Cards the API rejected come back without an id and were never created, so there is nothing to cancel for them. A failed
// cancel is logged by id only (a leftover card is also reported by the sandbox-test plan).
func cancelAll(t *testing.T, cards []IssuingCard.IssuingCard) {
	for _, card := range cards {
		if card.Id == "" {
			continue
		}
		if _, err := IssuingCard.Cancel(card.Id, nil); err.Errors != nil {
			t.Logf("could not cancel card %s", card.Id)
		}
	}
}

// The API rejection reaches the caller untouched, with the error code of the API.
func assertRejected(t *testing.T, err Error.StarkErrors, code string) {
	if err.Errors == nil {
		t.Errorf("expected the API to reject the call with %s", code)
		return
	}
	codes := []string{}
	for _, e := range err.Errors {
		codes = append(codes, e.Code)
	}
	assert.Contains(t, codes, code)
}

func holderFields(holder IssuingHolder.IssuingHolder) map[string]interface{} {
	return map[string]interface{}{"holderName": holder.Name, "holderTaxId": holder.TaxId, "holderExternalId": holder.ExternalId}
}

// pendingCards creates pending physical cards without holder, always cancels them at the end and skips the test when the
// sandbox has no holder that can be linked to the product.
func pendingCards(t *testing.T, n int) ([]IssuingCard.IssuingCard, IssuingHolder.IssuingHolder) {
	product := firstProduct(t)
	holder, ok := findExistingHolderForProduct(t, product)
	if !ok {
		t.Skip("no existing holder compatible with the sandbox product")
	}
	cards, err := IssuingCard.Create(cardsWithoutHolder(n, "physical", product.Id), nil, nil)
	if err.Errors != nil {
		t.Fatalf("could not create the cards: %v", err.Errors)
	}
	t.Cleanup(func() { cancelAll(t, cards) })
	return cards, holder
}

func TestIssuingCardWithoutHolderPhysicalIsCreatedPending(t *testing.T) {
	starkinfra.User = utils.ExampleProject
	product := firstProduct(t)

	cards, err := IssuingCard.Create(cardsWithoutHolder(1, "physical", product.Id), nil, nil)
	defer cancelAll(t, cards)

	assert.Nil(t, err.Errors)
	assert.Equal(t, 1, len(cards))
	assert.Equal(t, "pending", cards[0].Status)
	assert.Equal(t, "physical", cards[0].Type)
	assert.Equal(t, "", cards[0].HolderId)
	assert.Equal(t, "", cards[0].HolderName)
}

func TestIssuingCardWithoutHolderApiRejectsVirtualCard(t *testing.T) {
	starkinfra.User = utils.ExampleProject

	cards, err := IssuingCard.Create(cardsWithoutHolder(1, "virtual", firstProduct(t).Id), nil, nil)
	defer cancelAll(t, cards)

	assertRejected(t, err, "invalidHolder")
}

func TestIssuingCardWithoutHolderApiRejectsIncompleteHolderFields(t *testing.T) {
	starkinfra.User = utils.ExampleProject
	cards := cardsWithoutHolder(1, "physical", firstProduct(t).Id)
	cards[0].HolderName = "Tony Stark"

	created, err := IssuingCard.Create(cards, nil, nil)
	defer cancelAll(t, created)

	assertRejected(t, err, "invalidHolder")
}

func TestIssuingCardAssignHolderToAPendingCard(t *testing.T) {
	cards, holder := pendingCards(t, 1)

	card, err := IssuingCard.Update(cards[0].Id, holderFields(holder), nil)

	assert.Nil(t, err.Errors)
	assert.Equal(t, "pending", card.Status)
	assert.Equal(t, holder.Id, card.HolderId)
}

func TestIssuingCardAssignHolderAndActivateInTheSameCall(t *testing.T) {
	cards, holder := pendingCards(t, 1)
	patchData := holderFields(holder)
	patchData["status"] = "active"
	patchData["pin"] = "1234"

	card, err := IssuingCard.Update(cards[0].Id, patchData, nil)

	assert.Nil(t, err.Errors)
	assert.Equal(t, "active", card.Status)
	assert.Equal(t, holder.Id, card.HolderId)
}

func TestIssuingCardAssignHolderApiRejectsIncompleteHolderFields(t *testing.T) {
	cards, holder := pendingCards(t, 1)

	_, err := IssuingCard.Update(cards[0].Id, map[string]interface{}{"holderName": holder.Name}, nil)

	assertRejected(t, err, "invalidHolder")
}

func TestIssuingCardAssignHolderApiRejectsActivationWithoutHolder(t *testing.T) {
	cards, _ := pendingCards(t, 1)

	_, err := IssuingCard.Update(cards[0].Id, map[string]interface{}{"status": "active", "pin": "1234"}, nil)

	assertRejected(t, err, "invalidHolder")
}

func TestIssuingCardUpdateMultiAssignsHolderToSeveralCards(t *testing.T) {
	cards, holder := pendingCards(t, 2)
	items := []map[string]interface{}{}
	for _, card := range cards {
		item := holderFields(holder)
		item["id"] = card.Id
		items = append(items, item)
	}

	updated, err := IssuingCard.UpdateMulti(items, nil)

	assert.Nil(t, err.Errors)
	assert.Equal(t, 2, len(updated))
	ids := []string{}
	for _, card := range updated {
		assert.Equal(t, holder.Id, card.HolderId)
		ids = append(ids, card.Id)
	}
	assert.ElementsMatch(t, []string{cards[0].Id, cards[1].Id}, ids)
}

func TestIssuingCardUpdateMultiApiRejectsTheWholeListWhenOneItemIsInvalid(t *testing.T) {
	cards, holder := pendingCards(t, 2)
	valid := holderFields(holder)
	valid["id"] = cards[0].Id
	invalid := holderFields(holder)
	invalid["id"] = "0000000000000001"

	_, err := IssuingCard.UpdateMulti([]map[string]interface{}{valid, invalid}, nil)

	assertRejected(t, err, "invalidCard")
	card, getErr := IssuingCard.Get(cards[0].Id, nil, nil)
	assert.Nil(t, getErr.Errors)
	assert.Equal(t, "", card.HolderId)
}

func TestIssuingCardUpdateMultiApiRejectsIncompleteHolderFields(t *testing.T) {
	cards, holder := pendingCards(t, 1)

	_, err := IssuingCard.UpdateMulti([]map[string]interface{}{{"id": cards[0].Id, "holderName": holder.Name}}, nil)

	assertRejected(t, err, "invalidHolder")
}
