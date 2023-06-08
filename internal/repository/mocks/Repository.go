// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	repository "github.com/doffy007/user-login-register/internal/repository"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// UserRepository provides a mock function with given fields:
func (_m *Repository) UserRepository() repository.UserRepository {
	ret := _m.Called()

	var r0 repository.UserRepository
	if rf, ok := ret.Get(0).(func() repository.UserRepository); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(repository.UserRepository)
		}
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
