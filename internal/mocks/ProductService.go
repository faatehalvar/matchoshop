// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	dto "github.com/danisbagus/matchoshop/internal/dto"

	mock "github.com/stretchr/testify/mock"
)

// ProductService is an autogenerated mock type for the ProductService type
type ProductService struct {
	mock.Mock
}

// Create provides a mock function with given fields: data
func (_m *ProductService) Create(data *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(data)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(*dto.CreateProductRequest) *dto.ResponseData); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(*dto.CreateProductRequest) *errs.AppError); ok {
		r1 = rf(data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Delete provides a mock function with given fields: productID
func (_m *ProductService) Delete(productID int64) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(productID)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(int64) *dto.ResponseData); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(productID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetDetail provides a mock function with given fields: productID
func (_m *ProductService) GetDetail(productID int64) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(productID)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(int64) *dto.ResponseData); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(productID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// GetList provides a mock function with given fields:
func (_m *ProductService) GetList() (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called()

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func() *dto.ResponseData); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func() *errs.AppError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: productID, data
func (_m *ProductService) Update(productID int64, data *dto.CreateProductRequest) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(productID, data)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(int64, *dto.CreateProductRequest) *dto.ResponseData); ok {
		r0 = rf(productID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64, *dto.CreateProductRequest) *errs.AppError); ok {
		r1 = rf(productID, data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
