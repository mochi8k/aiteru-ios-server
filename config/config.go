package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	Port  string `default:"8000"`
	MySQL mySQL
}

type mySQL struct {
	User       string `default:"root"`
	Password   string `default:""`
	DB         string `default:"aiteru"`
	Connection string
}

var Config config

func init() {
	err := envconfig.Process("app", &Config)

	if err != nil {
		panic(err)

	}

	Config.MySQL.Connection =
		Config.MySQL.User + ":" + Config.MySQL.Password + "@/" + Config.MySQL.DB

	fmt.Printf("Config: %+v\n", Config)
}
