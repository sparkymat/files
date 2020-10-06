package config

import "github.com/caarlos0/env"

type API interface {
	Username() string
	Password() string
	Port() int
}

func New() API {
	var envValues envValues

	if err := env.Parse(&envValues); err != nil {
		panic(err)
	}

	return &service{
		envValues: envValues,
	}
}

type service struct {
	envValues envValues
}

type envValues struct {
	Username string `env:"USERNAME,required"`
	Password string `env:"PASSWORD,required"`
	Port     int    `env:"PORT" envDefault:"8080"`
}

func (s *service) Username() string {
	return s.envValues.Username
}

func (s *service) Password() string {
	return s.envValues.Password
}

func (s *service) Port() int {
	return s.envValues.Port
}
