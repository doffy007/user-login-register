package response

import (
	"net/http"

	constant "github.com/doffy007/user-login-register/internal/constants"
)

type BaseResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      []string
}

func (b BaseResponse) InternalServerError(err string) BaseResponse {
	b.StatusCode = http.StatusInternalServerError
	b.Message = constant.INTERNAL_SERVER_ERROR
	b.Data = false
	b.Error = []string{err}

	return b
}
