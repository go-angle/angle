package auth

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-angle/angle/config"
	"github.com/go-angle/angle/gh"
	"github.com/go-angle/angle/secure"
)

// cookieState state saved in cookie and encoded by JWT
type cookieState struct {
	jwt secure.Signer
	app *gh.App

	claims secure.SignClaims

	cookieMaxAge int
}

func newCookieState(t secure.Signer, app *gh.App, c *config.Config) StateKeeper {
	cookieMaxAge := c.HTTP.CookieMaxAge
	if cookieMaxAge <= 0 {
		cookieMaxAge = 7200
	}

	return &cookieState{
		jwt:          t,
		app:          app,
		claims:       make(secure.SignClaims),
		cookieMaxAge: cookieMaxAge,
	}
}

// String fmt.Stringer
func (cookieState) String() string {
	return "cookieState"
}

// Set set state item
func (state *cookieState) Set(c *gin.Context, key string, value interface{}) {
	state.claims[key] = value

	// write to context duplicated
	if c != nil {
		c.Set(key, value)
	}
}

// Get get state item
func (state *cookieState) Get(key string) (interface{}, bool) {
	v, ok := state.claims[key]
	return v, ok
}

// GetInt get state item and convert it to int
func (state *cookieState) GetInt(key string) (int, bool) {
	if v, ok := state.Get(key); ok {
		v, ok := v.(int)
		return v, ok
	}
	return 0, false
}

// GetString get state item and convert it to string
func (state *cookieState) GetString(key string) (string, bool) {
	if v, ok := state.Get(key); ok {
		v, ok := v.(string)
		return v, ok
	}
	return "", false
}

// GetFloat64 get float64 state item and convert it to float64
func (state *cookieState) GetFloat64(key string) (float64, bool) {
	if v, ok := state.Get(key); ok {
		v, ok := v.(float64)
		return v, ok
	}
	return 0.0, false
}

// Restore restore state from gin.Context
func (state *cookieState) Restore(c *gin.Context) error {
	cookie, err := c.Cookie(state.CookieName())
	if err != nil {
		return err
	}
	claims, err := state.jwt.Validate(cookie)
	if err != nil {
		return err
	}

	if claims != nil {
		state.claims = claims
	}

	for k := range state.claims {
		c.Set(k, state.claims[k])
	}

	return err

}

// Save save state
func (state *cookieState) Save(c *gin.Context) error {
	value, err := state.EncodeCookieValue()
	if err != nil {
		return err
	}

	c.SetCookie(state.CookieName(), value, state.cookieMaxAge, "/", state.cookieDomain(c), false, true)
	return nil
}

// EncodeCookieValue returns cookie value singed with JWT
func (state *cookieState) EncodeCookieValue() (string, error) {
	return state.jwt.Sign(state.claims)
}

// Clear clear state
func (state *cookieState) Clear(c *gin.Context) error {
	c.SetCookie(state.CookieName(), "", -1, "/", state.cookieDomain(c), false, true)
	return nil
}

// cookieDomain returns cookieDomain
func (state *cookieState) cookieDomain(c *gin.Context) string {
	if c.Request.Referer() != "" {
		r := c.Request.Referer()
		parsed, err := url.Parse(r)
		if err != nil {
			return ""
		}
		parts := strings.Split(parsed.Host, ":")
		if len(parts) > 0 {
			stop := len(parts) - 1
			return strings.Join(parts[:stop], ":")
		}
		return ""
	}
	return ""
}

// CookieName returns name of cookie
func (state *cookieState) CookieName() string {
	return fmt.Sprintf("__angle_%s__", state.app.Name())
}
