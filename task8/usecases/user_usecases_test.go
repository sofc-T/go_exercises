package usecases_test

import (
	"context"
	"testing"
	"time"

	"github.com/sofc-t/task_manager/task8/models"
	"github.com/sofc-t/task_manager/task8/usecases"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) Login(ctx context.Context, user models.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockUserRepository) FetchAllUsers(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserRepository) FetchByID(ctx context.Context, id string) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserRepository) PromoteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// UserUsecaseSuite is a test suite for the UserUsecase
type UserUsecaseSuite struct {
	suite.Suite
	mockRepo *MockUserRepository
	usecase  models.UserUsecase // Use the interface, not the concrete struct
}

func (suite *UserUsecaseSuite) SetupTest() {
	suite.mockRepo = new(MockUserRepository)
	suite.usecase = usecases.NewUserUsecase(suite.mockRepo, 1000*time.Second) // Use the factory function
}

func (suite *UserUsecaseSuite) TestCreate() {
	name, email := "John Doe", "john.doe@example.com"
	password:= mock.Anything
	user := models.User{Name: &name, Email: &email, Password: &password}

	suite.mockRepo.On("CreateUser", mock.Anything, user).Return(nil)

	err := suite.usecase.Create(context.Background(), user)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestLogin() {

	name, email := "John Doe", "john.doe@example.com"
	password:= mock.Anything
	user := models.User{Name: &name, Email: &email, Password: &password}
	suite.mockRepo.On("Login", mock.Anything, user).Return("mockToken", nil)

	token, err := suite.usecase.Login(context.Background(), user)

	suite.NoError(err)
	suite.Equal("mockToken", token)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestFetchAll() {

	name1, name2 := "John Doe", "Jane Doe"
	id1, id2 := primitive.NewObjectID(), primitive.NewObjectID()
	users := []models.User{
		{ID: id1, Name: &name1},
		{ID: id2, Name: &name2},
	}
	suite.mockRepo.On("FetchAllUsers", mock.Anything).Return(users, nil)

	result, err := suite.usecase.FetchAll(context.Background())

	suite.NoError(err)
	suite.ElementsMatch(users, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestFetchById() {
	name, email := "John Doe", "john.doe@example.com"
	password:= mock.Anything
	user := models.User{Name: &name, Email: &email, Password: &password}

	suite.mockRepo.On("FetchByID", mock.Anything, "1").Return(user, nil)

	result, err := suite.usecase.FetchById(context.Background(), "1")

	suite.NoError(err)
	suite.Equal(user, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *UserUsecaseSuite) TestPromoteUser() {
	suite.mockRepo.On("PromoteUser", mock.Anything, "1").Return(nil)

	err := suite.usecase.PromoteUser(context.Background(), "1")

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestUserUsecaseSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}
