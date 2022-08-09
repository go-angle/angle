package config

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"time"

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
	Name      string `yaml:"name"`
	Version   string `yaml:"version"`
	Stage     string `yaml:"stage"`
	SecretKey string `yaml:"secret-key"`

	HTTP HTTPConfig `yaml:"http"`

	App interface{} `yaml:"app"`
}

// HTTPConfig configuration of HTTP
type HTTPConfig struct {
	Listen              string `yaml:"listen"`
	ReadHeaderTimeoutMS uint32 `yaml:"read-header-timeout-ms"`
	ReadTimeoutMS       uint32 `yaml:"read-timeout-ms"`
	WriteTimeoutMS      uint32 `yaml:"write-timeout-ms"`

	CookieMaxAge int `yaml:"cookie-max-age"`
}

func (hc *HTTPConfig) Timeouts() (time.Duration, time.Duration, time.Duration) {
	rht := 10 * time.Second
	rt := time.Minute
	wt := time.Minute
	if hc.ReadHeaderTimeoutMS > 0 {
		rht = time.Duration(hc.ReadHeaderTimeoutMS) * time.Millisecond
	}
	if hc.ReadTimeoutMS > 0 {
		rt = time.Duration(hc.ReadTimeoutMS) * time.Millisecond
	}
	if hc.WriteTimeoutMS > 0 {
		wt = time.Duration(hc.WriteTimeoutMS) * time.Millisecond
	}
	return rht, rt, wt
}

func (hc *HTTPConfig) ListenAddr() string {
	host := "127.0.0.1"
	port := "8080"

	if hc.Listen != "" {
		parts := strings.Split(hc.Listen, ":")
		var c_port string
		var c_host string
		if len(parts) > 1 {
			c_port = parts[len(parts)-1]
			c_host = strings.Join(parts[0:len(parts)-1], ":")
		} else {
			c_host = hc.Listen
		}
		if c_host != "" {
			host = c_host
		}
		if c_port != "" {
			port = c_port
		}
	}
	env_host := os.Getenv("ANGLE_HTTP_HOST")
	env_port := os.Getenv("ANGLE_HTTP_PORT")
	if env_host != "" {
		host = env_host
	}
	if env_port != "" {
		port = env_port
	}
	return fmt.Sprintf("%s:%s", host, port)
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
