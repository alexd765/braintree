package braintree

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestCreatePaymentMethod(t *testing.T) {
	t.Parallel()

	tests := []struct {
		Name         string
		Input        PaymentMethodInput
		WantCardType string
	}{
		{
			Name: "CardMinimal",
			Input: PaymentMethodInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
			},
			WantCardType: CardTypeVisa,
		},
		{
			Name: "PaypalMinimal",
			Input: PaymentMethodInput{
				PaymentMethodNonce: "fake-paypal-future-nonce",
			},
		},
		{
			Name: "CardWithOptions",
			Input: PaymentMethodInput{
				PaymentMethodNonce: "fake-valid-mastercard-nonce",
				Options: &PaymentMethodOptions{
					MakeDefault: true,
				},
			},
			WantCardType: CardTypeMasterCard,
		},
		{
			Name: "CardWithRiskData",
			Input: PaymentMethodInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
				RiskData: &RiskData{
					CustomerIP: "123.123.123.123",
				},
			},
			WantCardType: CardTypeVisa,
		},
		{
			Name: "CardWithAddress",
			Input: PaymentMethodInput{
				PaymentMethodNonce: "fake-valid-mastercard-nonce",
				BillingAddress: &AddressInput{
					StreetAddress: "street",
				},
			},
			WantCardType: CardTypeMasterCard,
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
			pm, err := bt.PaymentMethod().Create(test.Input)
			if err != nil {
				t.Fatalf("unexpected err: %s", err)
			}
			switch pmi := pm.(type) {
			case *CreditCard:
				if test.WantCardType == "" {
					t.Fatal("payment method type: got *CreditCard, want *Paypal")
				}
				if pmi.CardType != test.WantCardType {
					t.Errorf("card type: got %s, want: %s", pmi.CardType, test.WantCardType)
				}
			case *Paypal:
				if test.WantCardType != "" {
					t.Fatal("payment method type: got *Paypal, want *CreditCard")
				}
				if pmi.Token == "" {
					t.Errorf("expected nonzero token")
				}
				fmt.Printf("paypal token: %s\n", pmi.Token)
			}
		})
	}

	t.Run("noCustomerID", func(t *testing.T) {
		t.Parallel()
		_, err := bt.PaymentMethod().Create(PaymentMethodInput{PaymentMethodNonce: "fake-valid-visa-nonce"})
		apiErr, ok := err.(*APIError)
		if !ok {
			t.Error("expected APIError")
		}
		if apiErr == nil || apiErr.Code != 91704 {
			t.Errorf("got %v, want error code 91704", apiErr)
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

		pm, err := bt.PaymentMethod().Create(PaymentMethodInput{CustomerID: customer.ID, PaymentMethodNonce: "fake-valid-visa-nonce"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		card, ok := pm.(*CreditCard)
		if !ok {
			t.Fatalf("payment method type: got %T, want CreditCard", pm)
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

	t.Run("visa", func(t *testing.T) {
		t.Parallel()

		pm, err := bt.PaymentMethod().Find("j9jjzj")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		card, ok := pm.(*CreditCard)
		if !ok {
			t.Fatalf("payment method type: got %T, want *CreditCard", pm)
		}
		if card.CardType != CardTypeVisa {
			t.Errorf("card type: got %s, want: %s", card.CardType, CardTypeVisa)
		}
	})

	t.Run("paypal", func(t *testing.T) {
		t.Parallel()
		pm, err := bt.PaymentMethod().Find("7wxmmp")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		pp, ok := pm.(*Paypal)
		if !ok {
			t.Fatalf("payment method type: got %T, want *Paypal", pm)
		}
		if pp.Token == "" {
			t.Errorf("got empty token, want non-empty")
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

		pm, err := bt.PaymentMethod().Create(PaymentMethodInput{CustomerID: customer.ID, PaymentMethodNonce: "fake-valid-visa-nonce"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		card, ok := pm.(*CreditCard)
		if !ok {
			t.Fatalf("payment method type: got %T, want CreditCard", pm)
		}

		pm, err = bt.PaymentMethod().Update(PaymentMethodInput{Token: card.Token, CardholderName: "name"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		card, ok = pm.(*CreditCard)
		if !ok {
			t.Fatalf("payment method type: got %T, want CreditCard", pm)
		}
		if card.CardholderName != "name" {
			t.Errorf("cardholder name: got: %s, want: name", card.CardholderName)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.PaymentMethod().Update(PaymentMethodInput{Token: "nonExisting"}); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
	})
}

func TestProtoPaymentMethodUnmarshalXML(t *testing.T) {
	t.Parallel()

	data := []byte("<abc></abc>")
	ppm := protoPaymentMethod{}
	wantErr := "unmarshal xml: unexpected start element: abc"
	if err := xml.Unmarshal(data, &ppm); err == nil || err.Error() != wantErr {
		t.Errorf("unmarshal protoPaymentMethod err: got %v, want %v", err, wantErr)
	}
}
