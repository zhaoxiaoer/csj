package c4

import (
	"html"
	"net/http"

	"../util"
)

type BadCode struct {
}

func (bc *BadCode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = bc.handleGet(w, r)
	case "POST":
		err = bc.handlePost(w, r)
	case "PUT":
		err = bc.handlePut(w, r)
	case "DELETE":
		err = bc.handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bc *BadCode) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	//	w.Write([]byte("get"))
	w.Header().Add("Content-Type", "text/html")
	title := "Code Sample"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>\n" +
		"<pre>\n" +
		getCode(r) +
		"</pre>\n" +
		"Now, wasn't that an interesting sample\n" +
		"of code?\n" +
		"</body></html>"))

	return
}

func (bc *BadCode) handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//	w.Write([]byte("post"))
	err = bc.handleGet(w, r)

	return
}

func (bc *BadCode) handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func (bc *BadCode) handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}

func getCode(r *http.Request) string {
	return html.EscapeString(r.FormValue("code"))
}
