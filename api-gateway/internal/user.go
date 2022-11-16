package internal

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	RoleUser  = 0
	RoleAdmin = 1
)

var (
	roles = map[int]string{
		0: "user",
		1: "admin",
	}
)

type User struct {
	BaseURL string
	helper  UserHTTPHelper
}

type UserHTTPHelper interface {
	Post(url string, body io.ReadCloser) (*http.Response, error)
	Put(url string, body io.ReadCloser) (*http.Response, error)
	Delete(url string) (*http.Response, error)
}

func NewUser(baseURL string, helper UserHTTPHelper) *User {
	return &User{
		BaseURL: baseURL,
		helper:  helper,
	}
}

func (u *User) GetAll(c *gin.Context) {
	url := fmt.Sprintf("%s/api/v1/user", u.BaseURL)

	resp, err := http.Get(url)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	c.JSON(resp.StatusCode, respBody)
}

func (u *User) GetByID(c *gin.Context) {
	ID := c.Param("id")
	url := fmt.Sprintf("%s/api/v1/user/%s", u.BaseURL, ID)

	resp, err := http.Get(url)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	c.JSON(resp.StatusCode, respBody)
}

func (u *User) Create(c *gin.Context) {
	url := fmt.Sprintf("%s/api/v1/user", u.BaseURL)

	resp, err := u.helper.Post(url, c.Request.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	c.JSON(resp.StatusCode, respBody)
}

func (u *User) Update(c *gin.Context) {
	ID := c.Param("id")
	url := fmt.Sprintf("%s/api/v1/user/%s", u.BaseURL, ID)

	resp, err := u.helper.Put(url, c.Request.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	c.JSON(resp.StatusCode, respBody)
}

func (u *User) Delete(c *gin.Context) {
	ID := c.Param("id")
	url := fmt.Sprintf("%s/api/v1/user/%s", u.BaseURL, ID)

	resp, err := u.helper.Delete(url)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	var respBody map[string]interface{}
	err = json.Unmarshal(body, &respBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	c.JSON(resp.StatusCode, respBody)
}
