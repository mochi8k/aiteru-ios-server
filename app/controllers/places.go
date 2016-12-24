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
	"strings"
)

func init() {
	rest.Register("/v1/places", map[string]rest.Handler{
		"GET": getPlaces,
	})

	rest.Register("/v1/places/:place-id", map[string]rest.Handler{
		"GET": getPlace,
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

func getPlaces(ps httprouter.Params, queries url.Values, body io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
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

func getPlace(ps httprouter.Params, queries url.Values, body io.Reader, session *models.Session) (rest.APIStatus, interface{}) {
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

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
