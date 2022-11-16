package service

import (
	"fmt"

	"github.com/rafimuhammad01/user-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
)

type Hash struct{}

func NewHash() *Hash {
	return &Hash{}
}

func (h *Hash) Hash(password string, cost int) (string, []error) {
	if password == "" {
		return "", []error{fmt.Errorf("[%w] password cannot be empty", domain.ErrUserInvalidData)}
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", []error{fmt.Errorf("[%w] bcrypt error %s", domain.ErrUserInternal, err.Error())}
	}

	return string(bytes), nil
}
