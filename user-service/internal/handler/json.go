package handler

type JSONResp struct {
	Message string `json:"message"`
	Data interface{} `json:"data,omitempty"`
}
