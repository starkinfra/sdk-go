package brcodepreview

import (
	"encoding/json"
	Error "github.com/starkinfra/core-go/starkcore/error"
	"github.com/starkinfra/core-go/starkcore/user/user"
	"github.com/starkinfra/sdk-go/starkinfra/utils"
	"time"
)

//	BrcodePreview struct
//
//	The BrcodePreview struct is used to preview information from a BR Code before paying it.
//
//	Parameters (required):
//	- Id [string]: BR Code from a Pix payment. This is also de information directly encoded in a QR Code. ex: "00020126580014br.gov.bcb.pix0136a629532e-7693-4846-852d-1bbff817b5a8520400005303986540510.005802BR5908T'Challa6009Sao Paulo62090505123456304B14A"
//
//	Attributes (return-only):
//	- AccountNumber [string]: Payment receiver account number. ex: "1234567"
//	- AccountType [string]: Payment receiver account type. ex: "checking"
//	- Amount [int]: Value in cents that this payment is expecting to receive. If 0, any value is accepted. ex: 123 (= R$1,23)
//	- AmountType [string]: amount type of the BR Code. If the amount type is "custom" the BR Code's amount can be changed by the sender at the moment of payment. Options: "fixed" or "custom"
//	- BankCode [string]: Payment receiver bank code. ex: "20018183"
//	- BranchCode [string]: Payment receiver branch code. ex: "0001"
//	- CashAmount [int]: Amount to be withdrawn from the cashier in cents. ex: 1000 (= R$ 10.00)
//	- CashierBankCode [string]: Cashier's bank code. ex: "20018183"
//	- CashierType [string]: Cashier's type. Options: "merchant", "participant" and "other"
//	- DiscountAmount [int]: Discount value calculated over nominalAmount. ex: 3000
//	- FineAmount [int]: Fine value calculated over nominalAmount. ex: 20000
//	- InterestAmount [int]: Interest value calculated over nominalAmount. ex: 10000
//	- KeyId [string]: Receiver's PixKey id. ex: "+5511989898989"
//	- Name [string]: Payment receiver name. ex: "Tony Stark"
//	- NominalAmount [int]: BR Code emission amount, without fines, fees and discounts. ex: 1234 (= R$ 12.34)
//	- ReconciliationId [string]: Reconciliation ID linked to this payment. If the brcode is dynamic, the reconciliationId will have from 26 to 35 alphanumeric characters, ex: "cd65c78aeb6543eaaa0170f68bd741ee". If the brcode is static, the ReconciliationId will have up to 25 alphanumeric characters "ah27s53agj6493hjds6836v49"
//	- ReductionAmount [int]: Reduction value to discount from nominalAmount. ex: 1000
//	- Scheduled [time.Time]: Date of payment execution. ex: time.Date(2023, 03, 10, 0, 0, 0, 0, time.UTC)
//	- Status [string]: Payment status. ex: "active", "paid", "canceled" or "unknown"
//	- TaxId [string]: Payment receiver tax ID. ex: "012.345.678-90"

type BrcodePreview struct {
	Id               string     `json:",omitempty"`
	AccountNumber    string     `json:",omitempty"`
	AccountType      string     `json:",omitempty"`
	Amount           int        `json:",omitempty"`
	AmountType       string     `json:",omitempty"`
	BankCode         string     `json:",omitempty"`
	BranchCode       string     `json:",omitempty"`
	CashAmount       int        `json:",omitempty"`
	CashierBankCode  string     `json:",omitempty"`
	CashierType      string     `json:",omitempty"`
	DiscountAmount   int        `json:",omitempty"`
	FineAmount       int        `json:",omitempty"`
	InterestAmount   int        `json:",omitempty"`
	KeyId            string     `json:",omitempty"`
	Name             string     `json:",omitempty"`
	NominalAmount    int        `json:",omitempty"`
	ReconciliationId string     `json:",omitempty"`
	ReductionAmount  int        `json:",omitempty"`
	Scheduled        *time.Time `json:",omitempty"`
	Status           string     `json:",omitempty"`
	TaxId            string     `json:",omitempty"`
}

var resource = map[string]string{"name": "BrcodePreview"}

func Create(previews []BrcodePreview, user user.User) ([]BrcodePreview, Error.StarkErrors) {
	//	Retrieve BrcodePreviews
	//
	//	Process BR Codes before paying them.
	//
	//	Parameters (required):
	//	- previews [slice of BrcodePreview structs]: Slice of BrcodePreview structs to preview. ex: []string{]brcodepreview.BrcodePreview("00020126580014br.gov.bcb.pix0136a629532e-7693-4846-852d-1bbff817b5a8520400005303986540510.005802BR5908T'Challa6009Sao Paulo62090505123456304B14A")
	//
	//	Parameters (optional):
	//	- user [Organization/Project struct, default nil]: Organization or Project struct. Not necessary if starkinfra.User was set before function call
	//
	//	Return:
	//	- Slice of BrcodePreview structs with updated attributes
	create, err := utils.Multi(resource, previews, nil, user)
	unmarshalError := json.Unmarshal(create, &previews)
	if unmarshalError != nil {
		return previews, err
	}
	return previews, err
}
