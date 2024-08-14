// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/sofc-t/task_manager/task8/models"
	mock "github.com/stretchr/testify/mock"
)

// TaskRepository is an autogenerated mock type for the TaskRepository type
type TaskRepository struct {
	mock.Mock
}

// DeleteTask provides a mock function with given fields: ctx, id
func (_m *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FetchTasks provides a mock function with given fields: ctx
func (_m *TaskRepository) FetchTasks(ctx context.Context) ([]models.Task, error) {
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

// FindTask provides a mock function with given fields: ctx, id
func (_m *TaskRepository) FindTask(ctx context.Context, id int) (models.Task, error) {
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

// InsertTask provides a mock function with given fields: ctx, task
func (_m *TaskRepository) InsertTask(ctx context.Context, task models.Task) (models.Task, error) {
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

// UpdateTask provides a mock function with given fields: ctx, id, title
func (_m *TaskRepository) UpdateTask(ctx context.Context, id int, title string) (models.Task, error) {
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

type mockConstructorTestingTNewTaskRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewTaskRepository creates a new instance of TaskRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewTaskRepository(t mockConstructorTestingTNewTaskRepository) *TaskRepository {
	mock := &TaskRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}