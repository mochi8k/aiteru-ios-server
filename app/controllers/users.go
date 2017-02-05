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
	rest "github.com/mochi8k/aiteru-server/app/http"
	"github.com/mochi8k/aiteru-server/app/models"
	. "github.com/mochi8k/aiteru-server/config"
)

func init() {
	rest.Register("/v1/users", map[string]rest.Handler{
		"POST": createUser,
		"GET":  getUsers,
	})

	rest.Register("/v1/users/:user-id", map[string]rest.Handler{
		"GET":    getUser,
		"PUT":    updateUser,
		"DELETE": deleteUser,
	})
}

type bodyParam struct {
	UserName string `json:"name"`
}

func createUser(_ httprouter.Params, _ url.Values, reader io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	var bodyParam bodyParam
	body, _ := ioutil.ReadAll(reader)

	if err := json.Unmarshal(body, &bodyParam); err != nil {
		return rest.Fail(http.StatusBadRequest, err.Error()), err
	}

	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	user := session.GetUser()
	createUserID := user.GetID()

	sq.
		Insert("users").
		Columns("user_name, created_at, created_by").
		Values(bodyParam.UserName, time.Now(), createUserID).
		RunWith(db).
		QueryRow()

	createdUser := toUser(
		sq.
			Select("*").
			From("users").
			Where(sq.Eq{"users.user_name": bodyParam.UserName}).
			RunWith(db).QueryRow(),
	)

	fmt.Printf("User: %+v\n", createdUser)

	return rest.Success(http.StatusCreated), map[string]*models.User{
		"user": createdUser,
	}
}

func updateUser(ps httprouter.Params, _ url.Values, reader io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	var bodyParam bodyParam
	body, _ := ioutil.ReadAll(reader)

	if err := json.Unmarshal(body, &bodyParam); err != nil {
		return rest.Fail(http.StatusBadRequest, err.Error()), err
	}

	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("user-id")

	sq.
		Update("users").
		SetMap(sq.Eq{
			"user_name":  bodyParam.UserName,
			"updated_at": time.Now(),
			"updated_by": session.GetUser().GetID(),
		}).
		Where(sq.Eq{"users.id": id}).
		RunWith(db).
		QueryRow()

	updatedUser := toUser(
		sq.
			Select("*").
			From("users").
			Where(sq.Eq{"users.id": id}).
			RunWith(db).
			QueryRow(),
	)

	if updatedUser.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil
	}

	fmt.Printf("User: %+v\n", updatedUser)

	return rest.Success(http.StatusOK), map[string]*models.User{
		"user": updatedUser,
	}
}

func getUsers(_ httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", Config.MySQL.Connection)
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
	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("user-id")
	fmt.Printf("user-id: %s\n", id)

	user := selectUser(db, id)

	if user.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil
	}

	return rest.Success(http.StatusOK), map[string]*models.User{
		"user": user,
	}
}

func deleteUser(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("user-id")
	fmt.Printf("user-id: %s\n", id)

	if user := selectUser(db, id); user.GetID() == "" {
		return rest.FailByCode(http.StatusNotFound), nil
	}

	// TODO: transaction
	_, err = sq.Delete("").From("users").Where(sq.Eq{"id": id}).Exec()
	errorChecker(err)

	_, err = sq.Delete("").From("place_owners").Where(sq.Eq{"owner_id": id}).Exec()
	errorChecker(err)

	_, err = sq.Delete("").From("place_collaborators").Where(sq.Eq{"collaborator_id": id}).Exec()
	errorChecker(err)

	return rest.Success(http.StatusNoContent), nil
}

func selectUser(db *sql.DB, userID string) *models.User {
	user := toUser(
		sq.
			Select("*").
			From("users").
			Where(sq.Eq{"users.id": userID}).
			RunWith(db).
			QueryRow(),
	)

	fmt.Printf("User: %+v\n", user)
	return user
}
