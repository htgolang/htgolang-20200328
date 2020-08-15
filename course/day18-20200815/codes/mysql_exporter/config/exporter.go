package config

type AuthConfig struct {
	User     string
	Password string
}

type WebConfig struct {
	Addr string
	Auth *AuthConfig
}

type MySQLConfig struct{}

type ExporterConfig struct {
	Web   *WebConfig
	MySQL *MySQLConfig
}
