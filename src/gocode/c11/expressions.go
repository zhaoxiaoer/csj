package c11

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime"
	"time"

	"../util"
)

type Expressions struct{}

func (ex *Expressions) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = ex.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (ex *Expressions) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>JSP Expressions</title>
</head>
<body>
<h2>JSP Expressions</h2>
<ul>
  <li>Current time: {{.CurrTime}}</li>
  <li>Server: {{.ServerInfo}}</li>
  <li>Session ID: {{.SessionID}}</li>
  <li>The <code>testParam</code> form parameter: {{.TestParam}}</li>
</ul>
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}

	data := struct {
		CurrTime   string
		ServerInfo string
		SessionID  string
		TestParam  string
	}{
		CurrTime:   time.Now().Format(time.UnixDate),
		ServerInfo: runtime.Version(),
		SessionID:  util.SM.GetSession(w, r).GetId(),
		TestParam:  r.FormValue("testParam"),
	}
	return t.Execute(w, data)
}
