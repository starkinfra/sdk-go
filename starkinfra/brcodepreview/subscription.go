package brcodepreview

import (
	"time"
)

// Subscription struct
//
// Subscription is a recurring payment that can be used to charge a user periodically.
// Returned as a subresource when previewing BR Codes that include subscription data.
//
// Attributes (return-only):
// - Amount [int]: Amount to be charged in cents. ex: 1000 = R$ 10.00
// - AmountMinLimit [int]: Minimum amount limit for the subscription. ex: 500 = R$ 5.00
// - BacenId [string]: BACEN (Brazilian Central Bank) identifier.
// - Created [time.Time]: Creation datetime for the subscription. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC)
// - Description [string]: Description of the subscription.
// - InstallmentEnd [time.Time]: End datetime for the installments. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC)
// - InstallmentStart [time.Time]: Start datetime for the installments. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC)
// - Interval [string]: Interval for the recurring charge. ex: "monthly"
// - PullRetryLimit [int]: Maximum number of retries for pulling the payment.
// - ReceiverBankCode [string]: Bank code of the receiver.
// - ReceiverName [string]: Name of the receiver.
// - ReceiverTaxId [string]: Tax ID of the receiver.
// - ReferenceCode [string]: Reference code for the subscription.
// - SenderFinalName [string]: Final sender name.
// - SenderFinalTaxId [string]: Final sender tax ID.
// - Status [string]: Current status of the subscription.
// - Type [string]: Type of the subscription.
// - Updated [time.Time]: Last update datetime for the subscription. ex: time.Date(2020, 3, 10, 0, 0, 0, 0, time.UTC)
type Subscription struct {
	Amount            int        `json:",omitempty"`
	AmountMinLimit    int        `json:",omitempty"`
	BacenId           string     `json:",omitempty"`
	Created           *time.Time `json:",omitempty"`
	Description       string     `json:",omitempty"`
	InstallmentEnd    *time.Time `json:",omitempty"`
	InstallmentStart  *time.Time `json:",omitempty"`
	Interval          string     `json:",omitempty"`
	PullRetryLimit    int        `json:",omitempty"`
	ReceiverBankCode  string     `json:",omitempty"`
	ReceiverName      string     `json:",omitempty"`
	ReceiverTaxId     string     `json:",omitempty"`
	ReferenceCode     string     `json:",omitempty"`
	SenderFinalName   string     `json:",omitempty"`
	SenderFinalTaxId  string     `json:",omitempty"`
	Status            string     `json:",omitempty"`
	Type              string     `json:",omitempty"`
	Updated           *time.Time `json:",omitempty"`
}
