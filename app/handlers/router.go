package handlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router *mux.Router = mux.NewRouter()

func GetRouter() http.Handler {
	return router
}

func HandleFunc(pattern string, handler http.HandlerFunc) {
	fmt.Printf("EndPoint: %s", pattern)
	router.HandleFunc(pattern, handler)
}
