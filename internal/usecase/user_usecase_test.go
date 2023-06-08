package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/doffy007/user-login-register/config"
	"github.com/doffy007/user-login-register/internal/repository"
	"github.com/doffy007/user-login-register/internal/request"
	"github.com/doffy007/user-login-register/internal/response"
)

func Test_userUsecase_UserLogin(t *testing.T) {
	var conf *config.Config
	type fields struct {
		ctx            context.Context
		config         *config.Config
		userRepository repository.UserRepository
	}
	type args struct {
		params request.ParamUserLogin
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  response.BaseResponse
	}{
		{
			name: "Test UserLogin",
			fields: fields{
				ctx:            context.Background(),
				config:         conf,
				userRepository: nil,
			},
			args: args{
				params: request.ParamUserLogin{
					Username: "ojannady",
					Password: "123456789",
				},
			},
			want: true,
			want1: response.BaseResponse{
				StatusCode: 200,
				Message:    "",
				Data:       nil,
				Error:      []string{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userUsecase{
				ctx:            tt.fields.ctx,
				config:         tt.fields.config,
				userRepository: tt.fields.userRepository,
			}
			got, got1 := u.UserLogin(tt.args.params)
			if got != tt.want {
				t.Errorf("userUsecase.UserLogin() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("userUsecase.UserLogin() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_userUsecase_UserRegister(t *testing.T) {
	type fields struct {
		ctx            context.Context
		config         *config.Config
		userRepository repository.UserRepository
	}
	type args struct {
		params request.ParamUserRegister
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
		want1  response.BaseResponse
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := userUsecase{
				ctx:            tt.fields.ctx,
				config:         tt.fields.config,
				userRepository: tt.fields.userRepository,
			}
			got, got1 := u.UserRegister(tt.args.params)
			if got != tt.want {
				t.Errorf("userUsecase.UserRegister() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("userUsecase.UserRegister() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
