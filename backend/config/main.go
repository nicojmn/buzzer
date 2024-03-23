package config

import (
	"log"
	"os"
	"github.com/go-yaml/yaml"
)

type Config struct {
	Teams struct {
		MaxNumber int
	}
}

func LoadConfig (filename string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(filename)

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	data := yaml.NewDecoder(file)
	if err := data.Decode(config); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}