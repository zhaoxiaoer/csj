package c8

import (
	"fmt"
	"net/http"
	"strconv"

	"../util"
)

type ClientAccessCount struct{}

func (cac *ClientAccessCount) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = cac.handleGet(w, r)
	case "POST":
		err = cac.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (cac *ClientAccessCount) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	countString := util.GetCookieValue(r, "accessCount", "1")
	count := 1
	i, err := strconv.Atoi(countString)
	if err == nil {
		count = i
	}
	cookie := http.Cookie{
		Name:   "accessCount",
		Value:  strconv.Itoa(count + 1),
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, &cookie)
	title := "Access Count Servlet"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<center>\n" +
		"<h1>" + title + "</h1>\n" +
		"<h2>This is visit number " +
		strconv.Itoa(count) + " by this browser.</h2>\n" +
		"</center></body></html>\n"))

	return
}

func (cac *ClientAccessCount) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = cac.handleGet(w, r)

	return
}
