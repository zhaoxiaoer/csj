package c6

import (
	"fmt"
	"net/http"
	"strings"
)

type WrongDestination struct{}

func (wd *WrongDestination) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = wd.handleGet(w, r)
	case "POST":
		err = wd.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wd *WrongDestination) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	if strings.Index(r.UserAgent(), "MSIE") != -1 {
		http.Redirect(w, r, "http://home.netscape.com", http.StatusFound)
	} else {
		http.Redirect(w, r, "http://www.microsoft.com", http.StatusFound)
	}

	return
}

func (wd *WrongDestination) handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	err = wd.handleGet(w, r)
	return
}
