package braintree

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
