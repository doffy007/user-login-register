package usecase

import (
	"context"
	"encoding/json"
	"net/http"

	constant "github.com/doffy007/user-login-register/internal/constants"

	"github.com/doffy007/user-login-register/config"
	"github.com/doffy007/user-login-register/internal/entity"
	"github.com/doffy007/user-login-register/internal/helper"
	"github.com/doffy007/user-login-register/internal/repository"
	"github.com/doffy007/user-login-register/internal/request"
	"github.com/doffy007/user-login-register/internal/response"
)

type UserUsecase interface {
	UserRegister(request.ParamUserRegister) (bool, response.BaseResponse)
	UserLogin(request.ParamUserLogin) (bool, response.BaseResponse)
}

type userUsecase struct {
	ctx            context.Context
	config         *config.Config
	userRepository repository.UserRepository
}

func (u usecase) UserHandler() UserUsecase {
	return &userUsecase{
		ctx:            u.ctx,
		config:         u.config,
		userRepository: u.repository.UserRepository(),
	}
}

// UserLogin implements UserUsecase.
func (u userUsecase) UserLogin(params request.ParamUserLogin) (bool, response.BaseResponse) {
	var filter entity.FindOneUser

	// Check if email or user
	if helper.CheckEmail(params.Username) {
		filter.Email = &params.Username
	} else {
		filter.Username = &params.Username
	}

	user, err := u.userRepository.FindOneUser(filter, []string{})
	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}

	if user == nil {
		return false, response.BaseResponse{StatusCode: http.StatusUnauthorized, Message: "username or password is incorrect"}
	}

	err = helper.ComparePassword(*user.Password, params.Password)

	if err != nil {
		return false, response.BaseResponse{StatusCode: http.StatusBadRequest, Message: "username or password is incorrect"}
	}

	payloadToken := map[string]interface{}{
		"user_id": user.ID,
	}

	token, _ := helper.CreateToken(constant.DIR_PRIVATE_KEY_JWT_DEV, payloadToken)

	res := map[string]interface{}{
		"username":     params.Username,
		"access_token": token,
	}

	return true, response.BaseResponse{StatusCode: http.StatusCreated, Message: "success login user", Data: res}

}

// UserRegister implements UserUsecase.
func (u userUsecase) UserRegister(params request.ParamUserRegister) (bool, response.BaseResponse) {

	if params.Username != "" {
		checkUsername, err := u.userRepository.FindOneUser(entity.FindOneUser{Username: &params.Username}, []string{})
		if err != nil {
			return false, response.BaseResponse{}.InternalServerError(err.Error())
		}

		if checkUsername != nil {
			return false, response.BaseResponse{StatusCode: http.StatusUnprocessableEntity, Message: "username already used"}
		}
	}

	if params.Email != "" {
		checkEmail, err := u.userRepository.FindOneUser(entity.FindOneUser{Email: &params.Email}, []string{})
		if err != nil {
			return false, response.BaseResponse{}.InternalServerError(err.Error())
		}

		if checkEmail != nil {
			return false, response.BaseResponse{StatusCode: http.StatusUnprocessableEntity, Message: "email already used"}
		}
	}

	hashedPassword, err := helper.HashPassword(params.Password)

	if err != nil {
		return false, response.BaseResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()}
	}

	params.Password = hashedPassword

	var payload entity.CreateUser
	rec, _ := json.Marshal(params)
	json.Unmarshal(rec, &payload)

	err = u.userRepository.CreateUser(payload)

	if err != nil {
		return false, response.BaseResponse{}.InternalServerError(err.Error())
	}

	return true, response.BaseResponse{StatusCode: http.StatusCreated, Message: "success created user", Data: true}
}
