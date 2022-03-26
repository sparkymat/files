package config

import "github.com/caarlos0/env"

func New() (*Service, error) {
	var envValues envValues

	if err := env.Parse(&envValues); err != nil {
		return nil, err
	}

	return &Service{
		envValues: envValues,
	}, nil
}

type Service struct {
	envValues envValues
}

type envValues struct {
	Username      string `env:"USERNAME,required"`
	Password      string `env:"PASSWORD,required"`
	Port          int    `env:"PORT" envDefault:"8080"`
	RootFolder    string `env:"ROOT_FOLDER,required"`
	SessionSecret string `env:"SESSION_SECRET" envDefault:"changeme"`
}

func (s *Service) Username() string {
	return s.envValues.Username
}

func (s *Service) Password() string {
	return s.envValues.Password
}

func (s *Service) Port() int {
	return s.envValues.Port
}

func (s *Service) RootFolder() string {
	return s.envValues.RootFolder
}

func (s *Service) SessionSecret() string {
	return s.envValues.SessionSecret
}
