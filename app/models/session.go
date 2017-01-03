package models

import (
	"github.com/satori/go.uuid"
)

type Session struct {
	AccessToken string // `json:"accessToken"`
	User        User   `json:"user"`
}

func (s Session) GetAccessToken() string {
	return s.AccessToken
}

func (s Session) GetUser() User {
	return s.User
}

func NewSession(user User) *Session {
	session := &Session{
		AccessToken: uuid.NewV4().String(),
		User:        user,
	}
	return session
}
