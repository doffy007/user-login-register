package usecase

import (
	"context"
	"reflect"
	"testing"

	"github.com/doffy007/user-login-register/config"
)

func TestNewUsecase(t *testing.T) {
	var conf *config.Config
	type args struct {
		ctx  context.Context
		conf *config.Config
	}
	tests := []struct {
		name string
		args args
		want Usecase
	}{
		{
			name: "Test NewUsecase",
			args: args{},
			want: NewUsecase(context.Background(), conf),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsecase(tt.args.ctx, tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}
