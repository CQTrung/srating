package repositories

import (
	"context"
	"reflect"
	"srating/domain"
	"testing"
)

func Test_userRepository_GetAllEmployee(t *testing.T) {
	type args struct {
		c     context.Context
		input domain.GetAllUserRequest
	}
	tests := []struct {
		name    string
		r       *userRepository
		args    args
		want    int64
		want1   int64
		want2   []*domain.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := tt.r.GetAllEmployee(tt.args.c, tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("userRepository.GetAllEmployee() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("userRepository.GetAllEmployee() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("userRepository.GetAllEmployee() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("userRepository.GetAllEmployee() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
