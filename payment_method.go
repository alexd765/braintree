package braintree

import "net/http"

// PaymentMethodInput is used to create a payment method on braintree.
// CustomerID and PaymentMethodNonce are required.
type PaymentMethodInput struct {
	// BillingAddress     Address
	BillingAddressID   string
	CardholderName     string
	CustomerID         string
	Options            PaymentMethodOptions
	PaymentMethodNonce string
	RiskData           PaymentMethodRiskData
	Token              string
}

// PaymentMethodOptions can be used as part of a PaymentMethodInput.
type PaymentMethodOptions struct {
	FailOnDuplicatePaymentMethod bool   `xml:"fail-on-duplicate-payment-method,omitempty"`
	MakeDefault                  bool   `xml:"make-default,omitempty"`
	VerificationAmount           string `xml:"verification-amount,omitempty"`
	VerificationMerchantID       string `xml:"verification-merchant-id,omitempty"`
	VerifyCard                   bool   `xml:"verify-card,omitempty"`
}

// PaymentMethodRiskData can be used as part of a PaymentMethodInput.
type PaymentMethodRiskData struct {
	CustomerBrowser string `xml:"customer-browser,omitempty"`
	CustomerIP      string `xml:"customer-ip,omitempty"`
}

// PaymentMethodGW is a payment method gateway.
type PaymentMethodGW struct {
	bt *Braintree
}

// Create a payment method on braintree.
//
// Todo: return other payment methods like paypal as well.
func (pgw PaymentMethodGW) Create(input *PaymentMethodInput) (*CreditCard, error) {
	card := &CreditCard{}
	if err := pgw.bt.execute(http.MethodPost, "payment_methods", card, input.sanitized()); err != nil {
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

type paymentMethodInputSanitized struct {
	XMLName string `xml:"payment-method"`
	// BillingAddress     *Address               `xml:"billing-address,omitempty"`
	BillingAddressID   string                 `xml:"billing-address-id,omitempty"`
	CardholderName     string                 `xml:"cardholder-name,omitempty"`
	CustomerID         string                 `xml:"customer-id"`
	Options            *PaymentMethodOptions  `xml:"options,omitempty"`
	PaymentMethodNonce string                 `xml:"payment-method-nonce"`
	RiskData           *PaymentMethodRiskData `xml:"risk-data,omitempty"`
	Token              string                 `xml:"token,omitempty"`
}

func (pmi PaymentMethodInput) sanitized() paymentMethodInputSanitized {
	sanitized := paymentMethodInputSanitized{
		BillingAddressID:   pmi.BillingAddressID,
		CardholderName:     pmi.CardholderName,
		CustomerID:         pmi.CustomerID,
		PaymentMethodNonce: pmi.PaymentMethodNonce,
		Token:              pmi.Token,
	}
	//	if pmi.BillingAddress != (Address{}) {
	//		sanitized.BillingAddress = &pmi.BillingAddress
	//	}
	if pmi.Options != (PaymentMethodOptions{}) {
		sanitized.Options = &pmi.Options
	}
	if pmi.RiskData != (PaymentMethodRiskData{}) {
		sanitized.RiskData = &pmi.RiskData
	}
	return sanitized
}
