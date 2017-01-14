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

func (p Place) GetID() string {
	return p.ID
}

func (p *Place) SetStatus(placeStatus *PlaceStatus) {
	p.Status = placeStatus
}

func (p *Place) IsOpen() bool {
	return p.Status.IsOpen
}
