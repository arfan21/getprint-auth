package helpers

import (
	"net/http"
	"strings"

	"github.com/arfan21/getprint-service-auth/app/constants"
)

func GetStatusCode(err error) int {
	if strings.Contains(err.Error(), "Duplicate") {
		return http.StatusConflict
	}
	if strings.Contains(err.Error(), "not found") {
		return http.StatusNotFound
	}
	if strings.Contains(err.Error(), "password not match") {
		return http.StatusBadRequest
	}
	if strings.Contains(err.Error(), "expired") {
		return http.StatusUnauthorized
	}

	switch err {
	case constants.ErrBadParamInput:
		return http.StatusBadRequest
	case constants.ErrConflict:
		return http.StatusConflict
	case constants.ErrNotFound:
		return http.StatusNotFound
	case constants.ErrEmailConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
