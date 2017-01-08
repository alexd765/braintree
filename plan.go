package braintree

import (
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// A Plan for a braintree subscription.
type Plan struct {
	// AddOns
	// BillingDayOfMonth int       `xml:"billing-day-of-month"`
	BillingFrequency int       `xml:"billing-frequency"`
	CreatedAt        time.Time `xml:"created-at"`
	CurrencyISOCode  string    `xml:"currency-iso-code"`
	Description      string    `xml:"description"`
	// Discounts
	ID   string `xml:"id"`
	Name string `xml:"name"`
	// NumberOfBillingCycles int    `xml:"number-of-billing-cycles"`
	Price decimal.Decimal `xml:"price"`
	// TrialDuration     int       `xml:"trial-duration"`
	// TrialDurationUnit string    `xml:"trial-duration_unit"`
	TrialPeriod bool      `xml:"trial-period"`
	UpdatedAt   time.Time `xml:"updated-at"`
}

// PlanGW is a Plan Gateway.
type PlanGW struct {
	bt *Braintree
}

// All returns all Plans in the merchant accounts.
func (pwg PlanGW) All() ([]Plan, error) {

	plansWrapper := struct {
		Plans []Plan `xml:"plan"`
	}{}
	if err := pwg.bt.execute(http.MethodGet, "plans", &plansWrapper, nil); err != nil {
		return nil, err
	}
	return plansWrapper.Plans, nil
}
