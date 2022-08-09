// Package gh is an abbreviation of gin http.
package gh

import (
	"context"
	"net"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/di"
	"github.com/go-angle/angle/log"
)

var routerGroupProvided bool

// ProvideRouterGroup define a router group
func ProvideRouterGroup(groupName string, fn interface{}) {
	t := reflect.TypeOf(fn)

	if t.Kind() != reflect.Func {
		panic("must use a function to return *gin.RouterGroup to provide HTTP Router Group")
	}
	if t.NumOut() < 1 {
		panic("to provide HTTP Router Group, a ptr of gin.RouterGroup must be returned")
	}
	switch reflect.New(t.Out(0)).Elem().Interface().(type) {
	case *gin.RouterGroup:
	default:
		panic("to provide HTTP Router Group, a ptr of gin.RouterGroup must be returned")
	}

	di.Provide(fx.Annotated{
		Name:   groupName,
		Target: fn,
	})
	routerGroupProvided = true
}

type params struct {
	fx.In

	Log     log.Logger
	Default *gin.RouterGroup `name:"default" optional:"true"`
}

// provideDefault provide default router group
func provideDefault(p params) *gin.RouterGroup {
	if p.Default == nil {
		p.Log.Warn("NO_DEFAULT_HTTP_ROUTER_GROUP_PROVIDED")
	}
	return p.Default
}

type middlewareParams struct {
	fx.In

	Middlewares []gin.HandlerFunc `group:"middleware"`
}

// ProvideMiddleware provide middleware
func ProvideMiddleware(fn interface{}) {
	di.Provide(fx.Annotated{
		Group:  "middleware",
		Target: fn,
	})
}

// serveHTTPIfNeeded listen & serve HTTP
func serveHTTPIfNeeded(lc fx.Lifecycle, app *App, log log.Logger, c *config.Config) {
	if !routerGroupProvided {
		return
	}

	addr := c.HTTP.ListenAddr()
	rht, rt, wt := c.HTTP.Timeouts()
	srv := &http.Server{
		Addr:    addr,
		Handler: app.Engine,

		ReadHeaderTimeout: rht,
		ReadTimeout:       rt,

		// WriteTimeout compute from header readed(https compute from accept).
		// When exceed write timeout, Nginx will raised a 502 Bad Gateway error.
		WriteTimeout: wt,
	}
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			tcpListner, err := net.Listen("tcp", addr)
			if err != nil {
				return err
			}
			log.Infof("Serving HTTP Service on %s ...", addr)
			go func() {
				if err := srv.Serve(tcpListner); err != nil && err != http.ErrServerClosed {
					log.Fatalf("Serving HTTP Service failed with error %s", err.Error())
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Infof("Shutdown HTTP Service on %s...", addr)
			err := srv.Shutdown(ctx)
			if err != nil && err != http.ErrServerClosed {
				log.Errorf("HTTP Service Shutdown failed with error %s", err.Error())
			} else {
				log.Info("HTTP Service has shutdown")
			}
			return err
		},
	})
}

func init() {
	di.Invoke(serveHTTPIfNeeded)
	di.Provide(provideDefault)
}
