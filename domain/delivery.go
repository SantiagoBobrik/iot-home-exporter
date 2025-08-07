package domain

import "fmt"

type ResponseError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%d - %s", e.Code, e.Message)
}

type Response struct {
	Data []Data `json:"data"`
}
