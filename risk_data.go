package braintree

// RiskData can be used as part of some input structs.
type RiskData struct {
	CustomerBrowser string `xml:"customer-browser,omitempty"`
	CustomerIP      string `xml:"customer-ip,omitempty"`
}
