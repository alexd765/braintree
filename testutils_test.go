package braintree

import (
	"reflect"
	"testing"
)

func compareErrors(t *testing.T, got, want error) {
	if got == want {
		return
	}
	if g, w := reflect.TypeOf(got), reflect.TypeOf(want); g != w {
		t.Errorf("error types: got '%v'; want '%v'", g, w)
	}
	if got == nil || want == nil || got.Error() != want.Error() {
		t.Errorf("error: got '%v'; want '%v'", got, want)
	}
}

func createTestCustomer(t *testing.T) *Customer {
	customer, err := bt.Customer().Create(CustomerInput{
		FirstName: "test",
		LastName:  "customer",
		CreditCard: &CreditCardInput{
			PaymentMethodNonce: "fake-valid-visa-nonce",
		},
	})
	if err != nil {
		t.Fatalf("unexpected error while creating test customer: %s", err)
	}
	return customer
}
