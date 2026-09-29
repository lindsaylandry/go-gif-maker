package config

import (
  "gopkg.in/yaml.v2"
  "os"
)

type Config struct {
	Team string `yaml:"team"`
}

func NewConfig() (*Config, error) {
  c := Config{}

  data, err := os.ReadFile("./configs/config.yaml")
  if err != nil {
    return &c, err
  }

  err = yaml.Unmarshal([]byte(data), &c)
  return &c, err
}
