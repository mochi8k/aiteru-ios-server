package controllers

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
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

func toUser(scanner sq.RowScanner) *models.User {
	var id, name, createdAt, createdUserID, updatedAt, updatedUserID string
	scanner.Scan(&id, &name, &createdAt, &createdUserID, &updatedAt, &updatedUserID)
	return &models.User{
		ID:            id,
		Name:          name,
		CreatedAt:     createdAt,
		CreatedUserID: createdUserID,
		UpdatedAt:     updatedAt,
		UpdatedUserID: updatedUserID,
	}
}

func getUsers(ps httprouter.Params, queries url.Values, body io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	res, err := sq.Select("*").From("users").RunWith(db).Query()
	errorChecker(err)

	var users []*models.User

	for res.Next() {
		user := toUser(res)
		users = append(users, user)
		fmt.Printf("User: %+v\n", user)
	}

	return rest.Success(http.StatusOK), users
}
