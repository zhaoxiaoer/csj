package c11

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

type BGColor struct{}

func (bgc *BGColor) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = bgc.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bgc *BGColor) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>{{.Title}}</title>
</head>
<body bgcolor={{.BgColor}}>
<h2 align="center">Testing a Background of "{{.BgColor}}"</h2>
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}

	bgColor := r.FormValue("bgColor")
	bgColor = strings.Trim(bgColor, " ")
	if strings.Compare(bgColor, "") == 0 {
		bgColor = "white"
	}
	data := struct {
		Title   string
		BgColor string
	}{
		Title:   "Color Testing",
		BgColor: bgColor,
	}
	return t.Execute(w, data)
}
