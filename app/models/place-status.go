package models

type PlaceStatus struct {
	PlaceID       string `json:"placeId"`
	IsOpen        bool   `json:"isOpen"`
	UpdatedAt     string `json:"updatedAt"`
	UpdatedUserID string `json:"updatedUserId"`
}

func (p PlaceStatus) GetPlaceID() string {
	return p.PlaceID
}
