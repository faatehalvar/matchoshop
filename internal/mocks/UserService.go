// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	errs "github.com/danisbagus/go-common-packages/errs"
	dto "github.com/danisbagus/matchoshop/internal/dto"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Login provides a mock function with given fields: req
func (_m *UserService) Login(req dto.LoginRequest) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(req)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(dto.LoginRequest) *dto.ResponseData); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(dto.LoginRequest) *errs.AppError); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Refresh provides a mock function with given fields: request
func (_m *UserService) Refresh(request dto.RefreshTokenRequest) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(request)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(dto.RefreshTokenRequest) *dto.ResponseData); ok {
		r0 = rf(request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(dto.RefreshTokenRequest) *errs.AppError); ok {
		r1 = rf(request)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}