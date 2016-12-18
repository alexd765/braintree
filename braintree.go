package braintree

import (
	"log"
	"os"
)

// Braintree client
type Braintree struct {
	Logger *log.Logger

	environment string
	merchantID  string
	publicKey   string
	privateKey  string
	customerGW  CustomerGW
	addressGW   AddressGW
}

// New returns a braintree client with credentials from env
func New() *Braintree {

	bt := &Braintree{
		environment: "sandbox",
		merchantID:  mustGetenv("BRAINTREE_MERCH_ID"),
		publicKey:   mustGetenv("BRAINTREE_PUB_KEY"),
		privateKey:  mustGetenv("BRAINTREE_PRIV_KEY"),
	}
	bt.customerGW = CustomerGW{bt: bt}
	bt.addressGW = AddressGW{bt: bt}

	return bt
}

// Customer provides the customer gateway for this braintree client
func (bt *Braintree) Customer() CustomerGW {
	return bt.customerGW
}

// Address provides the address gateway for this braintree client
func (bt *Braintree) Address() AddressGW {
	return bt.addressGW
}

func mustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("env: %s not set", key)
	}
	return value
}
