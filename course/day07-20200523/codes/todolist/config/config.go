package config

type config struct {
	LoginRetry int
}

func New() *config {
	return &config{
		LoginRetry: 3,
	}
}

var Config = New()
