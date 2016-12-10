package braintree

import "net/http"

//Customer is a braintree customer
type Customer struct {
	// Addresses
	// AndroidPayCards
	// ApplePayCards
	Company string `xml:"company"`
	// CreatedAt `xml:"created-at"`
	// CreditCards
	// CustomFields `xml:"custom-fields"`
	Email     string `xml:"email"`
	Fax       string `xml:"fax"`
	FirstName string `xml:"first-name"`
	ID        string `xml:"id"`
	LastName  string `xml:"last-name"`
	// PaymentMethods
	// PaypalAccounts
	Phone string `xml:"phone"`
	// UpdatedAt `xml:"updated-at"`
	Website string `xml:"website"`
}

// FindCustomer with a given id on braintree
func (bt *Braintree) FindCustomer(id string) (*Customer, error) {
	customer := &Customer{}
	err := bt.execute(http.MethodGet, "customers/"+id, customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
