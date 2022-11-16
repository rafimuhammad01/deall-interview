package domain

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var (
	ErrInternal = fmt.Errorf("internal server error")
)

type ErrorStruct struct {
	Code    int      `json:"-"`
	Message string   `json:"message"`
	Errors  []string `json:"errors,omitempty"`
}

func HandleError(arrErr []error) ErrorStruct {
	var (
		statusCode int
		message    string
		finalErr   []string
	)
	for _, v := range arrErr {
		parsed, err := parse(v.Error())
		if errors.Is(v, ErrInternal) || err != nil {
			statusCode = http.StatusInternalServerError
			log.Default().Println(v)
			return ErrorStruct{Code: statusCode, Message: "internal server error"}
		}

		finalErr = append(finalErr, parsed)
	}

	if errors.Is(arrErr[0], ErrInvalidDataLogin) {
		statusCode = http.StatusBadRequest
		message = ErrInvalidDataLogin.Error()
	} else if errors.Is(arrErr[0], ErrUserNotFound) {
		statusCode = http.StatusNotFound
		message = ErrUserNotFound.Error()
	} else if errors.Is(arrErr[0], ErrRefreshTokenInvalid) {
		statusCode = http.StatusNotFound
		message = ErrRefreshTokenInvalid.Error()
	}

	return ErrorStruct{
		Code:    statusCode,
		Message: message,
		Errors:  finalErr,
	}
}

func parse(err string) (string, error) {
	res := strings.SplitAfter(err, "]")
	if len(res) != 2 {
		return "", fmt.Errorf("[%w] %s", ErrInternal, "error format not matched")
	}
	return strings.TrimSpace(res[1]), nil
}
