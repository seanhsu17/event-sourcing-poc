package render

import "net/http"

type ErrCode int

var (
	// ReqArgumentErr request argument error
	ReqArgumentErr ErrCode = 1000

	// UnknownErr unknown error type
	UnknownErr ErrCode = 5000

	// ProcessErr process error type
	ProcessErr ErrCode = 5001
)

type AppError struct {
	// Code is http status code
	Code int
	// ErrCode is system error code
	ErrCode ErrCode
	// Message is error message
	Message string
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

func NewValidationError(message string) *AppError {
	return &AppError{
		Message: message,
		ErrCode: ReqArgumentErr,
		Code:    http.StatusUnprocessableEntity,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		ErrCode: UnknownErr,
		Code:    http.StatusInternalServerError,
	}
}

func NewProcessError(message string) *AppError {
	return &AppError{
		Message: message,
		ErrCode: ProcessErr,
		Code:    http.StatusInternalServerError,
	}
}
