package braintree

import "time"

// Card Types accepted by braintree.
const (
	CardTypeAmericanExpress = "American Express"
	CardTypeCarteBlanche    = "Card Blanche"
	CardTypeChinaUnionPay   = "China UnionPay"
	CardTypeDinersClub      = "Diners Club"
	CardTypeDiscover        = "Discover"
	CardTypeJCB             = "JCB"
	CardTypeLaser           = "Laser"
	CardTypeMaestro         = "Maestro"
	CardTypeMastercard      = "Mastercard"
	CardTypeSolo            = "Solo"
	CardTypeSwitch          = "Switch"
	CardTypeVisa            = "Visa"
	CardTypeUnknown         = "Unknown"
)

// A CreditCard is a braintree credit card.
type CreditCard struct {
	BillingAddress    Address   `xml:"billing-address"`
	BIN               string    `xml:"bin"`
	CardType          string    `xml:"card-type"`
	CardholderName    string    `xml:"cardholder-name"`
	Commercial        string    `xml:"commercial"`
	CountryOfIssuance string    `xml:"country-of-issuance"`
	CreatedAt         time.Time `xml:"created-at"`
	CustomerID        string    `xml:"customer-id"`
	CustomerLocation  string    `xml:"customer-location"`
	Debit             string    `xml:"debit"`
	Default           bool      `xml:"default"`
	DurbinRegulated   string    `xml:"durbin-regulated"`
	ExpirationDate    string    `xml:"expiration-date"`
	ExpirationMonth   string    `xml:"expiration-month"`
	ExpirationYear    string    `xml:"expiration-year"`
	Expired           bool      `xml:"expired"`
	Healthcare        string    `xml:"healthcare"`
	ImageURL          string    `xml:"image-url"`
	IssuingBank       string    `xml:"issuing-bank"`
	Last4             string    `xml:"last-4"`
	MaskedNumber      string    `xml:"masked-number"`
	Payroll           string    `xml:"payroll"`
	Prepaid           string    `xml:"prepaid"`
	ProductID         string    `xml:"product-id"`
	// Subscriptions
	Token                  string    `xml:"token"`
	UniqueNumberIdentifier string    `xml:"unique-number-identifier"`
	UpdatedAt              time.Time `xml:"updated-at"`
}
