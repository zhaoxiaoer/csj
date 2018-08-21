package c5

import (
	"fmt"
	"net/http"
	"strings"

	"../util"
)

type BrowserInsult struct {
}

func (bi *BrowserInsult) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = bi.handleGet(w, r)
	case "POST":
		err = bi.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bi *BrowserInsult) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	var title, msg string
	if strings.Index(r.UserAgent(), "MSIE") != -1 {
		title = "Microsoft Minion"
		msg = "Welcome, O spineless slave to the mighty empire."
	} else {
		title = "Hopeless Netscape Rebel"
		msg = "Enjoy it while you can. You <i>will</i> be assimilated!"
	}
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>\n" +
		msg + "\n" +
		"</body></html>"))

	return
}

func (bi *BrowserInsult) handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	err = bi.handleGet(w, r)

	return
}
