package braintree

import (
	"time"
)

//Kind of webhooks.
const (
	WebhookAccountUpdaterDailyReport = "account_updater_daily_report"

	WebhookCheck = "check"

	WebhookDisbursement          = "disbursement"
	WebhookDisbursementException = "disbursement_exception"

	WebhookDisputeLost   = "dispute_lost"
	WebhookDisputeOpened = "dispute_opened"
	WebhookDisputeWon    = "dispute_won"

	WebhookSubscriptionCanceled              = "subscription_canceled"
	WebhookSubscriptionChargedSuccessfully   = "subscription_charged_successfully"
	WebhookSubscriptionChargedUnsuccessfully = "subscription_charged_unsuccessfully"
	WebhookSubscriptionExpired               = "subscription_expired"
	WebhookSubscriptionTrialEnded            = "subscription_trial_ended"
	WebhookSubscriptionWentActive            = "subscription_went_active"
	WebhookSubscriptionWentPastDue           = "subscription_went_past_due"

	WebhookSubMerchantAccountApproved = "sub_merchant_account_approved"
	WebhookSubMerchantAccountDeclined = "sub_merchant_account_declined"

	WebhookPartnerMerchantAccountConnected    = "partner_merchant_account_connected"
	WebhookPartnerMerchantAccountDisconnected = "partner_merchant_account_disconnected"
	WebhookPartnerMerchantAccountDeclined     = "partner_merchant_account_declined"

	WebhookTransactionSettled            = "transaction_settled"
	WebhookTransactionSettlementDeclined = "transaction_settlement_declined"
	WebhookTransactionDispursed          = "transaction_dispursed"
)

// WebhookNotification is an automated notifications from braintree via webhook.
type WebhookNotification struct {
	Kind      string          `xml:"kind"`
	Subject   *WebhookSubject `xml:"subject"`
	Timestamp time.Time       `xml:"timestamp"`
}

// WebhookSubject will be implemented later.
type WebhookSubject interface {
	privateWebhook()
}
