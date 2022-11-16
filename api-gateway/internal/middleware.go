package internal

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strings"
)

var (
	ErrInvalidHeader = fmt.Errorf("invalid header")
	ErrInvalidToken  = fmt.Errorf("invalid token")
	ErrForbidden     = fmt.Errorf("forbidden access")
)

const (
	UserDataKey = "UserData"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) ExtractToken(r *http.Request) (string, error) {
	bearToken := r.Header.Get("Authorization")
	strArr := strings.Split(bearToken, " ")
	if len(strArr) != 2 {
		return "", fmt.Errorf("[%w] authorization token is needed", ErrInvalidHeader)
	}
	return strArr[1], nil
}

func (a *AuthMiddleware) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString, err := a.ExtractToken(r)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("[%w] unexpected signing method: %v", ErrInternal, token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		if !token.Valid {
			return nil, fmt.Errorf("[%w] token is invalid or expired", ErrInvalidToken)
		}
		return nil, fmt.Errorf("[%w] internal error when parsing jwt %s", ErrInternal, err.Error())
	}

	return token, nil
}

func (a *AuthMiddleware) TokenAuthMiddleware(role ...int) gin.HandlerFunc {
	return func(c *gin.Context) {
		payload, err := a.GrantAccess(c.Request, role...)
		if err != nil {
			errStruct := HandleError([]error{err})
			c.JSON(errStruct.Code, errStruct)
			c.Abort()
			return
		}
		c.Set(UserDataKey, payload)
		c.Next()
	}
}

func (a *AuthMiddleware) GrantAccess(r *http.Request, expectedRole ...int) (map[string]interface{}, error) {
	token, err := a.VerifyToken(r)
	if err != nil {
		return nil, err
	}

	payload, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("[%w] internal error when claims jwt token", ErrInvalidToken)
	}

	role := int(payload["role"].(float64))
	for _, v := range expectedRole {
		if role == v {
			return payload, nil
		}
	}

	return nil, fmt.Errorf("[%w] forbidden acces for role %s", ErrForbidden, roles[role])
}
