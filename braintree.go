package braintree

import (
	"log"
	"os"
)

// Braintree client
type Braintree struct {
	environment string
	merchantID  string
	publicKey   string
	privateKey  string
}

// New returns a braintree client with credentials from env
func New() *Braintree {
	return &Braintree{
		environment: "sandbox",
		merchantID:  mustGetenv("BRAINTREE_MERCH_ID"),
		publicKey:   mustGetenv("BRAINTREE_PUB_KEY"),
		privateKey:  mustGetenv("BRAINTREE_PRIV_KEY"),
	}
}

func mustGetenv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("env: %s not set", key)
	}
	return value
}
