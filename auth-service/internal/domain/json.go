package domain

type JSONResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
