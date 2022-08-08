package bootstrap

import (
	"context"
	"os"
	"time"

	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/di"
	"go.uber.org/fx"
)

const ConfigEnvKey = "GO_ANGLE_CONF_PATH"

var app *fx.App

// Set config path of angle.
func SetConfigPath(path string) {
	os.Setenv(ConfigEnvKey, path)
}

// Get config path of angle.
func GetConfigPath() string {
	ret := os.Getenv(ConfigEnvKey)
	if ret != "" {
		return ret
	}
	return "conf/config.yml"
}

// Start the angle.
func Start(timeout ...time.Duration) (context.CancelFunc, error) {
	gOptions := di.Options()
	options := make([]fx.Option, len(gOptions)+1, 0)
	options = append(options, configOptionProvider())
	for _, o := range gOptions {
		options = append(options, o)
	}
	app = fx.New(options...)
	ctx, cancel := getCtx(timeout...)
	return cancel, app.Start(ctx)
}

func configOptionProvider() fx.Option {
	configPath := GetConfigPath()
	reader, err := os.Open(configPath)
	if err != nil {
		panic(err)
	}
	// Provide ConfigOption to roll the world
	return fx.Provide(func() *config.Option {
		return &config.Option{
			Reader: reader,
		}
	})
}

// Stop the angle.
func Stop(timeout ...time.Duration) (context.CancelFunc, error) {
	ctx, cancel := getCtx(timeout...)
	return cancel, app.Stop(ctx)
}

func getCtx(timeout ...time.Duration) (context.Context, context.CancelFunc) {
	if len(timeout) > 0 {
		return context.WithTimeout(context.Background(), timeout[0])
	} else {
		return context.WithCancel(context.Background())
	}
}
