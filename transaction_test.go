package braintree

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestCreateTransaction(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{
			FirstName: "first",
			CreditCard: &CreditCardInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
			},
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		transaction, err := bt.Transaction().Create(TransactionInput{
			Amount: decimal.NewFromFloat(3),
			Options: &TransactionOptions{
				StoreInVaultOnSuccess: true,
			},
			PaymentMethodToken: customer.CreditCards[0].Token,
			Type:               TransactionTypeSale,
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if !transaction.Amount.Equals(decimal.NewFromFloat(3)) {
			t.Errorf("transaction.Amount: got %s, want 3", transaction.Amount)
		}

	})

	t.Run("withoutToken", func(t *testing.T) {
		t.Parallel()

		_, err := bt.Transaction().Create(TransactionInput{
			Amount: decimal.NewFromFloat(3),
			Type:   TransactionTypeSale,
		})
		if err == nil || err.Error() != "422 Unprocessable Entity" {
			t.Errorf("got: %v, want: 422 Unprocessable Entity", err)
		}
	})
}

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

func TestVoidTransaction(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{
			FirstName: "first",
			CreditCard: &CreditCardInput{
				PaymentMethodNonce: "fake-valid-visa-nonce",
			},
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		transaction, err := bt.Transaction().Create(TransactionInput{
			Amount: decimal.NewFromFloat(3.5),
			Options: &TransactionOptions{
				SubmitForSettlement: true,
			},
			PaymentMethodToken: customer.CreditCards[0].Token,
			Type:               TransactionTypeSale,
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != StatusSubmittedForSettlement {
			t.Fatalf("transaction.Status: expected %s, got %s", StatusSubmittedForSettlement, transaction.Status)
		}

		transaction, err = bt.Transaction().Void(transaction.ID)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != StatusVoided {
			t.Errorf("transaction.Status: expected %s, got %s", StatusVoided, transaction.Status)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Transaction().Void("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}
