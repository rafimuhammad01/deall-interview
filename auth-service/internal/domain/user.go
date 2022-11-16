package domain

import (
	"fmt"
)

var (
	ErrInvalidDataLogin    = fmt.Errorf("invalid login data")
	ErrUserNotFound        = fmt.Errorf("user not found")
	ErrRefreshTokenInvalid = fmt.Errorf("refresh token not found or expired")
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
	Role     int    `json:"role" bson:"role"`
}

func (u *User) ValidateLogin() []error {
	var arrErr []error
	if u.Username == "" {
		arrErr = append(arrErr, fmt.Errorf("[%w] username should not be empty", ErrInvalidDataLogin))
	}

	if u.Password == "" {
		arrErr = append(arrErr, fmt.Errorf("[%w] password should not be empty", ErrInvalidDataLogin))
	}

	return arrErr
}
