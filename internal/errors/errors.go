package errors

import "fmt"

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CustomError struct {
	Message string
	Err     error
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *CustomError) Unwrap() error {
	return e.Err
}
