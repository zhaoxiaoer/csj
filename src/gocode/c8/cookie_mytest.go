package c8

import (
	"fmt"
	"net/http"
	"strconv"

	"../util"
)

type CookieTest struct{}

func (ct *CookieTest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ct.handleGet(w, r)
	case "POST":
		err = ct.handlePost(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ct *CookieTest) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	for i := 0; i < 3; i++ {
		cookie := http.Cookie{
			Name:  "Session-Cookie-" + strconv.Itoa(i),
			Value: "Cookie-Value-S" + strconv.Itoa(i),
		}
		http.SetCookie(w, &cookie)
		cookie2 := http.Cookie{
			Name:   "Persistent-Cookie-" + strconv.Itoa(i),
			Value:  "Cookie-Value-P" + strconv.Itoa(i),
			MaxAge: 3600,
		}
		http.SetCookie(w, &cookie2)
	}
	title := "Active Cookies"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>" +
		"<table border=1 align=\"center\">\n" +
		"<tr bgcolor=\"#FFAD00\">\n" +
		"  <th>Cookie Name</th>\n" +
		"  <th>Cookie Value</th>\n"))
	cookies := r.Cookies()
	if len(cookies) == 0 {
		w.Write([]byte("<tr><th colspan=2>No cookies</th></tr>\n"))
	} else {
		for i := 0; i < len(cookies); i++ {
			cookie := cookies[i]
			w.Write([]byte("<tr>\n" +
				"  <td>" + cookie.Name + "</td>\n" +
				"  <td>" + cookie.Value + "</td>\n" +
				"</tr>\n"))
		}
	}
	w.Write([]byte("</table></body></html>\n"))

	return
}

func (ct *CookieTest) handlePost(w http.ResponseWriter,
	r *http.Request) (err error) {
	err = ct.handleGet(w, r)

	return
}
