package stores

import (
	"fmt"
	"github.com/mochi8k/aiteru-ios-server/app/models"
)

type sessionStore struct {
	sessionsMap map[string]*models.Session
}

var ss = &sessionStore{
	sessionsMap: map[string]*models.Session{},
}

func AddSession(s *models.Session) {
	fmt.Println(s.GetAccessToken())
	ss.sessionsMap[s.GetAccessToken()] = s
}

func GetSession(accessToken string) *models.Session {
	return ss.sessionsMap[accessToken]
}
