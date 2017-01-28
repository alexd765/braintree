package braintree

import (
	"encoding/xml"
	"time"
)

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
	CardTypeMasterCard      = "MasterCard"
	CardTypeSolo            = "Solo"
	CardTypeSwitch          = "Switch"
	CardTypeVisa            = "Visa"
	CardTypeUnknown         = "Unknown"
)

// A CreditCard is a braintree credit card.
type CreditCard struct {
	BillingAddress         Address        `xml:"billing-address"`
	BIN                    string         `xml:"bin"`
	CardType               string         `xml:"card-type"`
	CardholderName         string         `xml:"cardholder-name"`
	Commercial             string         `xml:"commercial"`
	CountryOfIssuance      string         `xml:"country-of-issuance"`
	CreatedAt              time.Time      `xml:"created-at"`
	CustomerID             string         `xml:"customer-id"`
	CustomerLocation       string         `xml:"customer-location"`
	Debit                  string         `xml:"debit"`
	Default                bool           `xml:"default"`
	DurbinRegulated        string         `xml:"durbin-regulated"`
	ExpirationDate         string         `xml:"expiration-date"`
	ExpirationMonth        string         `xml:"expiration-month"`
	ExpirationYear         string         `xml:"expiration-year"`
	Expired                bool           `xml:"expired"`
	Healthcare             string         `xml:"healthcare"`
	ImageURL               string         `xml:"image-url"`
	IssuingBank            string         `xml:"issuing-bank"`
	Last4                  string         `xml:"last-4"`
	MaskedNumber           string         `xml:"masked-number"`
	Payroll                string         `xml:"payroll"`
	Prepaid                string         `xml:"prepaid"`
	ProductID              string         `xml:"product-id"`
	Subscriptions          []Subscription `xml:"subscriptions>subscription"`
	Token                  string         `xml:"token"`
	UniqueNumberIdentifier string         `xml:"unique-number-identifier"`
	UpdatedAt              time.Time      `xml:"updated-at"`
}

// CreditCardInput is used to create or update a credit card on braintree.
type CreditCardInput struct {
	XMLName            xml.Name
	BillingAddress     *AddressInput         `xml:"billing-address,omitempty"`
	BillingAddressID   string                `xml:"billing-address-id,omitempty"`
	CardholderName     string                `xml:"cardholder-name,omitempty"`
	CustomerID         string                `xml:"customerID,omitempty"`
	Options            *PaymentMethodOptions `xml:"options,omitempty"`
	PaymentMethodNonce string                `xml:"payment-method-nonce,omitempty"`
	RiskData           *RiskData             `xml:"risk-data,omitempty"`
	Token              string                `xml:"token,omitempty"`
}

func (cc CreditCard) private() {}
