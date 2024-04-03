// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"
	dto "src/internal/payment/dto"

	mock "github.com/stretchr/testify/mock"

	model "src/internal/payment/model"
)

// IPaymentRepository is an autogenerated mock type for the IPaymentRepository type
type IPaymentRepository struct {
	mock.Mock
}

// CreatePayment provides a mock function with given fields: ctx, Payment
func (_m *IPaymentRepository) CreatePayment(ctx context.Context, Payment *model.Payment) error {
	ret := _m.Called(ctx, Payment)

	if len(ret) == 0 {
		panic("no return value specified for CreatePayment")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Payment) error); ok {
		r0 = rf(ctx, Payment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMyPayments provides a mock function with given fields: ctx, req
func (_m *IPaymentRepository) GetMyPayments(ctx context.Context, req *dto.ListPaymentReq) ([]*model.Payment, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetMyPayments")
	}

	var r0 []*model.Payment
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListPaymentReq) ([]*model.Payment, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *dto.ListPaymentReq) []*model.Payment); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Payment)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *dto.ListPaymentReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPaymentByOrderID provides a mock function with given fields: ctx, orderID
func (_m *IPaymentRepository) GetPaymentByOrderID(ctx context.Context, orderID string) (*model.Payment, error) {
	ret := _m.Called(ctx, orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetPaymentByOrderID")
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

// UpdatePayment provides a mock function with given fields: ctx, Payment
func (_m *IPaymentRepository) UpdatePayment(ctx context.Context, Payment *model.Payment) error {
	ret := _m.Called(ctx, Payment)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePayment")
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
