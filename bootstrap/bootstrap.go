package bootstrap

import (
	"context"

	"github.com/go-angle/angle/di"
	"go.uber.org/fx"
)

func Start(ctx context.Context) error {
	app := fx.New(di.Options()...)
	return app.Start(ctx)
}
