package config

import (
	"io"
	"io/ioutil"

	"github.com/go-angle/angle/di"
	"gopkg.in/yaml.v3"
)

const (
	// StageDev development environment.
	StageDev = "development"
	// StageTest test environment.
	StageTest = "test"
	// StageStaging staging environment.
	StageStaging = "staging"
	// StageProd production environment.
	StageProd = "production"
)

// Config configuration.
type Config struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Stage   string `yaml:"stage"`

	App interface{} `yaml:"app"`
}

// IsDevelopment returns true if current stage is development.
func (c *Config) IsDevelopment() bool {
	return c.Stage == StageDev
}

// UnmarshalApp configuration to out.
func (c *Config) UnmarshalApp(outPtr interface{}) error {
	byts, err := yaml.Marshal(c.App)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(byts, outPtr)
}

// Option that provide by applications
type Option struct {
	Reader io.ReadCloser
}

// G global config instance
var G *Config

// newConfig create application config via `ConfigOption`
func newConfig(option *Option) (*Config, error) {
	var err error
	G, err = parseConfig(option.Reader)
	return G, err
}

func parseConfig(r io.ReadCloser) (*Config, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	c := &Config{}
	if err := yaml.Unmarshal(data, c); err != nil {
		return nil, err
	}
	return c, nil
}

func init() {
	di.Provide(newConfig)
}
