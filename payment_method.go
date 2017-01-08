package braintree

import (
	"encoding/xml"
	"net/http"
)

// PaymentMethodInput is used to create a payment method on braintree.
//
// CustomerID and PaymentMethodNonce are required.
type PaymentMethodInput struct {
	XMLName            xml.Name
	BillingAddress     *AddressInput         `xml:"billing-address,omitempty"`
	BillingAddressID   string                `xml:"billing-address-id,omitempty"`
	CardholderName     string                `xml:"cardholder-name,omitempty"`
	CustomerID         string                `xml:"customer-id"`
	Options            *PaymentMethodOptions `xml:"options,omitempty"`
	PaymentMethodNonce string                `xml:"payment-method-nonce"`
	RiskData           *RiskData             `xml:"risk-data,omitempty"`
	Token              string                `xml:"token,omitempty"`
}

// PaymentMethodOptions can be used as part of a PaymentMethodInput.
type PaymentMethodOptions struct {
	FailOnDuplicatePaymentMethod bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
	MakeDefault                  bool   `xml:"make-default,omitempty"`
	VerificationAmount           string `xml:"verification-amount,omitempty"`
	VerificationMerchantID       string `xml:"verification-merchant-id,omitempty"`
	VerifyCard                   bool   `xml:"verify-card,omitempty"`
}

// PaymentMethodGW is a payment method gateway.
type PaymentMethodGW struct {
	bt *Braintree
}

// Create a payment method on braintree.
//
// Todo: return other payment methods like paypal as well.
func (pgw PaymentMethodGW) Create(input *PaymentMethodInput) (*CreditCard, error) {
	input.XMLName = xml.Name{Local: "payment-method"}
	card := &CreditCard{}
	if err := pgw.bt.execute(http.MethodPost, "payment_methods", card, input); err != nil {
		return nil, err
	}
	return card, nil
}

// Delete a payment method on braintree.
func (pgw PaymentMethodGW) Delete(token string) error {
	return pgw.bt.execute(http.MethodDelete, "payment_methods/any/"+token, nil, nil)
}

// Find a payment method on braintree.
func (pgw PaymentMethodGW) Find(token string) (*CreditCard, error) {
	card := &CreditCard{}
	if err := pgw.bt.execute(http.MethodGet, "payment_methods/any/"+token, card, nil); err != nil {
		return nil, err
	}
	return card, nil
}

// Update a payment method on braintree.
//
// Token is required.
func (pgw PaymentMethodGW) Update(input *PaymentMethodInput) (*CreditCard, error) {
	input.XMLName = xml.Name{Local: "payment-method"}
	card := &CreditCard{}
	if err := pgw.bt.execute(http.MethodPut, "payment_methods/any/"+input.Token, card, input); err != nil {
		return nil, err
	}
	return card, nil
}
