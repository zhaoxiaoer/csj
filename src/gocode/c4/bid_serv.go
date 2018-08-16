package c4

import (
	"net/http"
	"strconv"

	"../util"
	"./schestr"
)

type BidServ struct {
}

func (bs *BidServ) ServeHTTP(w http.ResponseWriter, h *http.Request) {
	var err error
	switch h.Method {
	case "GET":
		err = bs.handleGet(w, h)
	case "POST":
		err = bs.handlePost(w, h)
	case "PUT":
		err = bs.handlePut(w, h)
	case "DELETE":
		err = bs.handleDelete(w, h)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (bs *BidServ) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	//	w.Write([]byte("get"))
	bi := &schestr.BidInfo{}
	util.PopulateScheStr(bi, r)
	if bi.IsComplete() {
		bs.showBid(w, r, bi)
	} else {
		bs.showEntryForm(w, r, bi)
	}

	return
}

func (bs *BidServ) handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//	w.Write([]byte("post"))
	err = bs.handleGet(w, r)

	return
}

func (bs *BidServ) handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func (bs *BidServ) handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}

func (bs *BidServ) showBid(w http.ResponseWriter, h *http.Request, bi *schestr.BidInfo) {
	bs.submitBid(bi)
	w.Header().Set("Content-Type", "text/html")
	title := "Bid Submitted"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\"><center>\n" +
		"<h1>" + title + "</h1>\n" +
		"Your bid is now active. If your bid is successful,\n" +
		"you will be notified within 24 hours of the close\n" +
		"of bidding.\n" +
		"<p>\n" +
		"<table border=1>\n" +
		"  <tr><th bgcolor=\"black\"><font color=\"white\">" +
		bi.ItemName + "</font></th></tr>\n" +
		"  <tr><td>Item ID: " +
		bi.ItemID + "</td></tr>\n" +
		"  <tr><td>Name: " +
		bi.BidderName + "</td></tr>\n" +
		"  <tr><td>Email address: " +
		bi.EmailAddress + "</td></tr>\n" +
		"  <tr><td>Bid price: $" +
		strconv.FormatFloat(bi.BidPrice, 'f', -1, 64) + "</td></tr>\n" +
		"  <tr><td>Auto-increment price: " +
		strconv.FormatBool(bi.AutoIncrement) + "</td></tr>\n" +
		"</table></center></body></html>"))

	return
}

func (bs *BidServ) showEntryForm(w http.ResponseWriter, h *http.Request, bi *schestr.BidInfo) {
	isPartlyComplete := bi.IsPartlyComplete()
	w.Header().Set("Content-Type", "text/html")
	title := "Welcome to Auctions-R-Us. Please Enter Bid."
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\"><center>\n" +
		"<h1>" + title + "</h1>\n" +
		bs.warning(isPartlyComplete) +
		"<form>\n" +
		bs.inputElement1("Item ID", "itemID", bi.ItemID, isPartlyComplete) +
		bs.inputElement1("Item Name", "itemName", bi.ItemName, isPartlyComplete) +
		bs.inputElement1("Your Name", "bidderName", bi.BidderName, isPartlyComplete) +
		bs.inputElement1("Your Email Address", "emailAddress", bi.EmailAddress, isPartlyComplete) +
		bs.inputElement2("Amount Bid", "bidPrice", bi.BidPrice, isPartlyComplete) +
		bs.checkbox("Auto-increment bid to match other bidders?", "autoIncrement", bi.AutoIncrement) +
		"<input type=\"submit\" value=\"Submit Bid\">\n" +
		"</form></center></body></html>"))

	return
}

func (bs *BidServ) submitBid(bi *schestr.BidInfo) {

}

func (bs *BidServ) warning(isPartlyComplete bool) string {
	if isPartlyComplete {
		return "<h2>Required Data Missing! Enter and Resubmit.</h2>"
	} else {
		return ""
	}
}

func (bs *BidServ) inputElement1(prompt, name, value string, shouldPrompt bool) string {
	msg := ""
	if shouldPrompt && (value == "") {
		msg = "<b>Required field!</b> "
	}
	return (msg + prompt + ": " +
		"<input type=\"text\" name=\"" + name + "\"" +
		" value=\"" + value + "\"><br>\n")
}

func (bs *BidServ) inputElement2(prompt, name string, value float64, shouldPrompt bool) string {
	num := ""
	if value == 0 {
		num = ""
	} else {
		num = strconv.FormatFloat(value, 'f', -1, 64)
	}
	return bs.inputElement1(prompt, name, num, shouldPrompt)
}

func (bs *BidServ) checkbox(prompt, name string, isChecked bool) string {
	result := prompt + ": " + "<input type=\"checkbox\" name=\"" + name + "\""
	if isChecked {
		result = result + " checked"
	}
	result += "><br>\n"
	return result
}
