package issuingrule

import (
	CardMethod "github.com/starkinfra/sdk-go/starkinfra/cardmethod"
	MerchantCategory "github.com/starkinfra/sdk-go/starkinfra/merchantcategory"
	MerchantCountry "github.com/starkinfra/sdk-go/starkinfra/merchantcountry"
)

//	IssuingRule struct
//
//	The IssuingRule struct displays the spending rules of IssuingCards and IssuingHolders created in your Workspace.
//
//	Parameters (required):
//	- Name [string]: Rule name. ex: "Travel" or "Food"
//	- Amount [int]: Maximum amount that can be spent in the informed interval. ex: 200000 (= R$ 2000.00)
//
//	Parameters (optional):
//	- Id [string, default nil]: Unique id returned when an IssuingRule is created, used to update a specific IssuingRule. ex: "5656565656565656"
//	- Interval [string, default "lifetime"]: Interval after which the rule amount counter will be reset to 0. ex: "instant", "day", "week", "month", "year" or "lifetime"
//	- CurrencyCode [string, default "BRL"]: Code of the currency that the rule amount refers to. ex: "BRL" or "USD"
//	- Categories [slice of MerchantCategory structs, default nil]: Merchant categories accepted by the rule. ex: []string{MerchantCategory(code="fastFoodRestaurants")]
//  - Countries [slice of MerchantCountry structs, default nil]: Countries accepted by the rule. ex: []string{MerchantCountry(code="BRA")]
//  - Methods [slice of CardMethod structs, default nil]: Card purchase methods accepted by the rule. ex: []string{CardMethod(code="magstripe")]
//
//	Attributes (expanded return-only):
//	- CounterAmount [int]: Current rule spent amount. ex: 1000
//	- CurrencySymbol [string]: Currency symbol. ex: "R$"
//	- CurrencyName [string]: Currency name. ex: "Brazilian Real"

type IssuingRule struct {
	Name           string                              `json:",omitempty"`
	Amount         int                                 `json:",omitempty"`
	Id             string                              `json:",omitempty"`
	Interval       string                              `json:",omitempty"`
	CurrencyCode   string                              `json:",omitempty"`
	Categories     []MerchantCategory.MerchantCategory `json:",omitempty"`
	Countries      []MerchantCountry.MerchantCountry   `json:",omitempty"`
	Methods        []CardMethod.CardMethod             `json:",omitempty"`
	CounterAmount  int                                 `json:",omitempty"`
	CurrencySymbol string                              `json:",omitempty"`
	CurrencyName   string                              `json:",omitempty"`
}
