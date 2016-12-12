package braintree

import (
	"net/http"
	"time"
)

//Customer is a braintree customer
type Customer struct {
	Addresses []Address `xml:"addresses>address,omitempty"`
	// AndroidPayCards
	// ApplePayCards
	Company   string     `xml:"company,omitempty"`
	CreatedAt *time.Time `xml:"created-at,omitempty"`
	// CreditCards
	CustomFields CustomFields `xml:"custom-fields,omitempty"`
	Email        string       `xml:"email,omitempty"`
	Fax          string       `xml:"fax,omitempty"`
	FirstName    string       `xml:"first-name,omitempty"`
	ID           string       `xml:"id,omitempty"`
	LastName     string       `xml:"last-name,omitempty"`
	// PaymentMethods
	// PaypalAccounts
	Phone     string     `xml:"phone,omitempty"`
	UpdatedAt *time.Time `xml:"updated-at,omitempty"`
	Website   string     `xml:"website,omitempty"`
}

// CustomerGW is a Customer Gateway
type CustomerGW struct {
	bt *Braintree
}

// Create a Customer on braintree
func (cgw CustomerGW) Create(customer *Customer) (*Customer, error) {
	updated := &Customer{}
	if err := cgw.bt.execute(http.MethodPost, "customers", updated, customer.sanitized()); err != nil {
		return nil, err
	}
	return updated, nil
}

// Delete a Customer on braintree
func (cgw CustomerGW) Delete(id string) error {
	return cgw.bt.execute(http.MethodDelete, "customers/"+id, nil, nil)
}

// Find a Customer with a given id on braintree
func (cgw CustomerGW) Find(id string) (*Customer, error) {
	customer := &Customer{}
	if err := cgw.bt.execute(http.MethodGet, "customers/"+id, customer, nil); err != nil {
		return nil, err
	}
	return customer, nil
}

// Update a Customer on braintree
// Only non-empty fields are updated
// ID is required
func (cgw CustomerGW) Update(customer *Customer) (*Customer, error) {
	updatedCustomer := &Customer{}
	if err := cgw.bt.execute(http.MethodPut, "customers/"+customer.ID, updatedCustomer, customer.sanitized()); err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}

func (c Customer) sanitized() Customer {
	c.CreatedAt = nil
	c.UpdatedAt = nil
	return c
}
