package pixinternaltransactionreport

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	PixInternalTransactionReport struct
//
//	Transactions that happen internally — outside of the SPI — must be reported to
//	the Central Bank so they are reflected in the participant's statements. A
//	PixInternalTransactionReport is the report you create for each such transaction.
//	When you initialize a PixInternalTransactionReport, the entity will not be
//	automatically created in the Stark Infra API. The 'create' function sends the
//	structs to the Stark Infra API and returns the slice of created structs.
//
//	Parameters (required):
//	- Amount [int]: Amount in cents of the reported transaction. ex: 1234 (= R$ 12.34)
//	- Created [time.Time]: Datetime when the reported transaction occurred. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),
//	- EndToEndId [string]: Central Bank's unique transaction id. ex: "E20018183202201201213u34sav898j"
//	- Method [string]: Method used to create the reported Pix. ex: "manual", "dict", "staticQrcode", "dynamicQrcode"
//	- ReferenceType [string]: Type of the reported transaction. ex: "request", "reversal"
//	- SenderAccountNumber [string]: Sender's bank account number. ex: "876543-2"
//	- SenderBranchCode [string]: Sender's bank account branch code. ex: "1357-9"
//	- SenderAccountType [string]: Sender's bank account type. ex: "checking", "savings", "salary", "payment"
//	- SenderBankCode [string]: Sender's participant code (ISPB). ex: "20018183"
//	- SenderTaxId [string]: Sender's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//	- ReceiverAccountNumber [string]: Receiver's bank account number. ex: "876543-2"
//	- ReceiverBranchCode [string]: Receiver's bank account branch code. ex: "1357-9"
//	- ReceiverAccountType [string]: Receiver's bank account type. ex: "checking", "savings", "salary", "payment"
//	- ReceiverBankCode [string]: Receiver's participant code (ISPB). ex: "20018183"
//	- ReceiverTaxId [string]: Receiver's tax ID (CPF or CNPJ) with or without formatting. ex: "01234567890" or "20.018.183/0001-80"
//
//	Parameters (optional):
//	- ReceiverKeyId [string, default nil]: Receiver's Pix Key. ex: "+5511989898989"
//	- ReturnId [string, default nil]: Central Bank's unique reversal id. Required when ReferenceType is "reversal". ex: "D20018183202201201213u34sav898j"
//
//	Attributes (return-only):
//	- Id [string]: Unique id returned when the PixInternalTransactionReport is created. ex: "5656565656565656"
//	- Status [string]: Current PixInternalTransactionReport status. ex: "created", "failed", "sent", "success"
//	- Updated [time.Time]: Latest update datetime for the PixInternalTransactionReport. ex: time.Date(2020, 3, 10, 10, 30, 10, 0, time.UTC),

type PixInternalTransactionReport struct {
	Amount                int        `json:",omitempty"`
	Created               *time.Time `json:",omitempty"`
	EndToEndId            string     `json:",omitempty"`
	Method                string     `json:",omitempty"`
	ReferenceType         string     `json:",omitempty"`
	SenderAccountNumber   string     `json:",omitempty"`
	SenderBranchCode      string     `json:",omitempty"`
	SenderAccountType     string     `json:",omitempty"`
	SenderBankCode        string     `json:",omitempty"`
	SenderTaxId           string     `json:",omitempty"`
	ReceiverAccountNumber string     `json:",omitempty"`
	ReceiverBranchCode    string     `json:",omitempty"`
	ReceiverAccountType   string     `json:",omitempty"`
	ReceiverBankCode      string     `json:",omitempty"`
	ReceiverTaxId         string     `json:",omitempty"`
	ReceiverKeyId         string     `json:",omitempty"`
	ReturnId              string     `json:",omitempty"`
	Id                    string     `json:",omitempty"`
	Status                string     `json:",omitempty"`
	Updated               *time.Time `json:",omitempty"`
}

var resource = map[string]string{"name": "PixInternalTransactionReport"}

