package utils

import (
	"github.com/BurntSushi/toml"
)

type mysqlConfig struct {
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
	Mysql      mysqlConfig
	HttpServer httpServerConfig
}

func ParseTomlConfig() (TomlConfig, error) {
	var tomlConfig TomlConfig
	_, err := toml.DecodeFile("ect/todolist.toml", &tomlConfig)
	return tomlConfig, err
}
