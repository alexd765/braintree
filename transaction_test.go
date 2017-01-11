package braintree

import "testing"

func TestFindTransaction(t *testing.T) {
	t.Parallel()

	t.Run("existing", func(t *testing.T) {
		t.Parallel()

		transaction, err := bt.Transaction().Find("bx9a7av8")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != StatusSettled {
			t.Errorf("transaction.Status: got %s, want %s", transaction.Status, StatusSettled)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Transaction().Find("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}
