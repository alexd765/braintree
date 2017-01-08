package braintree

import (
	"os"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAllPlans(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()
		plans, err := bt.Plan().All()
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		if size := len(plans); size != 1 {
			t.Fatalf("expected %d plans, got %d", 1, size)
		}
		if plans[0].ID != "plan1" {
			t.Errorf("expected plan.ID == plan1, got plan.ID == %s", plans[0].ID)
		}
		if !plans[0].Price.Equals(decimal.NewFromFloat(5)) {
			t.Errorf("plan.Price: expected 5, got %s", plans[0].Price)
		}
	})

	t.Run("wrongMerchantID", func(t *testing.T) {
		t.Parallel()
		bt2 := &Braintree{
			environment: "sandbox",
			merchantID:  "blablub",
			publicKey:   os.Getenv("BRAINTREE_PUB_KEY"),
			privateKey:  os.Getenv("BRAINTREE_PRIV_KEY"),
		}
		bt2.planGW = PlanGW{bt: bt2}
		if _, err := bt2.Plan().All(); err == nil || err.Error() != "401 Unauthorized" {
			t.Errorf("got: %v, want: 401: Unauthorized", err)
		}
	})
}
