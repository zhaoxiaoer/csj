package c9

import (
	"fmt"
	"net/http"
	"strconv"

	"../util"
	"./shoppingcart"
)

type CatalogPage struct {
	items   []*shoppingcart.CatalogItem
	itemIDs []string
	title   string
}

func (cp *CatalogPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = cp.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (cp *CatalogPage) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	if cp.items == nil {
		http.Error(w, "Missing Items.", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(util.HeadWithTitle(cp.title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + cp.title + "</h1>\n"))
	for i := 0; i < len(cp.items); i++ {
		w.Write([]byte("<hr>\n"))
		if cp.items[i] == nil {
			w.Write([]byte("<font color=\"red\">" +
				"Unknown item ID " + cp.itemIDs[i] + "</font>\n"))
		} else {
			formURL := "/csj/c9/orderpage"
			//			formURL = url.PathEscape(formURL)
			w.Write([]byte("\n" +
				"<form action=\"" + formURL + "\">\n" +
				"<input type=\"hidden\" name=\"itemID\" " +
				"  value=\"" + cp.items[i].ItemID + "\">\n" +
				"<h2>" + cp.items[i].ShortDescription +
				" ($" + strconv.FormatFloat(cp.items[i].Cost, 'f', -1, 64) + ")</h2>\n" +
				cp.items[i].LongDescription + "\n" +
				"<p>\n<center>\n" +
				"<input type=\"submit\" " +
				"value=\"Add to Shopping Cart\">\n" +
				"</center>\n</p>\n</form>\n"))
		}
	}
	w.Write([]byte("<hr>\n</body></html>\n"))

	return
}

func (cp *CatalogPage) SetItem(itemIDs []string) {
	cp.itemIDs = itemIDs
	cp.items = make([]*shoppingcart.CatalogItem, len(cp.itemIDs))
	for i := 0; i < len(cp.items); i++ {
		cp.items[i] = shoppingcart.GetItem(cp.itemIDs[i])
	}
}

func (cp *CatalogPage) SetTitle(title string) {
	cp.title = title
}
