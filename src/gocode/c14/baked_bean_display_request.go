package c14

import (
	"fmt"
	"html/template"
	"net/http"

	"../util"
	"./bean"
)

type BakedBeanDisplayRequest struct{}

func (bbdr *BakedBeanDisplayRequest) ServeHTTP(w http.ResponseWriter,
	r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = bbdr.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bbdr *BakedBeanDisplayRequest) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	bakedBean := bean.NewBakedBean()
	err = util.PopulateScheStr(bakedBean, r)
	if err != nil {
		return
	}

	t, err := template.ParseFiles("c14/tpl/request.tpl", "c14/tpl/snippet.tpl")
	if err != nil {
		return
	}
	return t.ExecuteTemplate(w, "request", bakedBean)
}
