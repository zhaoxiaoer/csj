package c8

import (
	"fmt"
	"net/http"

	"../util"
)

type RegistrationForm struct{}

func (rf *RegistrationForm) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = rf.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rf *RegistrationForm) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	actionURL := "/csj/c8/registrationservlet"
	firstName := util.GetCookieValue(r, "firstName", "")
	lastName := util.GetCookieValue(r, "lastName", "")
	emailAddress := util.GetCookieValue(r, "emailAddress", "")
	title := "Please Register"
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body bgcolor=\"#FDF5E6\">\n" +
		"<center>\n" +
		"<h1>" + title + "</h1>" +
		"<form action=\"" + actionURL + "\">\n" +
		"First Name:\n" +
		"  <input type=\"text\" name=\"firstName\" " +
		"value=\"" + firstName + "\"><br>\n" +
		"Last Name:\n" +
		"  <input type=\"text\" name=\"lastName\" " +
		"value=\"" + lastName + "\"><br>\n" +
		"Email Address: \n" +
		"  <input type=\"text\" name=\"emailAddress\" " +
		"value=\"" + emailAddress + "\"><br>\n" +
		"<input type=\"submit\" value=\"Register\">\n" +
		"</form></center></body></html>"))

	return
}
