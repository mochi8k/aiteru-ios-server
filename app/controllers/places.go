package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	"github.com/mochi8k/aiteru-ios-server/app/models"
)

func init() {
	rest.Register("/v1/places", map[string]rest.Handler{
		"GET": getPlaces,
	})

	rest.Register("/v1/places/:place-id", map[string]rest.Handler{
		"GET": getPlace,
	})

	rest.Register("/v1/places/:place-id/status", map[string]rest.Handler{
		"POST": postStatus,
		"GET":  getStatus,
	})
}

func getDefaultSelectBuilder() sq.SelectBuilder {
	return sq.
		Select(`p.id,
      p.place_name,
      group_concat(distinct u1.id) as owner_ids,
      group_concat(distinct u2.id) as collaborator_ids,
      u3.id as created_by,
      p.created_at,
      u4.id as updated_by,
      p.updated_at
    `).
		From("places as p").
		Join("place_owners as po on p.id = po.place_id").
		Join("users as u1 on u1.id = po.owner_id").
		Join("place_collaborators as pc on p.id = pc.place_id").
		Join("users as u2 on u2.id = pc.collaborator_id").
		Join("users as u3 on p.created_by=u3.id").
		LeftJoin("users as u4 on p.updated_by=u4.id").
		GroupBy("p.id, u3.user_name, u4.user_name")
}

func toPlace(scanner sq.RowScanner) *models.Place {
	var id, placeName, ownerIDs, collaboratorIDs, createdAt, createdUserID, updatedAt, updatedUserID string
	scanner.Scan(&id, &placeName, &ownerIDs, &collaboratorIDs, &createdUserID, &createdAt, &updatedUserID, &updatedAt)
	return &models.Place{
		ID:              id,
		Name:            placeName,
		OwnerIDs:        strings.Split(ownerIDs, ","),
		CollaboratorIDs: strings.Split(collaboratorIDs, ","),
		CreatedAt:       createdAt,
		CreatedUserID:   createdUserID,
		UpdatedAt:       updatedAt,
		UpdatedUserID:   updatedUserID,
	}
}

// TODO: define interface
func toPlaceStatus(scanner sq.RowScanner) *models.PlaceStatus {
	var placeID, updatedAt, updatedUserID string
	var isOpen bool
	scanner.Scan(&placeID, &isOpen, &updatedAt, &updatedUserID)
	return &models.PlaceStatus{
		PlaceID:       placeID,
		IsOpen:        isOpen,
		UpdatedAt:     updatedAt,
		UpdatedUserID: updatedUserID,
	}
}

func getPlaces(_ httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	res, err := getDefaultSelectBuilder().RunWith(db).Query()

	errorChecker(err)

	var places []*models.Place

	for res.Next() {
		place := toPlace(res)
		places = append(places, place)
		fmt.Printf("Place: %+v\n", place)
	}

	defer res.Close()

	return rest.Success(http.StatusOK), places
}

func getPlace(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("place-id")
	fmt.Printf("place-id: %s\n", id)

	rowScanner := getDefaultSelectBuilder().Where(sq.Eq{"p.id": id}).RunWith(db).QueryRow()
	place := toPlace(rowScanner)

	fmt.Printf("Place: %+v\n", place)

	if place.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil

	}

	return rest.Success(http.StatusOK), place
}

type statusParam struct {
	IsOpen bool `json:"isOpen"`
}

func postStatus(ps httprouter.Params, _ url.Values, reader io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	var statusParam statusParam
	body, _ := ioutil.ReadAll(reader)

	if err := json.Unmarshal(body, &statusParam); err != nil {
		return rest.Fail(http.StatusBadRequest, err.Error()), err
	}

	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	placeID := ps.ByName("place-id")
	fmt.Printf("place-id: %s\n", placeID)

	user := session.GetUser()
	updatedUserID := user.GetID()

	// TODO: 権限チェック
	sq.
		Insert("place_status").
		Columns("place_id, is_open, updated_at, updated_by").
		Values(placeID, statusParam.IsOpen, time.Now(), updatedUserID).
		RunWith(db).
		QueryRow()

	placeStatus := toPlaceStatus(
		sq.
			Select("*").
			From("place_status as ps").
			Where(sq.Eq{"ps.place_id": placeID}).
			OrderBy("ps.updated_at DESC").
			Limit(1).
			RunWith(db).
			QueryRow(),
	)

	fmt.Printf("PlaceStatus: %+v\n", placeStatus)

	return rest.Success(http.StatusCreated), placeStatus
}

func getStatus(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	placeID := ps.ByName("place-id")
	fmt.Printf("place-id: %s\n", placeID)

	placeStatus := toPlaceStatus(
		sq.
			Select("*").
			From("place_status").
			Where(sq.Eq{"place_id": placeID}).
			OrderBy("updated_at DESC").
			Limit(1).
			RunWith(db).
			QueryRow(),
	)

	fmt.Printf("PlaceStatus: %+v\n", placeStatus)

	if placeStatus.PlaceID == "" {
		return rest.FailByCode(http.StatusNotFound), nil
	}

	return rest.Success(http.StatusOK), placeStatus
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
