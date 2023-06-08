package repository

import (
	"context"
	"reflect"
	"testing"

	"github.com/doffy007/user-login-register/config"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		ctx  context.Context
		conf *config.Config
	}
	tests := []struct {
		name string
		args args
		want Repository
	}{
		{
			name: "new repository",
			args: args{
				ctx:  context.TODO(),
				conf: &config.Config{},
			},
			want: NewRepository(context.Background(), &config.ConfigureApp),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.ctx, tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
