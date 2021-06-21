package helpers

import (
	"strings"

	"github.com/arfan21/getprint-service-auth/app/constants"
)

func CustomErrors(err error) string {
	if strings.Contains(err.Error(), "expired") {
		return constants.ErrTokenExpired.Error()
	}
	return err.Error()
}
