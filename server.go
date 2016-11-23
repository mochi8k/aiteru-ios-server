package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/")

	tpl, _ := template.ParseFiles("templates/attendance.tpl")
	tpl.Execute(w, "")
}

func AttendanceHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/attendance")

	if r.Method != "POST" {
		return
	}

	r.ParseForm()

	userName := strings.Join(r.Form["userName"], "")
	currentTime := strings.Join(r.Form["currentTime"], "")

	log.Printf("ユーザー名: %s", userName)
	log.Printf("出勤時間: %s", currentTime)

	db, err := sql.Open("mysql", "root@/gosample")
	errorChecker(err)

	stmt, err := db.Prepare("insert attendance set username=?, time=?")
	errorChecker(err)

	res, err := stmt.Exec(userName, currentTime)
	errorChecker(err)

	log.Println(res)
}

func main() {

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/attendance", AttendanceHandler)

	log.Println("activated the web server on port 8000")
	http.ListenAndServe(":8000", nil)

}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
