package mocks

import (
	context "context"
	"github.com/stretchr/testify/mock"
	"wallet/model"
	"wallet/store"
)

// TransRepo is an autogenerated mock type for the TransRepo type
type TransRepo struct {
	mock.Mock
}

// CreateTransfer provides a mock function with given fields: ctx, arg
func (_m *TransRepo) CreateTransfer(ctx context.Context, arg store.CreateTransferParams) (model.Trans, error) {

	ret := _m.Called(ctx, arg)

	var r0 model.Trans
	if rf, ok := ret.Get(0).(func(context.Context, store.CreateTransferParams) model.Trans); ok {
		r0 = rf(ctx, arg)
	} else {
		r0 = ret.Get(0).(model.Trans)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, store.CreateTransferParams) error); ok {
		r1 = rf(ctx, arg)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
