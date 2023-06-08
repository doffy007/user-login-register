package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/doffy007/user-login-register/config"
	"github.com/doffy007/user-login-register/internal/entity"
	"github.com/jmoiron/sqlx"
)

func Test_userRepository_CreateUser(t *testing.T) {
	type fields struct {
		ctx    context.Context
		config *config.Config
		db     *sqlx.DB
	}
	type args struct {
		payload entity.CreateUser
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test create user",
			fields: fields{
				ctx:    context.Background(),
				config: &config.Config{},
				db:     &sqlx.DB{},
			},
			args: args{
				payload: entity.CreateUser{
					Username: new(string),
					Email:    new(string),
					Password: new(string),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				ctx:    tt.fields.ctx,
				config: tt.fields.config,
				db:     tt.fields.db,
			}
			if err := r.CreateUser(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("userRepository.CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_userRepository_FindOneUser(t *testing.T) {
	type fields struct {
		ctx    context.Context
		config *config.Config
		db     *sqlx.DB
	}
	type args struct {
		payload entity.FindOneUser
		fields  []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *entity.Users
		wantErr bool
	}{
		{
			name: "Test FindOneUser",
			fields: fields{
				ctx:    context.Background(),
				config: &config.Config{},
				db:     &sqlx.DB{},
			},
			want: &entity.Users{
				Username: new(string),
				Email:    new(string),
				Password: new(string),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := userRepository{
				ctx:    tt.fields.ctx,
				config: tt.fields.config,
				db:     tt.fields.db,
			}
			got, err := r.FindOneUser(tt.args.payload, tt.args.fields)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.FindOneUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("userRepository.FindOneUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
