package tax_id_generator

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"time"

	"github.com/starkinfra/sdk-go/starkinfra"
	"github.com/starkinfra/sdk-go/starkinfra/brcodepreview"
	"github.com/starkinfra/sdk-go/starkinfra/creditholmes"
	"github.com/starkinfra/sdk-go/starkinfra/creditnote"
	"github.com/starkinfra/sdk-go/starkinfra/creditnote/invoice"
	"github.com/starkinfra/sdk-go/starkinfra/creditpreview"
	"github.com/starkinfra/sdk-go/starkinfra/creditsigner"
	"github.com/starkinfra/sdk-go/starkinfra/dynamicbrcode"
	"github.com/starkinfra/sdk-go/starkinfra/individualdocument"
	"github.com/starkinfra/sdk-go/starkinfra/individualidentity"
	"github.com/starkinfra/sdk-go/starkinfra/issuingcard"
	"github.com/starkinfra/sdk-go/starkinfra/issuingembossingrequest"
	"github.com/starkinfra/sdk-go/starkinfra/issuingholder"
	"github.com/starkinfra/sdk-go/starkinfra/issuinginvoice"
	"github.com/starkinfra/sdk-go/starkinfra/issuingrestock"
	"github.com/starkinfra/sdk-go/starkinfra/issuingwithdrawal"
	"github.com/starkinfra/sdk-go/starkinfra/pixchargeback"
	"github.com/starkinfra/sdk-go/starkinfra/pixclaim"
	"github.com/starkinfra/sdk-go/starkinfra/pixdirector"
	"github.com/starkinfra/sdk-go/starkinfra/pixfraud"
	"github.com/starkinfra/sdk-go/starkinfra/pixinfraction"
	"github.com/starkinfra/sdk-go/starkinfra/pixkey"
	"github.com/starkinfra/sdk-go/starkinfra/pixrequest"
	"github.com/starkinfra/sdk-go/starkinfra/pixreversal"
	"github.com/starkinfra/sdk-go/starkinfra/pixstatement"
	"github.com/starkinfra/sdk-go/starkinfra/staticbrcode"
	Utils "github.com/starkinfra/sdk-go/starkinfra/utils"
	"github.com/starkinfra/sdk-go/starkinfra/webhook"
	"github.com/starkinfra/sdk-go/tests/utils"
)

func BrcodePreview() []brcodepreview.BrcodePreview {

	starkinfra.User = utils.ExampleProject

	limit := 20
	var params = map[string]interface{}{}
	params["limit"] = limit

	var previews []brcodepreview.BrcodePreview

	brcodes, errorChannel := dynamicbrcode.Query(params, nil)

	loop:
	for {
		select {
		case err := <-errorChannel:
			if err.Errors != nil {
				for _, e := range err.Errors {
					fmt.Println(e.Code, e.Message)
				}
			}
		case brcode, ok := <-brcodes:
			if !ok {
				break loop
			}
			preview := brcodepreview.BrcodePreview{
				Id:      brcode.Id,
				PayerId: "012.345.678-90",
			}
			previews = append(previews, preview)
		
		}
	}

	return previews
}

func CreditNote() []creditnote.CreditNote {

	scheduled := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 2)
	amount := rand.Intn(100000 - 1)

	notes := []creditnote.CreditNote{
		{
			TemplateId:  "5706627130851328",
			Amount:      1000,
			Name:        "Jamie Lannister",
			TaxId:       Cpf(),
			Scheduled:   &scheduled,
			Invoices:    Invoice(amount),
			Payment:     Payment(),
			Signers:     Signer(),
			ExternalId:  ExternalId(),
			StreetLine1: "Rua ABC",
			StreetLine2: "Ap 123",
			District:    "Jardim Paulista",
			City:        "São Paulo",
			StateCode:   "SP",
			ZipCode:     "01234-567",
			PaymentType: "transfer",
		},
	}
	return notes
}

func Signer() []creditsigner.CreditSigner {

	signer := []creditsigner.CreditSigner{
		{
			Name:    "Jaime Lannister",
			Contact: "jaime.lannister@gmail.com",
			Method:  "link",
		},
	}
	return signer
}

func Invoice(amount int) []invoice.Invoice {

	due := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).AddDate(0, 0, 10).Format("2006-01-02")
	invoices := []invoice.Invoice{
		{
			Due:    due,
			Amount: amount,
		},
	}
	return invoices
}

func Payment() creditnote.Transfer {

	payment := creditnote.Transfer{
		BranchCode:    "1234",
		BankCode:      "00000000",
		AccountNumber: "129340-1",
		Name:          "Jaime Lannister",
		TaxId:         "012.345.678-90",
	}
	return payment
}

func Sac() creditpreview.CreditNotePreview {

	scheduled := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 1)
	initialDue := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 10)

	sac := creditpreview.CreditNotePreview{
		Type:            "sac",
		NominalAmount:   100000,
		Scheduled:       &scheduled,
		TaxId:           Cpf(),
		InitialDue:      &initialDue,
		NominalInterest: 10.0,
		Count:           3,
		Interval:        "week",
	}
	return sac
}