func Create(reports []PixInternalTransactionReport, user user.User) ([]PixInternalTransactionReport, Error.StarkErrors) {
	//	Create PixInternalTransactionReports
	//
	//	Send a slice of PixInternalTransactionReport structs for creation at the Stark Infra API
	//
	//	Parameters (required):
	//	- reports [slice of PixInternalTransactionReport structs]: Slice of PixInternalTransactionReport structs to be created in the API
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixInternalTransactionReport structs with updated attributes
	create, err := utils.Multi(resource, reports, nil, user)
	unmarshalError := json.Unmarshal(create, &reports)
	if unmarshalError != nil {
		return reports, err
	}
	return reports, err
}

func Get(id string, user user.User) (PixInternalTransactionReport, Error.StarkErrors) {
	//	Retrieve a specific PixInternalTransactionReport
	//
	//	Receive a single PixInternalTransactionReport struct previously created in the Stark Infra API by its id
	//
	//	Parameters (required):
	//	- id [string]: Struct unique id. ex: "5656565656565656"
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- pixInternalTransactionReport struct with updated attributes
	var pixInternalTransactionReport PixInternalTransactionReport
	get, err := utils.Get(resource, id, nil, user)
	unmarshalError := json.Unmarshal(get, &pixInternalTransactionReport)
	if unmarshalError != nil {
		return pixInternalTransactionReport, err
	}
	return pixInternalTransactionReport, err
}

func Query(params map[string]interface{}, user user.User) (chan PixInternalTransactionReport, chan Error.StarkErrors) {
	//	Retrieve PixInternalTransactionReports
	//
	//	Receive a channel of PixInternalTransactionReport structs previously created in the Stark Infra API
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- limit [int, default nil]: Maximum number of structs to be retrieved. Unlimited if nil. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "sent", "success"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- channel of PixInternalTransactionReport structs with updated attributes
	var pixInternalTransactionReport PixInternalTransactionReport
	reports := make(chan PixInternalTransactionReport)
	reportsError := make(chan Error.StarkErrors)
	query, errorChannel := utils.Query(resource, params, user)
	go func() {
		for content := range query {
			contentByte, _ := json.Marshal(content)
			err := json.Unmarshal(contentByte, &pixInternalTransactionReport)
			if err != nil {
				reportsError <- Error.UnknownError(err.Error())
				continue
			}
			reports <- pixInternalTransactionReport
		}
		for err := range errorChannel {
			reportsError <- err
		}
		close(reports)
		close(reportsError)
	}()
	return reports, reportsError
}

func Page(params map[string]interface{}, user user.User) ([]PixInternalTransactionReport, string, Error.StarkErrors) {
	//	Retrieve paged PixInternalTransactionReport structs
	//
	//	Receive a slice of up to 100 PixInternalTransactionReport structs previously created in the Stark Infra API and the cursor to the next page.
	//	Use this function instead of query if you want to manually page your requests.
	//
	//	Parameters (optional):
	//  - params [map[string]interface{}, default nil]: map of parameters for the query
	//		- cursor [string, default nil]: Cursor returned on the previous page function call
	//		- limit [int, default 100]: Maximum number of structs to be retrieved. Max = 100. ex: 35
	//		- after [string, default nil]: Date filter for structs created only after specified date. ex: "2020-03-10"
	//		- before [string, default nil]: Date filter for structs created only before specified date.  ex: "2022-11-10"
	//		- status [slice of strings, default nil]: Filter for status of retrieved structs. ex: []string{"created", "failed", "sent", "success"}
	//		- ids [slice of strings, default nil]: Slice of ids to filter retrieved structs. ex: []string{"5656565656565656", "4545454545454545"}
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- slice of PixInternalTransactionReport structs with updated attributes
	//	- cursor to retrieve the next page of PixInternalTransactionReport structs
	var pixInternalTransactionReports []PixInternalTransactionReport
	page, cursor, err := utils.Page(resource, params, user)
	unmarshalError := json.Unmarshal(page, &pixInternalTransactionReports)
	if unmarshalError != nil {
		return pixInternalTransactionReports, cursor, err
	}
	return pixInternalTransactionReports, cursor, err
}
