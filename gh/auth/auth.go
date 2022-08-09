package auth

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"

	"github.com/go-angle/angle/di"
)

// Authenticator defines how to authenticate users.
type Authenticator interface {
	// EnableGlobal to enable gin middleware globally to reject all requests that are not authencated.
	EnableGlobal(option *Option)

	// SkipPathGlobal to ignore then given request path for checking their authencation globally.
	SkipPathGlobal(path string)

	// EnableGroup enable middleware on the given router group to reject all requests that are not authencated.
	EnableGroup(r *gin.RouterGroup, option *Option)

	// SkipPath to ignore the given request for checking their authencation on the given router group.
	SkipPath(r *gin.RouterGroup, path string)

	// Username returns current username.
	Username(c *gin.Context) (string, bool)
}

// StateKeeper declare how to keep the authentication state.
type StateKeeper interface {
	fmt.Stringer

	// Set set state.
	Set(c *gin.Context, key string, value interface{})

	// Get item from state.
	Get(key string) (interface{}, bool)

	// GetString get string item from state
	GetString(key string) (string, bool)

	// GetInt get int itm from state
	GetInt(key string) (int, bool)

	// GetFloat64 get float64 item from state
	GetFloat64(key string) (float64, bool)

	// Restore restore state and sync all state to gin.Context
	Restore(c *gin.Context) error

	// Save state, such as cookie or something else.
	Save(c *gin.Context) error

	// Clear clear state
	Clear(c *gin.Context) error

	// CookieName returns cookie name
	CookieName() string

	// EncodeCookieValue encode cookie value
	EncodeCookieValue() (string, error)
}

// Option option to enable Authenticator
type Option struct {
	LoginEndpoint  string
	LogoutEndpoint string
}

// EnableCookieStateKeeper enable jwt cookie to keep state
func EnableCookieStateKeeper() {
	ProvideStateKeeper(newCookieStateKeeper)
}

// ProvideAuthenticator provide auther
func ProvideAuthenticator(name string, f interface{}) {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("must use a function to provide Authenticator")
	}

	if t.NumOut() == 0 {
		panic("must returns a value from provide function")
	}
	in := reflect.TypeOf((*Authenticator)(nil)).Elem()
	if !t.Out(0).Implements(in) {
		panic("provide function's first return value must implements Authenticator")
	}
	di.Provide(fx.Annotated{
		Name:   name,
		Target: f,
	})
}

// ProvideStateKeeper provide custom StateKeeper implementation
func ProvideStateKeeper(newFn interface{}) {
	t := reflect.TypeOf(newFn)
	if t.Kind() != reflect.Func {
		panic("must use a function to provide Authenticator")
	}

	if t.NumOut() == 0 {
		panic("must returns a value from provide function")
	}
	in := reflect.TypeOf((*StateKeeper)(nil)).Elem()
	if !t.Out(0).Implements(in) {
		panic("provide function's first return value must implements StateKeeper")
	}
	di.Provide(newFn)
}
