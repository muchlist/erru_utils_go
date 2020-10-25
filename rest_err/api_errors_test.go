package rest_err

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is the message")
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status())
	assert.EqualValues(t, "this is the message", err.Message())
	assert.EqualValues(t, "internal_server_error", err.Error())
}

func TestNewBadRequestError(t *testing.T) {

}

func TestNewNotFoundError(t *testing.T) {

}

func TestNewAPIError(t *testing.T) {

}
