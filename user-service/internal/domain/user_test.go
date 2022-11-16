package domain

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	type fields struct {
		ID       string
		Name     string
		Username string
		Password string
		Role     int
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		{
			name: "success no error",
			fields: fields{
				ID:       "ID",
				Name:     "Name",
				Username: "username",
				Password: "testpassword",
			},
			want: nil,
		},
		{
			name: "failed validate one field",
			fields: fields{
				ID:       "",
				Name:     "test name",
				Username: "test username",
				Password: "testpassword",
			},
			want: []error{
				fmt.Errorf(
					"[%w] id should not be empty", ErrUserInvalidData,
				),
			},
		},
		{
			name: "failed in multiple fields",
			fields: fields{
				ID:       "test",
				Name:     "",
				Username: "",
				Password: "",
				Role:     3,
			},
			want: []error{
				fmt.Errorf("[%w] name should not be empty", ErrUserInvalidData),
				fmt.Errorf("[%w] username should not be empty", ErrUserInvalidData),
				fmt.Errorf("[%w] password should not be empty", ErrUserInvalidData),
				fmt.Errorf("[%w] role should be only 0 (user) or 1 (admin)", ErrUserInvalidData),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Username: tt.fields.Username,
				Password: tt.fields.Password,
				Role:     tt.fields.Role,
			}
			if got := u.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("User.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
