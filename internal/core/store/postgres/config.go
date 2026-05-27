package core_postgres

import "github.com/kelseyhightower/envconfig"

type Config struct {
	User     string `envocnfig:"USER" required:"true"`
	Password string `envocnfig:"PASSWORD" required:"true"`
	DB       string `envocnfig:"DB" required:"true"`
	Host     string `envocnfig:"HOST" required:"true"`
	Port     string `envocnfig:"PORT" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config
	if err := envconfig.Process("POSTGRES", &config); err != nil {
		return Config{}, err
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		panic(err)
	}

	return config
}
