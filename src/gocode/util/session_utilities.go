package util

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("httpsessiontest"))

type MyHttpSession struct {
	*sessions.Session
}

func GetSession(w http.ResponseWriter, r *http.Request) *MyHttpSession {
	session, _ := store.Get(r, "JSESSIONID")
	session.Options = &sessions.Options{
		Path:     "/csj",
		MaxAge:   30,
		HttpOnly: true,
	}
	return &MyHttpSession{session}
	//	return session
}
