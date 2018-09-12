package c11

import (
	"fmt"
	"html/template"
	"net/http"
)

type ThreeParams struct{}

func (tp *ThreeParams) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = tp.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (tp *ThreeParams) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>{{.Title}}</title>
</head>
<body>
<h2>{{.Title}}</h2>
<ul>
  <li><b>param1</b>: {{.Param1}}</li>
  <li><b>param2</b>: {{.Param2}}</li>
  <li><b>param3</b>: {{.Param3}}</li>
</ul>
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}

	data := struct {
		Title  string
		Param1 string
		Param2 string
		Param3 string
	}{
		Title:  "Reading Three Request Parameters",
		Param1: r.FormValue("param1"),
		Param2: r.FormValue("param2"),
		Param3: r.FormValue("param3"),
	}
	return t.Execute(w, data)
}
