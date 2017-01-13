package braintree

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/shopspring/decimal"
)

// Payment instrument types of a transaction.
const (
	PaymentInstrumentAndroidPayCard = "android_pay_card"
	PaymentInstrumentApplePayCard   = "apple_pay_card"
	PaymentInstrumentCreditCard     = "credit_card"
	PaymentInstrumentPaypalAccount  = "paypal_account"
	PaymentInstrumentVenmoAccount   = "venmo_account"
)

// Status types of a transaction.
const (
	StatusAuthorisationExpired   = "authorisation_expired"
	StatusAuthorized             = "authorized"
	StatusAuthorizing            = "authorizing"
	StatusSettlementPending      = "settlement_pending"
	StatusSettlementConfirmed    = "settlement_confirmed"
	StatusSettlementDeclined     = "settlement_declined"
	StatusFailed                 = "failed"
	StatusGatewayRejected        = "gateway_reject"
	StatusProcessorDeclined      = "status_processor_declined"
	StatusSettled                = "settled"
	StatusSettling               = "settling"
	StatusSubmittedForSettlement = "submitted_for_settlement"
	StatusVoided                 = "status_voided"
)

// Types of a transaction.
const (
	TransactionTypeSale   = "sale"
	TransactionTypeCredit = "credit"
)

// A Transaction on braintree.
type Transaction struct {
	// AddOns
	AdditionalProcessorResponse string          `xml:"additional-processor-response"`
	Amount                      decimal.Decimal `xml:"amount"`
	// AndroidPayCard
	// ApplePayDetails
	// AVSErrorResponseCode
	// AVSPostalCodeResponseCode
	// AVSStreetAddressResponseCode
	BillingDetails Address   `xml:"billing-details"`
	Channel        string    `xml:"channel"`
	CreatedAt      time.Time `xml:"created-at"`
	// CreditCardDetails
	CurrencyISOCode string
	CustomFields    CustomFields `xml:"custom-fields"`
	// CustomerDetails
	// CVVResponseCode
	// Descriptor
	// DisbursementDetails
	// Discounts
	// Disputes
	EscrowStatus string `xml:"escrow-status"`
	// GatewayRejectionReason
	ID                    string `xml:"id"`
	MerchantAccountID     string `xml:"merchant-account-id"`
	OrderID               string `xml:"order-id"`
	PaymentInstrumentType string `xml:"payment-instrument-type"`
	// PaypalDetails
	PlanID string `xml:"plan-id"`
	// ProcessorAuthoriationCode
	ProcessorResponseCode           string `xml:"processor-response-code"`
	ProcessorResponseText           string `xml:"processor-response-text"`
	ProcessorSettlementResponseCode string `xml:"processor-settlement-response-code"`
	ProcessorSettlementResponseText string `xml:"processor-settlement-response-text"`
	PurchaseOrderNumber             string `xml:"purchase-order-number"`
	// Recurring
	// RefundIDs
	RefundedTransactionID string `xml:"refunded-transaction-id"`
	// RiskData
	// ServiceFeeAmount`
	SettlementBatchID string `xml:"settlement-batch-id"`
	// ShippingDetails
	Status        string   `xml:"status"`
	StatusHistory []string `xml:"status-history"`
	// SubscriptionDetails
	SubscriptionID string `xml:"subscription-id"`
	// TaxAmount
	TaxExempt bool `xml:"tax-exempt"`
	// ThreeDSecureInfo
	Type      string    `xml:"type"`
	UpdatedAt time.Time `xml:"updated-at"`
	// VemnoAccount
	// VoiceRefferalNumber
}

// TransactionInput is used to do a sale.
type TransactionInput struct {
	XMLName          xml.Name
	Amount           decimal.Decimal `xml:"amount"`
	Billing          *AddressInput   `xml:"billing,omitempty"`
	BillingAddressID string          `xml:"billing-address-id,omitempty"`
	Channel          string          `xml:"channel,omitempty"`
	CustomFields     CustomFields    `xml:"custom-fields,omitempty"`
	Customer         *CustomerInput  `xml:"customer,omitempty"`
	CustomerID       string          `xml:"customer-id,omitempty"`
	// Descriptor
	DeviceData        string `xml:"device-date,omitempty"`
	DeviceSessionID   string `xml:"device-session-id,omitempty"`
	MerchantAccountID string `xml:"merchant-account-id,omitempty"`
	// Options
	OrderID             string `xml:"order-id,omitempty"`
	PaymentMethodNonce  string `xml:"payment-method-nonce,omitempty"`
	PaymentMethodToken  string `xml:"payment-method_token,omitempty"`
	PurchaseOrderNumber string `xml:"purchase-order-number,omitempty"`
	Recurring           bool   `xml:"recurring,omitempty"`
	// RiskData
	// ServiceFeeAmount
	Shipping          *AddressInput `xml:"shipping,omitempty"`
	ShippingAddressID string        `xml:"shipping-address-id,omitempty"`
	// TaxAmount
	TaxExempt bool `xml:"tax-exempt,omitempty"`
	// ThreeDSecurePassThru
	TransactionSource string `xml:"transaction-source,omitempty"`
	Type              string `xml:"type,omitempty"`
}

// TransactionGW is a transaction gateway.
type TransactionGW struct {
	bt *Braintree
}

// Create a transaction on braintree.
//
// One of PaymentMethodNonce or PaymentMethodToken is required.
// Amount and Type are required.
func (tgw TransactionGW) Create(transaction TransactionInput) (*Transaction, error) {
	transaction.XMLName = xml.Name{Local: "transaction"}
	resp := &Transaction{}
	if err := tgw.bt.execute(http.MethodPost, "transactions", resp, transaction); err != nil {
		return nil, err
	}
	return resp, nil
}

// Find a transaction with a given transaction id on braintree.
func (tgw TransactionGW) Find(id string) (*Transaction, error) {
	transaction := &Transaction{}
	if err := tgw.bt.execute(http.MethodGet, "transactions/"+id, transaction, nil); err != nil {
		return nil, err
	}
	return transaction, nil
}
