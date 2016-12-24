package models

type User struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	CreatedAt     string `json:"createdAt"`
	CreatedUserID string `json:"createdUserId"`
	UpdatedAt     string `json:"updatedAt"`
	UpdatedUserID string `json:"updatedUserId"`
}

func (u User) GetID() string {
	return u.ID
}
