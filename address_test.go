package braintree

import (
	"reflect"
	"testing"
)

func TestCreateAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		customer, err := bt.Customer().Create(CustomerInput{FirstName: "test", LastName: "create address"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address := &Address{CustomerID: customer.ID, StreetAddress: "street"}
		got, err := bt.Address().Create(address)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		address.ID = got.ID
		address.CreatedAt = got.CreatedAt
		address.UpdatedAt = got.UpdatedAt
		if !reflect.DeepEqual(got, address) {
			t.Errorf("got: %+v\nwant: %+v", got, address)
		}
	})

	t.Run("withoutID", func(t *testing.T) {
		t.Parallel()

		address := &Address{StreetAddress: "street"}
		if _, err := bt.Address().Create(address); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestDeleteAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()
		customer, err := bt.Customer().Create(CustomerInput{FirstName: "test", LastName: "create address"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address := &Address{CustomerID: customer.ID, StreetAddress: "street"}
		address, err = bt.Address().Create(address)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		if err := bt.Address().Delete(customer.ID, address.ID); err != nil {
			t.Errorf("unexpected err: %s", err)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if err := bt.Address().Delete("cus1", "bla"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestFindAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()

		got, err := bt.Address().Find("cus1", "qd")
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		want := &Address{CustomerID: "cus1", ID: "qd", FirstName: "first", LastName: "last"}
		want.CreatedAt = got.CreatedAt
		want.UpdatedAt = got.UpdatedAt
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %+v\nwant: %+v", got, want)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		if _, err := bt.Address().Find("cus1", "bla"); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}

func TestUpdateAddress(t *testing.T) {
	t.Parallel()

	t.Run("shouldWork", func(t *testing.T) {
		t.Parallel()
		customer, err := bt.Customer().Create(CustomerInput{FirstName: "test", LastName: "update address"})
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		address := &Address{CustomerID: customer.ID, StreetAddress: "street"}
		want, err := bt.Address().Create(address)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}

		want.FirstName = "test2"
		got, err := bt.Address().Update(want)
		if err != nil {
			t.Fatalf("unexpected err: %s", err)
		}
		want.CreatedAt = got.CreatedAt
		want.UpdatedAt = got.UpdatedAt
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %+v\nwant: %+v", got, want)
		}
	})

	t.Run("nonExisting", func(t *testing.T) {
		t.Parallel()

		address := &Address{CustomerID: "cus1", ID: "bla", FirstName: "first"}
		if _, err := bt.Address().Update(address); err == nil || err.Error() != "404 Not Found" {
			t.Errorf("got: %v, want: 404 Not Found", err)
		}
	})
}
