package dto

import (
	"time"
	"wallet/model"
)

type TransResultDto struct {
	FromWalletAdd string    `json:"from_wallet_address"`
	ToWalletAdd   string    `json:"to_wallet_address"`
	Amount        int64     `json:"amount"`
	CreatedAt     time.Time `json:"created_at"`
}

type CreditDto struct {
	WalletAddress string `json:"wallet_address" binding:"required"`
	Amount        int64  `json:"amount" validate:"required"`
}

func NewTransResultDto(trans model.Trans) TransResultDto {
	return TransResultDto{
		FromWalletAdd: trans.FromWalletAdd,
		ToWalletAdd:   trans.ToWalletAdd,
		Amount:        trans.Amount,
		CreatedAt:     trans.CreatedAt,
	}
}
