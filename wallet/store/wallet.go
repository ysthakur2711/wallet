package store

import (
	"context"
	"errors"
	"fmt"
	"time"
	"wallet/model"
)

type WalletRepo interface {
	CreateWallet(ctx context.Context, arg CreateWalletParams) (model.Wallet, error)
	GetWalletByUsername(ctx context.Context, username string) (model.Wallet, error)
	GetWalletByAddress(ctx context.Context, address string) (model.Wallet, error)
	UpdateWalletStatus(ctx context.Context, arg UpdateWalletStatusParams) (model.Wallet, error)
	SendMoney(ctx context.Context, arg SendMoneyParams) (WalletTransferResult, error)
	AddWalletBalance(ctx context.Context, params AddWalletBalanceParams) (model.Wallet, error)
}

type walletRepository struct {
	db        *gorm.DB
	transRepo TransRepo
}

func NewWalletRepo(client *gorm.DB, transferRepo TransRepo) WalletRepo {
	return &walletRepository{
		db:        client,
		transRepo: transferRepo,
	}
}

type CreateWalletParams struct {
	Username string `json:"username"`
	Currency string `json:"currency"`
}

func (q *walletRepository) CreateWallet(ctx context.Context, arg CreateWalletParams) (model.Wallet, error) {
	logrus.Println("log  CreateWallet in store/wallet/CreateWallet ")
	var w model.Wallet

	//// CHECK IF THERE  EXIST A USER OR NOT WITH THIS USERNAME
	var u model.User
	res1 := q.db.Where("username = ?", arg.Username).Take(&u)
	// SELECT * FROM users WHERE name = "jinzhu";

	// check error ErrRecordNotFound
	if errors.Is(res1.Error, gorm.ErrRecordNotFound) {
		logrus.Println("No user exist with this !! ")
		return w, res1.Error
	}
	wa, err := uuid.NewV4()
	if err != nil {
		logrus.Println("error in creating new uuid for wa(wallet_address) !!")
		return w, err
	}
	w = model.Wallet{
		Username:      arg.Username,
		WalletAddress: wa.String(),
		Status:        model.WalletStatusACTIVE,
		Balance:       0,
		Currency:      arg.Currency,
		CreatedAt:     time.Now(),
	}
	res := q.db.Create(&w) // pass pointer of data to Create

	if res.Error != nil {
		return w, fmt.Errorf("Something wrong happend could not create entry in DB")
	}

	return w, nil
}

func (q *walletRepository) GetWalletByUsername(ctx context.Context, username string) (model.Wallet, error) {

	logrus.Println("log  GetWallet in store/wallet/GetWallet")

	var i model.Wallet
	res := q.db.Where("username = ?", username).Take(&i)
	// SELECT * FROM wallets WHERE username = "jinzhu";

	// check error ErrRecordNotFound
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		logrus.Println("wallet with this id not found !! ")
		return i, res.Error
	}

	return i, nil
}
func (q *walletRepository) GetWalletByAddress(ctx context.Context, address string) (model.Wallet, error) {

	logrus.Println("log  GetWalletByAddress in store/wallet/GetWalletByAddress")

	var i model.Wallet

	res := q.db.Where("wallet_address = ?", address).Take(&i)
	// SELECT * FROM wallets WHERE wallet_address = "jinzhu";

	// check error ErrRecordNotFound
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		logrus.Println("wallet with this wallet_address not found !! ")
		return i, fmt.Errorf("wallet_address does not exists")
	}
	return i, nil
}

type UpdateWalletStatusParams struct {
	Status model.WalletStatus `json:"status"`
	ID     int64              `json:"id"`
}

func (q *walletRepository) UpdateWalletStatus(ctx context.Context, arg UpdateWalletStatusParams) (model.Wallet, error) {

	logrus.Println("log  UpdateWalletStatus in store/wallet/UpdateWalletStatus ")

	var i model.Wallet
	q.db.Model(&i).Where("id = ?", arg.ID).Update("status", arg.Status)

	return i, nil
}

type WalletTransferResult struct {
	Wallet model.Wallet `json:"wallet"`
	Trans  model.Trans  `json:"trans"`
}

type SendMoneyParams struct {
	FromWalletAddress string `json:"from_wallet_address"`
	ToWalletAddress   string `json:"to_wallet_address"`
	Amount            int64  `json:"amount"`
}

func (q *walletRepository) SendMoney(ctx context.Context, arg SendMoneyParams) (WalletTransferResult, error) {

	logrus.Println("log  SendMoney in store/wallet/SendMoney")

	var res WalletTransferResult

	//create a new transaction and handle rollback/commit based on the
	err := q.db.Transaction(func(tx *gorm.DB) error {

		var err error

		//logrus.Println("log  Tx in store/wallet/SendMoney")

		//TODO: CREATE TRANSFER ENTITY then // STORE ENTRY FOR TRANSFER in CreateTransefer function
		trans, err := q.transRepo.CreateTransfer(ctx, arg)

		if err != nil {
			return err
		}

		res.Trans = trans

		//TODO: UPDATE BALANCE OF ACCOUNTS
		fromWallet, _, err := addMoney(ctx, q, arg.FromWalletAddress, -arg.Amount, arg.ToWalletAddress, arg.Amount)

		res.Wallet = fromWallet

		return err
	})
}

type AddWalletBalanceParams struct {
	WalletAddress string `json:"wallet_address"`
	Amount        int64  `json:"amount"`
}

func (q *walletRepository) AddWalletBalance(ctx context.Context, params AddWalletBalanceParams) (model.Wallet, error) {

	logrus.Println("log  AddWalletBalance in store/wallet/AddWalletBalance ")

	var i model.Wallet

	q.db.Where("wallet_address = ?", params.WalletAddress).Take(&i)
	q.db.Model(&i).Where("wallet_address = ?", params.WalletAddress).Update("balance", i.Balance+params.Amount)

	return i, nil
}

func addMoney(ctx context.Context, q *walletRepository, address1 string, amount1 int64, address2 string, amount2 int64) (wallet1 model.Wallet, wallet2 model.Wallet, err error) {
	wallet1, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
		WalletAddress: address1,
		Amount:        amount1,
	})

	if err != nil {
		return
	}

	wallet2, err = q.AddWalletBalance(ctx, AddWalletBalanceParams{
		WalletAddress: address2,
		Amount:        amount2,
	})

	if err != nil {
		return
	}

	return
}
