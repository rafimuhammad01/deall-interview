package domain

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestHandleError(t *testing.T) {
	type args struct {
		arrErr []error
	}
	tests := []struct {
		name string
		args args
		want ErrorStruct
	}{
		{
			name: "success handle error internal",
			args: args{
				arrErr: []error{
					fmt.Errorf("[%w] error example", ErrInternal),
				},
			},
			want: ErrorStruct{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
			},
		},
		{
			name: "success handle error invalid data login",
			args: args{
				arrErr: []error{
					fmt.Errorf("[%w] error example", ErrInvalidDataLogin),
				},
			},
			want: ErrorStruct{
				Code:    http.StatusBadRequest,
				Message: ErrInvalidDataLogin.Error(),
				Errors: []string{
					"error example",
				},
			},
		},
		{
			name: "success handle error user not found",
			args: args{
				arrErr: []error{
					fmt.Errorf("[%w] error example", ErrUserNotFound),
				},
			},
			want: ErrorStruct{
				Code:    http.StatusNotFound,
				Message: ErrUserNotFound.Error(),
				Errors: []string{
					"error example",
				},
			},
		},
		{
			name: "success handle error refresh token invalid",
			args: args{
				arrErr: []error{
					fmt.Errorf("[%w] error example", ErrRefreshTokenInvalid),
				},
			},
			want: ErrorStruct{
				Code:    http.StatusNotFound,
				Message: ErrRefreshTokenInvalid.Error(),
				Errors: []string{
					"error example",
				},
			},
		},
		{
			name: "invalid error format",
			args: args{
				arrErr: []error{fmt.Errorf("test error")},
			},
			want: ErrorStruct{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HandleError(tt.args.arrErr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HandleError() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		err string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.err)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
