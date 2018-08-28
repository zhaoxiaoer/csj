package c8

import (
	"fmt"
	"net/http"
	"strings"

	"../util"
)

type RepeatVisitor struct{}

func (rv *RepeatVisitor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = rv.handleGet(w, r)
	case "POST":
		err = rv.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rv *RepeatVisitor) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")

	newbie := true
	cookies := r.Cookies()
	if len(cookies) > 0 {
		for i := 0; i < len(cookies); i++ {
			cookie := cookies[i]
			if (strings.Compare(cookie.Name, "repeatVisitor") == 0) &&
				(strings.Compare(cookie.Value, "yes") == 0) {
				newbie = false
				break
			}
		}
	}
	title := ""
	if newbie {
		cookie := http.Cookie{
			Name:   "repeatVisitor",
			Value:  "yes",
			MaxAge: 60 * 60 * 24 * 365,
		}
		http.SetCookie(w, &cookie)
		title = "Welcome Aboard"
	} else {
		title = "Welcome Back"
	}
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>\n" +
		"</body></html>\n"))

	return
}

func (rv *RepeatVisitor) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = rv.handleGet(w, r)
	return
}
