// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	domain "github.com/danisbagus/matchoshop/internal/core/domain"

	mock "github.com/stretchr/testify/mock"
)

// OrderRepo is an autogenerated mock type for the OrderRepo type
type OrderRepo struct {
	mock.Mock
}

// GetAllByUserID provides a mock function with given fields: userID
func (_m *OrderRepo) GetAllByUserID(userID int64) ([]domain.OrderDetail, *errs.AppError) {
	ret := _m.Called(userID)

	var r0 []domain.OrderDetail
	if rf, ok := ret.Get(0).(func(int64) []domain.OrderDetail); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.OrderDetail)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(userID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetOneByID provides a mock function with given fields: ID
func (_m *OrderRepo) GetOneByID(ID int64) (*domain.OrderDetail, *errs.AppError) {
	ret := _m.Called(ID)

	var r0 *domain.OrderDetail
	if rf, ok := ret.Get(0).(func(int64) *domain.OrderDetail); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.OrderDetail)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(ID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Insert provides a mock function with given fields: form
func (_m *OrderRepo) Insert(form *domain.OrderDetail) (int64, *errs.AppError) {
	ret := _m.Called(form)

	var r0 int64
	if rf, ok := ret.Get(0).(func(*domain.OrderDetail) int64); ok {
		r0 = rf(form)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(*domain.OrderDetail) *errs.AppError); ok {
		r1 = rf(form)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// UpdatePaid provides a mock function with given fields: form
func (_m *OrderRepo) UpdatePaid(form *domain.PaymentResult) *errs.AppError {
	ret := _m.Called(form)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(*domain.PaymentResult) *errs.AppError); ok {
		r0 = rf(form)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}
