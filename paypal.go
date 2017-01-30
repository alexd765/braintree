package braintree

import "time"

// A Paypal account on braintree.
type Paypal struct {
	BillingAgreementID string         `xml:"billing-agreement-id"`
	CreatedAt          time.Time      `xml:"created-at"`
	CustomerID         string         `xml:"customer-id"`
	Default            bool           `xml:"default"`
	Email              string         `xml:"email"`
	ImageURL           string         `xml:"image-url"`
	Subscriptions      []Subscription `xml:"subscriptions>subscription"`
	Token              string         `xml:"token"`
	UpdatedAt          time.Time      `xml:"updated-at"`
}

func (pp Paypal) private() {}
