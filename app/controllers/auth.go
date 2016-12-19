package controllers

import (
	"database/sql"
	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"fmt"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	"github.com/mochi8k/aiteru-ios-server/app/models"
	"github.com/mochi8k/aiteru-ios-server/app/stores"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type loginParam struct {
	UserName string `json:"name"`
	Address  string `json:"address"`
}

type auth struct {
	rest.APIResourceBase
}

func init() {
	http.Handle("/auth", rest.APIResourceHandler(auth{}))
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

func (auth) Post(url string, queries url.Values, reader io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {

	var loginParam loginParam

	body, _ := ioutil.ReadAll(reader)

	if err := json.Unmarshal(body, &loginParam); err != nil {
		return rest.Fail(http.StatusBadRequest, err.Error()), err
	}

	fmt.Printf("Login: %v\n", loginParam)

	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	rowScanner := sq.
		Select("*").
		From("users").
		Where(sq.Eq{"users.user_name": loginParam.UserName}).
		RunWith(db).QueryRow()

	user := toUser(rowScanner)

	if user.ID == "" {
		return rest.FailByCode(http.StatusBadRequest), nil
	}

	session := models.NewSession(*user)
	stores.AddSession(session)

	return rest.Success(http.StatusOK), session
}
