package controllers

import (
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/mochi8k/aiteru-ios-server/app/models"
)

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

func errorChecker(err error) {
	if err != nil {
		panic(err)
	}
}
