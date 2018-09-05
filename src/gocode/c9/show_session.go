package c9

import (
	"fmt"
	"net/http"
	"strconv"

	"../util"
	//	"github.com/gorilla/sessions"
)

type ShowSession struct{}

func (ss *ShowSession) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ss.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ss *ShowSession) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	session := util.GetSession(r)
	var heading string
	var accessCount int
	var ok bool
	if accessCount, ok = session.Values["accessCount"].(int); !ok {
		accessCount = 0
		heading = "Welcome, Newcomer"
	} else {
		heading = "Welcome Back"
		accessCount += 1
	}
	session.Values["accessCount"] = accessCount
	session.Save(r, w)
	title := "Session Tracking Example"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<center>\n" +
		"<h1>" + heading + "</h1>\n" +
		"<h2>Information on Your Session:</h2>\n" +
		"<table border=1>\n" +
		"<tr bgcolor=\"#FFAD00\">\n" +
		"  <th>Info Type</th><th>Value</th>\n" +
		"</tr>\n" +
		"<tr>\n" +
		"  <td>ID</td><td>" + session.ID + "</td>\n" +
		"</tr>\n" +
		"<tr>\n" +
		"  <td>Creation Time</td><td>" + "" + "</td>\n" +
		"</tr>\n" +
		"<tr>\n" +
		"  <td>Time of Last Access</td><td>" + "" + "</td>\n" +
		"</tr>\n" +
		"<tr>\n" +
		"  <td>Number of Previous Accesses</td><td>" + strconv.Itoa(accessCount) + "</td>\n" +
		"</tr>\n" +
		"</table>\n" +
		"</center></body></html>"))

	return
}
