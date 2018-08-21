package c6

import (
	"fmt"
	"net/http"
	"net/url"

	"../util"
)

type SearchEngines struct{}

func (se *SearchEngines) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = se.handleGet(w, r)
	case "POST":
		err = se.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (se *SearchEngines) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	searchString := r.FormValue("searchString")
	if searchString == "" {
		se.reportProblem(w, "Missing search string")
		return
	}
	searchString = url.QueryEscape(searchString)

	searchEngineName := r.FormValue("searchEngine")
	if searchEngineName == "" {
		se.reportProblem(w, "Missing search engine name")
		return
	}

	searchURL := util.MakeURL(searchEngineName, searchString)
	if searchURL != "" {
		http.Redirect(w, r, searchURL, http.StatusFound)
	} else {
		se.reportProblem(w, "Unrecognized search engine")
	}

	return
}

func (se *SearchEngines) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = se.handleGet(w, r)

	return
}

func (se *SearchEngines) reportProblem(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusNotFound)
}
