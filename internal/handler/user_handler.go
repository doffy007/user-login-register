package handler

import (
	"fmt"
	"net/http"

	"github.com/doffy007/user-login-register/internal/constants"
	"github.com/doffy007/user-login-register/internal/helper"
	"github.com/doffy007/user-login-register/internal/request"
	"github.com/doffy007/user-login-register/internal/usecase"
	"github.com/doffy007/user-login-register/internal/utils"
	"github.com/doffy007/user-login-register/internal/utils/validations"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	UserRegister(c *gin.Context)
	UserLogin(c *gin.Context)
}

type userHandler struct {
	usecase usecase.UserUsecase
}

func (h handler) UserHandler() UserHandler {
	return userHandler{
		usecase: h.usecase.UserHandler(),
	}
}

// UserRegister implements UserHandler.

// UserRegister 	godoc
// @Summary 		Register for user
// @Description 	Register User
// @Param 			tags body request.ParamUserRegister true "Create tags"
// @Produce 		application/json
// @Tags 			tags
// @Success 		200
// @Router			/v1/user/register [post]
func (h userHandler) UserRegister(c *gin.Context) {
	var params request.ParamUserRegister

	if err := c.ShouldBind(&params); err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler UserRegister Parse Param: %v\n", err.Error())
		utils.NewErrorResponse(c, http.StatusBadRequest, constants.ERROR_PARSE_PARAM, err.Error())
		return
	}

	err := validations.ValidateUserRegisterParams(params)

	if err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler UserRegister Validation: %v\n", err.Error())
		utils.NewErrorResponse(c, http.StatusUnprocessableEntity, constants.ERROR_VALIDATION, helper.FormatError(err))
		return
	}
	ok, res := h.usecase.UserRegister(params)

	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler UserRegister From Usecase: %v\n", res.Message)
		utils.NewErrorResponse(c, res.StatusCode, res.Message, []string{res.Message})
		return
	}

	utils.NewSuccessResponse(c, res.StatusCode, res.Message, res.Data)

}

// UserLogin implements UserHandler.

// UserLogin	 	godoc
// @Summary 		Login for user
// @Description 	Login User
// @Param 			tags body request.ParamUserLogin true "Login tags"
// @Produce 		application/json
// @Tags 			tags
// @Success 		200
// @Router			/v1/user/login [post]
func (h userHandler) UserLogin(c *gin.Context) {
	var params request.ParamUserLogin

	if err := c.ShouldBind(&params); err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler UserLogin Parse Param: %v\n", err.Error())
		utils.NewErrorResponse(c, http.StatusBadRequest, constants.ERROR_PARSE_PARAM, err.Error())
		return
	}

	ok, res := h.usecase.UserLogin(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler UserLogin From Usecase: %v\n", res.Message)
		utils.NewErrorResponse(c, res.StatusCode, res.Message, []string{res.Message})
		return
	}

	utils.NewSuccessResponse(c, res.StatusCode, res.Message, res.Data)
}
