package util

import (
	"net/http"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

func PopulateScheStr(obj interface{}, r *http.Request) (err error) {
	err = r.ParseForm()
	if err != nil {
		return
	}

	err = decoder.Decode(obj, r.Form)
	return
}
