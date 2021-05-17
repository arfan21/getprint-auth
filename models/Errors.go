package models

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrNotFound            = errors.New("error not found")
	ErrConflict            = errors.New("error conflict")
	ErrBadParamInput       = errors.New("error bad request")
	ErrEmailConflict       = errors.New("error duplicate email")
)
