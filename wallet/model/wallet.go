package model

import "time"

type WalletStatus string

const (
	WalletStatusACTIVE   WalletStatus = "ACTIVE"
	WalletStatusINACTIVE WalletStatus = "INACTIVE"
)

type Wallet struct {
	ID            int64        `gorm:"primary_key;AUTO_INCREMENT;not_null" json:"id"`
	Username      string       `json:"username;unique"`
	WalletAddress string       `json:"wallet_address;unique"`
	Status        WalletStatus `json:"status"`
	//UserID               int64        `json:"user_id"`
	Balance   int64     `json:"balance"`
	Currency  string    `json:"currency"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (e *Wallet) IsBalanceSufficient(expectedAmount int64) bool {
	return e.Balance >= expectedAmount
}
