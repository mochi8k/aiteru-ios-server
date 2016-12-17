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
	    group_concat(distinct u1.user_name) as owner_names,
	    group_concat(distinct u2.user_name) as collaborator_names,
	    u3.user_name as created_by,
	    p.created_at,
	    u4.user_name as updated_by,
	    p.updated_at
	  from places p
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
      users as u4 on p.updated_by=u3.id
	  group by
      p.id, u3.user_name, u4.user_name
	`)

	errorChecker(err)

	var places []*models.Place

	for res.Next() {
		var id int
		var placeName, ownerNames, collaboratorNames, createdBy, createdAt, updatedBy, updatedAt string
		res.Scan(&id, &placeName, &ownerNames, &collaboratorNames, &createdBy, &createdAt, &updatedBy, &updatedAt)
		place := &models.Place{
			ID:              id,
			Name:            placeName,
			Owners:          strings.Split(ownerNames, ","),
			Collaborators:   strings.Split(collaboratorNames, ","),
			CreatedAt:       createdAt,
			CreatedUserName: createdBy,
			UpdatedAt:       updatedAt,
			UpdatedUserName: updatedBy,
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
