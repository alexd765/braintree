package braintree

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
)

// APIError is an error returned from braintree for calls.
type APIError struct {
	Attribute string `xml:"attribute"`
	Code      int    `xml:"code"`
	Message   string `xml:"message"`
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func parseError(resp *http.Response) error {

	if resp.StatusCode != http.StatusUnprocessableEntity {
		return errors.New(resp.Status)
	}

	apiErr := struct {
		XMLName     xml.Name   `xml:"api-error-response"`
		Address     []APIError `xml:"errors>address>errors>error"`
		Customer    []APIError `xml:"errors>customer>errors>error"`
		Transaction []APIError `xml:"errors>transaction>errors>error"`
	}{}

	if err := xml.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
		return err
	}

	errs := append(apiErr.Address, apiErr.Customer...)
	errs = append(errs, apiErr.Transaction...)
	if len(errs) == 0 {
		// Return status code for errors we aren't handling yet.
		return errors.New(resp.Status)
	}

	return &errs[0]
}
