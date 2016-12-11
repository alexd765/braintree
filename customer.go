package braintree

import "net/http"

//Customer is a braintree customer
type Customer struct {
	Addresses []Address `xml:"addresses>address,omitempty"`
	// AndroidPayCards
	// ApplePayCards
	Company string `xml:"company,omitempty"`
	// CreatedAt `xml:"created-at"`
	// CreditCards
	CustomFields CustomFields `xml:"custom-fields,omitempty"`
	Email        string       `xml:"email,omitempty"`
	Fax          string       `xml:"fax,omitempty"`
	FirstName    string       `xml:"first-name,omitempty"`
	ID           string       `xml:"id,omitempty"`
	LastName     string       `xml:"last-name,omitempty"`
	// PaymentMethods
	// PaypalAccounts
	Phone string `xml:"phone,omitempty"`
	// UpdatedAt `xml:"updated-at"`
	Website string `xml:"website,omitempty"`
}

// CreateCustomer on braintree
func (bt *Braintree) CreateCustomer(customer *Customer) (*Customer, error) {
	updatedCustomer := &Customer{}
	if err := bt.execute(http.MethodPost, "customers", updatedCustomer, customer); err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}

// DeleteCustomer on braintree
func (bt *Braintree) DeleteCustomer(id string) error {
	return bt.execute(http.MethodDelete, "customers/"+id, nil, nil)
}

// FindCustomer with a given id on braintree
func (bt *Braintree) FindCustomer(id string) (*Customer, error) {
	customer := &Customer{}
	if err := bt.execute(http.MethodGet, "customers/"+id, customer, nil); err != nil {
		return nil, err
	}
	return customer, nil
}

// UpdateCustomer on braintree
// Only non-empty fields are updated
// ID is required
func (bt *Braintree) UpdateCustomer(customer *Customer) (*Customer, error) {
	updatedCustomer := &Customer{}
	if err := bt.execute(http.MethodPut, "customers/"+customer.ID, updatedCustomer, customer); err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}
