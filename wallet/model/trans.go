package model

import "time"

type Trans struct {
	ID            int64     `gorm:"primary_key;AUTO_INCREMENT;not_null" json:"id"`
	FromWalletAdd string    `json:"from_wallet_address"`
	ToWalletAdd   string    `json:"to_wallet_address"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}
type PaymentRequestStatus string

const (
	PaymentRequestStatusWAITINGAPPROVAL PaymentRequestStatus = "WAITING_APPROVAL"
	PaymentRequestStatusAPPROVED        PaymentRequestStatus = "APPROVED"
	PaymentRequestStatusREFUSED         PaymentRequestStatus = "REFUSED"
	PaymentRequestStatusPAYMENTSUCCESS  PaymentRequestStatus = "PAYMENT_SUCCESS"
	PaymentRequestStatusPAYMENTFAILED   PaymentRequestStatus = "PAYMENT_FAILED"
)

type PaymentRequest struct {
	ID            int64                `json:"id"`
	FromWalletAdd string               `json:"from_wallet_add"`
	ToWalletAdd   string               `json:"to_wallet_add"`
	Amount        int64                `json:"amount"`
	Status        PaymentRequestStatus `json:"status"`
	CreatedAt     time.Time            `json:"created_at"`
}
