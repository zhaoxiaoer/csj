package c6

import (
	"fmt"
	"net/http"

	"../util"
)

type SearchEngineForm struct{}

func (sef *SearchEngineForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = sef.handleGet(w, r)
	case "POST":
		err = sef.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (sef *SearchEngineForm) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	title := "One-Stop Web Search!"
	actionURL := "/csj/c6/searchengines"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<center>\n" +
		"<h1>" + title + "</h1>\n" +
		"<form action=\"" + actionURL + "\">\n" +
		"  Search keywords: \n" +
		"  <input type=\"text\" name=\"searchString\"><p>\n\n"))
	specs := util.GetCommonSpecs()
	for i := 0; i < len(specs); i++ {
		w.Write([]byte("<input type=\"radio\" name=\"searchEngine\" " +
			"value=\"" + specs[i].Name + "\">\n\n"))
		w.Write([]byte(specs[i].Name + "<br>\n\n"))
	}
	w.Write([]byte("<br>  <input type=\"submit\">\n" +
		"</form>\n" +
		"</center></body></html>\n\n"))

	return
}

func (sef *SearchEngineForm) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = sef.handleGet(w, r)
	return
}
