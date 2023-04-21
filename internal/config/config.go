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

type Default struct {
	AppPort int    `json:"app_port" env:"APP_PORT" validate:"required|numeric"`
	AppHost string `json:"app_host" env:"APP_HOST" validate:"required|string"`
	AppEnv  string `json:"app_env" env:"APP_ENV" validate:"required|string"`
	AppName string `json:"app_name" env:"APP_NAME" validate:"required|string"`
}

type Config struct {
	Defaults *Default `json:"defaults"`
}

func (c *Config) Validate() error {
	v := validate.Struct(c)
	if v.Validate() {
		return nil
	}
	return v.Errors
}

func Load(file string) (*Config, error) {
	c := Config{
		Defaults: &Default{
			AppPort: defaultAppPort,
			AppHost: defaultAppHost,
			AppEnv:  defaultAppEnv,
			AppName: defaultAppName,
		},
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	if err = json.Unmarshal(bytes, &c); err != nil {
		return nil, err
	}

	if err = c.Validate(); err != nil {
		return nil, err
	}

	return &c, nil

}
