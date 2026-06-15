package sdk

import (
	"github.com/starkinfra/sdk-go/starkinfra"
	IssuingStock "github.com/starkinfra/sdk-go/starkinfra/issuingstock"
	IssuingStockRule "github.com/starkinfra/sdk-go/starkinfra/issuingstockrule"
	"github.com/starkinfra/sdk-go/tests/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIssuingStockRuleQuery(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 10
	var params = map[string]interface{}{}
	params["limit"] = limit

	rules, errorChannel := IssuingStockRule.Query(params, nil)
	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case rule, ok := <-rules:
			if !ok {
				break loop
			}
			assert.NotEmpty(t, rule.Id)
		}
	}
}

func TestIssuingStockRulePage(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	limit := 3
	var params = map[string]interface{}{}
	params["limit"] = limit

	rules, cursor, err := IssuingStockRule.Page(params, nil)
	if err.Errors != nil {
		for _, e := range err.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	for _, rule := range rules {
		assert.NotEmpty(t, rule.Id)
	}

	assert.NotEmpty(t, cursor)
}

func TestIssuingStockRuleCreateUpdateCancel(t *testing.T) {

	starkinfra.User = utils.ExampleProject

	var stockParams = map[string]interface{}{}
	stockParams["limit"] = 1

	var stockId string
	stocks, stockErrorChannel := IssuingStock.Query(stockParams, nil)
	stockLoop:
	for {
		select {
		case err := <-stockErrorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case stock, ok := <-stocks:
			if !ok {
				break stockLoop
			}
			stockId = stock.Id
		}
	}

	assert.NotEmpty(t, stockId)

	var activeParams = map[string]interface{}{}
	activeParams["stockIds"] = []string{stockId}
	activeParams["status"] = []string{"active"}

	activeRules, activeErrorChannel := IssuingStockRule.Query(activeParams, nil)
	activeLoop:
	for {
		select {
		case err := <-activeErrorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		case rule, ok := <-activeRules:
			if !ok {
				break activeLoop
			}
			_, errFree := IssuingStockRule.Cancel(rule.Id, nil)
			if errFree.Errors != nil {
				for _, e := range errFree.Errors {
					t.Errorf("code: %s, message: %s", e.Code, e.Message)
				}
			}
		}
	}

	rules := []IssuingStockRule.IssuingStockRule{
		{
			MinimumBalance: 10000,
			StockId:        stockId,
			Tags:           []string{"card", "corporate"},
			Emails:         []string{"john.doe@enterprise.com"},
			Phones:         []string{"+5511912345678"},
		},
	}

	created, errCreate := IssuingStockRule.Create(rules, nil)
	if errCreate.Errors != nil {
		for _, e := range errCreate.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.NotEmpty(t, created[0].Id)
	id := created[0].Id

	var patchData = map[string]interface{}{}
	patchData["minimumBalance"] = 20000

	updated, errUpdate := IssuingStockRule.Update(id, patchData, nil)
	if errUpdate.Errors != nil {
		for _, e := range errUpdate.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.Equal(t, 20000, updated.MinimumBalance)

	canceled, errCancel := IssuingStockRule.Cancel(id, nil)
	if errCancel.Errors != nil {
		for _, e := range errCancel.Errors {
			t.Errorf("code: %s, message: %s", e.Code, e.Message)
		}
	}

	assert.Equal(t, "canceled", canceled.Status)
}
