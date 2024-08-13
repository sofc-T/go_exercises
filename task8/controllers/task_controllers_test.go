package controllers_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sofc-t/task_manager/task8/controllers"
	"github.com/sofc-t/task_manager/task8/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// MockTaskUsecase is a mock implementation of the TaskUsecase interface
type MockTaskUsecase struct {
	mock.Mock
}

func (m *MockTaskUsecase) Fetch(ctx context.Context) ([]models.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskUsecase) Find(ctx context.Context, id int) (models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskUsecase) Update(ctx context.Context, id int, title string) (models.Task, error) {
	args := m.Called(ctx, id, title)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskUsecase) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskUsecase) Create(ctx context.Context, task models.Task) (models.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(models.Task), args.Error(1)
}

// TaskControllerTestSuite is the test suite for TaskController
type TaskControllerTestSuite struct {
	suite.Suite
	router       *gin.Engine
	mockUsecase  *MockTaskUsecase
	taskController *controllers.TaskController
}

func (suite *TaskControllerTestSuite) SetupTest() {
	suite.mockUsecase = new(MockTaskUsecase)
	suite.taskController = &controllers.TaskController{TaskUsecase: suite.mockUsecase}
	suite.router = gin.Default()

	// Set up routes
	suite.router.GET("/tasks", suite.taskController.GetAllTasksHandler)
	suite.router.GET("/tasks/:id", suite.taskController.GetTaskHandler)
	suite.router.PUT("/tasks/:id", suite.taskController.UpdateTaskHandler)
	suite.router.DELETE("/tasks/:id", suite.taskController.DeleteTaskHandler)
	suite.router.POST("/tasks", suite.taskController.CreateTaskHandler)
}

func (suite *TaskControllerTestSuite) TestGetAllTasks() {
	mockTasks := []models.Task{
		{Id: 1, Title: "Task 1"},
		{Id: 2, Title: "Task 2"},
	}

	suite.mockUsecase.On("Fetch", mock.Anything).Return(mockTasks, nil)

	req, _ := http.NewRequest("GET", "/tasks", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusAccepted, w.Code)
	
}

func (suite *TaskControllerTestSuite) TestGetTask() {
	mockTask := models.Task{Id: 1, Title: "Task 1"}

	suite.mockUsecase.On("Find", mock.Anything, 1).Return(mockTask, nil)

	req, _ := http.NewRequest("GET", "/tasks/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	
}

func (suite *TaskControllerTestSuite) TestUpdateTask() {
	mockTask := models.Task{Id: 1, Title: "Updated Task"}

	suite.mockUsecase.On("Update", mock.Anything, 1, "Updated Task").Return(mockTask, nil)

	req, _ := http.NewRequest("PUT", "/tasks/1", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(strings.NewReader(`{"Title":"Updated Task"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusAccepted, w.Code)
	
}

func (suite *TaskControllerTestSuite) TestDeleteTask() {
	suite.mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/tasks/1", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
}

func (suite *TaskControllerTestSuite) TestCreateTask() {
	mockTask := models.Task{Id: 1, Title: "New Task"}

	suite.mockUsecase.On("Create", mock.Anything, mockTask).Return(mockTask, nil)

	req, _ := http.NewRequest("POST", "/tasks", nil)
	req.Header.Set("Content-Type", "application/json")
	req.Body = ioutil.NopCloser(strings.NewReader(`{"Title":"New Task"}`))
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

}

func TestTaskControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TaskControllerTestSuite))
}
