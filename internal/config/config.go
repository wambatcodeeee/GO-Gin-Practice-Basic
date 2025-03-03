package config

import (
	"github.com/naoina/toml"
	"os"
)

type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	conf := &Config{}

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(conf); err != nil {
		panic(err)
	} else {
		return conf
	}
}
