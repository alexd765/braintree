package braintree

import (
	"encoding/xml"
	"net/http"
	"time"
)

// Address is a braintree address
type Address struct {
	Company            string    `xml:"company,omitempty"`
	CountryCodeAlpha2  string    `xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string    `xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string    `xml:"country-code-numeric,omitempty"`
	CountryName        string    `xml:"country-name,omitempty"`
	CreatedAt          time.Time `xml:"created-at"`
	CustomerID         string    `xml:"customer-id,omitempty"`
	ExtendedAddress    string    `xml:"extended-address,omitempty"`
	FirstName          string    `xml:"first-name,omitempty"`
	ID                 string    `xml:"id,omitempty"`
	LastName           string    `xml:"last-name,omitempty"`
	Locality           string    `xml:"locality,omitempty"`
	PostalCode         string    `xml:"postal-code,omitempty"`
	Region             string    `xml:"region,omitempty"`
	StreetAddress      string    `xml:"street-address,omitempty"`
	UpdatedAt          time.Time `xml:"updated-at"`
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

// Delete an address on braintree
func (agw AddressGW) Delete(customerID, addressID string) error {
	return agw.bt.execute(http.MethodDelete, "customers/"+customerID+"/addresses/"+addressID, nil, nil)
}

// Find gets a specific address for a customer
func (agw AddressGW) Find(customerID, addressID string) (*Address, error) {
	address := &Address{}
	if err := agw.bt.execute(http.MethodGet, "customers/"+customerID+"/addresses/"+addressID, address, nil); err != nil {
		return nil, err
	}
	return address, nil
}

type addressSanitized struct {
	XMLName            xml.Name
	Company            string `xml:"company,omitempty"`
	CountryCodeAlpha2  string `xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string `xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string `xml:"country-code-numeric,omitempty"`
	CountryName        string `xml:"country-name,omitempty"`
	ExtendedAddress    string `xml:"extended-address,omitempty"`
	FirstName          string `xml:"first-name,omitempty"`
	ID                 string `xml:"id,omitempty"`
	LastName           string `xml:"last-name,omitempty"`
	Locality           string `xml:"locality,omitempty"`
	PostalCode         string `xml:"postal-code,omitempty"`
	Region             string `xml:"region,omitempty"`
	StreetAddress      string `xml:"street-address,omitempty"`
}

// sanitize returns an address without CreatedAt, CustomerID, UpdatedAt
func (a Address) sanitize() addressSanitized {
	return addressSanitized{
		XMLName:            xml.Name{Local: "address"},
		Company:            a.Company,
		CountryCodeAlpha2:  a.CountryCodeAlpha2,
		CountryCodeAlpha3:  a.CountryCodeAlpha3,
		CountryCodeNumeric: a.CountryCodeNumeric,
		CountryName:        a.CountryName,
		ExtendedAddress:    a.ExtendedAddress,
		FirstName:          a.FirstName,
		ID:                 a.ID,
		LastName:           a.LastName,
		Locality:           a.Locality,
		PostalCode:         a.PostalCode,
		Region:             a.Region,
		StreetAddress:      a.StreetAddress,
	}
}
