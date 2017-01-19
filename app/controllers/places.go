package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"

	"io"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	rest "github.com/mochi8k/aiteru-server/app/http"
	"github.com/mochi8k/aiteru-server/app/models"
	. "github.com/mochi8k/aiteru-server/config"
)

func init() {
	rest.Register("/v1/places", map[string]rest.Handler{
		"POST": createPlace,
		"GET":  getPlaces,
	})

	rest.Register("/v1/places/:place-id", map[string]rest.Handler{
		"GET":    getPlace,
		"DELETE": deletePlace,
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

type createPlaceParam struct {
	Name            string   `json:"name"`
	OwnerIDs        []string `json:"owners"`
	CollaboratorIDs []string `json:"collaborators"`
}

func createPlace(_ httprouter.Params, _ url.Values, reader io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
	var createPlaceParam createPlaceParam
	body, _ := ioutil.ReadAll(reader)

	if err := json.Unmarshal(body, &createPlaceParam); err != nil {
		return rest.Fail(http.StatusBadRequest, err.Error()), err
	}

	if createPlaceParam.Name == "" {
		return rest.FailByCode(http.StatusBadRequest), nil
	}

	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	user := session.GetUser()
	createUserID := user.GetID()
	createPlaceParam.OwnerIDs = append(createPlaceParam.OwnerIDs, createUserID)
	createPlaceParam.CollaboratorIDs = append(createPlaceParam.CollaboratorIDs, createUserID)

	// TODO: transaction
	sq.
		Insert("places").
		Columns("place_name, created_at, created_by").
		Values(createPlaceParam.Name, time.Now(), createUserID).
		RunWith(db).
		QueryRow()

	var createdPlaceID string

	sq.
		Select("id").
		From("places").
		Where(sq.Eq{"place_name": createPlaceParam.Name}).
		RunWith(db).QueryRow().Scan(&createdPlaceID)

	for _, ownerID := range createPlaceParam.OwnerIDs {
		sq.
			Insert("place_owners").
			Columns("place_id, owner_id").
			Values(createdPlaceID, ownerID).
			RunWith(db).
			QueryRow()
	}

	for _, collaboratorID := range createPlaceParam.CollaboratorIDs {
		sq.
			Insert("place_collaborators").
			Columns("place_id, collaborator_id").
			Values(createdPlaceID, collaboratorID).
			RunWith(db).
			QueryRow()
	}

	createdPlace := selectPlace(db, createdPlaceID)

	fmt.Printf("CreatedPlace: %+v\n", createdPlace)

	return rest.Success(http.StatusCreated), map[string]*models.Place{
		"place": createdPlace,
	}
}

func getPlaces(_ httprouter.Params, query url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	res, err := getDefaultSelectBuilder().RunWith(db).Query()

	errorChecker(err)

	var places []*models.Place

	for res.Next() {
		place := toPlace(res)

		// TODO: should be accessed once
		placeStatus := toPlaceStatus(
			sq.
				Select("*").
				From("place_status").
				Where(sq.Eq{"place_id": place.GetID()}).
				OrderBy("updated_at DESC").
				Limit(1).
				RunWith(db).
				QueryRow(),
		)
		place.SetStatus(placeStatus)
		fmt.Printf("Place: %+v\n", place)

		places = append(places, place)
	}

	defer res.Close()

	if param := query.Get("isOpen"); param == "true" || param == "false" {
		isOpen, _ := strconv.ParseBool(param)
		places = placesFilter(places, isOpen)
	}

	return rest.Success(http.StatusOK), map[string][]*models.Place{
		"places": places,
	}
}

func getPlace(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("place-id")
	fmt.Printf("place-id: %s\n", id)

	place := selectPlace(db, id)

	if place.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil

	}

	return rest.Success(http.StatusOK), map[string]*models.Place{
		"place": place,
	}
}

func deletePlace(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", Config.MySQL.Connection)
	errorChecker(err)

	defer db.Close()

	id := ps.ByName("place-id")
	fmt.Printf("place-id: %s\n", id)

	place := selectPlace(db, id)

	if place.ID == "" {
		return rest.FailByCode(http.StatusNotFound), nil

	}

	// TODO: transaction
	_, err = sq.Delete("").From("places").Where(sq.Eq{"id": id}).RunWith(db).Exec()
	errorChecker(err)

	_, err = sq.Delete("").From("place_owners").Where(sq.Eq{"place_id": id}).RunWith(db).Exec()
	errorChecker(err)

	_, err = sq.Delete("").From("place_collaborators").Where(sq.Eq{"place_id": id}).RunWith(db).Exec()
	errorChecker(err)

	_, err = sq.Delete("").From("place_status").Where(sq.Eq{"place_id": id}).RunWith(db).Exec()
	errorChecker(err)

	return rest.Success(http.StatusNoContent), nil
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

	db, err := sql.Open("mysql", Config.MySQL.Connection)
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

	return rest.Success(http.StatusCreated), map[string]*models.PlaceStatus{
		"status": placeStatus,
	}
}

func getStatus(ps httprouter.Params, _ url.Values, _ io.Reader, _ *models.Session) (rest.APIStatus, interface{}) {
	db, err := sql.Open("mysql", Config.MySQL.Connection)
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

	return rest.Success(http.StatusOK), map[string]*models.PlaceStatus{
		"status": placeStatus,
	}

}

func selectPlace(db *sql.DB, placeID string) *models.Place {
	place := toPlace(
		getDefaultSelectBuilder().
			Where(sq.Eq{"p.id": placeID}).
			RunWith(db).
			QueryRow(),
	)

	// TODO: should be accessed once
	placeStatus := toPlaceStatus(
		sq.
			Select("*").
			From("place_status").
			Where(sq.Eq{"place_id": place.GetID()}).
			OrderBy("updated_at DESC").
			Limit(1).
			RunWith(db).
			QueryRow(),
	)
	place.SetStatus(placeStatus)

	fmt.Printf("Place: %+v\n", place)
	return place
}

func placesFilter(places []*models.Place, isOpen bool) []*models.Place {
	var filteredPlaces []*models.Place

	for _, place := range places {
		if place.IsOpen() == isOpen {
			filteredPlaces = append(filteredPlaces, place)
		}
	}
	return filteredPlaces
}
