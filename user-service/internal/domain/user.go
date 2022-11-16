package domain

import "fmt"

var (
	ErrUserInvalidData = fmt.Errorf("invalid data")
	ErrUserExist       = fmt.Errorf("user already exist")
	ErrUserNotFound    = fmt.Errorf("user not found")
	ErrUserInternal    = fmt.Errorf("internal server error")
)

type User struct {
	ID       string `json:"id" bson:"_id"`
	Name     string `json:"name" bson:"name"`
	Username string `json:"username" bson:"username"`
	Role     int    `json:"role"`
	Password string `json:"-" bson:"password"`
}

func (u *User) Validate() []error {
	var arrErr []error

	if u.ID == "" {
		arrErr = append(arrErr, fmt.Errorf("[%w] id should not be empty", ErrUserInvalidData))
	}

	if u.Name == "" {
		arrErr = append(arrErr, fmt.Errorf("[%w] name should not be empty", ErrUserInvalidData))
	}

	if u.Username == "" {
		arrErr = append(arrErr, fmt.Errorf("[%w] username should not be empty", ErrUserInvalidData))
	}

	if u.Password == "" {
		arrErr = append(arrErr, fmt.Errorf("[%w] password should not be empty", ErrUserInvalidData))
	}

	if u.Role != 0 && u.Role != 1 {
		arrErr = append(arrErr, fmt.Errorf("[%w] role should be only 0 (user) or 1 (admin)", ErrUserInvalidData))
	}

	return arrErr
}
