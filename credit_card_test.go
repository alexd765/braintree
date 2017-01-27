package braintree

import "testing"

func TestCreditCardPrivate(t *testing.T) {
	CreditCard{}.private()
}
