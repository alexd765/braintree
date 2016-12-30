package braintree

import "testing"

func TestAllPlans(t *testing.T) {
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
}
