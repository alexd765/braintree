package braintree

import "net/http"

// Address is a braintree address
type Address struct {
	Company            string `xml:"company,omitempty"`
	CountryCodeAlpha2  string `xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string `xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string `xml:"country-code-numeric,omitempty"`
	CountryName        string `xml:"country-name,omitempty"`
	// CreatedAt
	CustomerID      string `xml:"customer-id,omitempty"`
	ExtendedAddress string `xml:"extended-address,omitempty"`
	FirstName       string `xml:"first-name,omitempty"`
	ID              string `xml:"id,omitempty"`
	LastName        string `xml:"last-name,omitempty"`
	Locality        string `xml:"locality,omitempty"`
	PostalCode      string `xml:"postal-code,omitempty"`
	Region          string `xml:"region,omitempty"`
	StreetAddress   string `xml:"street-address,omitempty"`
	// UpdatedAt
}

// CreateAddress creates an address on braintree.
// CustomerID is required.
func (bt *Braintree) CreateAddress(address *Address) (*Address, error) {

	// braintree only wants the customerID in the url and not in the payload
	// workaround:
	customerID := address.CustomerID
	tempAddress := *address
	tempAddress.CustomerID = ""

	updatedAddress := &Address{}
	if err := bt.execute(http.MethodPost, "customers/"+customerID+"/addresses", updatedAddress, &tempAddress); err != nil {
		return nil, err
	}
	return updatedAddress, nil
}
