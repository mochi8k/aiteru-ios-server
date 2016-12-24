package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var router *httprouter.Router = httprouter.New()

func GetInstance() http.Handler {
	return router
}

func POST(pattern string, handler httprouter.Handle) {
	fmt.Printf("EndPoint: POST %s\n", pattern)
	router.POST(pattern, handler)
}

func GET(pattern string, handler httprouter.Handle) {
	fmt.Printf("EndPoint: GET %s\n", pattern)
	router.GET(pattern, handler)
}

func PUT(pattern string, handler httprouter.Handle) {
	fmt.Printf("EndPoint: PUT %s\n", pattern)
	router.PUT(pattern, handler)
}

func DELETE(pattern string, handler httprouter.Handle) {
	fmt.Printf("EndPoint: DELETE %s\n", pattern)
	router.DELETE(pattern, handler)
}
