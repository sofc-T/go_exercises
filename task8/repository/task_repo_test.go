package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sofc-t/task_manager/task8/models"
	"github.com/sofc-t/task_manager/task8/repository"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type TaskRepositorySuite struct {
	suite.Suite
	mtestInstance *mtest.T
	repo          models.TaskRepository
	db            *mongo.Database
}

func (suite *TaskRepositorySuite) SetupSuite() {
	suite.mtestInstance = mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock))
	suite.db = suite.mtestInstance.Client.Database("taskmanager")
	suite.repo = repository.NewTaskRepository(*suite.db, "tasks")
}

func (suite *TaskRepositorySuite) TearDownSuite() {
	fmt.Println("TEaring Down")
}


func (suite *TaskRepositorySuite) TestFetchTasks() {
	suite.mtestInstance.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.tasks", mtest.FirstBatch, bson.D{
		{Key: "id", Value: 1},
		{Key: "title", Value: "Test Task"},
	}))

	tasks, err := suite.repo.FetchTasks(context.TODO())
	suite.NoError(err)
	suite.Len(tasks, 1)
	suite.Equal("Test Task", tasks[0].Title)
}

func (suite *TaskRepositorySuite) TestInsertTask() {
	suite.mtestInstance.AddMockResponses(mtest.CreateSuccessResponse())

	task := models.Task{Id: 1, Title: "Test Task"}
	insertedTask, err := suite.repo.InsertTask(context.TODO(), task)
	suite.NoError(err)
	suite.Equal(task, insertedTask)
}

func (suite *TaskRepositorySuite) TestFindTask() {
	suite.mtestInstance.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.tasks", mtest.FirstBatch, bson.D{
		{Key: "id", Value: 1},
		{Key: "title", Value: "Test Task"},
	}))

	task, err := suite.repo.FindTask(context.TODO(), 1)
	suite.NoError(err)
	suite.Equal(1, task.Id)
	suite.Equal("Test Task", task.Title)
}

func (suite *TaskRepositorySuite) TestUpdateTask() {
	suite.mtestInstance.AddMockResponses(mtest.CreateSuccessResponse())

	updatedTask, err := suite.repo.UpdateTask(context.TODO(), 1, "Updated Title")
	suite.NoError(err)
	suite.Equal("Updated Title", updatedTask.Title)
}

func (suite *TaskRepositorySuite) TestDeleteTask() {
	suite.mtestInstance.AddMockResponses(mtest.CreateSuccessResponse())

	err := suite.repo.DeleteTask(context.TODO(), 1)
	suite.NoError(err)
}

func TestTaskRepositorySuite(t *testing.T) {
	suite.Run(t, new(TaskRepositorySuite))
}
