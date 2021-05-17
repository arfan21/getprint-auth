package utils

import (
	"net/http"
	"strings"

	"github.com/arfan21/getprint-service-auth/models"
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

	switch err {
	case models.ErrBadParamInput:
		return http.StatusBadRequest
	case models.ErrConflict:
		return http.StatusConflict
	case models.ErrNotFound:
		return http.StatusNotFound
	case models.ErrEmailConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
