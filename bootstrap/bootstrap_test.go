package bootstrap

import (
	"testing"
	"time"
)

func TestBootstrap(t *testing.T) {
	SetConfigPath("config.yml")
	_, err := Start(time.Second)
	if err != nil {
		t.Fatalf("start failed %s", err)
	}
	_, err = Stop(time.Second)
	if err != nil {
		t.Fatalf("Stop failed %s", err)
	}
}
