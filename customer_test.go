package braintree

import "testing"

func TestFindCustomer(t *testing.T) {
	t.Parallel()

	bt := New()

	t.Run("found", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.FindCustomer("cus1")
		if err != nil {
			t.Errorf("unexpected err: %s", err)
		}
		if customer == nil {
			t.Error("customer unexpected nil")
		}
		// fmt.Printf("customer: %+v\n", customer)
	})

	t.Run("notFound", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.FindCustomer("cus2")
		if err == nil || err.Error() != "404: not found" {
			t.Errorf("got: %v, want: 404: not found", err)
		}
		if customer != nil {
			t.Errorf("got: %+v, want: <nil>", customer)
		}
	})
}
