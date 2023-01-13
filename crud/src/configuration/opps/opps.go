package opps

import (
	"net/http"
)

type Opps struct {
	Message string `json:"message"`
	Err string `json:"error"`
	Code int	`json:"code"`
	Causes []Causes `json:"causes,omitempy"`
}

type Causes struct {
	Field string	`json:"field"`
	Message string	`json:"message"`
}

func (o *Opps) Error() string {
	return o.Message
}

func NewOpps(message, err string, code int, causes []Causes) *Opps {
	return &Opps{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBadRequestError(message string) *Opps {
	return &Opps{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Causes) *Opps {
	return &Opps{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *Opps {
	return &Opps{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusInternalServerError,
	}
}

func NewNotFoundError(message string) *Opps {
	return &Opps{
		Message: message,
		Err: "not_found",
		Code: http.StatusNotFound,
	}
}

func NewForbiddenError(message string) *Opps {
	return &Opps{
		Message: message,
		Err: "forbidden",
		Code: http.StatusForbidden,
	}
}