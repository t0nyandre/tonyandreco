package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gookit/validate"
)

const (
	defaultAppPort = 4000
	defaultAppHost = "localhost"
	defaultAppEnv  = "development"
	defaultAppName = "tonyandreco"
)

type Config struct {
	Port        int    `json:"port" env:"APP_PORT" validate:"required|numeric"`
	Hostname    string `json:"hostname" env:"APP_HOST" validate:"required|string"`
	Environment string `json:"environment" env:"APP_ENV" validate:"required|string"`
	Name        string `json:"name" env:"APP_NAME" validate:"required|string"`
}

func (c *Config) Validate() error {
	v := validate.Struct(c)
	if v.Validate() {
		return nil
	}
	return v.Errors
}

func Load(file string) (*Config, error) {
	c := &Config{
		Port:        defaultAppPort,
		Hostname:    defaultAppHost,
		Environment: defaultAppEnv,
		Name:        defaultAppName,
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bytes, c); err != nil {
		return nil, err
	}

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return c, nil

}
