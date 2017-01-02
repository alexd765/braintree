package braintree

import "testing"

func TestCreatePaymentMethod(t *testing.T) {
	t.Parallel()

	t.Run("noToken", func(t *testing.T) {
		t.Parallel()
		customer := &Customer{FirstName: "first"}
		customer, err := bt.Customer().Create(customer)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		input := &PaymentMethodInput{CustomerID: customer.ID, PaymentMethodNonce: "fake-valid-visa-nonce"}
		card, err := bt.PaymentMethod().Create(input)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if card.CardType != CardTypeVisa {
			t.Errorf("card type: got %s, want: %s", card.CardType, CardTypeVisa)
		}
	})

	t.Run("noCustomerID", func(t *testing.T) {
		t.Parallel()
		input := &PaymentMethodInput{PaymentMethodNonce: "fake-valid-visa-nonce"}
		if _, err := bt.PaymentMethod().Create(input); err == nil || err.Error() != "422 Unprocessable Entity" {
			t.Errorf("got: %v, want: 422 Unprocessable Entity", err)
		}

	})
}
