# angle

> angle is a Go framework with battery included.

[![build](https://github.com/go-angle/angle/actions/workflows/build.yml/badge.svg)](https://github.com/go-angle/angle/actions/workflows/build.yml) [![Go Reference](https://pkg.go.dev/badge/github.com/go-angle/angle.svg)](https://pkg.go.dev/github.com/go-angle/angle)

## Highlight Features

- Dependency Injection via [fx](https://pkg.go.dev/go.uber.org/fx).
- Fast structure log via [zerolog](https://github.com/rs/zerolog#benchmarks).
- And more...

## Getting Started

Install angle:

``` shell
go get github.com/go-angle/angle
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


More details please see [examples](examples).
