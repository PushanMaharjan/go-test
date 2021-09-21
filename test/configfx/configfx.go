package configfx

import (
	"io/ioutil"

	"go.uber.org/fx"
	"gopkg.in/yaml.v2"
)

// application config
type ApplicationConfig struct {
	Address string `yaml:"address"`
}

// config
type Config struct {
	ApplicationConfig `yaml:"application"`
}

// config to fx
func ProvideConfig() *Config {
	conf := Config{}
	data, err := ioutil.ReadFile("config/base.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(data), &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}

var Module = fx.Options(
	fx.Provide(ProvideConfig),
)
