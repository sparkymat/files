package config

import (
	"errors"
	"fmt"

	"github.com/caarlos0/env"
)

var ErrAuthDetailsMissing = errors.New("auth details missing")

func New() (*Service, error) {
	var envValues envValues

	if err := env.Parse(&envValues); err != nil {
		return nil, fmt.Errorf("failed to parse env values. err: %w", err)
	}

	// If auth is not disabled, ensure username and password exists
	if !envValues.AuthDisabled && (envValues.Username == "" || envValues.Password == "") {
		return nil, ErrAuthDetailsMissing
	}

	return &Service{
		envValues: envValues,
	}, nil
}

type Service struct {
	envValues envValues
}

type envValues struct {
	AuthDisabled  bool   `env:"AUTH_DISABLED"`
	Username      string `env:"USERNAME"`
	Password      string `env:"PASSWORD"`
	Port          int    `env:"PORT" envDefault:"8080"`
	RootFolder    string `env:"ROOT_FOLDER,required"`
	SessionSecret string `env:"SESSION_SECRET" envDefault:"changeme"`
}

func (s *Service) AuthDisabled() bool {
	return s.envValues.AuthDisabled
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
