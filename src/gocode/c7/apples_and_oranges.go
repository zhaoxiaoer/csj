package c7

import (
	"fmt"
	"net/http"
)

type ApplesAndOranges struct{}

func (aao *ApplesAndOranges) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = aao.handleGet(w, r)
	case "POST":
		err = aao.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (aao *ApplesAndOranges) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "application/vnd.ms-excel")

	w.Write([]byte("\tQ1\tQ2\tQ3\tQ4\tTotal\n"))
	w.Write([]byte("Apples\t78\t87\t92\t29\t=sum(B2:E2)\n"))
	w.Write([]byte("Oranges\t77\t86\t93\t30\t=sum(B3:E3)\n"))
	return
}

func (aao *ApplesAndOranges) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = aao.handleGet(w, r)
	return
}
