package resterrors

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

//RestErr is interface to expose restErr strut to outside package
type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []interface{}
}

type restErr struct {
	ErrMessage string        `json:"message"`
	ErrStatus  int           `json:"status"`
	ErrError   string        `json:"error"`
	ErrCauses  []interface{} `json:"causes"`
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: %v",
		e.ErrMessage, e.ErrStatus, e.ErrError, e.ErrCauses)
}

func (e restErr) Message() string {
	return e.ErrMessage
}

func (e restErr) Status() int {
	return e.ErrStatus
}

func (e restErr) Causes() []interface{} {
	return e.ErrCauses
}

// NewRestError returns json formarted error message
func NewRestError(message string, status int, err string, causes []interface{}) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  status,
		ErrError:   err,
		ErrCauses:  causes,
	}
}

//NewRestErrorFromBytes retruns json formated error message for byte slice input
func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, errors.New("invalid json")
	}
	return apiErr, nil
}

//NewBadRequestError retruns json formated error message for "Bad Request"
func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError:   "bad_request",
	}
}

//NewNotFoundError retruns json formated error message for "resource not found"
func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusNotFound,
		ErrError:   "not_found",
	}
}

//NewUnauthorizedError retruns json formated error message for "authorization failuer"
func NewUnauthorizedError(message string) RestErr {
	return restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusUnauthorized,
		ErrError:   "unauthorized",
	}
}

//NewInternalServerError retruns json formated error message for "internal error"
func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrMessage: message,
		ErrStatus:  http.StatusInternalServerError,
		ErrError:   "internal_server_error",
	}
	if err != nil {
		result.ErrCauses = append(result.ErrCauses, err.Error())
	}
	return result
}
