package braintree

import (
	"encoding/xml"
	"net/http"
	"time"
)

// Address is a braintree address.
type Address struct {
	Company            string    `xml:"company"`
	CountryCodeAlpha2  string    `xml:"country-code-alpha2"`
	CountryCodeAlpha3  string    `xml:"country-code-alpha3"`
	CountryCodeNumeric string    `xml:"country-code-numeric"`
	CountryName        string    `xml:"country-name"`
	CreatedAt          time.Time `xml:"created-at"`
	CustomerID         string    `xml:"customer-id"`
	ExtendedAddress    string    `xml:"extended-address"`
	FirstName          string    `xml:"first-name"`
	ID                 string    `xml:"id"`
	LastName           string    `xml:"last-name"`
	Locality           string    `xml:"locality"`
	PostalCode         string    `xml:"postal-code"`
	Region             string    `xml:"region"`
	StreetAddress      string    `xml:"street-address"`
	UpdatedAt          time.Time `xml:"updated-at"`
}

// AddressInput is used to create or update an address on braintree.
type AddressInput struct {
	XMLName            xml.Name
	Company            string `xml:"company,omitempty"`
	CountryCodeAlpha2  string `xml:"country-code-alpha2,omitempty"`
	CountryCodeAlpha3  string `xml:"country-code-alpha3,omitempty"`
	CountryCodeNumeric string `xml:"country-code-numeric,omitempty"`
	CountryName        string `xml:"country-name,omitempty"`
	ExtendedAddress    string `xml:"extended-address,omitempty"`
	FirstName          string `xml:"first-name,omitempty"`
	ID                 string `xml:"id,omitempty,omitempty"`
	LastName           string `xml:"last-name,omitempty"`
	Locality           string `xml:"locality,omitempty"`
	PostalCode         string `xml:"postal-code,omitempty"`
	Region             string `xml:"region,omitempty"`
	StreetAddress      string `xml:"street-address,omitempty"`
}

// AddressGW is an Address Gateway.
type AddressGW struct {
	bt *Braintree
}

// Create an address on braintree.
func (agw AddressGW) Create(customerID string, address AddressInput) (*Address, error) {
	address.XMLName = xml.Name{Local: "address"}
	resp := &Address{}
	if err := agw.bt.execute(http.MethodPost, "customers/"+customerID+"/addresses", resp, address); err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete an address on braintree.
func (agw AddressGW) Delete(customerID, addressID string) error {
	return agw.bt.execute(http.MethodDelete, "customers/"+customerID+"/addresses/"+addressID, nil, nil)
}

// Find gets a specific address for a customer.
func (agw AddressGW) Find(customerID, addressID string) (*Address, error) {
	address := &Address{}
	if err := agw.bt.execute(http.MethodGet, "customers/"+customerID+"/addresses/"+addressID, address, nil); err != nil {
		return nil, err
	}
	return address, nil
}

// Update an address in braintree.
func (agw AddressGW) Update(customerID, addressID string, address AddressInput) (*Address, error) {
	address.XMLName = xml.Name{Local: "address"}
	resp := &Address{}
	if err := agw.bt.execute(http.MethodPut, "customers/"+customerID+"/addresses/"+addressID, resp, address); err != nil {
		return nil, err
	}
	return resp, nil
}
