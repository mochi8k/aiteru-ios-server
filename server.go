package main

import (
	"fmt"
	"net/http"

	_ "github.com/mochi8k/aiteru-ios-server/app/controllers"
)

func main() {
	fmt.Println("activated the web server on port 8000")
	http.ListenAndServe(":8000", nil)
}
