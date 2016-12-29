package braintree

import (
	"net/http"
	"time"
)

const (
	// SubscriptionStatusActive means the subscription is active.
	SubscriptionStatusActive = "active"

	// SubscriptionStatusCanceled means the subscription was canceled.
	SubscriptionStatusCanceled = "canceled"

	// SubscriptionStatusExpired means the subscription has expired.
	SubscriptionStatusExpired = "expired"

	// SubscriptionStatusPastDue means the subscription is past due.
	SubscriptionStatusPastDue = "past due"

	// SubscriptionStatusPending means the the subscription will begin in the future.
	SubscriptionStatusPending = "pending"
)

// A Subscription on braintree
type Subscription struct {
	// AddOns
	// Balance
	// BillingPeriodEndDate
	BillingDayOfMonth int `xml:"billing-day-of-month"`
	// BillingPeriodStartDate
	CreatedAt           time.Time `xml:"created-at"`
	CurrentBillingCycle int       `xml:"current-billing-cycle"`
	// DaysPastDue         int       `xml:"days-past-due"`
	// Descriptor
	// Discounts
	// FailureCount int    `xml:"failure-count"`
	ID string `xml:"id"`
	// MerchantAccountID string `xml:"merchant-account-id"`
	// NeverExpires      bool   `xml:"never-expires"`
	// NextBillAmount
	// NextBillingDate
	// NextBillingPeriodAmount
	// NumberOfBillingCycles int `xml:"number-of-billing-cycles"`
	// PaidThroughDate
	PaymentMethodToken string `xml:"payment-method-token"`
	PlanID             string `xml:"plan-id"`
	// Price
	Status string `xml:"status"`
	// StatusHistory
	// Transactions
	// TrialDuration     int       `xml:"trial-duration"`
	// TrialDurationUnit string    `xml:"trial-duration-unit"`
	// TrialPeriod bool      `xml:"trial-period"`
	UpdatedAt time.Time `xml:"updated-at"`
}

// SubscriptionGW is a Subscription Gateway.
type SubscriptionGW struct {
	bt *Braintree
}

// Find a subscription with a given subscription id on braintree.
func (sgw SubscriptionGW) Find(id string) (*Subscription, error) {
	subscription := &Subscription{}
	if err := sgw.bt.execute(http.MethodGet, "subscriptions/"+id, subscription, nil); err != nil {
		return nil, err
	}
	return subscription, nil
}
