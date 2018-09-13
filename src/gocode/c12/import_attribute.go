package c12

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"../util"
)

var NO_VALUE string = "<i>No Value</i>"

type ImportAttribute struct{}

func (ia *ImportAttribute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ia.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ia *ImportAttribute) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	oldID := util.GetCookieValue(r, "userID", NO_VALUE)
	if oldID == NO_VALUE {
		newID := ia.randomID()
		cookie := &http.Cookie{
			Name:   "userID",
			Value:  newID,
			MaxAge: 60 * 60 * 24 * 365,
		}
		http.SetCookie(w, cookie)
	}

	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>{{.Title}}</title>
</head>
<body>
<h2>{{.Title}}</h2>
This page was accessed on {{.Time}} with a userID cookie of {{.OldID}}
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}
	data := struct {
		Title string
		Time  string
		OldID string
	}{
		Title: "The import Attribute",
		Time:  time.Now().Format(time.UnixDate),
		OldID: oldID,
	}
	return t.Execute(w, data)
}

func (ia *ImportAttribute) randomID() string {
	num := rand.Intn(10000000)
	return "id" + strconv.FormatInt(int64(num), 10)
}
