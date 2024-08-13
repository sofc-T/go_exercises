package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/sofc-t/task_manager/task8/models"
	"github.com/sofc-t/task_manager/task8/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) FetchTasks(ctx context.Context) ([]models.Task, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskRepository) FindTask(ctx context.Context, id int) (models.Task, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskRepository) UpdateTask(ctx context.Context, id int, title string) (models.Task, error) {
	args := m.Called(ctx, id, title)
	return args.Get(0).(models.Task), args.Error(1)
}

func (m *MockTaskRepository) DeleteTask(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockTaskRepository) InsertTask(ctx context.Context, task models.Task) (models.Task, error) {
	args := m.Called(ctx, task)
	return args.Get(0).(models.Task), args.Error(1)
}

type TaskUsecaseSuite struct {
	suite.Suite
	mockRepo *MockTaskRepository
	usecase  models.TaskUsecase
}

func (suite *TaskUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockTaskRepository)
	suite.usecase = usecases.NewTaskUsecase(suite.mockRepo, 1000*time.Second)
}

func (suite *TaskUsecaseSuite) TestFetch() {
	task1 := models.Task{Id: 1, Title: "Task 1"}
	task2 := models.Task{Id: 2, Title: "Task 2"}
	tasks := []models.Task{task1, task2}

	suite.mockRepo.On("FetchTasks", mock.Anything).Return(tasks, nil)

	result, err := suite.usecase.Fetch(context.Background())

	suite.NoError(err)
	suite.ElementsMatch(tasks, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestFind() {
	task := models.Task{Id: 1, Title: "Task 1"}

	suite.mockRepo.On("FindTask", mock.Anything, 1).Return(task, nil)

	result, err := suite.usecase.Find(context.Background(), 1)

	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestUpdate() {
	task := models.Task{Id: 1, Title: "Updated Task"}

	suite.mockRepo.On("UpdateTask", mock.Anything, 1, "Updated Task").Return(task, nil)

	result, err := suite.usecase.Update(context.Background(), 1, "Updated Task")

	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestDelete() {
	suite.mockRepo.On("DeleteTask", mock.Anything, 1).Return(nil)

	err := suite.usecase.Delete(context.Background(), 1)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *TaskUsecaseSuite) TestCreate() {
	task := models.Task{Id: 1, Title: "New Task"}

	suite.mockRepo.On("InsertTask", mock.Anything, task).Return(task, nil)

	result, err := suite.usecase.Create(context.Background(), task)

	suite.NoError(err)
	suite.Equal(task, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestTaskUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TaskUsecaseSuite))
}
