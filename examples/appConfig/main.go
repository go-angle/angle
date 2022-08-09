package main

import (
	"github.com/go-angle/angle"
	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/log"
)

type appConfig struct {
	MyConf string `yaml:"my-conf"`
}

var app = &appConfig{}

func main() {
	ch, _, err := angle.Start("config.yml", func(c *config.Config, logger log.Logger /* log may be not ready during bootstrap */) {
		if err := c.UnmarshalApp(app); err != nil {
			logger.Fatalf("parse app config failed with error: %v", err)
		}
		logger.Infof("parsed app config: %v", app)
		logger.Infof("and my-conf is: %s", app.MyConf)
	})
	if err != nil {
		log.Fatalf("bootstrap failed with error: %v", err)
	}
	<-ch
	angle.Stop()
}
