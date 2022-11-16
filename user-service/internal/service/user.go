package service

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rafimuhammad01/user-service/internal/domain"
)

type UserRepository interface {
	GetAll() ([]*domain.User, []error)
	GetByID(id string) (*domain.User, []error)
	GetByUsername(username string) (*domain.User, []error)
	Create(domain.User) []error
	Update(domain.User) []error
	Delete(id string) []error
}

type HashAlgo interface {
	Hash(password string, cost int) (string, []error)
}

type User struct {
	repository UserRepository
	hashAlgo   HashAlgo
}

func NewUser(repository UserRepository, hashAlgo HashAlgo) *User {
	return &User{
		repository: repository,
		hashAlgo:   hashAlgo,
	}
}

func (u *User) GetAll() ([]*domain.User, []error) {
	return u.repository.GetAll()
}

func (u *User) GetByID(id string) (*domain.User, []error) {
	return u.repository.GetByID(id)
}

func (u *User) Create(user domain.User) []error {
	var arrErr []error

	// assign id
	user.ID = uuid.New().String()

	// validate error
	err := user.Validate()
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	// validate username
	existingUser, err := u.repository.GetByUsername(user.Username)
	if existingUser != nil {
		arrErr = append(arrErr, fmt.Errorf("[%w] username %s already exist", domain.ErrUserExist, user.Username))
	}
	for _, v := range err {
		if errors.Is(v, domain.ErrUserNotFound) {
			continue
		} else {
			arrErr = append(arrErr, err...)
		}
	}

	// hashPassword
	hashPassword, err := u.hashAlgo.Hash(user.Password, 14)
	if err != nil {
		arrErr = append(arrErr, err...)
	}
	user.Password = hashPassword

	if arrErr != nil {
		return arrErr
	}

	// create user
	err = u.repository.Create(user)
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	return arrErr
}

func (u *User) Update(user domain.User) []error {
	// validate user
	var arrErr []error

	// validate error
	err := user.Validate()
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	// check if user id exist
	existingUser, err := u.repository.GetByID(user.ID)
	for _, v := range err {
		if existingUser == nil && errors.Is(v, domain.ErrUserNotFound) {
			arrErr = append(arrErr, fmt.Errorf("[%w] user with id %s not found", domain.ErrUserNotFound, user.ID))
		} else {
			arrErr = append(arrErr, err...)
		}
	}

	// validate username
	existingUser, err = u.repository.GetByUsername(user.Username)
	if existingUser != nil {
		arrErr = append(arrErr, fmt.Errorf("[%w] username %s already exist", domain.ErrUserExist, user.Username))
	}
	for _, v := range err {
		if errors.Is(v, domain.ErrUserNotFound) {
			continue
		} else {
			arrErr = append(arrErr, err...)
		}
	}

	// hashPassword
	hashPassword, err := u.hashAlgo.Hash(user.Password, 14)
	if err != nil {
		arrErr = append(arrErr, err...)
	}
	user.Password = hashPassword

	if arrErr != nil {
		return arrErr
	}

	// update user
	err = u.repository.Update(user)
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	return arrErr
}

func (u *User) Delete(id string) []error {
	var arrErr []error

	// check for existing user
	existingUser, err := u.repository.GetByID(id)
	for _, v := range err {
		if existingUser == nil && errors.Is(v, domain.ErrUserNotFound) {
			arrErr = append(arrErr, fmt.Errorf("[%w] user with id %s not found", domain.ErrUserNotFound, id))
		} else {
			arrErr = append(arrErr, err...)
		}
	}

	if arrErr != nil {
		return arrErr
	}

	// delete user
	err = u.repository.Delete(id)
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	return arrErr
}
