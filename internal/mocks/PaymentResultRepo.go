// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	mock "github.com/stretchr/testify/mock"
)

// PaymentResultRepo is an autogenerated mock type for the PaymentResultRepo type
type PaymentResultRepo struct {
	mock.Mock
}

// CheckByID provides a mock function with given fields: PaymentResultID
func (_m *PaymentResultRepo) CheckByID(PaymentResultID string) (bool, *errs.AppError) {
	ret := _m.Called(PaymentResultID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(PaymentResultID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(string) *errs.AppError); ok {
		r1 = rf(PaymentResultID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// CheckByOrderIDAndStatus provides a mock function with given fields: OrderID, status
func (_m *PaymentResultRepo) CheckByOrderIDAndStatus(OrderID int64, status string) (bool, *errs.AppError) {
	ret := _m.Called(OrderID, status)

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, string) bool); ok {
		r0 = rf(OrderID, status)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64, string) *errs.AppError); ok {
		r1 = rf(OrderID, status)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
