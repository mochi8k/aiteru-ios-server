package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/")

	tpl, _ := template.ParseFiles("templates/attendance.tpl")
	tpl.Execute(w, "")
}

func AttendanceHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/attendance")

	if r.Method != "POST" {
		return
	}

	r.ParseForm()

	userName := strings.Join(r.Form["userName"], "")
	currentTime := strings.Join(r.Form["currentTime"], "")
	fmt.Println(userName)
	fmt.Println(currentTime)

	db, err := sql.Open("mysql", "root@/gosample")
	errorChecker(err)

	stmt, err := db.Prepare("insert attendance set username=?, time=?")
	errorChecker(err)

	res, err := stmt.Exec(userName, currentTime)
	errorChecker(err)
	fmt.Println(res)
}

func main() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/attendance", AttendanceHandler)

	fmt.Println("activated the web server on port 8000")
	http.ListenAndServe(":8000", nil)
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
