package c10

import (
	"fmt"
	"html/template"
	"net/http"
)

type OrderConfirmation struct{}

func (oc *OrderConfirmation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = oc.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (oc *OrderConfirmation) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Order Confirmation</title>
</head>
<body>
<h2>Order Confirmation</h2>
Thanks for ordering <i>{{.}}</i>!
</body>
</html>`
	t, err := template.New("order").Parse(tpl)
	if err != nil {
		return fmt.Errorf("template error")
	}

	return t.Execute(w, r.FormValue("title"))
}
