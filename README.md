# angle

> angle is a Go framework with battery included.

[![build](https://github.com/go-angle/angle/actions/workflows/build.yml/badge.svg)](https://github.com/go-angle/angle/actions/workflows/build.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/go-angle/angle.svg)](https://pkg.go.dev/github.com/go-angle/angle)
![GitHub](https://img.shields.io/github/license/go-angle/angle)

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

### Automatically Binding

``` go
package main

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/go-angle/angle"
	"github.com/go-angle/angle/di"
	"github.com/go-angle/angle/gh"
	"github.com/go-angle/angle/log"
	"go.uber.org/fx"
)

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
	gh.ProvideRouterGroup("api", func(app *gh.App) *gin.RouterGroup {
		return app.Engine.Group("api")
	})

	di.Invoke(func(r routerParams) {
		r.Default.POST("/users", gh.MustBind(createUser).HandlerFunc())
	})
}

type UserReq struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type UserResponse struct {
	Ok bool `json:"ok"`
	ID int  `json:"id"`
}

func createUser(req *UserReq) (*UserResponse, error) {
	if req.Age <= 0 {
		return nil, errors.New("no, it's impossible")
	}
	return &UserResponse{
		Ok: true,
		ID: 1,
	}, nil
}

```

### More

More details please see [examples](examples).
