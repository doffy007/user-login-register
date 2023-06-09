// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	entity "github.com/doffy007/user-login-register/internal/entity"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: _a0
func (_m *UserRepository) CreateUser(_a0 entity.CreateUser) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(entity.CreateUser) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindOneUser provides a mock function with given fields: _a0, _a1
func (_m *UserRepository) FindOneUser(_a0 entity.FindOneUser, _a1 []string) (*entity.Users, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *entity.Users
	var r1 error
	if rf, ok := ret.Get(0).(func(entity.FindOneUser, []string) (*entity.Users, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(entity.FindOneUser, []string) *entity.Users); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Users)
		}
	}

	if rf, ok := ret.Get(1).(func(entity.FindOneUser, []string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
