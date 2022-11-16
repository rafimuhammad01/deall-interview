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

var (
	ApplicationJson = "application/json"
)

type Auth struct {
	BaseURL string
	helper  AuthHTTPHelper
}

type AuthHTTPHelper interface {
	Post(url string, body io.ReadCloser) (*http.Response, error)
}

func NewAuth(baseURL string, helper AuthHTTPHelper) *Auth {
	return &Auth{
		BaseURL: baseURL,
		helper:  helper,
	}
}

func (a *Auth) Login(c *gin.Context) {
	url := fmt.Sprintf("%s/api/v1/auth/login", a.BaseURL)

	resp, err := a.helper.Post(url, c.Request.Body)
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

func (a *Auth) RefreshToken(c *gin.Context) {
	url := fmt.Sprintf("%s/api/v1/auth/refresh", a.BaseURL)
	resp, err := a.helper.Post(url, c.Request.Body)
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
