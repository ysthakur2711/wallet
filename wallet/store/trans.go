package store

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"time"
	"wallet/model"
)

type TransRepo interface {
	CreateTransfer(ctx context.Context, arg SendMoneyParams) (model.Trans, error)
}

type transRepository struct {
	db *gorm.DB
}

func NewTransRepo(client *gorm.DB) TransRepo {
	return &transRepository{
		db: client,
	}
}

type CreateTransferParams struct {
	FromWalletAddress string `json:"from_wallet_address"`
	ToWalletAddress   string `json:"to_wallet_address"`
	Amount            int64  `json:"amount"`
}

func (q *transRepository) CreateTransfer(ctx context.Context, arg SendMoneyParams) (model.Trans, error) {

	logrus.Println("log  CreateTransfer in store/trans/CreateTransfer ")
	// STORE ENTRY FOR TRANSFER
	var i model.Trans

	i = model.Trans{
		FromWalletAdd: arg.FromWalletAddress,
		ToWalletAdd:   arg.ToWalletAddress,
		Amount:        arg.Amount,
		CreatedAt:     time.Now(),
	}
	res := q.db.Create(&i) // pass pointer of data to Create

	if res.Error != nil {
		return i, fmt.Errorf("Something wrong happend could not create entry in DB")
	}

	return i, nil
}
