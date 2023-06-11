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

	handler := userHandler{
		usecase: mockUsecase,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	params := request.ParamUserRegister{
		Username:        "ojannady",
		Email:           "ojannady@gmail.com",
		Password:        "123456789",
		ConfirmPassword: "123456789",
	}
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/register", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("request_body", params)

	handler.UserRegister(c)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")
}

func TestUserHandler_UserLogin(t *testing.T) {

	mockUsecase := &mockUserUsecase{}

	handler := userHandler{
		usecase: mockUsecase,
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	params := request.ParamUserLogin{
		Username: "ojannady",
		Password: "123456789",
	}
	c.Request, _ = http.NewRequest(http.MethodPost, "/v1/user/login", nil)
	c.Request.Header.Set("Content-Type", "application/json")
	c.Set("request_body", params)

	handler.UserLogin(c)

	assert.Equal(t, http.StatusBadRequest, w.Code, "Unexpected status code")
}
