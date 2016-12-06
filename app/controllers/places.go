package api

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	"io"
	"net/http"
	"net/url"
)

func init() {
	http.Handle("/places/", rest.APIResourceHandler(places{}))
}

type places struct {
	rest.APIResourceBase
}

func (p places) Get(url string, queries url.Values, body io.Reader) (rest.APIStatus, interface{}) {
	fmt.Println("GET: /places")

	fmt.Println(url)
	fmt.Println(queries)
	fmt.Println(body)

	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	res, err := db.Query("select * from places")
	errorChecker(err)

	defer res.Close()

	fmt.Println(res)

	return rest.Success(http.StatusOK), nil
}

func (p places) Post(url string, queries url.Values, body io.Reader) (rest.APIStatus, interface{}) {
	fmt.Println(url)
	fmt.Println(queries)
	fmt.Println(body)
	return rest.Success(http.StatusOK), nil
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
