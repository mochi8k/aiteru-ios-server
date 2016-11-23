package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
	"strings"
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

	fmt.Printf("ユーザー名: %s", userName)
	fmt.Printf("出勤時間: %s", currentTime)

	db, err := sql.Open("mysql", "root@/gosample")
	errorChecker(err)

	defer db.Close()

	stmt, err := db.Prepare("insert attendance set username=?, time=?")
	errorChecker(err)

	res, err := stmt.Exec(userName, currentTime)
	errorChecker(err)

	fmt.Println(res)

	w.Header().Set("Location", "/history")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

type row struct {
	id       int
	userName string
	time     string
}

func (r *row) toString() string {
	return fmt.Sprintf("%d %s %s", r.id, r.userName, r.time)
}

func HistoryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/history")
	db, err := sql.Open("mysql", "root@/gosample")
	errorChecker(err)

	defer db.Close()

	res, err := db.Query("select * from attendance")
	errorChecker(err)

	var rows []*row

	for res.Next() {
		var id int
		var userName string
		var time string
		err := res.Scan(&id, &userName, &time)
		errorChecker(err)
		fmt.Printf("id: %d", id)
		fmt.Printf("userName: %s", userName)
		fmt.Printf("time: %s", time)

		rows = append(rows, &row{
			id:       id,
			userName: userName,
			time:     time,
		})

	}

	var messages []string

	for _, row := range rows {
		messages = append(messages, row.toString())
	}

	n, e := w.Write([]byte(strings.Join(messages, "\n")))
	fmt.Println(n, e)
}

func main() {

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/attendance", AttendanceHandler)
	http.HandleFunc("/history", HistoryHandler)

	fmt.Println("activated the web server on port 8000")
	http.ListenAndServe(":8000", nil)

}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
