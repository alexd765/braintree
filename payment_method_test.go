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
			customer, err := bt.Customer().Create(CustomerInput{FirstName: "first"})
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

func TestDeletePaymentMethod(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{FirstName: "first"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		input := &PaymentMethodInput{CustomerID: customer.ID, PaymentMethodNonce: "fake-valid-visa-nonce"}
		card, err := bt.PaymentMethod().Create(input)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if err := bt.PaymentMethod().Delete(card.Token); err != nil {
			t.Errorf("unexpected err: %s", err)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if err := bt.PaymentMethod().Delete("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
	})
}

func TestFindPaymentMethod(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		card, err := bt.PaymentMethod().Find("j9jjzj")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if card.CardType != CardTypeVisa {
			t.Errorf("card type: got %s, want: %s", card.CardType, CardTypeVisa)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.PaymentMethod().Find("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
	})
}

func TestUpdatePaymentMethod(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{FirstName: "first"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		input := &PaymentMethodInput{CustomerID: customer.ID, PaymentMethodNonce: "fake-valid-visa-nonce"}
		card, err := bt.PaymentMethod().Create(input)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		input = &PaymentMethodInput{Token: card.Token, CardholderName: "name"}
		card, err = bt.PaymentMethod().Update(input)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if card.CardholderName != "name" {
			t.Errorf("cardholder name: got: %s, want: name", card.CardholderName)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		input := &PaymentMethodInput{Token: "nonExisting"}
		if _, err := bt.PaymentMethod().Update(input); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
	})
}
