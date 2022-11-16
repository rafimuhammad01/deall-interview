package internal

import (
	"bytes"
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
	userData, _ := c.Get(UserDataKey)
	ID := userData.(map[string]interface{})["user_id"].(string)
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

func (u *User) CreateAdmin(c *gin.Context) {
	url := fmt.Sprintf("%s/api/v1/user", u.BaseURL)

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	var reqBody map[string]interface{}
	err = json.Unmarshal(body, &reqBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	reqBody["role"] = 1

	reqBodyByte, err := json.Marshal(reqBody)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	reqBodyBuffer := bytes.NewBuffer(reqBodyByte)
	reqBodyReadClose := io.NopCloser(reqBodyBuffer)

	resp, err:= u.helper.Post(url, reqBodyReadClose)
	if err != nil {
		log.Default().Println(err)
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "internal server error",
		})
		return
	}

	body, err = ioutil.ReadAll(resp.Body)
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
