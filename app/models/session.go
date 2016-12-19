package models

import (
	"github.com/satori/go.uuid"
)

type Session struct {
	AccessToken string
	User        User
}

func (s Session) GetAccessToken() string {
	return s.AccessToken
}

func NewSession(user User) *Session {
	session := &Session{
		AccessToken: uuid.NewV4().String(),
		User:        user,
	}
	return session
}
