// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/sofc-t/task_manager/task8/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskUsecase is an autogenerated mock type for the TaskUsecase type
type TaskUsecase struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, task
func (_m *TaskUsecase) Create(ctx context.Context, task models.Task) (models.Task, error) {
	ret := _m.Called(ctx, task)

	var r0 models.Task
	if rf, ok := ret.Get(0).(func(context.Context, models.Task) models.Task); ok {
		r0 = rf(ctx, task)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, models.Task) error); ok {
		r1 = rf(ctx, task)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *TaskUsecase) Delete(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Fetch provides a mock function with given fields: ctx
func (_m *TaskUsecase) Fetch(ctx context.Context) ([]models.Task, error) {
	ret := _m.Called(ctx)

	var r0 []models.Task
	if rf, ok := ret.Get(0).(func(context.Context) []models.Task); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Task)
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

// Find provides a mock function with given fields: ctx, id
func (_m *TaskUsecase) Find(ctx context.Context, id int) (models.Task, error) {
	ret := _m.Called(ctx, id)

	var r0 models.Task
	if rf, ok := ret.Get(0).(func(context.Context, int) models.Task); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, title
func (_m *TaskUsecase) Update(ctx context.Context, id int, title string) (models.Task, error) {
	ret := _m.Called(ctx, id, title)

	var r0 models.Task
	if rf, ok := ret.Get(0).(func(context.Context, int, string) models.Task); ok {
		r0 = rf(ctx, id, title)
	} else {
		r0 = ret.Get(0).(models.Task)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, string) error); ok {
		r1 = rf(ctx, id, title)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewTaskUsecase interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskUsecase creates a new instance of TaskUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskUsecase(t mockConstructorTestingTNewTaskUsecase) *TaskUsecase {
	mock := &TaskUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
