package util

import (
	"net/http"
	"strings"
)

func GetCookieValue(r *http.Request, cookieName, defaultValue string) string {
	cookies := r.Cookies()
	if len(cookies) > 0 {
		for i := 0; i < len(cookies); i++ {
			cookie := cookies[i]
			if strings.Compare(cookieName, cookie.Name) == 0 {
				return cookie.Value
			}
		}
	}
	return defaultValue
}

func GetCookie(r *http.Request, cookieName string) *http.Cookie {
	cookies := r.Cookies()
	if len(cookies) > 0 {
		for i := 0; i < len(cookies); i++ {
			cookie := cookies[i]
			if strings.Compare(cookieName, cookie.Name) == 0 {
				return cookie
			}
		}
	}
	return nil
}
