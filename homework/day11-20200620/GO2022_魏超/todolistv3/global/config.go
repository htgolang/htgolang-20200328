package global

import (
	"log"

	"github.com/BurntSushi/toml"
)

type databaseConfig struct {
	Device   string
	UserName string
	Password string
	Host     string
	Port     int
	DBName   string
}

type httpServerConfig struct {
	Host string
	Port int
}

type TomlConfig struct {
	Database   databaseConfig
	HttpServer httpServerConfig
}

var Config TomlConfig

func init() {
	var err error
	Config, err = ParseTomlConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func ParseTomlConfig() (TomlConfig, error) {
	var tomlConfig TomlConfig
	_, err := toml.DecodeFile("config/app.toml", &tomlConfig)
	return tomlConfig, err
}
