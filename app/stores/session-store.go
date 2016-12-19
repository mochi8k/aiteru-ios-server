package stores

import (
	"fmt"
	"github.com/mochi8k/aiteru-ios-server/app/models"
	"github.com/satori/go.uuid"
)

type sessionStore struct {
	sessionsMap map[uuid.UUID]*models.Session
}

var ss = &sessionStore{
	sessionsMap: map[uuid.UUID]*models.Session{},
}

func AddSession(s *models.Session) {
	fmt.Println(s.GetAccessToken())
	ss.sessionsMap[s.GetAccessToken()] = s
}

func GetSession(accessToken uuid.UUID) *models.Session {
	return ss.sessionsMap[accessToken]
}
