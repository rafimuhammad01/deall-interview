package service

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/rafimuhammad01/auth-service/internal/domain"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type JWT struct {
	secretKey      string
	userRepository UserRepository
	expiredTime    int
}

type UserRepository interface {
	GetByUsername(username string) (*domain.User, []error)
	SaveToken(userid string, role int, refreshToken string) []error
	GetUserDataFromToken(refreshToken string) (map[string]interface{}, []error)
}

func NewJWT(secretKey string, repository UserRepository, expiredTime int) *JWT {
	return &JWT{
		secretKey:      secretKey,
		userRepository: repository,
		expiredTime:    expiredTime,
	}
}

func (j *JWT) Login(u domain.User) (string, string, []error) {
	var arrErr []error

	err := u.ValidateLogin()
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	user, err := j.userRepository.GetByUsername(u.Username)
	for _, v := range err {
		if errors.Is(v, domain.ErrUserNotFound) {
			return "", "", []error{fmt.Errorf("[%w] invalid username or password", domain.ErrInvalidDataLogin)}
		}

		arrErr = append(arrErr, err...)
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return "", "", []error{fmt.Errorf("[%w] invalid username or password", domain.ErrInvalidDataLogin)}
		}

		arrErr = append(arrErr, fmt.Errorf("[%w] error when compare hash password with bcrypt %s", domain.ErrInternal, err.Error()))
	}

	if arrErr != nil {
		return "", "", arrErr
	}

	accessToken, err := j.CreateAccessToken(user.ID, user.Role)
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	refreshToken, err := j.CreateRefreshToken(user.ID, user.Role)
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	if arrErr != nil {
		return "", "", arrErr
	}

	return accessToken, refreshToken, nil
}

func (j *JWT) CreateRefreshToken(userid string, role int) (string, []error) {
	var arrErr []error

	refreshToken := uuid.New().String()
	err := j.userRepository.SaveToken(userid, role, refreshToken)
	if err != nil {
		arrErr = append(arrErr, err...)
	}

	if arrErr != nil {
		return "", arrErr
	}

	return refreshToken, nil
}

func (j *JWT) CreateAccessTokenWithRefresh(token string) (string, []error) {
	data, err := j.userRepository.GetUserDataFromToken(token)
	if err != nil {
		return "", err
	}

	userID := data["user_id"].(string)
	role := int(data["role"].(float64))

	accessToken, err := j.CreateAccessToken(userID, role)
	if err != nil {
		return "", err
	}

	return accessToken, nil
}

func (j *JWT) CreateAccessToken(userid string, role int) (string, []error) {
	var arrErr []error

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["role"] = role
	atClaims["exp"] = time.Now().Add(time.Minute * time.Duration(j.expiredTime)).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	// sign token
	token, err := at.SignedString([]byte(j.secretKey))
	if err != nil {
		arrErr = append(arrErr, fmt.Errorf("[%w] internal error when sign jwt %s", domain.ErrInternal, err.Error()))
	}

	if arrErr != nil {
		return "", arrErr
	}

	return token, nil
}
