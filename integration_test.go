package braintree

import (
	"reflect"
	"testing"
)

func TestIntegration(t *testing.T) {
	t.Parallel()

	cus1 := &Customer{
		Company:      "a",
		CustomFields: CustomFields{"testcustom1": "b"},
		Email:        "ccc@example.org",
		Fax:          "d",
		FirstName:    "e",
		LastName:     "f",
		Phone:        "g",
		Website:      "www.example.org",
	}

	cus2, err := bt.Customer().Create(cus1)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	cus1.ID = cus2.ID
	cus1.CreatedAt = cus2.CreatedAt
	cus1.UpdatedAt = cus2.UpdatedAt
	if !reflect.DeepEqual(cus1, cus2) {
		t.Errorf("got: %+v\nwant: %+v", cus2, cus1)
	}

	cus3, err := bt.Customer().Find(cus2.ID)
	if err != nil {
		t.Fatalf("unexpected err: %s", err)
	}
	cus2.UpdatedAt = cus3.UpdatedAt
	if !reflect.DeepEqual(cus2, cus3) {
		t.Errorf("got: %+v\nwant: %+v", cus3, cus2)
	}

	if err := bt.Customer().Delete(cus3.ID); err != nil {
		t.Fatalf("unexpected err: %s", err)
	}

	if _, err := bt.Customer().Find(cus3.ID); err == nil || err.Error() != "404 Not Found" {
		t.Errorf("got: %v, want: 404 Not Found", err)
	}

}
