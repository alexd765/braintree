package braintree

import (
	"bytes"
	"encoding/xml"
	"net/http"
)

func (bt *Braintree) execute(method, path string, v interface{}, payload interface{}) error {

	url := bt.baseURL() + path
	buf := new(bytes.Buffer)
	if err := xml.NewEncoder(buf).Encode(payload); err != nil {
		return err
	}

	if bt.Logger != nil {
		bt.Logger.Printf(">>> %s %s with payload: %s\n", method, url, buf)
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return err
	}
	req.Header.Set("X-ApiVersion", "4")
	req.Header.Set("Content-Type", "application/xml")
	req.SetBasicAuth(bt.publicKey, bt.privateKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	switch resp.StatusCode {

	case http.StatusOK, http.StatusCreated:
		if v == nil {
			return nil
		}
		return xml.NewDecoder(resp.Body).Decode(v)

	default:
		return parseError(resp)
	}
}

func (bt *Braintree) baseURL() string {
	if bt.environment == EnvironmentProduction {
		return "https://www.braintreegateway.com/merchants/" + bt.merchantID + "/"
	}
	return "https://sandbox.braintreegateway.com/merchants/" + bt.merchantID + "/"
}
