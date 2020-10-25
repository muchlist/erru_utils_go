package rest_err

import (
	"encoding/json"
	"errors"
	"net/http"
)

//APIError interface untuk mengembalikan error
type APIError interface {
	Message() string
	Status() int
	Error() string
}

type apiError struct {
	AStatus  int           `json:"status"`
	AMessage string        `json:"message"`
	Anerror  string        `json:"error"`
	ACauses  []interface{} `json:"causes"`
}

func (e *apiError) Status() int {
	return e.AStatus
}

func (e *apiError) Message() string {
	return e.AMessage
}

func (e *apiError) Error() string {
	return e.Anerror
}

//NewAPIError membuat error yang belum terdifinisikan
func NewAPIError(statusCode int, message string) APIError {
	return &apiError{
		AStatus:  statusCode,
		AMessage: message,
	}
}

//NewAPIErrorFromBytes membuat api error dari []byte
func NewAPIErrorFromBytes(body []byte) (APIError, error) {
	var result apiError
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.New("invalid json body")
	}

	return &result, nil
}

//NewNotFoundError membuat error tidak ditemukan
func NewNotFoundError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusNotFound,
		AMessage: message,
	}
}

//NewInternalServerError membuat error 500 internal
func NewInternalServerError(message string, err error) APIError {
	result := &apiError{
		AStatus:  http.StatusInternalServerError,
		AMessage: message,
		Anerror:  "internal_server_error",
		ACauses:  []interface{}{err.Error()},
	}
	if err != nil {
		result.ACauses = append(result.ACauses, err.Error())
	}
	return result
}

//NewBadRequestError membuat error jika kesalahan ada pada user
func NewBadRequestError(message string) APIError {
	return &apiError{
		AStatus:  http.StatusBadRequest,
		AMessage: message,
	}
}
