package controllers

import (
	"database/sql"
	_ "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"github.com/julienschmidt/httprouter"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	"github.com/mochi8k/aiteru-ios-server/app/models"
	"io"
	"net/http"
	"net/url"
	_ "time"
)

func init() {
	rest.Register("/v1/users/", map[string]rest.Handler{
		"GET": getUsers,
	})
}

func getUsers(ps httprouter.Params, queries url.Values, body io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	return rest.Success(http.StatusOK), nil
}

// func (u users) Post(url string, queries url.Values, body io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
// 	fmt.Println("Post")
// 	// sq.
// 	// 	Insert("users").
// 	// 	Columns("user_name", "created_at", "created_by").
// 	// 	Values()

// 	return rest.Success(http.StatusCreated), nil
// }
