package braintree

import (
	"encoding/xml"
	"net/http"
	"time"
)

//Customer is a braintree customer
type Customer struct {
	Addresses []Address `xml:"addresses>address,omitempty"`
	// AndroidPayCards
	// ApplePayCards
	Company      string       `xml:"company,omitempty"`
	CreatedAt    time.Time    `xml:"created-at"`
	CreditCards  []CreditCard `xml:"credit-cards>credit-card,omitempty"`
	CustomFields CustomFields `xml:"custom-fields,omitempty"`
	Email        string       `xml:"email,omitempty"`
	Fax          string       `xml:"fax,omitempty"`
	FirstName    string       `xml:"first-name,omitempty"`
	ID           string       `xml:"id,omitempty"`
	LastName     string       `xml:"last-name,omitempty"`
	// PaymentMethods
	// PaypalAccounts
	Phone     string    `xml:"phone,omitempty"`
	UpdatedAt time.Time `xml:"updated-at"`
	Website   string    `xml:"website,omitempty"`
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

type customerSanitized struct {
	XMLName   xml.Name
	Addresses []Address `xml:"addresses>address,omitempty"`
	// AndroidPayCards
	// ApplePayCards
	Company string `xml:"company,omitempty"`
	// CreditCards
	CustomFields CustomFields `xml:"custom-fields,omitempty"`
	Email        string       `xml:"email,omitempty"`
	Fax          string       `xml:"fax,omitempty"`
	FirstName    string       `xml:"first-name,omitempty"`
	ID           string       `xml:"id,omitempty"`
	LastName     string       `xml:"last-name,omitempty"`
	// PaymentMethods
	// PaypalAccounts
	Phone   string `xml:"phone,omitempty"`
	Website string `xml:"website,omitempty"`
}

// sanitized returns a customer without CreatedAt, UpdatedAt
func (c Customer) sanitized() customerSanitized {
	return customerSanitized{
		XMLName:      xml.Name{Local: "customer"},
		Addresses:    c.Addresses,
		Company:      c.Company,
		CustomFields: c.CustomFields,
		Email:        c.Email,
		Fax:          c.Fax,
		FirstName:    c.FirstName,
		ID:           c.ID,
		LastName:     c.LastName,
		Phone:        c.Phone,
		Website:      c.Website,
	}
}
