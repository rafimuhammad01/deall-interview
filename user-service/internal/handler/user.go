package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafimuhammad01/user-service/internal/domain"
	"github.com/rafimuhammad01/user-service/internal/dto"
)

type UserService interface {
	GetAll() ([]*domain.User, []error)
	Create(domain.User) []error
	GetByID(id string) (*domain.User, []error)
	Update(user domain.User) []error
	Delete(id string) []error
}

type User struct {
	service UserService
}

func NewUser(service UserService) *User {
	return &User{
		service: service,
	}
}

func (u *User) GetAll(c *gin.Context) {
	resp, err := u.service.GetAll()
	if err != nil {
		errStruct := domain.HandleError(err)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, JSONResp{Message: "OK", Data: resp})
}

func (u *User) GetByID(c *gin.Context) {
	id := c.Param("id")

	resp, arrErr := u.service.GetByID(id)
	if arrErr != nil {
		errStruct := domain.HandleError(arrErr)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, JSONResp{Message: "OK", Data: resp})
}

func (u *User) Create(c *gin.Context) {
	var reqBody dto.CreateAndUpdateUserReq

	err := c.Bind(&reqBody)
	if err != nil {
		errStruct := domain.HandleError([]error{fmt.Errorf("[%w] %s", domain.ErrUserInternal, err.Error())})
		c.JSON(errStruct.Code, errStruct)
		return
	}

	userDomainModel := domain.User{
		Username: reqBody.Username,
		Name:     reqBody.Name,
		Password: reqBody.Password,
		Role:     reqBody.Role,
	}

	arrErr := u.service.Create(userDomainModel)
	if arrErr != nil {
		errStruct := domain.HandleError(arrErr)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, JSONResp{Message: "OK"})
}

func (u *User) Update(c *gin.Context) {
	var reqBody dto.CreateAndUpdateUserReq

	err := c.Bind(&reqBody)
	if err != nil {
		errStruct := domain.HandleError([]error{fmt.Errorf("[%w] %s", domain.ErrUserInternal, err.Error())})
		c.JSON(errStruct.Code, errStruct)
		return
	}

	userDomainModel := domain.User{
		ID:       c.Param("id"),
		Username: reqBody.Username,
		Name:     reqBody.Name,
		Password: reqBody.Password,
		Role:     reqBody.Role,
	}

	arrErr := u.service.Update(userDomainModel)
	if arrErr != nil {
		errStruct := domain.HandleError(arrErr)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, JSONResp{Message: "OK"})
}

func (u *User) Delete(c *gin.Context) {
	ID := c.Param("id")

	arrErr := u.service.Delete(ID)
	if arrErr != nil {
		errStruct := domain.HandleError(arrErr)
		c.JSON(errStruct.Code, errStruct)
		return
	}

	c.JSON(http.StatusOK, JSONResp{Message: "OK"})
}
