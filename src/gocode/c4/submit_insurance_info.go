package c4

import (
	"net/http"
	"strconv"

	"../util"
	"./schestr"
)

type SubmitInsuranceInfo struct {
}

func (sii *SubmitInsuranceInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = sii.handleGet(w, r)
	case "POST":
		err = sii.handlePost(w, r)
	case "PUT":
		err = sii.handlePut(w, r)
	case "DELETE":
		err = sii.handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (sii *SubmitInsuranceInfo) handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	//	w.Write([]byte("get"))
	var ii schestr.InsuranceInfo
	util.PopulateScheStr(&ii, r)
	w.Header().Add("Content-Type", "text/html")
	title := "Insurance Info for " + ii.Name
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<center>\n" +
		"<h1>" + title + "</h1>\n" +
		"<ul>\n" +
		"  <li>Employee ID: " + ii.EmployeeID + "\n" +
		"  <li>Number of children: " + strconv.Itoa(ii.NumChildren) +
		"  <li>Married?: " + strconv.FormatBool(ii.IsMarried) +
		"</ul></center></body></html>"))

	return
}

func (sii *SubmitInsuranceInfo) handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//	w.Write([]byte("post"))
	err = sii.handleGet(w, r)

	return
}

func (sii *SubmitInsuranceInfo) handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func (sii *SubmitInsuranceInfo) handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}
