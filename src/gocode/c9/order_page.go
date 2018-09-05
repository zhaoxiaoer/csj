package c9

import (
	"fmt"
	"net/http"
	"strconv"

	"../util"
	"./shoppingcart"
)

type OrderPage struct{}

func (op *OrderPage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = op.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (op *OrderPage) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	//	session := util.GetSession(r)
	//	cart, ok := session.Values["shoppingCart"].(shoppingcart.ShoppingCart)
	//	fmt.Printf("%v, %v\n", cart, ok)
	//	if !ok {
	//		cart = *shoppingcart.NewShoppingCart()
	//	}

	session := util.SM.GetSession(w, r)
	//	fmt.Printf("1 %v\n", session)
	cart, ok := session.Get("shoppingCart").(*shoppingcart.ShoppingCart)
	//	fmt.Printf("0 cart: %v\n", cart)
	if !ok {
		cart = shoppingcart.NewShoppingCart()
		session.Set("shoppingCart", cart)
	}

	//	fmt.Printf("2%v\n", session)

	itemID := r.FormValue("itemID")
	//	fmt.Printf("itemID: %v\n", itemID)
	if itemID != "" {
		numItems := r.FormValue("numItems")
		//		fmt.Printf("numItems: %v\n", numItems)
		if numItems == "" {
			cart.AddItem(itemID)
		} else {
			n, err := strconv.ParseUint(numItems, 10, 64)
			if err != nil {
				n = 1
			}
			fmt.Printf("n: %v\n", n)
			cart.SetNumOrdered(itemID, int(n))
		}
	}

	//	fmt.Printf("3%v\n", session)

	//	session.Values["shoppingCart"] = cart
	//	session.Values["shoppingCart"] = "aaa"
	//	session.Save(r, w)

	//	fmt.Printf("1 cart: %v\n", cart)

	w.Header().Set("Content-Type", "text/html")
	title := "Status of Your Order"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<h1 align=\"center\">" + title + "</h1>\n"))
	itemsOrdered := cart.GetItemsOrdered()
	if len(itemsOrdered) == 0 {
		w.Write([]byte("<h2><i>No items in your cart...</i></h2>\n"))
	} else {
		w.Write([]byte("<table border=1 align=\"center\">\n" +
			"<tr bgcolor=\"#FFAD00\">\n" +
			"  <th>Item ID</th><th>Description\n</th>\n" +
			"  <th>Unit Cost</th><th>Number</th><th>Total Cost</th>" +
			"</tr>\n"))
		for i := 0; i < len(itemsOrdered); i++ {
			order := itemsOrdered[i]
			w.Write([]byte("<tr>\n" +
				"  <td>" + order.GetItemID() + "</td>\n" +
				"  <td>" + order.GetShortDescription() + "</td>\n" +
				"  <td>" +
				strconv.FormatFloat(order.GetUnitCost(), 'f', -1, 64) +
				"</td>\n" +
				"  <td>" +
				"<form>\n" +
				"<input type=\"hidden\" name=\"itemID\" " +
				"value=\"" + order.GetItemID() + "\">\n" +
				"<input type=\"text\" name=\"numItems\" size=3 value=\"" +
				strconv.FormatInt(int64(order.GetNumItems()), 10) + "\">\n" +
				"<small>\n" +
				"<input type=\"submit\" value=\"Update Order\">\n" +
				"</small>\n" +
				"</form>\n" +
				"</td>" +
				"<td>" +
				strconv.FormatFloat(order.GetTotalCost(), 'f', -1, 64) + "</td>\n"))
		}
		checkoutURL := "/csj/c9/Checkout.html"
		w.Write([]byte("</table>\n" +
			"<form action=\"" + checkoutURL + "\">\n" +
			"<big><center>\n" +
			"<input type=\"submit\" value=\"Proceed to Checkout\">\n" +
			"</center></big></form>\n"))
	}
	w.Write([]byte("</body></html>\n"))

	return
}
