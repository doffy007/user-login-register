package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/doffy007/user-login-register/internal/request"
	"github.com/doffy007/user-login-register/internal/response"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockUserUsecase struct {
	mock.Mock
}

func (m *mockUserUsecase) UserRegister(params request.ParamUserRegister) (bool, response.BaseResponse) {
	args := m.Called(params)
	return args.Bool(0), args.Get(1).(response.BaseResponse)
}

func (m *mockUserUsecase) UserLogin(params request.ParamUserLogin) (bool, response.BaseResponse) {
	args := m.Called(params)
	return args.Bool(0), args.Get(1).(response.BaseResponse)
}

func TestUserHandler_UserRegister(t *testing.T) {

	mockUsecase := &mockUserUsecase{}

	// Create the userHandler with the mock usecase
	handler := userHandler{
		usecase: mockUsecase,
	}

	// Create a mock Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request body and bind it to the context
	params := request.ParamUserRegister{
		// Set the request body fields here
	}
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/register", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("request_body", params)

	// Call the method being tested
	handler.UserRegister(c)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")
	// Assert other response attributes or values if needed
}

func TestUserHandler_UserLogin(t *testing.T) {

	mockUsecase := &mockUserUsecase{}

	// Create the userHandler with the mock usecase
	handler := userHandler{
		usecase: mockUsecase,
	}

	// Create a mock Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Set up the request body and bind it to the context
	params := request.ParamUserLogin{
		// Set the request body fields here
	}
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/login", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("request_body", params)

	// Call the method being tested
	handler.UserLogin(c)

	// Assert the response
	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")
	// Assert other response attributes or values if needed
}