func Price() creditpreview.CreditNotePreview {

	scheduled := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 1)
	initialDue := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 10)

	price := creditpreview.CreditNotePreview{
		Type:            "price",
		NominalAmount:   100000,
		TaxId:           Cpf(),
		Scheduled:       &scheduled,
		InitialDue:      &initialDue,
		NominalInterest: 10,
		Count:           3,
		Interval:        "year",
	}
	return price
}

func American() creditpreview.CreditNotePreview {

	scheduled := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 1)
	initialDue := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 10)

	american := creditpreview.CreditNotePreview{
		Type:            "american",
		NominalAmount:   100000,
		Scheduled:       &scheduled,
		TaxId:           Cpf(),
		InitialDue:      &initialDue,
		NominalInterest: 10,
		Count:           3,
		Interval:        "month",
	}
	return american
}

func Bullet() creditpreview.CreditNotePreview {

	scheduled := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 1)
	initialDue := time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.UTC).Add(time.Hour * 24 * 10)

	bullet := creditpreview.CreditNotePreview{
		TaxId:           Cpf(),
		Type:            "bullet",
		NominalAmount:   100000,
		Scheduled:       &scheduled,
		InitialDue:      &initialDue,
		NominalInterest: 10,
	}
	return bullet
}

func CreditHolmes() []creditholmes.CreditHolmes {

	holmes := []creditholmes.CreditHolmes{
		{
			TaxId:      Cpf(),
			Competence: "2022-09",
		},
	}
	return holmes
}

func CreditPreview() []creditpreview.CreditPreview {

	sac := Sac()
	price := Price()
	american := American()
	bullet := Bullet()
	types := "credit-note"

	previews := []creditpreview.CreditPreview{
		{
			Credit: sac,
			Type:   types,
		},
		{
			Credit: price,
			Type:   types,
		},
		{
			Credit: american,
			Type:   types,
		},
		{
			Credit: bullet,
			Type:   types,
		},
	}
	return previews
}

func DynamicBrcode() []dynamicbrcode.DynamicBrcode {

	brcodes := []dynamicbrcode.DynamicBrcode{
		{
			Name:       "Jamie Lannister",
			City:       "Rio de Janeiro",
			ExternalId: ExternalId(),
			Type:       "instant",
		},
	}
	return brcodes
}

func IssuingCard() []issuingcard.IssuingCard {

	cards := []issuingcard.IssuingCard{
		{
			HolderName:       "Tony Stark",
			HolderTaxId:      "66.673.705/0001-88",
			HolderExternalId: ExternalId(),
		},
	}
	return cards
}

func IssuingCardEmbossing(holder issuingholder.IssuingHolder) []issuingcard.IssuingCard {

	cards := []issuingcard.IssuingCard{
		{
			HolderName:       holder.Name,
			HolderTaxId:      holder.TaxId,
			HolderExternalId: holder.ExternalId,
			ProductId:        "53810299",
			Type:             "physical",
		},
	}
	return cards
}

func IssuingRestock() []issuingrestock.IssuingRestock {

	restocks := []issuingrestock.IssuingRestock{
		{
			Count:   1000,
			StockId: "5152058940325888",
		},
		{
			Count:   2000,
			StockId: "6277958847168512",
		},
	}
	return restocks
}

func IssuingEmbossingRequest(cardId string) []issuingembossingrequest.IssuingEmbossingRequest {

	restocks := []issuingembossingrequest.IssuingEmbossingRequest{
		{
			CardId:                 cardId,
			KitId:                  "5659002855751680",
			DisplayName1:           "teste",
			ShippingCity:           "Sao Paulo",
			ShippingCountryCode:    "BRA",
			ShippingDistrict:       "Bela Vista",
			ShippingService:        "loggi",
			ShippingStateCode:      "SP",
			ShippingStreetLine1:    "teste",
			ShippingStreetLine2:    "teste",
			ShippingTrackingNumber: "5656565656565656",
			ShippingZipCode:        "12345-678",
			EmbosserId:             "5634161670881280",
		},
	}
	return restocks
}

func IssuingHolder() []issuingholder.IssuingHolder {

	holders := []issuingholder.IssuingHolder{
		{
			Name:       "Jannie Lanister",
			TaxId:      Cpf(),
			ExternalId: ExternalId(),
		},
	}
	return holders
}

func IssuingInvoice() issuinginvoice.IssuingInvoice {

	invoice := issuinginvoice.IssuingInvoice{
		Amount: 12345,
		TaxId:  Cpf(),
		Name:   "Jannie Lanister",
		Tags:   []string{"tony", "stark"},
	}
	return invoice
}

func IssuingWithdrawal() issuingwithdrawal.IssuingWithdrawal {

	withdrawal := issuingwithdrawal.IssuingWithdrawal{
		Amount:      123456,
		ExternalId:  ExternalId(),
		Description: "testeIssuingWithdrawal",
	}
	return withdrawal
}

