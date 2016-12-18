package controllers

import (
	_ "database/sql"
	_ "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	_ "github.com/mochi8k/aiteru-ios-server/app/models"
	"io"
	"net/http"
	"net/url"
	_ "time"
)

type users struct {
	rest.APIResourceBase
}

func init() {
	http.Handle("/users/", rest.APIResourceHandler(users{}))
}

func (u users) Get(url string, queries url.Values, body io.Reader) (rest.APIStatus, interface{}) {
	fmt.Println("Get")
	return rest.Success(http.StatusCreated), nil
}

func (u users) Post(url string, queries url.Values, body io.Reader) (rest.APIStatus, interface{}) {
	fmt.Println("Post")
	// sq.
	// 	Insert("users").
	// 	Columns("user_name", "created_at", "created_by").
	// 	Values()

	return rest.Success(http.StatusCreated), nil
}
