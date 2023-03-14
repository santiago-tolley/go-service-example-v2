package entity

import "fmt"

type ErrorType string

const (
	INPUT_ERROR   ErrorType = "input"
	SERVICE_ERROR ErrorType = "service"
	CLIENT_ERROR  ErrorType = "client"
	UNKNOWN_ERROR ErrorType = "unkown"
)

type TypedError struct {
	Type    ErrorType
	Code    int
	Message string
}

func (e TypedError) Error() string {
	return fmt.Sprintf("%v error (%v): %v", e.Type, e.Code, e.Message)
}
