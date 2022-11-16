package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/auth-service/internal/domain"
	"net/http"
)

type Auth struct {
	authService AuthService
}

type AuthService interface {
	Login(domain.User) (string, string, []error)
	CreateAccessTokenWithRefresh(string) (string, []error)
}

func NewUser(service AuthService) *Auth {
	return &Auth{
		authService: service,
	}
}

func (u *Auth) Login(c *gin.Context) {
	var reqBody domain.User

	err := c.BindJSON(&reqBody)
	if err != nil {
		errStruct := domain.HandleError([]error{fmt.Errorf("[%w] %s", domain.ErrInternal, err.Error())})
		c.JSON(errStruct.Code, errStruct)
		return
	}

	accessToken, refreshToken, arrErr := u.authService.Login(reqBody)
	if arrErr != nil {
		errStruct := domain.HandleError(arrErr)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, domain.JSONResponse{
		Message: "ok",
		Data: map[string]interface{}{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}

func (u *Auth) RefreshToken(c *gin.Context) {
	var reqBody map[string]interface{}

	err := c.BindJSON(&reqBody)
	if err != nil {
		errStruct := domain.HandleError([]error{fmt.Errorf("[%w] %s", domain.ErrInternal, err.Error())})
		c.JSON(errStruct.Code, errStruct)
		return
	}

	refreshToken := reqBody["refresh_token"].(string)
	newToken, arrErr := u.authService.CreateAccessTokenWithRefresh(refreshToken)
	if arrErr != nil {
		errStruct := domain.HandleError(arrErr)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, domain.JSONResponse{
		Message: "ok",
		Data: map[string]interface{}{
			"access_token":  newToken,
			"refresh_token": refreshToken,
		},
	})

}
