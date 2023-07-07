package helper

import (
	"net/http"
	"time"
)

func GenerateCookie(name, value string, expires time.Time) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.HttpOnly = true
	cookie.Expires = expires
	cookie.Name = name
	cookie.Value = value

	return cookie
}
