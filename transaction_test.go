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
		apiErr, ok := err.(*APIError)
		if !ok {
			t.Errorf("expected error of type APIError")
		}
		if apiErr == nil || apiErr.Code != 91508 {
			t.Errorf("api error code: got %v, want 91508", apiErr)
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
		if transaction.Status != TransactionStatusSettled {
			t.Errorf("transaction.Status: got %s, want %s", transaction.Status, TransactionStatusSettled)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Transaction().Find("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestRefundTransaction(t *testing.T) {
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
			Amount: decimal.NewFromFloat(3.7),
			Options: &TransactionOptions{
				SubmitForSettlement: true,
			},
			PaymentMethodToken: customer.CreditCards[0].Token,
			Type:               TransactionTypeSale,
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != TransactionStatusSubmittedForSettlement {
			t.Fatalf("transaction.Status: expected %s, got %s", TransactionStatusSubmittedForSettlement, transaction.Status)
		}

		transaction, err = bt.Transaction().Settle(transaction.ID)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != TransactionStatusSettled {
			t.Errorf("transaction.Status: expected %s, got %s", TransactionStatusSettled, transaction.Status)
		}

		transaction2, err := bt.Transaction().Refund(transaction.ID)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction2.Status != TransactionStatusSubmittedForSettlement {
			t.Errorf("transaction2.Status: expected %s, got %s", TransactionStatusSubmittedForSettlement, transaction.Status)
		}
		if transaction2.RefundedTransactionID != transaction.ID {
			t.Errorf("transaction2.RefundedTransactionID: expected %s, got %s", transaction.ID, transaction2.RefundedTransactionID)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Transaction().Refund("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestSettleTransaction(t *testing.T) {
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
			Amount: decimal.NewFromFloat(3.6),
			Options: &TransactionOptions{
				SubmitForSettlement: true,
			},
			PaymentMethodToken: customer.CreditCards[0].Token,
			Type:               TransactionTypeSale,
		})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != TransactionStatusSubmittedForSettlement {
			t.Fatalf("transaction.Status: expected %s, got %s", TransactionStatusSubmittedForSettlement, transaction.Status)
		}

		transaction, err = bt.Transaction().Settle(transaction.ID)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != TransactionStatusSettled {
			t.Errorf("transaction.Status: expected %s, got %s", TransactionStatusSettled, transaction.Status)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Transaction().Settle("nonExisting"); err == nil || err.Error() != "404 Not Found" {
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
		if transaction.Status != TransactionStatusSubmittedForSettlement {
			t.Fatalf("transaction.Status: expected %s, got %s", TransactionStatusSubmittedForSettlement, transaction.Status)
		}

		transaction, err = bt.Transaction().Void(transaction.ID)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		if transaction.Status != TransactionStatusVoided {
			t.Errorf("transaction.Status: expected %s, got %s", TransactionStatusVoided, transaction.Status)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Transaction().Void("nonExisting"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}
