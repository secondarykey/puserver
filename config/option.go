package config

import (
	"os"

	"golang.org/x/xerrors"
)

type Option func(*Config) error

func Host(server string) Option {
	return func(c *Config) error {
		c.Server = server
		return nil
	}
}

func Port(port int) Option {
	return func(c *Config) error {
		c.Port = port
		return nil
	}
}

func Context(ctx string) Option {
	return func(c *Config) error {
		c.Context = ctx
		return nil
	}
}

func Jar(path string) Option {
	return func(c *Config) error {
		c.Jar = path

		if _, err := os.Stat(path); err != nil {
			return xerrors.Errorf("Jar file not exists: %w", err)
		}

		return nil
	}
}
