package braintree

import (
	"encoding/xml"
	"net/http"
)

// ClientTokenInput is used to generate a client token.
type ClientTokenInput struct {
	XMLName           xml.Name
	CustomerID        *string `xml:"customer-id,omitempty"`
	MerchantAccountID *string `xml:"merchant-account-id,omitempty"`
	// Options
	Version int `xml:"version"`
}

// ClientTokenGW is a ClientToken Gateway.
type ClientTokenGW struct {
	bt *Braintree
}

// Generate a client token.
//
// Version is required.
func (ctgw ClientTokenGW) Generate(ct ClientTokenInput) (string, error) {
	ct.XMLName = xml.Name{Local: "client-token"}
	resp := &struct {
		Value string `xml:"value"`
	}{}
	if err := ctgw.bt.execute(http.MethodPost, "client_token", resp, ct); err != nil {
		return "", err
	}
	return resp.Value, nil
}
