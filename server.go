package main

import (
	"fmt"
	"net/http"

	_ "github.com/mochi8k/aiteru-server/app/controllers"
	"github.com/mochi8k/aiteru-server/app/handlers/router"
	. "github.com/mochi8k/aiteru-server/config"
)

func main() {
	fmt.Println("activated the web server on port " + Config.Port)
	http.ListenAndServe(":"+Config.Port, router.GetInstance())
}
