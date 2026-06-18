package issuingbillingtransaction

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	IssuingBillingTransaction struct
//
//	The IssuingBillingTransaction structs represent the transactions that compose the invoices charged for your issuing usage.
//
//	Attributes (return-only):
//	- Id [string]: unique id returned when the IssuingBillingTransaction is created. ex: "5656565656565656"
//	- Amount [int]: transaction amount in cents. ex: 1234 (= R$ 12.34)
//	- InvoiceId [string]: id of the parent IssuingBillingInvoice. May be empty when the transaction is not tied to an invoice. ex: "5656565656565656"
//	- Installment [int]: installment number of the transaction. ex: 1
//	- InstallmentCount [int]: total number of installments of the transaction. ex: 12
//	- Balance [int]: remaining balance in cents. ex: 1234 (= R$ 12.34)
//	- HolderName [string]: card holder name. ex: "Tony Stark"
//	- Source [string]: source of the transaction. ex: "issuing-purchase/5656565656565656"
//	- ExternalId [string]: external id of the transaction. ex: "my_external_id"
//	- Description [string]: transaction description. ex: "Buying food"
//	- CardEnding [string]: last 4 digits of the card. ex: "1234"
//	- Tax [int]: tax amount in cents. ex: 1234 (= R$ 12.34)
//	- Rate [float64]: tax rate as a percentage. ex: 1.5
//	- MerchantAmount [int]: merchant amount in cents. ex: 1234 (= R$ 12.34)
//	- MerchantCurrencyCode [string]: merchant currency code in ISO 4217 format. ex: "USD"
//	- Created [time.Time]: creation datetime for the IssuingBillingTransaction. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC)

type IssuingBillingTransaction struct {
	Id                   string     `json:",omitempty"`
	Amount               int        `json:",omitempty"`
	InvoiceId            string     `json:",omitempty"`
	Installment          int        `json:",omitempty"`
	InstallmentCount     int        `json:",omitempty"`
	Balance              int        `json:",omitempty"`
	HolderName           string     `json:",omitempty"`
	Source               string     `json:",omitempty"`
	ExternalId           string     `json:",omitempty"`
	Description          string     `json:",omitempty"`
	CardEnding           string     `json:",omitempty"`
	Tax                  int        `json:",omitempty"`
	Rate                 float64    `json:",omitempty"`
	MerchantAmount       int        `json:",omitempty"`
	MerchantCurrencyCode string     `json:",omitempty"`
	Created              *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "IssuingBillingTransaction"}

func Query(params map[string]interface{}, user user.User) (chan IssuingBillingTransaction, chan Error.StarkErrors) {
	//	Retrieve IssuingBillingTransactions
	//
	//	Receive a channel of IssuingBillingTransaction structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- invoiceId [string, default nil]: id of the IssuingBillingInvoice the transactions belong to. ex: "5656565656565656"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of IssuingBillingTransaction structs with updated attributes
	var issuingBillingTransaction IssuingBillingTransaction
	transactions := make(chan IssuingBillingTransaction)
	transactionsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &issuingBillingTransaction)
			if err != nil {
				transactionsError <- Error.UnknownError(err.Error())
				continue
			}
			transactions <- issuingBillingTransaction
		}
		for err := range errorChannel {
			transactionsError <- err
		}
		close(transactions)
		close(transactionsError)
	}()
	return transactions, transactionsError
}

func Page(params map[string]interface{}, user user.User) ([]IssuingBillingTransaction, string, Error.StarkErrors) {
	//	Retrieve paged IssuingBillingTransactions
	//
	//	Receive a slice of up to 100 IssuingBillingTransaction structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//	- params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2022-11-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date. ex: "2022-11-10"
	//		- invoiceId [string, default nil]: id of the IssuingBillingInvoice the transactions belong to. ex: "5656565656565656"
	//		- tags [slice of strings, default nil]: tags to filter retrieved structs. ex: []string{"tony", "stark"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of IssuingBillingTransaction structs with updated attributes
	//	- cursor to retrieve the next page of IssuingBillingTransaction structs
	var issuingBillingTransactions []IssuingBillingTransaction
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &issuingBillingTransactions)
	if unmarshalError != nil {
		return issuingBillingTransactions, cursor, err
	}
	return issuingBillingTransactions, cursor, err
}
