package braintree

import (
	"encoding/xml"
	"net/http"
	"time"
)

//Customer is a braintree customer.
type Customer struct {
	Addresses []Address `xml:"addresses>address"`
	// AndroidPayCards
	// ApplePayCards
	Company      string       `xml:"company"`
	CreatedAt    time.Time    `xml:"created-at"`
	CreditCards  []CreditCard `xml:"credit-cards>credit-card"`
	CustomFields CustomFields `xml:"custom-fields"`
	Email        string       `xml:"email"`
	Fax          string       `xml:"fax"`
	FirstName    string       `xml:"first-name"`
	ID           string       `xml:"id"`
	LastName     string       `xml:"last-name"`
	// PaymentMethods
	// PaypalAccounts
	Phone     string    `xml:"phone"`
	UpdatedAt time.Time `xml:"updated-at"`
	Website   string    `xml:"website"`
}

// CustomerInput is used to create or update a customer.
type CustomerInput struct {
	XMLName            xml.Name
	Company            string           `xml:"company,omitempty"`
	CreditCard         *CreditCardInput `xml:"credit-card,omitempty"`
	CustomFields       CustomFields     `xml:"custom-fields,omitempty"`
	Email              string           `xml:"email,omitempty"`
	Fax                string           `xml:"fax,omitempty"`
	FirstName          string           `xml:"first-name,omitempty"`
	ID                 string           `xml:"id,omitempty"`
	LastName           string           `xml:"last-name,omitempty"`
	PaymentMethodNonce string           `xml:"payment-method-nonce,omitempty"`
	Phone              string           `xml:"phone,omitempty"`
	RiskData           *RiskData        `xml:"risk-data,omitempty"`
	Website            string           `xml:"website,omitempty"`
}

// CustomerGW is a Customer Gateway.
type CustomerGW struct {
	bt *Braintree
}

// Create a Customer on braintree.
func (cgw CustomerGW) Create(customer CustomerInput) (*Customer, error) {
	customer.XMLName = xml.Name{Local: "customer"}
	resp := &Customer{}
	if err := cgw.bt.execute(http.MethodPost, "customers", resp, customer); err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete a Customer on braintree.
func (cgw CustomerGW) Delete(id string) error {
	return cgw.bt.execute(http.MethodDelete, "customers/"+id, nil, nil)
}

// Find a Customer with a given id on braintree.
func (cgw CustomerGW) Find(id string) (*Customer, error) {
	customer := &Customer{}
	if err := cgw.bt.execute(http.MethodGet, "customers/"+id, customer, nil); err != nil {
		return nil, err
	}
	return customer, nil
}

// Update a Customer on braintree.
//
// ID is required.
func (cgw CustomerGW) Update(customer CustomerInput) (*Customer, error) {
	customer.XMLName = xml.Name{Local: "customer"}
	updatedCustomer := &Customer{}
	if err := cgw.bt.execute(http.MethodPut, "customers/"+customer.ID, updatedCustomer, customer); err != nil {
		return nil, err
	}
	return updatedCustomer, nil
}
