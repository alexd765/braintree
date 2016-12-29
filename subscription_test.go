package braintree

import "testing"

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
			t.Errorf("got: %v, want: 404: Not Found", err)
		}
		if subscription != nil {
			t.Errorf("got: %+v, want: <nil>", subscription)
		}
	})
}
