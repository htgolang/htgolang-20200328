package config

type config struct {
	LoginRetry int
}

func NewConfig() *config {
	return &config{
		LoginRetry: 3,
	}
}

var Config = NewConfig()