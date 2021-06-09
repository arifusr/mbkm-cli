package config

import (
	"errors"
	"fmt"
	"os"
)

type Config struct {
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) ValidateConfig() error {
	if os.Getenv("MIGRATION_DIRECTORY") == "" {
		fmt.Print("please setup MIGRATION_DIRECTORY on .mbkm-cli")
		return errors.New("no MIGRATION_DIRECTORY")
	}
	if _, err := os.Stat(os.Getenv("MIGRATION_DIRECTORY")); os.IsNotExist(err) {
		fmt.Print("directory " + os.Getenv("MIGRATION_DIRECTORY") + " not found")
		return errors.New("no MIGRATION_DIRECTORY")
	}

	return nil
}
