package c5

import (
	"bytes"
	"net/http"

	"../util"
)

type LongServ struct {
}

func (ls *LongServ) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ls.handleGet(w, r)
	case "POST":
		err = ls.handlePost(w, r)
	case "PUT":
		err = ls.handlePut(w, r)
	case "DELETE":
		err = ls.handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ls *LongServ) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	title := "Long Page"
	htm := util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>\n"
	line := "Blah, blah, blah, blah, blah. Yadda, yadda, yadda, yadda.\n"
	//	for i := 0; i < 10000; i++ {
	//		htm += line
	//	}
	//	htm += "</body></html>\n"
	var b bytes.Buffer
	b.WriteString(htm)
	for i := 0; i < 10000; i++ {
		b.WriteString(line)
	}
	b.WriteString("</body></html>\n")

	if util.IsGzipSupported(r) && !util.IsGzipDisabled(r) {
		w.Header().Set("Content-Encoding", "gzip")
		w.Write(util.GetGzipData(b.String()))
	} else {
		w.Write(b.Bytes())
	}

	return
}

func (ls *LongServ) handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	err = ls.handleGet(w, r)

	return
}

func (ls *LongServ) handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func (ls *LongServ) handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}
