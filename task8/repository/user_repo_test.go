package repository_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/sofc-t/task_manager/task8/models"
	"github.com/sofc-t/task_manager/task8/repository"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type UserRepositorySuite struct {
	suite.Suite
	mt   *mtest.T
	repo models.UserRepository
	db   *mongo.Database
}

func (suite *UserRepositorySuite) SetupTest() {
	suite.mt = mtest.New(suite.T(), mtest.NewOptions().ClientType(mtest.Mock))
	suite.db = suite.mt.Client.Database("taskmanager")
	suite.repo = repository.NewUserRepository(*suite.db, "users")
}

func (suite *UserRepositorySuite) TearDownTest() {
	fmt.Println("Tearing down")
}

func (suite *UserRepositorySuite) TestCreateUser() {
	suite.mt.Run("User Creation", func(mt *mtest.T) {
		user := models.User{
			ID:       primitive.NewObjectID(),
			Name:     stringPointer("Test User"),
			Email:    stringPointer("test@example.com"),
			Password: stringPointer("hashedpassword"),
			Role:     stringPointer("user"),
		}

		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := suite.repo.CreateUser(context.TODO(), user)
		suite.NoError(err)
	})
}

func (suite *UserRepositorySuite) TestLoginSuccess() {
	suite.mt.Run("Login Success", func(mt *mtest.T) {
		user := models.User{
			
			Email:    stringPointer("test@example.com"),
			Password: stringPointer("hashedpassword"),
			Name:     stringPointer("Test User"),
			Role:     stringPointer("user"),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.users", mtest.FirstBatch, bson.D{
			{Key: "email", Value: *user.Email},
			{Key: "password", Value: *user.Password},
			{Key: "name", Value: *user.Name},
			{Key: "role", Value: *user.Role},
		}))

		token, err := suite.repo.Login(context.TODO(), models.User{Email: user.Email, Password: user.Password})
		suite.NoError(err)
		suite.NotEmpty(token)
	})
}

func (suite *UserRepositorySuite) TestLoginFailure() {
	suite.mt.Run("Login Failure", func(mt *mtest.T) {
		mt.AddMockResponses(mtest.CreateCursorResponse(0, "taskmanager.users", mtest.FirstBatch, bson.D{}))

		_, err := suite.repo.Login(context.TODO(), models.User{Email: stringPointer("nonexistent@example.com")})
		suite.Error(err)
	})
}

func (suite *UserRepositorySuite) TestFetchAllUsers() {
	suite.mt.Run("Fetch All Users", func(mt *mtest.T) {
		users := []models.User{
			{
				ID:    primitive.NewObjectID(),
				Name:  stringPointer("User1"),
				Email: stringPointer("user1@example.com"),
			},
			{
				ID:    primitive.NewObjectID(),
				Name:  stringPointer("User2"),
				Email: stringPointer("user2@example.com"),
			},
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.users", mtest.FirstBatch, bson.D{
			{Key: "name", Value: *users[0].Name},
			{Key: "email", Value: *users[0].Email},
		}))
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.users", mtest.FirstBatch, bson.D{
			{Key: "name", Value: *users[1].Name},
			{Key: "email", Value: *users[1].Email},
		}))

		fetchedUsers, err := suite.repo.FetchAllUsers(context.TODO())
		suite.NoError(err)
		suite.Len(fetchedUsers, 2)
	})
}

func (suite *UserRepositorySuite) TestFetchByID() {
	suite.mt.Run("Fetch By ID", func(mt *mtest.T) {
		userID := primitive.NewObjectID()
		user := models.User{
			
			Name:  stringPointer("Test User"),
			Email: stringPointer("test@example.com"),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.users", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: userID},
			{Key: "name", Value: *user.Name},
			{Key: "email", Value: *user.Email},
		}))

		fetchedUser, err := suite.repo.FetchByID(context.TODO(), userID.Hex())
		suite.NoError(err)
		suite.Equal(*user.Name, *fetchedUser.Name)
		suite.Equal(*user.Email, *fetchedUser.Email)
	})
}

func (suite *UserRepositorySuite) TestPromoteUser() {
	suite.mt.Run("Promote User", func(mt *mtest.T) {
		userID := primitive.NewObjectID()
		user := models.User{
			
			Role: stringPointer("user"),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "taskmanager.users", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: userID},
			{Key: "role", Value: *user.Role},
		}))
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := suite.repo.PromoteUser(context.TODO(), userID.Hex())
		suite.NoError(err)
	})
}

// Helper function to create string pointers
func stringPointer(s string) *string {
	return &s
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositorySuite))
}
