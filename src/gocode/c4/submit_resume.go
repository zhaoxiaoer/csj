package c4

import (
	"fmt"
	"net/http"
)

type SubmitResume struct {
}

func (sr *SubmitResume) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("111"))

	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("post"))

	r.ParseForm()
	fmt.Println(r.PostForm.Get("name"))

	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}
