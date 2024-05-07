// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	model "src/internal/order/model"

	mock "github.com/stretchr/testify/mock"
)

// IPaymentRepository is an autogenerated mock type for the IPaymentRepository type
type IPaymentRepository struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, Payment
func (_m *IPaymentRepository) Create(ctx context.Context, Payment *model.Payment) error {
	ret := _m.Called(ctx, Payment)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Payment) error); ok {
		r0 = rf(ctx, Payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetPaymentByID provides a mock function with given fields: ctx, orderID
func (_m *IPaymentRepository) GetPaymentByID(ctx context.Context, orderID string) (*model.Payment, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetPaymentByID")
	}

	var r0 *model.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Payment, error)); ok {
		return rf(ctx, orderID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Payment); ok {
		r0 = rf(ctx, orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, Payment
func (_m *IPaymentRepository) Update(ctx context.Context, Payment *model.Payment) error {
	ret := _m.Called(ctx, Payment)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Payment) error); ok {
		r0 = rf(ctx, Payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewIPaymentRepository creates a new instance of IPaymentRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIPaymentRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IPaymentRepository {
	mock := &IPaymentRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}