package braintree

import "testing"

func TestCreatePaymentMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name  string
		Input PaymentMethodInput
	}{
		{
			Name: "minimal",
		},
		{
			Name: "withOptions",
			Input: PaymentMethodInput{
				Options: PaymentMethodOptions{
					MakeDefault: true,
				},
			},
		},
		{
			Name: "withRiskData",
			Input: PaymentMethodInput{
				RiskData: PaymentMethodRiskData{
					CustomerIP: "123.123.123.123",
				},
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()
			customer := &Customer{FirstName: "first"}
			customer, err := bt.Customer().Create(customer)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			test.Input.CustomerID = customer.ID
			test.Input.PaymentMethodNonce = "fake-valid-visa-nonce"
			card, err := bt.PaymentMethod().Create(&test.Input)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			if card.CardType != CardTypeVisa {
				t.Errorf("card type: got %s, want: %s", card.CardType, CardTypeVisa)
			}
		})
	}

	t.Run("noCustomerID", func(t *testing.T) {
		t.Parallel()
		input := &PaymentMethodInput{PaymentMethodNonce: "fake-valid-visa-nonce"}
		if _, err := bt.PaymentMethod().Create(input); err == nil || err.Error() != "422 Unprocessable Entity" {
			t.Errorf("got: %v, want: 422 Unprocessable Entity", err)
		}

	})
}
