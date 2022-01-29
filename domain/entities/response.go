package entities

import "net/http"

type Response struct {
	Status     string      `json:"status,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (r Response) Ok(data interface{}) Response {
	r.Status = "200 OK"
	r.StatusCode = http.StatusOK
	r.Message = "OK"
	r.Data = data

	return r
}

func (r Response) Bad() Response {
	r.Status = "400 Bad Request"
	r.StatusCode = http.StatusBadRequest

	return r
}

func (r Response) Unauthorized() Response {
	r.Status = "401 Unauthorized"
	r.StatusCode = http.StatusUnauthorized

	return r
}
