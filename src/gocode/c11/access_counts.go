package c11

import (
	"fmt"
	"html/template"
	"net/http"
)

type AccessCounts struct {
	ACounts int
}

func (ac *AccessCounts) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ac.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ac *AccessCounts) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>{{.Title}}</title>
</head>
<body>
<h1>{{.Title}}</h1>

<h2>Accesses to page since server reboot: </h2>
{{.ACounts}}
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}

	ac.ACounts += 1
	data := struct {
		Title   string
		ACounts int
	}{
		Title:   "JSP Declaration",
		ACounts: ac.ACounts,
	}
	return t.Execute(w, data)
}
