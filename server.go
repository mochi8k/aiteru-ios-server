package main

import (
	"fmt"
	"net/http"
)

func HTTPHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("/attendance")

	fmt.Fprintf(w, "I've been working late every day this week.")
}

func main() {

	http.HandleFunc("/attendance", HTTPHandler)

	fmt.Println("activated the web server on port 8000")
	http.ListenAndServe(":8000", nil)
}
