package c4

import (
	"fmt"
	"html"
	"net/http"
	"strconv"
	"strings"

	"../util"
)

type SubmitResume struct {
}

func (sr *SubmitResume) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleGet(w, r)
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("111"))

	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	r.ParseForm()

	if r.PostForm.Get("previewButton") != "" {
		showPreview(w, r)
	} else {
		storeResume(r)
		showConfirmation(w, r)
	}

	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("put"))

	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	w.Write([]byte("delete"))

	return
}

func showPreview(w http.ResponseWriter, r *http.Request) {
	headingFont := r.PostForm.Get("headingFont")
	headingFont = replaceIfMissingOrDefault(headingFont, "")
	//	headingSize := getSize(r.PostForm.Get("headingSize"), 32)
	bodyFont := r.PostForm.Get("bodyFont")
	bodyFont = replaceIfMissingOrDefault(bodyFont, "")
	//	bodySize := getSize(r.PostForm.Get("bodySize"), 18)
	fgColor := r.PostForm.Get("fgColor")
	fgColor = replaceIfMissing(fgColor, "black")
	bgColor := r.PostForm.Get("bgColor")
	bgColor = replaceIfMissing(bgColor, "white")
	name := r.PostForm.Get("name")
	name = replaceIfMissing(name, "Lou Zer")
	title := r.PostForm.Get("title")
	title = replaceIfMissing(title, "Loser")
	email := r.PostForm.Get("email")
	email = replaceIfMissing(email, "contact@hot-computer-jobs.com")
	languages := r.PostForm.Get("languages")
	languages = replaceIfMissing(languages, "<i>Nont</i>")
	languageList := makeList(languages)
	skills := r.PostForm.Get("skills")
	skills = replaceIfMissing(skills, "Not many, obviously.")
	w.Write([]byte(util.HeadWithTitle("Resume for "+html.EscapeString(name)) +
		"<body>\n" +
		"<center>\n" +
		"<span class=\"heading1\">" + html.EscapeString(name) + "</span><br>\n" +
		"<span class=\"heading2\">" + html.EscapeString(title) + "<br>\n" +
		"<a href=\"mailto:" + html.EscapeString(email) + "\">" + html.EscapeString(email) + "</a></span>\n" +
		"</center><br><br>\n" +
		"<span class=\"header3\">Programming Languages</span>\n" +
		languageList + "<br></br>\n" +
		"<span class=\"heading3\">Skills and Experience" + "</span><br><br>\n" +
		skills + "\n" +
		"</body></html>"))
}

func replaceIfMissing(orig, replacement string) string {
	if strings.TrimSpace(orig) == "" {
		return replacement
	} else {
		return orig
	}
}

func replaceIfMissingOrDefault(orig, replacement string) string {
	if (strings.TrimSpace(orig) == "") || (orig == "default") {
		return replacement
	} else {
		return orig
	}
}

func getSize(sizeStr string, defSize int) int {
	i, err := strconv.ParseUint(sizeStr, 10, 64)
	if (err != nil) || (i >= 100) {
		fmt.Printf("err: %v\n", err)
		return defSize
	} else {
		return int(i)
	}
}

func makeList(items string) string {
	ss := strings.Split(items, ",")
	list := "<ul>\n"
	for i := 0; i < len(ss); i++ {
		list += ("  <li>" + strings.TrimSpace(ss[i]) + "</li>\n")
	}
	list += "</ul>"
	return list
}

func showConfirmation(w http.ResponseWriter, r *http.Request) {
	title := "Submission Confirmed."
	w.Write([]byte(util.HeadWithTitle(title) +
		"<body>\n" +
		"<h1>" + title + "</h1>\n" +
		"Your resume should appear online within\n" +
		"24 hours. If it doesn't, try submitting\n" +
		"again with a different email address.\n" +
		"</body></html>"))
}

func storeResume(r *http.Request) {
	email := r.PostForm.Get("email")
	putInSpamList(email)
}

func putInSpamList(emailAddr string) {
	// Code removed to protect the guilty
}
