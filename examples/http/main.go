package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-angle/angle"
	"github.com/go-angle/angle/di"
	"github.com/go-angle/angle/gttp"
	"github.com/go-angle/angle/log"
	"go.uber.org/fx"
)

type appConfig struct {
	MyConf string `yaml:"my-conf"`
}

var app = &appConfig{}

func main() {
	ch, _, err := angle.Start("config.yml")
	if err != nil {
		log.Fatalf("bootstrap failed with error: %v", err)
	}
	<-ch
	angle.Stop()
}

type routerParams struct {
	fx.In

	Default *gin.RouterGroup `name:"api"`
}

func init() {
	gttp.ProvideRouterGroup("api", func(app *gttp.App) *gin.RouterGroup {
		return app.Engine.Group("api")
	})

	di.Invoke(func(r routerParams) {
		r.Default.GET("/", func(c *gin.Context) {
			c.Writer.WriteString("Hello, world!")
		})
	})
}
