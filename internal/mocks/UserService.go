// Code generated by mockery v2.10.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/danisbagus/matchoshop/internal/core/domain"
	dto "github.com/danisbagus/matchoshop/internal/dto"

	errs "github.com/danisbagus/go-common-packages/errs"

	mock "github.com/stretchr/testify/mock"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Delete provides a mock function with given fields: userID, roleID
func (_m *UserService) Delete(userID int64, roleID int64) *errs.AppError {
	ret := _m.Called(userID, roleID)

	var r0 *errs.AppError
	if rf, ok := ret.Get(0).(func(int64, int64) *errs.AppError); ok {
		r0 = rf(userID, roleID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*errs.AppError)
		}
	}

	return r0
}

// GetDetail provides a mock function with given fields: userID
func (_m *UserService) GetDetail(userID int64) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(userID)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(int64) *dto.ResponseData); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
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

// GetList provides a mock function with given fields: roldID
func (_m *UserService) GetList(roldID int64) ([]domain.UserDetail, *errs.AppError) {
	ret := _m.Called(roldID)

	var r0 []domain.UserDetail
	if rf, ok := ret.Get(0).(func(int64) []domain.UserDetail); ok {
		r0 = rf(roldID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.UserDetail)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64) *errs.AppError); ok {
		r1 = rf(roldID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
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

// RegisterCustomer provides a mock function with given fields: req
func (_m *UserService) RegisterCustomer(req *dto.RegisterCustomerRequest) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(req)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(*dto.RegisterCustomerRequest) *dto.ResponseData); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(*dto.RegisterCustomerRequest) *errs.AppError); ok {
		r1 = rf(req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}

// Update provides a mock function with given fields: userID, data
func (_m *UserService) Update(userID int64, data *dto.UpdateUserRequest) (*dto.ResponseData, *errs.AppError) {
	ret := _m.Called(userID, data)

	var r0 *dto.ResponseData
	if rf, ok := ret.Get(0).(func(int64, *dto.UpdateUserRequest) *dto.ResponseData); ok {
		r0 = rf(userID, data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.ResponseData)
		}
	}

	var r1 *errs.AppError
	if rf, ok := ret.Get(1).(func(int64, *dto.UpdateUserRequest) *errs.AppError); ok {
		r1 = rf(userID, data)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*errs.AppError)
		}
	}

	return r0, r1
}
