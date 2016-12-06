package main

import (
	"net/http"

	_ "github.com/mochi8k/aiteru-ios-server/app/controllers"
)

func main() {
	http.ListenAndServe(":8000", nil)
}
