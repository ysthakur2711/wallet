package mocks

import (
	context "context"
	"github.com/stretchr/testify/mock"
	"wallet/model"
	"wallet/store"
)

// WalletRepo is an autogenerated mock type for the WalletRepo type
type WalletRepo struct {
	mock.Mock
}

// CreateWallet provides a mock function with given fields: ctx, arg
func (_m *WalletRepo) CreateWallet(ctx context.Context, arg store.CreateWalletParams) (model.Wallet, error) {
	ret := _m.Called(ctx, arg)

	var r0 model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, store.CreateWalletParams) model.Wallet); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(model.Wallet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.CreateWalletParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1
}

// GetWalletByAddress provides a mock function with given fields: ctx, address
func (_m *WalletRepo) GetWalletByAddress(ctx context.Context, address string) (model.Wallet, error) {

	ret := _m.Called(ctx, address)

	var r0 model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Wallet); ok {
		r0 = rf(ctx, address)
	} else {
		r0 = ret.Get(0).(model.Wallet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, address)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetWalletByUsername provides a mock function with given fields: ctx, username
func (_m *WalletRepo) GetWalletByUsername(ctx context.Context, username string) (model.Wallet, error) {
	ret := _m.Called(ctx, username)

	var r0 model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, string) model.Wallet); ok {
		r0 = rf(ctx, username)
	} else {
		r0 = ret.Get(0).(model.Wallet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, username)
	} else {
		r1 = ret.Error(1)
	}
	return r0, r1

}

// SendMoney provides a mock function with given fields: ctx, arg
func (_m *WalletRepo) SendMoney(ctx context.Context, arg store.SendMoneyParams) (store.WalletTransferResult, error) {
	ret := _m.Called(ctx, arg)

	var r0 store.WalletTransferResult
	if rf, ok := ret.Get(0).(func(context.Context, store.SendMoneyParams) store.WalletTransferResult); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(store.WalletTransferResult)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.SendMoneyParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateWalletStatus provides a mock function with given fields: ctx, arg
func (_m *WalletRepo) UpdateWalletStatus(ctx context.Context, arg store.UpdateWalletStatusParams) (model.Wallet, error) {
	ret := _m.Called(ctx, arg)

	var r0 model.Wallet
	if rf, ok := ret.Get(0).(func(context.Context, store.UpdateWalletStatusParams) model.Wallet); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(model.Wallet)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.UpdateWalletStatusParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
