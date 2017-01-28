package braintree

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

// A PaymentMethod is currently a *CreditCard or *Paypal.
type PaymentMethod interface {
	private()
}

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
func (pgw PaymentMethodGW) Create(input PaymentMethodInput) (PaymentMethod, error) {
	input.XMLName = xml.Name{Local: "payment-method"}
	ppm := &protoPaymentMethod{}
	if err := pgw.bt.execute(http.MethodPost, "payment_methods", ppm, input); err != nil {
		return nil, err
	}
	return ppm.pm, nil
}

// Delete a payment method on braintree.
func (pgw PaymentMethodGW) Delete(token string) error {
	return pgw.bt.execute(http.MethodDelete, "payment_methods/any/"+token, nil, nil)
}

// Find a payment method on braintree.
func (pgw PaymentMethodGW) Find(token string) (PaymentMethod, error) {
	ppm := &protoPaymentMethod{}
	if err := pgw.bt.execute(http.MethodGet, "payment_methods/any/"+token, ppm, nil); err != nil {
		return nil, err
	}
	return ppm.pm, nil
}

// Update a payment method on braintree.
//
// Token is required.
func (pgw PaymentMethodGW) Update(input PaymentMethodInput) (PaymentMethod, error) {
	input.XMLName = xml.Name{Local: "payment-method"}
	ppm := &protoPaymentMethod{}
	if err := pgw.bt.execute(http.MethodPut, "payment_methods/any/"+input.Token, ppm, input); err != nil {
		return nil, err
	}
	return ppm.pm, nil
}

type protoPaymentMethod struct {
	pm PaymentMethod
}

func (ppm *protoPaymentMethod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	switch start.Name.Local {
	case "credit-card":
		(*ppm).pm = &CreditCard{}
		return d.DecodeElement(ppm.pm, &start)
	case "paypal-account":
		(*ppm).pm = &Paypal{}
		return d.DecodeElement(ppm.pm, &start)
	}
	return fmt.Errorf("unmarshal xml: unexpected start element: %s", start.Name.Local)
}
