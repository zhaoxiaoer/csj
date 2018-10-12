package v2c2

import (
	"fmt"
	"html/template"
	"net/http"
)

type NotFound struct{}

func (nf *NotFound) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		if r.URL.Path == "/" {
			http.Redirect(w, r, "/csj/", http.StatusFound)
			return
		}

		err = nf.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (nf *NotFound) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	t, err := template.ParseFiles("errorp/NotFound.tpl")
	if err != nil {
		return
	}

	return t.Execute(w, r.URL.Path)
}
