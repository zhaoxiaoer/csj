package c15

import (
	"fmt"
	"html/template"
	"net/http"
)

type ShowBalance struct{}

func (sb *ShowBalance) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = sb.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (sb *ShowBalance) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	var tplName string
	bankCustomer, err := GetCustomer(r.FormValue("id"))
	if err != nil {
		tplName = "c15/unknown_customer.tpl"
	} else if bankCustomer.Balance < 0 {
		tplName = "c15/negative_balance.tpl"
	} else if bankCustomer.Balance < 10000 {
		tplName = "c15/normal_balance.tpl"
	} else {
		tplName = "c15/high_balance.tpl"
	}

	t, err := template.ParseFiles(tplName)
	if err != nil {
		return
	}
	return t.Execute(w, bankCustomer)
}
