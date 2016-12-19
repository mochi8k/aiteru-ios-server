package models

import (
	"github.com/satori/go.uuid"
)

type sessionManger struct {
	sessionsMap map[uuid.UUID]*session
}

func (sm *sessionManger) addSession(s *session) {
	sm.sessionsMap[s.GetAccessToken()] = s
}

func (sm sessionManger) GetSession(accessToken uuid.UUID) *session {
	return sm.sessionsMap[accessToken]
}

type session struct {
	AccessToken uuid.UUID
	User        User
}

func (s session) GetAccessToken() uuid.UUID {
	return s.AccessToken
}

func newSession(user User) *session {
	session := &session{
		AccessToken: uuid.NewV4(),
		User:        user,
	}
	return session
}
