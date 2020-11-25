package config

type Config struct {
	Service  service
	LogLevel string `envconfig:"LOGLEVEL"`
}

type service struct {
	Parser string `envconfig:"SERVICE_PARSER"`
}
