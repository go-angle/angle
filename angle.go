// Package angle Go framework with battery included.
package angle

import (
	"context"
	"os"
	"time"

	"github.com/go-angle/angle/bootstrap"
)

const (
	// Version of current.
	Version = "v0.1.0"
)

// Err returns error of angle.
func Err() error {
	return bootstrap.Err()
}

// Start start app.
func Start(configPath string, invokes ...interface{}) (<-chan os.Signal, context.CancelFunc, error) {
	bootstrap.SetConfigPath(configPath)
	return bootstrap.Start(invokes...)
}

// Stop stop app.
func Stop(timeout ...time.Duration) (cancel context.CancelFunc, err error) {
	return bootstrap.Stop(timeout...)
}

// RunOnce running functions during angle enviroment context.
func RunOnce(configPath string, invokes ...interface{}) (cancel context.CancelFunc, err error) {
	_, cancel, err = Start(configPath, invokes...)
	if err != nil {
		return cancel, err
	}
	return Stop()
}
