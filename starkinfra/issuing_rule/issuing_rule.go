package issuing_rule

//	IssuingRule object
//	The IssuingRule object displays the spending rules of IssuingCards and IssuingHolders created in your Workspace.
//
//	Parameters (required):
//	- name [string]: rule name. ex: "Travel" or "Food"
//	- amount [integer]: maximum amount that can be spent in the informed interval. ex: 200000 (= R$ 2000.00)
//
//	Parameters (optional):
//	- id [string, default None]: unique id returned when an IssuingRule is created, used to update a specific IssuingRule. ex: "5656565656565656"
//	- interval [string, default "lifetime"]: interval after which the rule amount counter will be reset to 0. ex: "instant", "day", "week", "month", "year" or "lifetime"
//	- currency_code [string, default "BRL"]: code of the currency that the rule amount refers to. ex: "BRL" or "USD"
//	- categories [list of strings, default []]: merchant categories accepted by the rule. ex: ["eatingPlacesRestaurants", "travelAgenciesTourOperators"]
//	- countries [list of strings, default []]: countries accepted by the rule. ex: ["BRA", "USA"]
//	- methods [list of strings, default []]: card purchase methods accepted by the rule. ex: ["chip", "token", "server", "manual", "magstripe", "contactless"]
//
//	Attributes (expanded return-only):
//	- counter_amount [integer]: current rule spent amount. ex: 1000
//	- currency_symbol [string]: currency symbol. ex: "R$"
//	- currency_name [string]: currency name. ex: "Brazilian Real"

type IssuingRule struct {
	Name           string   `json:"name"`
	Amount         int      `json:"amount"`
	Id             string   `json:"id"`
	Interval       string   `json:"interval"`
	CurrencyCode   string   `json:"currencyCode"`
	Categories     []string `json:"categories"`
	Countries      []string `json:"countries"`
	Methods        []string `json:"methods"`
	CounterAmount  int      `json:"counterAmount"`
	CurrencySymbol string   `json:"currencySymbol"`
	CurrencyName   string   `json:"currencyName"`
}

var resource = map[string]string{"class": IssuingRule{}, "name": "IssuingRule"}

func ParseRules() {

}
