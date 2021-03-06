package braintree

import (
	"fmt"
	"testing"
	"time"

	"github.com/alexd765/braintree/btdate"
	"github.com/shopspring/decimal"
)

func TestCancelSubscription(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer := createTestCustomer(t)

		subscription, err := bt.Subscription().Create(
			SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
			},
		)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		subscription, err = bt.Subscription().Cancel(subscription.ID)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if subscription.Status != SubscriptionStatusCanceled {
			t.Errorf("subscription.Status: got: %s, want: %s", subscription.Status, SubscriptionStatusCanceled)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()
		if _, err := bt.Subscription().Cancel("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestCreateSubscription(t *testing.T) {
	t.Parallel()

	customer := createTestCustomer(t)
	twothousand := decimal.NewFromFloat(2000)

	tests := []struct {
		name                string
		input               SubscriptionInput
		wantErr             error
		wantStartDate       btdate.Date
		wantNextBillingDate btdate.Date
	}{
		{
			name: "normal",
			input: SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
			},
			wantStartDate:       btdate.Today(),
			wantNextBillingDate: btdate.FromTime(time.Now().UTC().AddDate(0, 1, 0)),
		},
		{
			name: "trial1day",
			input: SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
				TrialDuration:      1,
				TrialDurationUnit:  "day",
				TrialPeriod:        true,
			},
			wantNextBillingDate: btdate.FromTime(time.Now().UTC().AddDate(0, 0, 1)),
		},
		{
			name: "trial2months",
			input: SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
				TrialDuration:      2,
				TrialDurationUnit:  "month",
				TrialPeriod:        true,
			},
			wantNextBillingDate: btdate.FromTime(time.Now().UTC().AddDate(0, 2, 0)),
		},
		{
			name: "finite",
			input: SubscriptionInput{
				PlanID:                "plan1",
				PaymentMethodToken:    customer.CreditCards[0].Token,
				NumberOfBillingCycles: 3,
			},
			wantStartDate:       btdate.Today(),
			wantNextBillingDate: btdate.FromTime(time.Now().UTC().AddDate(0, 1, 0)),
		},
		{
			name: "invalidTrialDurationUnit",
			input: SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
				TrialDuration:      1,
				TrialDurationUnit:  "invalid",
				TrialPeriod:        true,
			},
			wantErr: &ValidationError{"", 81909, "Trial Duration Unit is invalid."},
		},
		{
			name: "withoutToken",
			input: SubscriptionInput{
				PlanID: "plan1",
			},
			wantErr: &ValidationError{"", 91903, "Payment method token is invalid."},
		},
		{
			name: "paymentFailed",
			input: SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
				Price:              &twothousand,
			},
			wantErr: &ProcessorError{2000, "Do Not Honor"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			subscription, err := bt.Subscription().Create(test.input)
			compareErrors(t, err, test.wantErr)
			if err != nil {
				return
			}
			if subscription.PlanID != "plan1" {
				t.Errorf("subscription.PlanID: got: %s, want: plan1", subscription.PlanID)
			}
			if subscription.BillingPeriodStartDate != test.wantStartDate {
				t.Errorf("subscription.BillingPeriodStartDate: got %s, want %s", subscription.BillingPeriodStartDate, test.wantStartDate)
			}
			if subscription.NextBillingDate != test.wantNextBillingDate {
				t.Errorf("subscription.NextBillingDate: got %s, want %s", subscription.NextBillingDate, test.wantNextBillingDate)
			}
		})
	}
}

func TestFindSubscription(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		subscription, err := bt.Subscription().Find("sub1")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if !subscription.Price.Equals(decimal.NewFromFloat(5)) {
			t.Errorf("subscription price: got %s, want 5", subscription.Price)
		}
		if len(subscription.Transactions) == 0 {
			t.Fatalf("subscription.Transactions: got 0, want more")
		}
		if planID := subscription.Transactions[0].PlanID; planID != "plan1" {
			t.Errorf("subscription.Transactions[0].PlanID: got %s, want plan1", planID)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		subscription, err := bt.Subscription().Find("sub2")
		if err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
		if subscription != nil {
			t.Errorf("got: %+v, want: <nil>", subscription)
		}
	})
}

func TestRetryChargeSubscription(t *testing.T) {

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()
		if err := bt.Subscription().RetryCharge("sub2"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})

	t.Run("notPastDue", func(t *testing.T) {
		t.Parallel()
		err := bt.Subscription().RetryCharge("sub1")
		valErr, ok := err.(*ValidationError)
		if !ok {
			t.Fatalf("expected ValidationError")
		}
		if valErr == nil || valErr.Code != 81531 {
			t.Errorf("got %v, want error code 81531", valErr)
		}
	})

	t.Run("shouldWork", func(t *testing.T) {
		t.Skip("manual intervention required")
		if err := bt.Subscription().RetryCharge(""); err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
	})
}

func TestUpdateSubscription(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer := createTestCustomer(t)

		subscription, err := bt.Subscription().Create(
			SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
			},
		)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		wantPrice := decimal.NewFromFloat(6)
		subscription, err = bt.Subscription().Update(
			SubscriptionInput{
				ID:    subscription.ID,
				Price: &wantPrice,
			},
		)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if !subscription.Price.Equals(decimal.NewFromFloat(6)) {
			t.Errorf("subscription.Price: got %s, want 6", subscription.Price)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		subscription, err := bt.Subscription().Update(SubscriptionInput{ID: "cus2"})
		if err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
		if subscription != nil {
			t.Errorf("got: %+v, want: <nil>", subscription)
		}
	})
}

// TestGeneratePastDueSubscriptions will generate subscriptions for a user
// "pastDue" with variying trial durations. When the trial is over the charge
// attempt will fail and the subscription state will be Past Due.
func TestGeneratePastDueSubscriptions(t *testing.T) {
	t.Skip("Not really a test.")

	customer, err := bt.Customer().Find("pastDue")
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	twoThousand := decimal.NewFromFloat(2000)
	for day := 1; day <= 14; day++ {
		fmt.Printf("day %d\n", day)
		for count := 0; count < 10; count++ {
			_, err := bt.Subscription().Create(SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
				Price:              &twoThousand,
				TrialDuration:      day,
				TrialDurationUnit:  "day",
				TrialPeriod:        true,
			})
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
		}
	}
}
