package c17

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"../util"

	_ "github.com/go-sql-driver/mysql"
)

type Northwind struct{}

func (nw *Northwind) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = nw.handleGet(w, r)
	default:
		err = fmt.Errorf("%v", r.Method+" may not be supported")
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (nw *Northwind) handleGet(w http.ResponseWriter,
	r *http.Request) (err error) {
	w.Header().Set("ContentType", "text/html")
	w.Write([]byte(util.HeadWithTitle("Northwind Results") +
		"<body bgcolor=\"#FDF5E6\"><center>\n" +
		"<h1>Northwind Results</h1>\n"))
	w.Write([]byte("<ul>\n" +
		"  <li><b>Database:</b> MySQL</li>\n" +
		"  <li><b>Version:</b> 5.7.13-log</li>\n" +
		"</ul>\n"))

	url := "root:123456@tcp(127.0.0.1:3306)/ztest"
	db, err := sql.Open("mysql", url)
	if err != nil {
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select * from employee")
	if err != nil {
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return
	}
	defer rows.Close()

	cs, err := rows.Columns()
	if err != nil {
		return
	}

	w.Write([]byte("<table border=1>\n"))
	w.Write([]byte("<tr>\n"))
	for i := 0; i < len(cs); i++ {
		w.Write([]byte("<th>" + cs[i] + "</th>\n"))
	}
	w.Write([]byte("</tr>\n"))

	var employee_id int
	var firstname, lastname string
	for rows.Next() {
		w.Write([]byte("<tr>\n"))

		err = rows.Scan(&employee_id, &firstname, &lastname)
		if err != nil {
			return
		}
		w.Write([]byte("  <td>" + strconv.FormatInt(int64(employee_id), 10) + "</td>\n"))
		w.Write([]byte("  <td>" + firstname + "</td>\n"))
		w.Write([]byte("  <td>" + lastname + "</td>\n"))

		w.Write([]byte("</tr>\n"))
	}
	w.Write([]byte("</table>\n"))
	w.Write([]byte("</center></body></html>\n"))

	return
}
