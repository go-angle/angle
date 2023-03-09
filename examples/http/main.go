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
		r.Default.GET("/", func(c *gin.Context) {
			c.Writer.WriteString("Hello, world!")
		})
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
		return nil, errors.New("no, it's impossiable")
	}
	return &UserResponse{
		Ok: true,
		ID: 1,
	}, nil
}
