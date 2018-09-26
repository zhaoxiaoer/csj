package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	url := "root:123456@tcp(127.0.0.1:3306)/ztest"
	db, err := sql.Open("mysql", url)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("select firstname, lastname from employee")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer rows.Close()

	var firstname, lastname string
	for rows.Next() {
		err := rows.Scan(&firstname, &lastname)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		fmt.Printf("%v %v\n", firstname, lastname)
	}
}
