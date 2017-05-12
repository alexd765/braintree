package braintree

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

// An APIError is returned from braintree in response to an invalid api call.
type APIError struct {
	Attribute string `xml:"attribute"`
	Code      int    `xml:"code"`
	Message   string `xml:"message"`
}

// A ProcessorError is returned from braintree if a payment failed.
type ProcessorError struct {
	Code    int
	Message string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func (e *ProcessorError) Error() string {
	return fmt.Sprintf("Code %d: %s", e.Code, e.Message)
}

func parseError(resp *http.Response) error {

	if resp.StatusCode != http.StatusUnprocessableEntity {
		return errors.New(resp.Status)
	}

	errs := struct {
		XMLName             xml.Name     `xml:"api-error-response"`
		Address             []APIError   `xml:"errors>address>errors>error"`
		ClientToken         []APIError   `xml:"errors>client-token>errors>error"`
		CreditCard          []APIError   `xml:"errors>credit-card>errors>error"`
		Customer            []APIError   `xml:"errors>customer>errors>error"`
		Subscription        []APIError   `xml:"errors>subscription>errors>error"`
		Transaction         []APIError   `xml:"errors>transaction>errors>error"`
		TransactionResponse *Transaction `xml:"transaction"`
	}{}

	if err := xml.NewDecoder(resp.Body).Decode(&errs); err != nil {
		return err
	}

	// api errors have the highest priority
	apiErrs := append(errs.Address, errs.ClientToken...)
	apiErrs = append(apiErrs, errs.CreditCard...)
	apiErrs = append(apiErrs, errs.Customer...)
	apiErrs = append(apiErrs, errs.Subscription...)
	apiErrs = append(apiErrs, errs.Transaction...)
	if len(apiErrs) > 0 {
		return &apiErrs[0]
	}

	// then we look for processor response errors
	if errs.TransactionResponse != nil && errs.TransactionResponse.ProcessorResponseCode != "" {
		code, err := strconv.Atoi(errs.TransactionResponse.ProcessorResponseCode)
		if err != nil {
			return err
		}
		return &ProcessorError{
			Code:    code,
			Message: errs.TransactionResponse.ProcessorResponseText,
		}
	}

	// return the status for errors we aren't handling yet
	return errors.New(resp.Status)
}
