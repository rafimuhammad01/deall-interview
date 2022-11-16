package service

import (
	"fmt"
	"github.com/rafimuhammad01/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"reflect"
	"testing"
)

func TestHash_Hash(t *testing.T) {
	type args struct {
		password string
		cost     int
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 []error
	}{
		{
			name:  "success",
			args:  args{password: "password", cost: 14},
			want:  "$2a$14$UH4mK8Kd3cSNs2RG7MwzG.kpmUA9O/w.B9fllYEgDrnBdT67sIeq2.",
			want1: nil,
		},
		{
			name: "handle empty string",
			args: args{password: "", cost: 14},
			want: "",
			want1: []error{
				fmt.Errorf("[%w] password cannot be empty", domain.ErrUserInvalidData),
			},
		},
		{
			name: "error from bcrypt",
			args: args{
				password: "test",
				cost:     34,
			},
			want:  "",
			want1: []error{fmt.Errorf("[%w] bcrypt error crypto/bcrypt: cost 34 is outside allowed range (4,31)", domain.ErrUserInternal)},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHash()
			got, got1 := h.Hash(tt.args.password, tt.args.cost)
			if err := bcrypt.CompareHashAndPassword([]byte(got), []byte(tt.args.password)); err != nil {
				if !reflect.DeepEqual(got1, tt.want1) {
					t.Errorf("Hash() got = %v, want %v", got, tt.want)
					t.Errorf("Hash() got1 = %v, want %v", got1, tt.want1)
				}
			}
		})
	}
}
