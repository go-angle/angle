package main

import (
	"github.com/go-angle/angle"
	"github.com/go-angle/angle/log"
)

func main() {
	ch, _, err := angle.Start("config.yml")
	if err != nil {
		log.Fatalf("bootstrap failed with error: %v", err)
	}
	<-ch
	angle.Stop()
}
