package controllers

import (
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/mochi8k/aiteru-ios-server/app/models"
)

const (
	parseFormat = "2006-01-02 15:04:05"
)

func toRFC3339(dateTime string) string {
	if dateTime == "" {
		return ""
	}
	t, _ := time.Parse(parseFormat, dateTime)
	return t.Format(time.RFC3339)
}

func toUser(scanner sq.RowScanner) *models.User {
	var id, name, createdAt, createdUserID, updatedAt, updatedUserID string
	scanner.Scan(&id, &name, &createdAt, &createdUserID, &updatedAt, &updatedUserID)
	return &models.User{
		ID:            id,
		Name:          name,
		CreatedAt:     toRFC3339(createdAt),
		CreatedUserID: createdUserID,
		UpdatedAt:     toRFC3339(updatedAt),
		UpdatedUserID: updatedUserID,
	}
}

func toPlace(scanner sq.RowScanner) *models.Place {
	var id, placeName, ownerIDs, collaboratorIDs, createdAt, createdUserID, updatedAt, updatedUserID string
	scanner.Scan(&id, &placeName, &ownerIDs, &collaboratorIDs, &createdUserID, &createdAt, &updatedUserID, &updatedAt)
	return &models.Place{
		ID:              id,
		Name:            placeName,
		OwnerIDs:        strings.Split(ownerIDs, ","),
		CollaboratorIDs: strings.Split(collaboratorIDs, ","),
		CreatedAt:       toRFC3339(createdAt),
		CreatedUserID:   createdUserID,
		UpdatedAt:       toRFC3339(updatedAt),
		UpdatedUserID:   updatedUserID,
	}
}

func toPlaceStatus(scanner sq.RowScanner) *models.PlaceStatus {
	var placeID, updatedAt, updatedUserID string
	var isOpen bool
	scanner.Scan(&placeID, &isOpen, &updatedAt, &updatedUserID)
	return &models.PlaceStatus{
		PlaceID:       placeID,
		IsOpen:        isOpen,
		UpdatedAt:     toRFC3339(updatedAt),
		UpdatedUserID: updatedUserID,
	}
}

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
