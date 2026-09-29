package config

import (
  "gopkg.in/yaml.v2"
  "os"
)

type Config struct {
	Path string `yaml:"path"`
	InputFiles []string `yaml:"input_files"`
	OutputFiles string `yaml:"output_files"`
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
