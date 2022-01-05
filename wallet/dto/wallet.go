package dto

import (
	"time"
	"wallet/model"
	"wallet/store"
)

type TransferMoneyDto struct {
	FromWalletAddress string `json:"from_wallet_address" binding:"required"`
	ToWalletAddress   string `json:"to_wallet_address" binding:"required"`
	Amount            int64  `json:"amount" validate:"required"`
}

type UpdatedWalletBalanceDto struct {
	UpdatedBalance int64     `json:"updated_balance" validate:"required" `
	Currency       string    `json:"currency" validate:"required" `
	UpdatedAt      time.Time `json:"updated_at" `
}

type WalletTransferResultDto struct {
	Wallet model.Wallet `json:"wallet" validate:"required"`
	Trans  model.Trans  `json:"trans" validate:"required"`
}

type CreateWalletDto struct {
	Username string `json:"username" validate:"required,alphanum"`
	Currency string `json:"currency" validate:"required" `
}

type WalletDto struct {
	ID            int64              `json:"id" validate:"required" `
	Username      string             `json:"username;unique" validate:"required" `
	WalletAddress string             `json:"address;unique" validate:"required" `
	Status        model.WalletStatus `json:"status" validate:"required" `
	//UserID               int64        `json:"user_id" validate:"required" `
	Balance   int64     `json:"balance" validate:"required" `
	Currency  string    `json:"currency" validate:"required" `
	CreatedAt time.Time `json:"created_at" validate:"required" `
	UpdatedAt time.Time `json:"updated_at" validate:"required" `
}

func NewWalletTransferDto(wtr store.WalletTransferResult) WalletTransferResultDto {
	return WalletTransferResultDto{
		Wallet: wtr.Wallet,
		Trans:  wtr.Trans,
	}
}

func NewWalletDto(wallet model.Wallet) WalletDto {
	return WalletDto{
		ID:            wallet.ID,
		Username:      wallet.Username,
		WalletAddress: wallet.WalletAddress,
		Status:        wallet.Status,
		Balance:       wallet.Balance,
		Currency:      wallet.Currency,
		CreatedAt:     wallet.CreatedAt,
		UpdatedAt:     wallet.UpdatedAt,
	}
}

func NewUpdatedWalletBalanceDto(wallet model.Wallet) UpdatedWalletBalanceDto {
	return UpdatedWalletBalanceDto{
		UpdatedBalance: wallet.Balance,
		Currency:       wallet.Currency,
		UpdatedAt:      wallet.UpdatedAt,
	}
}
