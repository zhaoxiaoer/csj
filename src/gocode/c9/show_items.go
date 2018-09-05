package c9

import (
	"fmt"
	"net/http"
	"strings"

	"../util"
)

type ShowItems struct{}

func (si *ShowItems) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = si.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (si *ShowItems) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	session := util.GetSession(r)
	previousItems, ok := session.Values["previousItems"].([]string)
	if !ok {
		previousItems = make([]string, 0)
	}
	newItem := r.FormValue("newItem")

	if strings.Compare(newItem, "") != 0 {
		previousItems = append(previousItems, newItem)
	}
	session.Values["previousItems"] = previousItems
	session.Save(r, w)
	// go语言的同步可以使用sync.Mutex

	w.Header().Set("Content-Type", "text/html")
	title := "Items Purchased"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1>" + title + "</h1>\n"))
	if len(previousItems) == 0 {
		w.Write([]byte("<i>No items</i>\n"))
	} else {
		w.Write([]byte("<ul>\n"))
		for i := 0; i < len(previousItems); i++ {
			w.Write([]byte("<li>" + previousItems[i] + "</li>\n"))
		}
		w.Write([]byte("</ul>\n"))
	}
	w.Write([]byte("</body></html>\n"))

	return
}
