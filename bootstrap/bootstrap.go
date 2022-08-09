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

// Err returns error of fx.App
func Err() error {
	return app.Err()
}

// Start the angle.
func Start(invokes ...interface{}) (<-chan os.Signal, context.CancelFunc, error) {
	gOptions := di.Options()
	options := make([]fx.Option, 0, len(gOptions)+len(invokes)+1)
	options = append(options, configOptionProvider())
	for _, o := range gOptions {
		options = append(options, o)
	}
	options = append(options, fx.Invoke(invokes...))
	app = fx.New(options...)
	ctx, cancel := getCtx()
	err := app.Start(ctx)
	return app.Done(), cancel, err
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
