package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUser_ValidateLogin(t *testing.T) {
	type fields struct {
		ID       string
		Username string
		Password string
		Role     int
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		// TODO: Add test cases.
		{
			name: "success",
			fields: fields{
				ID:       "1",
				Username: "rafi",
				Password: "rafi",
				Role:     0,
			},
			want: nil,
		},
		{
			name: "validation error",
			fields: fields{
				ID:       "",
				Username: "",
				Password: "",
				Role:     0,
			},
			want: []error{
				fmt.Errorf("[%w] username should not be empty", ErrInvalidDataLogin),
				fmt.Errorf("[%w] password should not be empty", ErrInvalidDataLogin),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Role:     tt.fields.Role,
			}
			if got := u.ValidateLogin(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateLogin() = %v, want %v", got, tt.want)
			}
		})
	}
}
