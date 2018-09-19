package c14

import (
	"fmt"
	"net/http"
)

type SaleEntry1Form struct{}

func (sef *SaleEntry1Form) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = sef.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (sef *SaleEntry1Form) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	tpl := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Invoking SaleEntry1.jsp</title>
</head>
<body>
<center>
  <table boder=5>
    <tr>
      <th>Invoking SaleEntry1.jsp</th>
    </tr>
  </table>
  <form action="/csj/c14/SaleEntry1.jsp">
    Item ID: <input type="text" name="itemID"><br>
    Number of Items: <input type="text" name="numItems"><br>
    Discount Code: <input type="text" name="discountCode"><br>
    <input type="submit" value="Show Price">
  </form>
</center>
</body>
</html>`
	w.Write([]byte(tpl))

	return
}
