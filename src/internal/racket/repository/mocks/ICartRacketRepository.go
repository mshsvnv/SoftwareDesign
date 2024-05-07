// Code generated by mockery v2.42.1. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// ICartRacketRepository is an autogenerated mock type for the ICartRacketRepository type
type ICartRacketRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, id
func (_m *ICartRacketRepository) Delete(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewICartRacketRepository creates a new instance of ICartRacketRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewICartRacketRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ICartRacketRepository {
	mock := &ICartRacketRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}