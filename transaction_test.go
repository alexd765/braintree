package braintree

import (
	"errors"
	"testing"

	"github.com/shopspring/decimal"
)

func TestCreateTransaction(t *testing.T) {
	t.Parallel()

	customer := createTestCustomer(t)

	tests := []struct {
		name    string
		txInput TransactionInput
		wantErr error
	}{
		{
			name: "shouldWork",
			txInput: TransactionInput{
				Amount: decimal.NewFromFloat(3),
				Options: &TransactionOptions{
					StoreInVaultOnSuccess: true,
				},
				PaymentMethodToken: customer.CreditCards[0].Token,
				Type:               TransactionTypeSale,
			},
			wantErr: nil,
		},
		{
			name: "withoutToken",
			txInput: TransactionInput{
				Amount: decimal.NewFromFloat(3),
				Type:   TransactionTypeSale,
			},
			wantErr: &ValidationError{"", 91508, "Cannot determine payment method."},
		},
		{
			name: "paymentFailed",
			txInput: TransactionInput{
				Amount: decimal.NewFromFloat(2000),
				Options: &TransactionOptions{
					StoreInVaultOnSuccess: true,
				},
				PaymentMethodToken: customer.CreditCards[0].Token,
				Type:               TransactionTypeSale,
			},
			wantErr: &ProcessorError{2000, "Do Not Honor"},
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := bt.Transaction().Create(test.txInput)
			compareErrors(t, err, test.wantErr)
		})
	}

	t.Run("duplicate", func(t *testing.T) {
		t.Parallel()

		txInput := TransactionInput{
			Amount: decimal.NewFromFloat(3.8),
			Options: &TransactionOptions{
				StoreInVaultOnSuccess: true,
			},
			PaymentMethodToken: customer.CreditCards[0].Token,
			Type:               TransactionTypeSale,
		}

		if _, err := bt.Transaction().Create(txInput); err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		_, err := bt.Transaction().Create(txInput)
		wantErr := &GatewayError{"duplicate"}
		compareErrors(t, err, wantErr)
	})
}

func TestFindTransaction(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		id      string
		wantErr error
	}{
		{name: "existing", id: "bx9a7av8", wantErr: nil},
		{name: "nonExisting", id: "nonExisting", wantErr: errors.New("404 Not Found")},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			_, err := bt.Transaction().Find(test.id)
			compareErrors(t, err, test.wantErr)
		})
	}
}

func TestRefundTransaction(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer := createTestCustomer(t)

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

		customer := createTestCustomer(t)

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

		customer := createTestCustomer(t)

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
