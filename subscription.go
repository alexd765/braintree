package braintree

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/alexd765/braintree/btdate"
	"github.com/shopspring/decimal"
)

// Possible subscription statuses on braintree.
const (
	SubscriptionStatusActive   = "Active"
	SubscriptionStatusCanceled = "Canceled"
	SubscriptionStatusExpired  = "Expired"
	SubscriptionStatusPastDue  = "Past Due"
	SubscriptionStatusPending  = "Pending"
)

// A Subscription on braintree.
type Subscription struct {
	// AddOns
	Balance                decimal.Decimal `xml:"balance"`
	BillingPeriodEndDate   btdate.Date     `xml:"billing-period-end-date"`
	BillingDayOfMonth      int             `xml:"billing-day-of-month"`
	BillingPeriodStartDate btdate.Date     `xml:"billing-period-start-date"`
	CreatedAt              time.Time       `xml:"created-at"`
	CurrentBillingCycle    int             `xml:"current-billing-cycle"`
	// DaysPastDue         int       `xml:"days-past-due"`
	// Descriptor
	// Discounts
	// FailureCount int    `xml:"failure-count"`
	ID                string `xml:"id"`
	MerchantAccountID string `xml:"merchant-account-id"`
	// NeverExpires      bool   `xml:"never-expires"`
	NextBillAmount          decimal.Decimal `xml:"next-bill-amount"`
	NextBillingDate         btdate.Date     `xml:"next-billing-date"`
	NextBillingPeriodAmount decimal.Decimal `xml:"next-billing-period-amount"`
	// NumberOfBillingCycles int `xml:"number-of-billing-cycles"`
	PaidThroughDate    btdate.Date     `xml:"paid-through-date"`
	PaymentMethodToken string          `xml:"payment-method-token"`
	PlanID             string          `xml:"plan-id"`
	Price              decimal.Decimal `xml:"price"`
	Status             string          `xml:"status"`
	// StatusHistory
	// Transactions
	// TrialDuration     int       `xml:"trial-duration"`
	// TrialDurationUnit string    `xml:"trial-duration-unit"`
	// TrialPeriod bool      `xml:"trial-period"`
	UpdatedAt time.Time `xml:"updated-at"`
}

// SubscriptionInput is used to create or update a subscription.
type SubscriptionInput struct {
	XMLName xml.Name
	// AddOns
	BillingDayOfMonth int `xml:"billing-day-of-month,omitempty"`
	// Descriptor
	// Discounts
	FirstBillingDate  *btdate.Date `xml:"first-billing-date"`
	ID                string       `xml:"id,omitempty"`
	MerchantAccountID string       `xml:"merchant-account-id,omitempty"`
	NeverExpires      bool         `xml:"never-expires,omitempty"`
	// Options
	PaymentMethodNonce string           `xml:"payment-method-nonce,omitempty"`
	PaymentMethodToken string           `xml:"payment-method-token,omitempty"`
	PlanID             string           `xml:"plan-id,omitempty"`
	Price              *decimal.Decimal `xml:"price,omitempty"`
	TrialDuration      int              `xml:"trial-duration,omitempty"`
	TrialDurationUnit  string           `xml:"trial-duration-unit,omitempty"`
	TrialPeriod        bool             `xml:"trial-period,omitempty"`
}

// SubscriptionGW is a Subscription Gateway.
type SubscriptionGW struct {
	bt *Braintree
}

// Cancel a subscription on braintree.
func (sgw SubscriptionGW) Cancel(id string) (*Subscription, error) {
	subscription := &Subscription{}
	if err := sgw.bt.execute(http.MethodPut, "subscriptions/"+id+"/cancel", subscription, nil); err != nil {
		return nil, err
	}
	return subscription, nil
}

// Create a subscription on braintree.
//
// One of PaymentMethodNonce or PaymentMethodToken is required.
// PlanID is required.
func (sgw SubscriptionGW) Create(subscription SubscriptionInput) (*Subscription, error) {
	subscription.XMLName = xml.Name{Local: "subscription"}
	resp := &Subscription{}
	if err := sgw.bt.execute(http.MethodPost, "subscriptions", resp, subscription); err != nil {
		return nil, err
	}
	return resp, nil
}

// Find a subscription with a given subscription id on braintree.
func (sgw SubscriptionGW) Find(id string) (*Subscription, error) {
	subscription := &Subscription{}
	if err := sgw.bt.execute(http.MethodGet, "subscriptions/"+id, subscription, nil); err != nil {
		return nil, err
	}
	return subscription, nil
}

// Update a subscription on braintree.
//
// ID is required.
func (sgw SubscriptionGW) Update(subscription SubscriptionInput) (*Subscription, error) {
	subscription.XMLName = xml.Name{Local: "subscription"}
	resp := &Subscription{}
	if err := sgw.bt.execute(http.MethodPut, "subscriptions/"+subscription.ID, resp, subscription); err != nil {
		return nil, err
	}
	return resp, nil
}
