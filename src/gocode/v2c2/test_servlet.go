package v2c2

import (
	"fmt"
	"net/http"
)

type TServlet1 struct{}

func (ts *TServlet1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ts.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ts *TServlet1) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	uri := r.URL.RequestURI()
	w.Write([]byte("<!DOCTYPE html>\n" +
		"<html>\n" +
		"<head>\n" +
		"<title>" + "Test Servlet 1" + "</title>\n" +
		"</head>\n" +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h2>" + "Servlet Name: Test1" + "</h2>\n" +
		"<h2>URI: " + uri + "</h2>\n" +
		"</body>\n" +
		"</html>"))

	return
}

type TServlet2 struct{}

func (ts *TServlet2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ts.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ts *TServlet2) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	uri := r.URL.RequestURI()
	w.Write([]byte("<!DOCTYPE html>\n" +
		"<html>\n" +
		"<head>\n" +
		"<title>" + "Test Servlet 1" + "</title>\n" +
		"</head>\n" +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h2>" + "Servlet Name: Test1" + "</h2>\n" +
		"<h2>URI: " + uri + "</h2>\n" +
		"</body>\n" +
		"</html>"))

	return
}
