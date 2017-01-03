package braintree

import "testing"

func TestCreateSubscription(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{FirstName: "first"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		paymentMethodInput := &PaymentMethodInput{
			CustomerID:         customer.ID,
			PaymentMethodNonce: "fake-valid-visa-nonce",
		}
		card, err := bt.PaymentMethod().Create(paymentMethodInput)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		subscription, err := bt.Subscription().Create(
			SubscriptionInput{
				PaymentMethodToken: card.Token,
				PlanID:             "plan1",
			},
		)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if subscription.PlanID != "plan1" {
			t.Errorf("subscription.PlanID: got: %s, want: plan1", subscription.PlanID)
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
		if subscription == nil {
			t.Error("subscription unexpected nil")
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
