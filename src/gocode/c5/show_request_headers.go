package c5

import (
	"fmt"
	"net/http"

	"../util"
)

type ShowRequestHeaders struct {
}

func (srh *ShowRequestHeaders) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = srh.handleGet(w, r)
	case "POST":
		err = srh.handlePost(w, r)
	case "PUT":
		err = srh.handlePut(w, r)
	case "DELETE":
		err = srh.handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (srh *ShowRequestHeaders) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	title := "Servlet Example: Showing Request Headers"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>\n" +
		"<b>Request Method: </b>" + r.Method + "<br>\n" +
		"<b>Request URI: </b>" + r.RequestURI + "<br>\n" +
		"<b>Request Protocol: </b>" + r.Proto + "<br>\n" +
		"<table border=1 align=\"center\">\n" +
		"<tr bgcolor=\"#FFAD00\">\n" +
		"  <th>Header Name</th>\n" +
		"  <th>Header Value</th>\n" +
		"</tr>\n"))
	for k, v := range r.Header {
		w.Write([]byte("<tr>\n  <td>" + k + "</td>\n"))
		w.Write([]byte("  <td>" + v[0] + "</td>\n</tr>\n"))
	}
	w.Write([]byte("</table>\n"))

	fmt.Printf("%s, %s\n", r.Header.Get("Host"), r.Host)

	w.Write([]byte("<br><br><br><table border=1 align=\"center\">\n" +
		"<tr bgcolor=\"#FFAD00\">\n" +
		"  <th>Header Name</th>\n" +
		"  <th>Header Value</th>\n" +
		"</tr>\n"))
	w.Header().Set("Name", "zhaoxiaoer")
	for k, v := range w.Header() {
		w.Write([]byte("<tr>\n  <td>" + k + "</td>\n"))
		w.Write([]byte("  <td>" + v[0] + "</td>\n</tr>\n"))
	}
	w.Write([]byte("</table>\n"))

	w.Write([]byte("</body></html>"))

	return
}

func (srh *ShowRequestHeaders) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = srh.handleGet(w, r)

	return
}

func (srh *ShowRequestHeaders) handlePut(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func (srh *ShowRequestHeaders) handleDelete(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}
