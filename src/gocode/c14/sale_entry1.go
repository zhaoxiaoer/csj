package c14

import (
	"fmt"
	"html/template"
	"net/http"

	"../util"
	"./bean"
)

type SaleEntry1 struct{}

func (se *SaleEntry1) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = se.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (se *SaleEntry1) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	saleEntry := &bean.SaleEntry{}
	err = util.PopulateScheStr(saleEntry, r)
	if err != nil {
		return
	}
	if saleEntry.NumItems <= 0 {
		saleEntry.NumItems = 1
	}
	if saleEntry.DiscountCode <= 0 {
		saleEntry.DiscountCode = 1.0
	}
	//	fmt.Printf("%v\n", saleEntry)

	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Using jsp:setProperty</title>
</head>
<body>
<center>
  <table border=5>
    <tr>
      <th>Using jsp:setProperty</th>
    </tr>
  </table>
  <table border=1>
    <tr>
      <th>Item ID</th>
      <th>Unit Price</th>
      <th>Number Ordered</th>
      <th>Total Price</th>
    </tr>
    <tr align="right">
      <td>{{.ItemID}}</td>
      <td>{{.ItemCost}}</td>
      <td>{{.NumItems}}</td>
      <td>{{.TotalCost}}</td>
    </tr>
  </table>
</center>
</body>
</html>`
	t, err := template.New("csj").Parse(tpl)
	if err != nil {
		return nil
	}

	data := &struct {
		ItemID    string
		ItemCost  float64
		NumItems  int
		TotalCost float64
	}{
		ItemID:    saleEntry.ItemID,
		ItemCost:  saleEntry.GetItemCost(),
		NumItems:  saleEntry.NumItems,
		TotalCost: saleEntry.GetTotalCost(),
	}
	return t.Execute(w, data)
}
