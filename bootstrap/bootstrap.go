package bootstrap

import (
	"context"
	"os"

	"github.com/go-angle/angle/di"
	"go.uber.org/fx"
)

const ConfigEnvKey = "GO_ANGLE_CONF_PATH"

// Set config path of angle.
func SetConfigPath(path string) {
	os.Setenv(ConfigEnvKey, path)
}

func Start(ctx context.Context) error {
	app := fx.New(di.Options()...)
	return app.Start(ctx)
}
