package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sofc-t/task_manager/task8/controllers"
	"github.com/sofc-t/task_manager/task8/models"
	
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockUserUsecase is a mock implementation of the UserUsecase interface
type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Create(ctx context.Context, user models.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserUsecase) Login(ctx context.Context, user models.User) (string, error) {
	args := m.Called(ctx, user)
	return args.String(0), args.Error(1)
}

func (m *MockUserUsecase) FetchById(ctx context.Context, id string) (models.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *MockUserUsecase) PromoteUser(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *MockUserUsecase) FetchAll(ctx context.Context) ([]models.User, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.User), args.Error(1)
}

// UserControllerSuite is a test suite for UserController
type UserControllerSuite struct {
	suite.Suite
	mockUsecase *MockUserUsecase
	uc          controllers.UserController
	router      *gin.Engine
}

// SetupSuite runs once before the suite's tests are run
func (suite *UserControllerSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)
	suite.mockUsecase = new(MockUserUsecase)
	suite.uc = controllers.UserController{UserUsecase: suite.mockUsecase}
	suite.router = gin.Default()
}

// TestSignUp tests the SignUp handler
func (suite *UserControllerSuite) TestSignUp() {
	suite.router.POST("/signup", suite.uc.SignUp)

	suite.mockUsecase.On("Create", mock.Anything, mock.AnythingOfType("models.User")).Return(nil)

	req, _ := http.NewRequest("POST", "/signup", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestLogin tests the Login handler
func (suite *UserControllerSuite) TestLogin() {
	suite.router.POST("/login", suite.uc.Login)

	suite.mockUsecase.On("Login", mock.Anything, mock.AnythingOfType("models.User")).Return("mockToken", nil)

	req, _ := http.NewRequest("POST", "/login", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusCreated, w.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestGetUserByID tests the GetUseryID handler
func (suite *UserControllerSuite) TestGetUserByID() {
	suite.router.GET("/users/:id", suite.uc.GetUseryID)
	id := primitive.NewObjectID() // Generate a new ObjectID
	name := "Mock User"  
	mockUser := models.User{ID: id, Name: &name}
	suite.mockUsecase.On("FetchById", mock.Anything, "mockID").Return(mockUser, nil)

	req, _ := http.NewRequest("GET", "/users/mockID", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusOK, w.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// TestPromoteUser tests the PromoteUser handler
func (suite *UserControllerSuite) TestPromoteUser() {
	suite.router.POST("/promote", suite.uc.PromoteUser)

	suite.mockUsecase.On("PromoteUser", mock.Anything, "mockID").Return(nil)

	req, _ := http.NewRequest("POST", "/promote", nil)
	w := httptest.NewRecorder()
	suite.router.ServeHTTP(w, req)

	suite.Equal(http.StatusAccepted, w.Code)
	suite.mockUsecase.AssertExpectations(suite.T())
}

// Run the test suite
func TestUserControllerSuite(t *testing.T) {
	suite.Run(t, new(UserControllerSuite))
}
