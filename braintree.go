package braintree

import (
	"errors"
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

	addressGW       AddressGW
	clientTokenGW   ClientTokenGW
	customerGW      CustomerGW
	paymentMethodGW PaymentMethodGW
	planGW          PlanGW
	subscriptionGW  SubscriptionGW
	transactionGW   TransactionGW
}

// New returns a braintree client with credentials from env.
//
// BRAINTREE_MERCHANT_ID, BRAINTREE_PUBLIC_KEY and
// BRAINTREE_PRIVATE_KEY have to be set.
func New() (*Braintree, error) {

	bt := &Braintree{
		environment: "sandbox",
		merchantID:  os.Getenv("BRAINTREE_MERCHANT_ID"),
		publicKey:   os.Getenv("BRAINTREE_PUBLIC_KEY"),
		privateKey:  os.Getenv("BRAINTREE_PRIVATE_KEY"),
	}
	bt.addressGW = AddressGW{bt: bt}
	bt.clientTokenGW = ClientTokenGW{bt: bt}
	bt.customerGW = CustomerGW{bt: bt}
	bt.paymentMethodGW = PaymentMethodGW{bt: bt}
	bt.planGW = PlanGW{bt: bt}
	bt.subscriptionGW = SubscriptionGW{bt: bt}
	bt.transactionGW = TransactionGW{bt: bt}

	if bt.merchantID == "" {
		return nil, errors.New("env BRAINTREE_MERCHANT_ID not set")
	}
	if bt.publicKey == "" {
		return nil, errors.New("env BRAINTREE_PUBLIC_KEY not set")
	}
	if bt.privateKey == "" {
		return nil, errors.New("env BRAINTREE_PRIVATE_KEY not set")
	}

	return bt, nil
}

// Address provides the address gateway for this braintree client.
func (bt *Braintree) Address() AddressGW {
	return bt.addressGW
}

// ClientToken provides the client token gateway for this braintree client.
func (bt *Braintree) ClientToken() ClientTokenGW {
	return bt.clientTokenGW
}

// Customer provides the customer gateway for this braintree client.
func (bt *Braintree) Customer() CustomerGW {
	return bt.customerGW
}

// PaymentMethod provides the payment method gateway for this braintree client.
func (bt *Braintree) PaymentMethod() PaymentMethodGW {
	return bt.paymentMethodGW
}

// Plan provides the plan gateway for this braintree client.
func (bt *Braintree) Plan() PlanGW {
	return bt.planGW
}

// Subscription provides the subscription gateway for this braintree client.
func (bt *Braintree) Subscription() SubscriptionGW {
	return bt.subscriptionGW
}

// Transaction provides the transaction gateway for this braintree client.
func (bt *Braintree) Transaction() TransactionGW {
	return bt.transactionGW
}
