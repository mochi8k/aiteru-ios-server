package models

type Place struct {
	ID              string       `json:"id"`
	Name            string       `json:"name"`
	OwnerIDs        []string     `json:"ownerIds"`
	CollaboratorIDs []string     `json:"collaboratorIds"`
	CreatedAt       string       `json:"createdAt"`
	CreatedUserID   string       `json:"createdUserId"`
	UpdatedAt       string       `json:"updatedAt"`
	UpdatedUserID   string       `json:"updatedUserId"`
	Status          *PlaceStatus `json:"status"`
}