func PixChargeback(e2e string) []pixchargeback.PixChargeback {

	chargebacks := []pixchargeback.PixChargeback{
		{
			Amount:      123456,
			ReferenceId: e2e,
			Reason:      "fraud",
		},
	}
	return chargebacks
}

func PixClaim() pixclaim.PixClaim {

	claim := pixclaim.PixClaim{
		AccountCreated: "2022-01-01",
		AccountNumber:  "76543",
		AccountType:    "checking",
		BranchCode:     "1234",
		Name:           "Jamie Lannister",
		TaxId:          "40.516.230/0001-22",
		KeyId:          fmt.Sprintf("+55119898671%v", rand.Intn(99)),
	}
	return claim
}

func PixDirector() pixdirector.PixDirector {

	director := pixdirector.PixDirector{
		Name:       "Edward Stark",
		TaxId:      "422.791.690-96",
		Phone:      "+5511999999999",
		Email:      "ned.stark@company.com",
		Password:   "12345678",
		TeamEmail:  "pix.team@company.com",
		TeamPhones: []string{"+5511989898989"},
	}
	return director
}

func PixInfraction() []pixinfraction.PixInfraction {

	infractions := []pixinfraction.PixInfraction{
		{
			ReferenceId:   "E35547753202201201450oo8sDGca066",
			Type:          "reversal",
			Method:        "scam",
			Description:   "testInfractionGolang",
			Tags:          []string{"tony", "stark"},
			OperatorEmail: "ned.stark@company.com",
			OperatorPhone: "+5511999999999",
		},
	}
	return infractions
}

func PixFraud() []pixfraud.PixFraud {

	frauds := []pixfraud.PixFraud{
		{
			ExternalId: "my_external_id_219",
			Type:       "scam",
			TaxId:      "01234567890",
		},
	}
	return frauds
}

func PixKey() pixkey.PixKey {

	accountCreated := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)

	key := pixkey.PixKey{
		AccountCreated: &accountCreated,
		AccountNumber:  "76543",
		AccountType:    "checking",
		BranchCode:     "1234",
		Name:           "Antony",
		TaxId:          Cpf(),
		Id:             fmt.Sprintf("+5511%v", rand.Intn(999999999-111111111)),
	}
	return key
}

func PixRequest() []pixrequest.PixRequest {

	requests := []pixrequest.PixRequest{
		{
			Amount:                12345,
			ExternalId:            ExternalId(),
			SenderName:            "Edward Stark",
			SenderTaxId:           "20.018.183/0001-80",
			SenderBranchCode:      "1357-9",
			SenderAccountNumber:   "876543-2",
			SenderAccountType:     "checking",
			ReceiverName:          "Edward Stark",
			ReceiverTaxId:         "01234567890",
			ReceiverBankCode:      "20018183",
			ReceiverAccountNumber: "876543-2",
			ReceiverBranchCode:    "1357-9",
			ReceiverAccountType:   "payment",
			EndToEndId:            Utils.EndToEndId("35547753"),
			Priority:              "low",
		},
	}
	return requests
}

func PixReversal() []pixreversal.PixReversal {

	reversals := []pixreversal.PixReversal{
		{
			Amount:     12345,
			ExternalId: ExternalId(),
			EndToEndId: "E34052649202204723420u34sDGd19l2",
			Reason:     "bankError",
		},
	}
	return reversals
}

func PixStatement() pixstatement.PixStatement {

	after := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)
	before := time.Date(2022, 01, 01, 0, 0, 0, 0, time.UTC)

	statement := pixstatement.PixStatement{
		After:  &after,
		Before: &before,
		Type:  "interchange",
	}
	return statement
}

func StaticBrcode() []staticbrcode.StaticBrcode {

	brcodes := []staticbrcode.StaticBrcode{
		{
			Name:            "Tony Stark",
			KeyId:           "+5511989898989",
			City:            "Rio de Janeiro",
			CashierBankCode: "20018183",
			Description:     "A Static Brcode",
		},
	}
	return brcodes
}

func IndividualIdentity() []individualidentity.IndividualIdentity {

	identities := []individualidentity.IndividualIdentity{
		{
			Name:  "Walter White",
			TaxId: "012.345.678-90",
			Tags:  []string{"breaking", "bad"},
		},
	}
	return identities
}

func IndividualDocument(identityId, documentType string, bytes []byte) []individualdocument.IndividualDocument {

	documents := []individualdocument.IndividualDocument{
		{
			Type:        documentType,
			ContentType: "image/png",
			Content:     base64.StdEncoding.EncodeToString(bytes),
			IdentityId:  identityId,
			Tags:        []string{"breaking", "bad"},
		},
	}

	return documents
}

func Webhook() webhook.Webhook {

	webhookObj := webhook.Webhook{
		Url:           fmt.Sprintf("https://webhook.site/%v", rand.Intn(20-10)),
		Subscriptions: []string{"boleto"},
	}
	return webhookObj
}
