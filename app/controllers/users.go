package controllers

import (
	"database/sql"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/julienschmidt/httprouter"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	"github.com/mochi8k/aiteru-ios-server/app/models"
)

func init() {
	rest.Register("/v1/users", map[string]rest.Handler{
		"POST": createUser,
		"GET":  getUsers,
	})

	rest.Register("/v1/users/:user-id", map[string]rest.Handler{
		"GET": getUser,
	})
}

type createParam struct {
	UserName string `json:"name"`
}

func createUser(_ httprouter.Params, _ url.Values, reader io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	var createParam createParam
	body, _ := ioutil.ReadAll(reader)

	if err := json.Unmarshal(body, &createParam); err != nil {
		return rest.Fail(http.StatusBadRequest, err.Error()), err
	}

	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	user := session.GetUser()
	createUserID := user.GetID()

	sq.
		Insert("users").
		Columns("user_name, created_at, created_by").
		Values(createParam.UserName, time.Now(), createUserID).
		RunWith(db).
		QueryRow()

	createdUser := toUser(
		sq.
			Select("*").
			From("users").
			Where(sq.Eq{"users.user_name": createParam.UserName}).
			RunWith(db).QueryRow(),
	)

	fmt.Printf("User: %+v\n", createdUser)

	return rest.Success(http.StatusCreated), map[string]*models.User{
		"user": createdUser,
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

	return rest.Success(http.StatusOK), map[string][]*models.User{
		"users": users,
	}
}

func getUser(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("user-id")
	fmt.Printf("user-id: %s\n", id)

	user := toUser(
		sq.
			Select("*").
			From("users").
			Where(sq.Eq{"users.id": id}).
			RunWith(db).
			QueryRow(),
	)

	if user.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil
	}

	return rest.Success(http.StatusOK), map[string]*models.User{
		"user": user,
	}
}
