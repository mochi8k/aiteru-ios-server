package stores

import (
	"encoding/json"
	"fmt"

	"gopkg.in/redis.v5"

	"github.com/mochi8k/aiteru-server/app/models"
	. "github.com/mochi8k/aiteru-server/config"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     Config.Redis.Host + ":" + Config.Redis.Port,
		Password: Config.Redis.Password,
		DB:       0,
	})

	fmt.Printf("Client: %v", client)

	pong, err := client.Ping().Result()
	fmt.Printf("Client Init: %v-%v\n", pong, err)
}

func GetSession(accessToken string) *models.Session {
	userString, err := client.Get(accessToken).Result()
	if userString == "" || err == redis.Nil {
		return nil
	}

	fmt.Printf("UserString by Redis: %v\n", userString)

	var user *models.User
	json.Unmarshal([]byte(userString), &user)

	if user.GetID() == "" {
		return nil
	}

	return &models.Session{
		AccessToken: accessToken,
		User:        *user,
	}
}

func AddSession(session *models.Session) {
	marshaledUser, _ := json.Marshal(session.GetUser())

	err := client.Set(session.GetAccessToken(), string(marshaledUser), 0).Err()

	if err != nil {
		fmt.Printf("Error add session: %v\n", err)
	} else {
		val, _ := client.Get(session.GetAccessToken()).Result()
		fmt.Printf("Add to Redis: %v\n", val)
	}
}
