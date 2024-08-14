// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/sofc-t/task_manager/task8/models"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) CreateUser(ctx context.Context, user models.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchAllUsers provides a mock function with given fields: ctx
func (_m *UserRepository) FetchAllUsers(ctx context.Context) ([]models.User, error) {
	ret := _m.Called(ctx)

	var r0 []models.User
	if rf, ok := ret.Get(0).(func(context.Context) []models.User); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FetchByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) FetchByID(ctx context.Context, id string) (models.User, error) {
	ret := _m.Called(ctx, id)

	var r0 models.User
	if rf, ok := ret.Get(0).(func(context.Context, string) models.User); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.User)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: ctx, user
func (_m *UserRepository) Login(ctx context.Context, user models.User) (string, error) {
	ret := _m.Called(ctx, user)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, models.User) string); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PromoteUser provides a mock function with given fields: ctx, id
func (_m *UserRepository) PromoteUser(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
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