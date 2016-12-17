package models

type Place struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Owners          []string `json:"owners"`
	Collaborators   []string `json:"collaborators"`
	CreatedAt       string   `json:"createdAt"`
	CreatedUserName string   `json:"createdUserName"`
	UpdatedAt       string   `json:"updatedAt"`
	UpdatedUserName string   `json:"updatedUserName"`
}
