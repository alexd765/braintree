package braintree

import (
	"testing"

	"github.com/alexd765/braintree/btdate"
	"github.com/shopspring/decimal"
)

func TestCancelSubscription(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()
		customer, err := bt.Customer().Create(CustomerInput{
			FirstName: "first",
			CreditCard: &CreditCardInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
			},
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

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

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{
			FirstName: "first",
			CreditCard: &CreditCardInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
			},
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		subscription, err := bt.Subscription().Create(
			SubscriptionInput{
				PlanID:             "plan1",
				PaymentMethodToken: customer.CreditCards[0].Token,
			},
		)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if subscription.PlanID != "plan1" {
			t.Errorf("subscription.PlanID: got: %s, want: plan1", subscription.PlanID)
		}
		if subscription.BillingPeriodStartDate != btdate.Today() {
			t.Errorf("subscription.BillingPeriodStartDate: got %s, want %s", subscription.BillingPeriodStartDate, btdate.Today())
		}
	})

	t.Run("withoutToken", func(t *testing.T) {
		if _, err := bt.Subscription().Create(SubscriptionInput{PlanID: "plan1"}); err == nil || err.Error() != "422 Unprocessable Entity" {
			t.Errorf("got: %v, want: 422 Unprocessable Entity", err)
		}
	})
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

func TestUpdateSubscription(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{
			FirstName: "first",
			CreditCard: &CreditCardInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
			},
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

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
