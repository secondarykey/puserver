package config

import "golang.org/x/xerrors"

var gConfig *Config

func init() {
	gConfig = defaultConfig()
}

type Config struct {
	Server  string
	Port    int
	Context string

	Jar string
}

func defaultConfig() *Config {
	var c Config
	c.Server = ""
	c.Port = 8080
	c.Context = ""
	c.Jar = "plantuml-nodot.jar"
	return &c
}

func Get() *Config {
	return gConfig
}

func Set(opts ...Option) error {
	if opts == nil {
		return nil
	}
	for _, opt := range opts {
		err := opt(gConfig)
		if err != nil {
			return xerrors.Errorf("Option.call() error: %w", err)
		}
	}
	return nil
}
