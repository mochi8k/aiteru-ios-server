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
	rest.Register("/v1/users", map[string]rest.Handler{
		"GET": getUsers,
	})

	rest.Register("/v1/users/:user-id", map[string]rest.Handler{
		"GET": getUser,
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

func getUsers(_ httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
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

func getUser(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("user-id")
	fmt.Printf("user-id: %s\n", id)

	rowScanner := sq.Select("*").From("users").Where(sq.Eq{"users.id": id}).RunWith(db).QueryRow()
	user := toUser(rowScanner)

	if user.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil
	}

	return rest.Success(http.StatusOK), user
}
