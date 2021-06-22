package helpers

import "net/http"

func RemoveRefreshTokenCookie() *http.Cookie {
	cookieRefreshToken := new(http.Cookie)
	cookieRefreshToken.Name = "getprint-refresh-token"
	cookieRefreshToken.Value = ""
	cookieRefreshToken.MaxAge = -1
	cookieRefreshToken.HttpOnly = true
	cookieRefreshToken.Secure = true
	cookieRefreshToken.Path = "/api/v1/auth/refresh-token"

	return cookieRefreshToken
}
