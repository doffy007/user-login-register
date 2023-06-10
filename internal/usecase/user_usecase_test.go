package usecase

import (
	"context"
	"net/http"
	"testing"

	"github.com/doffy007/user-login-register/config"
	"github.com/doffy007/user-login-register/internal/entity"
	"github.com/doffy007/user-login-register/internal/request"
	"github.com/stretchr/testify/assert"
)

type mockUserRepository struct {
	findOneUserResult *entity.Users
	findOneUserErr    error
	createUserErr     error
}

func (m *mockUserRepository) FindOneUser(filter entity.FindOneUser, fields []string) (*entity.Users, error) {
	return m.findOneUserResult, m.findOneUserErr
}

func (m *mockUserRepository) CreateUser(user entity.CreateUser) error {
	return m.createUserErr
}

type mockHelper struct{}

func (m *mockHelper) CheckEmail(username string) bool {
	return true
}

func (m *mockHelper) ComparePassword(savedPassword, providedPassword string) error {
	return nil
}

func (m *mockHelper) CreateToken(privateKeyPath string, payload map[string]interface{}) (string, error) {
	return "mocked_token", nil
}

func (m *mockHelper) HashPassword(password string) (string, error) {
	return "hashed_password", nil
}

func TestUserUsecase_UserLogin(t *testing.T) {
	// Set up the test
	usecase := userUsecase{
		ctx:            context.Background(),
		config:         &config.Config{},
		userRepository: &mockUserRepository{},
	}

	params := request.ParamUserLogin{
		Username: string("john@example.com"),
		Password: string("password"),
	}

	// Call the method being tested
	success, response := usecase.UserLogin(params)

	// Assert the results
	assert.NotNil(t, success, "Expected login to be successful")
	assert.NotNil(t, http.StatusCreated, response.StatusCode, "Unexpected status code")
	assert.NotNil(t, "success login user", response.Message, "Unexpected message")
	assert.NotNil(t, "john@example.com", response.Data, "Unexpected username")
	assert.NotNil(t, "mocked_token", response.Data, "Unexpected access token")
}

func TestUserUsecase_UserRegister(t *testing.T) {
	// Set up the test
	usecase := userUsecase{
		ctx:            context.Background(),
		config:         &config.Config{},
		userRepository: &mockUserRepository{},
	}

	params := request.ParamUserRegister{
		Username: "john",
		Email:    "john@example.com",
		Password: "password",
	}

	// Call the method being tested
	success, response := usecase.UserRegister(params)

	// Assert the results
	assert.True(t, success, "Expected user registration to be successful")
	assert.Equal(t, http.StatusCreated, response.StatusCode, "Unexpected status code")
	assert.Equal(t, "success created user", response.Message, "Unexpected message")
	assert.Equal(t, true, response.Data, "Unexpected data")
}
