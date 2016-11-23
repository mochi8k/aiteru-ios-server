package main

import (
	"fmt"
	"html/template"
	"net/http"
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

	fmt.Println(r.Method)
	fmt.Fprintf(w, "I've been working late every day this week.")
}

func main() {

	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/attendance", AttendanceHandler)

	fmt.Println("activated the web server on port 8000")
	http.ListenAndServe(":8000", nil)
}
