package api

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"fmt"
	rest "github.com/mochi8k/aiteru-ios-server/app/http"
	"github.com/mochi8k/aiteru-ios-server/app/models"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func init() {
	http.Handle("/places/", rest.APIResourceHandler(places{}))
}

type places struct {
	rest.APIResourceBase
}

func (p places) Get(url string, queries url.Values, body io.Reader) (rest.APIStatus, interface{}) {
	fmt.Println("GET: /places")

	fmt.Println(url)
	fmt.Println(queries)
	fmt.Println(body)

	db, err := sql.Open("mysql", "root@/aiteru")
	errorChecker(err)

	defer db.Close()

	res, err := db.Query(`
		select
		    p.id,
		    p.place_name,
		    group_concat(distinct u1.id) as owner_ids,
		    group_concat(distinct u2.id) as collaborator_ids,
		    u3.id as created_by,
		    p.created_at,
		    u4.id as updated_by,
		    p.updated_at
		from
		    places as p
		inner join
		    place_owners as po on p.id = po.place_id
		inner join
		    users as u1 on u1.id = po.owner_id
		inner join
		    place_collaborators as pc on p.id = pc.place_id
		inner join
		    users as u2 on u2.id = pc.collaborator_id
		inner join
		    users as u3 on p.created_by=u3.id
		left outer join
		    users as u4 on p.updated_by=u4.id
		group by
		    p.id, u3.user_name, u4.user_name
	`)

	errorChecker(err)

	var places []*models.Place

	for res.Next() {
		var id, placeName, ownerIDs, collaboratorIDs, createdAt, createdUserID, updatedAt, updatedUserID string
		res.Scan(&id, &placeName, &ownerIDs, &collaboratorIDs, &createdUserID, &createdAt, &updatedUserID, &updatedAt)
		place := &models.Place{
			ID:              id,
			Name:            placeName,
			OwnerIDs:        strings.Split(ownerIDs, ","),
			CollaboratorIDs: strings.Split(collaboratorIDs, ","),
			CreatedAt:       createdAt,
			CreatedUserID:   createdUserID,
			UpdatedAt:       updatedAt,
			UpdatedUserID:   updatedUserID,
		}
		places = append(places, place)

		fmt.Printf("Place: %+v\n", place)
	}

	defer res.Close()

	return rest.Success(http.StatusOK), places
}

func (p places) Post(url string, queries url.Values, body io.Reader) (rest.APIStatus, interface{}) {
	fmt.Println(url)
	fmt.Println(queries)
	fmt.Println(body)
	return rest.Success(http.StatusOK), nil
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
