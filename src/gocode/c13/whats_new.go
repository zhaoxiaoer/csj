package c13

import (
	"fmt"
	"html/template"
	"net/http"
)

type WhatsNew struct{}

func (wn *WhatsNew) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = wn.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (wn *WhatsNew) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	t, err := template.ParseFiles("c13/layout.tpl", "c13/item1.tpl",
		"c13/item2.tpl", "c13/item3.tpl")
	if err != nil {
		return
	}

	return t.ExecuteTemplate(w, "layout", nil)
}
