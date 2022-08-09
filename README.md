# angle

> angle is a Go framework with battery included.

[![build](https://github.com/go-angle/angle/actions/workflows/build.yml/badge.svg)](https://github.com/go-angle/angle/actions/workflows/build.yml)

## Get Started

Install angle:

``` shell
go get https://github.com/go-angle/angle
```

Create `config.yml`:

``` yaml
name: minimal
stage: development
```

Then here is the code:

``` go
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
```

Now we can run it:

``` shell
go run main.go
```
