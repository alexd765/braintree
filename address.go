package braintree

import (
	"net/http"
	"time"
)

// Address is a braintree address
type Address struct {
	Company            string     `xml:"company,omitempty"`
	CountryCodeAlpha2  string     `xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string     `xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string     `xml:"country-code-numeric,omitempty"`
	CountryName        string     `xml:"country-name,omitempty"`
	CreatedAt          *time.Time `xml:"created-at,omitempty"`
	CustomerID         string     `xml:"customer-id,omitempty"`
	ExtendedAddress    string     `xml:"extended-address,omitempty"`
	FirstName          string     `xml:"first-name,omitempty"`
	ID                 string     `xml:"id,omitempty"`
	LastName           string     `xml:"last-name,omitempty"`
	Locality           string     `xml:"locality,omitempty"`
	PostalCode         string     `xml:"postal-code,omitempty"`
	Region             string     `xml:"region,omitempty"`
	StreetAddress      string     `xml:"street-address,omitempty"`
	UpdatedAt          *time.Time `xml:"updated-at,omitempty"`
}

// AddressGW is an Address Gateway
type AddressGW struct {
	bt *Braintree
}

// Create an address on braintree.
// CustomerID is required.
func (agw AddressGW) Create(address *Address) (*Address, error) {

	updated := &Address{}
	if err := agw.bt.execute(http.MethodPost, "customers/"+address.CustomerID+"/addresses", updated, address.sanitize()); err != nil {
		return nil, err
	}
	return updated, nil
}

func (a Address) sanitize() Address {
	a.CreatedAt = nil
	a.CustomerID = ""
	a.UpdatedAt = nil
	return a
}
