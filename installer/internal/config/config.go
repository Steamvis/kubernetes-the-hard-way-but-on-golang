package config

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

var Cfg Config

type Config struct {
	JumpHost JumpHost `yaml:"jumpHost"`
}

func Parse(configPath string) error {
	f, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &Cfg)
	if err != nil {
		return err
	}

	return nil
}
