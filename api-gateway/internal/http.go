package internal

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

var (
	MethodDelete = "DELETE"
)

type HTTPHelper struct {
}

func NewHTTPHelper() *HTTPHelper {
	return &HTTPHelper{}
}

func (h *HTTPHelper) Post(url string, body io.ReadCloser) (*http.Response, error) {
	reqBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	reqBodyBuffer := bytes.NewBuffer(reqBody)
	response, err := http.Post(url, ApplicationJson, reqBodyBuffer)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (h *HTTPHelper) Put(url string, body io.ReadCloser) (*http.Response, error) {
	client := &http.Client{}

	reqBody, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}

	reqBodyBuffer := bytes.NewBuffer(reqBody)
	req, err := http.NewRequest(MethodDelete, url, reqBodyBuffer)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (h *HTTPHelper) Delete(url string) (*http.Response, error) {
	client := &http.Client{}

	req, err := http.NewRequest(MethodDelete, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
