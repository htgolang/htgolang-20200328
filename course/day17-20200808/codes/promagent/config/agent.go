package config

import "time"

type AgentConfig struct {
	UUID         string        `mapstructure:"-"`
	Addr         string        `mapstructure:"addr"`
	ServerConfig *ServerConfig `mapstructure:"server"`
	LogConfig    *LogConfig    `mapstructure:"log"`
	TaskConfig   *TaskConfig   `mapstructure:"task"`
}

type ServerConfig struct {
	Addr  string `mapstructure:"addr"`
	Token string `mapstructure:"token"`
}

type LogConfig struct {
	Filename   string `mapstructure:"filename"`
	Maxsize    int    `mapstructure:"maxsize"`
	Maxbackups int    `mapstructure:"maxbackups"`
	Compress   bool   `mapstructure:"compress"`
}

type TaskConfig struct {
	Register RegisterConfig `mapstructure:"register"`
	Profile  ProfileConfig  `mapstructure:"profile"`
}

type RegisterConfig struct {
	Interval time.Duration `mapstructure:"interval"`
}

type ProfileConfig struct {
	Interval time.Duration `mapstructure:"interval"`
	Tpl      string        `mapstructure:"tpl"`
	Output   string        `mapstructure:"output"`
}
