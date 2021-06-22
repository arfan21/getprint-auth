package helpers

import (
	"net/http"
)

func NewRefreshTokenCookie(token string) *http.Cookie {
	cookieRefreshToken := new(http.Cookie)
	cookieRefreshToken.Name = "getprint-refresh-token"
	cookieRefreshToken.Value = token
	cookieRefreshToken.MaxAge = 0
	cookieRefreshToken.HttpOnly = true
	cookieRefreshToken.Secure = true
	cookieRefreshToken.Path = "/api/v1/auth/refresh-token"

	return cookieRefreshToken
}
