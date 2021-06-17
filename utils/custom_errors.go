package utils

import (
	"strings"

	"github.com/arfan21/getprint-service-auth/models"
)

func CustomErrors(err error) string {
	if strings.Contains(err.Error(), "expired") {
		return models.ErrTokenExpired.Error()
	}
	return err.Error()
}
