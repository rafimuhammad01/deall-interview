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
			name: "handle err internal",
			args: args{
				arrErr: []error{fmt.Errorf("[%w] test error", ErrUserInternal)},
			},
			want: ErrorStruct{
				Code:    http.StatusInternalServerError,
				Message: "internal server error",
				Errors:  nil,
			},
		},
		{
			name: "error is user not exist",
			args: args{arrErr: []error{fmt.Errorf("[%w] test error", ErrUserExist)}},
			want: ErrorStruct{
				Code:    http.StatusBadRequest,
				Message: "user already exist",
				Errors:  []string{"test error"},
			},
		},
		{
			name: "error is user invalid data",
			args: args{arrErr: []error{fmt.Errorf("[%w] test error", ErrUserInvalidData)}},
			want: ErrorStruct{
				Code:    http.StatusBadRequest,
				Message: "invalid data",
				Errors:  []string{"test error"},
			},
		},
		{
			name: "error is user not found",
			args: args{arrErr: []error{fmt.Errorf("[%w] test error", ErrUserNotFound)}},
			want: ErrorStruct{
				Code:    http.StatusNotFound,
				Message: "user not found",
				Errors:  []string{"test error"},
			},
		},
		{
			name: "error multiple",
			args: args{arrErr: []error{fmt.Errorf("[%w] test error 1", ErrUserInvalidData), fmt.Errorf("[%w] test error 2", ErrUserNotFound)}},
			want: ErrorStruct{
				Code:    http.StatusBadRequest,
				Message: "invalid data",
				Errors:  []string{"test error 1", "test error 2"},
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
		{
			name:    "success parsed",
			args:    args{err: fmt.Sprintf("[%s] test error", ErrUserInvalidData.Error())},
			want:    "test error",
			wantErr: false,
		},
		{
			name:    "invalid format",
			args:    args{err: fmt.Sprintf("%s test error", ErrUserInvalidData.Error())},
			want:    "",
			wantErr: true,
		},
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
