package entities

import (
	"net/http"
)

type Response struct {
	Status     string      `json:"status,omitempty"`
	StatusCode int         `json:"statusCode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func (r Response) Build(message string, data interface{}) Response {
	r.Message = message
	r.Data = data
	return r
}

func (r Response) Ok() Response {
	r.Status = "200 OK"
	r.StatusCode = http.StatusOK

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

func (r Response) IntervalServerError() Response {
	r.Status = "500 Internal Server Error"
	r.StatusCode = http.StatusInternalServerError

	return r
}
