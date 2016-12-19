package models

import (
	"github.com/satori/go.uuid"
)

type Session struct {
	AccessToken uuid.UUID
	User        User
}

func (s Session) GetAccessToken() uuid.UUID {
	return s.AccessToken
}

func NewSession(user User) *Session {
	session := &Session{
		AccessToken: uuid.NewV4(),
		User:        user,
	}
	return session
}
