package c14

import (
	"fmt"
	"html/template"
	"net/http"

	//	"./bean"
)

type StringBean struct{}

func (sb *StringBean) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = sb.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (sb *StringBean) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Using JavaBeans with JSP</title>
</head>
<body>
<table border=5 align="center">
  <tr>
    <th>Using JavaBeans with JSP</th>
  </tr>
</table>
<ol>
  <li>Initial value (from jsp:getProperty):
    <i>{{.Message}}</i>
  </li>
  <li>Initial value (from JSP expression):
    <i>{{.Message}}</i>
  </li>
  <li>
    Value after setting property with jsp:setProperty:
    <i>{{.Message2}}</i>
  </li>
  <li>
    Value after setting property with scriptlet:
    <i>{{.Message3}}</i>
  </li>
</ol>
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return
	}

	data := struct {
		Message  string
		Message2 string
		Message3 string
	}{
		Message:  "No message specified",
		Message2: "Best string bean: Fortex",
		Message3: "My favorite: Kentucky Wonder",
	}

	return t.Execute(w, &data)
}
