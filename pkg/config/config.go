package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

var (
	config *Config
)

// option defines configuration option
type option struct {
	configFile string
}

// Init initializes `config` from the default config file.
// use `WithConfigFile` to specify the location of the config file
func Init(opts ...Option) error {
	opt := &option{}

	for _, optFunc := range opts {
		optFunc(opt)
	}

	out, err := ioutil.ReadFile(opt.configFile)
	if err != nil {
		return err
	}

	err = json.Unmarshal(out, &config)
	if err != nil {
		return err
	}

	return nil
}

// Option define an option for config package
type Option func(*option)

// WithConfigFile set `config` to use the given config file
func WithConfigFile(dev string, prod string) Option {
	return func(opt *option) {
		env := os.Getenv("APP_ENV")
		opt.configFile = prod
		if env != "staging" && env != "production" {
			opt.configFile = filepath.Join(filepath.Join(os.Getenv("GOPATH"), "src/github.com/reyhanfahlevi/poc-test"), dev)
		}
	}
}

// Get config
func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config
}
