package c8

import (
	"fmt"
	"net/http"
	"strings"

	"../util"
)

type RegistrationServlet struct{}

func (rs *RegistrationServlet) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = rs.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (rs *RegistrationServlet) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("Content-Type", "text/html")
	isMissingValue := false
	firstName := r.FormValue("firstName")
	if rs.isMissing(firstName) {
		firstName = "Missing first name"
		isMissingValue = true
	}
	lastName := r.FormValue("lastName")
	if rs.isMissing(lastName) {
		lastName = "Missing last name"
		isMissingValue = true
	}
	emailAddress := r.FormValue("emailAddress")
	if rs.isMissing(emailAddress) {
		emailAddress = "Missing email address"
		isMissingValue = true
	}
	cookie := http.Cookie{
		Name:   "firstName",
		Value:  firstName,
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, &cookie)
	cookie2 := http.Cookie{
		Name:   "lastName",
		Value:  lastName,
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, &cookie2)
	cookie3 := http.Cookie{
		Name:   "emailAddress",
		Value:  emailAddress,
		MaxAge: 60 * 60 * 24 * 365,
	}
	http.SetCookie(w, &cookie3)
	formAddress := "/csj/c8/registrationform"
	if isMissingValue {
		http.Redirect(w, r, formAddress, http.StatusFound)
	} else {
		title := "Thanks for Registering"
		w.Write([]byte(util.HeadWithTitle(title) +
			"<body bgcolor=\"#FDF5E6\">\n" +
			"<center>\n" +
			"<h1>" + title + "</h1>\n" +
			"<ul>\n" +
			"  <li><b>First Name</b>: " + firstName + "</li>\n" +
			"  <li><b>Last Name</b>: " + lastName + "</li>\n" +
			"  <li><b>Email address</b>: " + emailAddress + "</li>\n" +
			"</ul>\n" +
			"</center></body></html>"))
	}

	return
}

func (rs *RegistrationServlet) isMissing(param string) bool {
	return len(strings.Trim(param, " ")) == 0
}
