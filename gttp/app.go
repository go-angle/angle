package gttp

import (
	"github.com/gin-gonic/gin"

	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/di"
)

// App Global App
type App struct {
	config *config.Config
	Engine *gin.Engine
}

// newApp Create cli app
func newApp(config *config.Config, p middlewareParams) (*App, error) {
	var r *gin.Engine

	middlewares := make([]gin.HandlerFunc, 0, 3)
	middlewares = append(middlewares, p.Middlewares...)

	if !config.IsDevelopment() {
		gin.SetMode(gin.ReleaseMode)
	}

	if gin.Mode() == gin.DebugMode {
		r = gin.Default()
	} else {
		r = gin.New()
		middlewares = append(middlewares, gin.Recovery())
	}
	r.Use(middlewares...)

	return &App{
		config: config,
		Engine: r,
	}, nil
}

// Name returns app name
func (app *App) Name() string {
	return app.config.Name
}

// ParseConfig parse `app` field in config file to an variable
func (app *App) ParseConfig(ptr interface{}) error {
	return app.config.UnmarshalApp(ptr)
}

func init() {
	di.Provide(newApp)
}
